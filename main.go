//go:generate goagen bootstrap -d github.com/psavelis/goa-pos-poc/design

package main

import (
	"crypto/tls"
	"net"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/psavelis/goa-pos-poc/app"
	mgo "gopkg.in/mgo.v2"
)

// DataStore holds mgo's session pool
var (
	Database *mgo.Database
)

func main() {
	// Create service
	service := goa.New("pos")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = true

	dialInfo, err := mgo.ParseURL("mongodb://c:c@dev-cluster-shard-00-00-ozch3.mongodb.net:27017,dev-cluster-shard-00-01-ozch3.mongodb.net:27017,dev-cluster-shard-00-02-ozch3.mongodb.net:27017/test?replicaSet=dev-cluster-shard-0&authSource=admin")

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

	Database = session.DB("services-pos")

	// Mount "Purchase" controller
	c := NewPurchaseController(service)
	app.MountPurchaseController(service, c)

	cs := NewSwaggerController(service)
	app.MountSwaggerController(service, cs)

	// Start service
	if err := service.ListenAndServe(":5001"); err != nil {
		service.LogError("startup", "err", err)
	}
}
