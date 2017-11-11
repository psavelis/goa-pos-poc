//go:generate goagen bootstrap -d github.com/psavelis/goa-pos-poc/design

package main

import (
	"crypto/tls"
	"net"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/psavelis/goa-pos-poc/app"
	"github.com/psavelis/goa-pos-poc/controllers"
	mgo "gopkg.in/mgo.v2"
)

// DataStore holds mgo's session pool

func main() {

	// REVIEW: workaround
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

	dialInfo, err := mgo.ParseURL("mongodb://psavelis:psavelis@development-shard-00-00-ozch3.mongodb.net:27017,development-shard-00-01-ozch3.mongodb.net:27017,development-shard-00-02-ozch3.mongodb.net:27017/test?replicaSet=development-shard-0&authSource=admin")

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
	c := controllers.NewPurchaseController(service, database)
	app.MountPurchaseController(service, c)

	cs := controllers.NewSwaggerController(service)
	app.MountSwaggerController(service, cs)

	// Start service
	if err := service.ListenAndServe(":5000"); err != nil {
		service.LogError("startup", "err", err)
	}
}
