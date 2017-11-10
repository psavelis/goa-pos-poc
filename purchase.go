package main

import (
	"github.com/goadesign/goa"
	"github.com/psavelis/goa-pos-poc/app"
	"gopkg.in/mgo.v2/bson"
)

// PurchaseController implements the Purchase resource.
type PurchaseController struct {
	*goa.Controller
}

// NewPurchaseController creates a Purchase controller.
func NewPurchaseController(service *goa.Service) *PurchaseController {
	return &PurchaseController{Controller: service.NewController("PurchaseController")}
}

// Create runs the create action.
func (c *PurchaseController) Create(ctx *app.CreatePurchaseContext) error {

	newId := bson.NewObjectId()
	ctx.Payload.ID = &newId
	Database.C("Purchase").With(Database.Session.Copy()).Insert(ctx.Payload)

	ctx.ResponseData.Header().Set("Location", app.PurchaseHref(newId))

	return ctx.Created()
}

// Show runs the show action.
func (c *PurchaseController) Show(ctx *app.ShowPurchaseContext) error {
	// PurchaseController_Show: start_implement

	// Put your logic here

	// PurchaseController_Show: end_implement
	//Database.C("").Count()
	// Mock
	res := &app.Purchase{
		TransactionID: "9BN1kXMNdEb8dql",
		Locator:       "POS-TEST-001",
		PurchaseValue: 159.99,
	}

	return ctx.OK(res)
}
