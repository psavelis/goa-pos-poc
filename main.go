//go:generate goagen bootstrap -d github.com/psavelis/goa-pos-poc/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/psavelis/goa-pos-poc/app"
	mgo "gopkg.in/mgo.v2"
    bson "gopkg.in/mgo.v2/bson"
)

type DataStore struct {
	session  *mgo.Session
}

func main() {
	// Create service
	service := goa.New("pos")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// MongoDB connection pool setup
	session, err := mgo.Dial("localhost")
	if(err != nil){
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic)

	DataStore.session = session

	// Mount "Purchase" controller
	c := NewPurchaseController(service)
	app.MountPurchaseController(service, c)

	cs := NewSwaggerController(service)
	app.MountSwaggerController(service, cs)

	// Start service
	if err := service.ListenAndServe(":5001"); err != nil {
		service.LogError("startup", "err", err)
	}
,
}
