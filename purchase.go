package main

import (
	"github.com/goadesign/goa"
	"github.com/psavelis/goa-pos-poc/app"
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
	// PurchaseController_Create: start_implement

	// Put your logic here

	// PurchaseController_Create: end_implement
	return nil
}

// Find runs the find action.
func (c *PurchaseController) Find(ctx *app.FindPurchaseContext) error {
	// PurchaseController_Find: start_implement

	// Put your logic here

	// PurchaseController_Find: end_implement
	res := &app.Purchase{}
	return ctx.OK(res)
}
