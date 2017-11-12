//go:generate goagen bootstrap -d github.com/psavelis/goa-pos-poc/design

package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/psavelis/goa-pos-poc/app"
	"github.com/psavelis/goa-pos-poc/controllers"
	mgo "gopkg.in/mgo.v2"
)

type Database struct {
}

func main() {

	// REVIEW: goa media types
	goa.ErrorMediaIdentifier = "application/json"

	// Create service
	service := goa.New("pos")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// MongoDB (Atlas) setup
	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = true

	mgoUser := os.Getenv("MONGO_USR")

	if mgoUser == "" {
		service.LogError("$MONGO_USR must be set")
	}

	mgoPassword := os.Getenv("MONGO_PWD")

	if mgoPassword == "" {
		service.LogError("$MONGO_PWD must be set")
	}

	dialInfo, err := mgo.ParseURL(fmt.Sprintf("mongodb://%s:%s@development-shard-00-00-ozch3.mongodb.net:27017,development-shard-00-01-ozch3.mongodb.net:27017,development-shard-00-02-ozch3.mongodb.net:27017/test?replicaSet=development-shard-0&authSource=admin", mgoUser, mgoPassword))

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	// services-pos database
	database := session.DB("services-pos")

	//Database.C("Purchase").RemoveAll(bson.M{})

	// Purchases collection index
	index := mgo.Index{
		Key:        []string{"locator"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = database.C("Purchase").EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	// Mount "Purchase" controller
	purchaseController := controllers.NewPurchaseController(service, database)
	app.MountPurchaseController(service, purchaseController)

	swaggerController := controllers.NewSwaggerController(service)
	app.MountSwaggerController(service, swaggerController)

	publicController := controllers.NewPublicController(service)
	app.MountPublicController(service, publicController)

	jsController := controllers.NewJsController(service)
	app.MountJsController(service, jsController)

	port := os.Getenv("PORT")

	if port == "" {
		service.LogError("$PORT must be set")
	}

	// Start service
	if err := service.ListenAndServe(fmt.Sprintf(":%s", port)); err != nil {
		service.LogError("startup", "err", err)
	}
}
