package controllers

import (
	"github.com/goadesign/goa"
	"github.com/psavelis/goa-pos-poc/app"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PurchaseController implements the Purchase resource.
type PurchaseController struct {
	*goa.Controller
}

var (
	Database *mgo.Database
)

// NewPurchaseController creates a Purchase controller.
func NewPurchaseController(service *goa.Service, database *mgo.Database) *PurchaseController {
	Database = database
	return &PurchaseController{Controller: service.NewController("PurchaseController")}
}

// Create runs the create action.
func (c *PurchaseController) Create(ctx *app.CreatePurchaseContext) error {

	newID := bson.NewObjectId()

	ctx.Payload.ID = &newID

	session := Database.Session.Copy()
	defer session.Close()

	err := session.DB("services-pos").C("Purchase").Insert(ctx.Payload)

	if err != nil {
		if mgo.IsDup(err) {
			return ctx.Conflict()
		}
		// TODO: log
		return ctx.Err()
	}

	ctx.ResponseData.Header().Set("Location", app.PurchaseHref(newID.Hex()))

	return ctx.Created()
}

// Show runs the show action.
func (c *PurchaseController) Show(ctx *app.ShowPurchaseContext) error {

	session := Database.Session.Copy()
	defer session.Close()

	result := app.Purchase{}

	err := session.DB("services-pos").C("Purchase").FindId(bson.ObjectIdHex(ctx.TransactionID)).One(&result)

	if err != nil {
		return ctx.NotFound()
	}

	result.TransactionID = ctx.TransactionID
	result.Href = app.PurchaseHref(ctx.TransactionID)

	return ctx.OK(&result)
}
