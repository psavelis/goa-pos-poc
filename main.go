//go:generate goagen bootstrap -d github.com/psavelis/goa-pos-poc/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/psavelis/goa-pos-poc/app"
)

func main() {
	// Create service
	service := goa.New("pos")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

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
