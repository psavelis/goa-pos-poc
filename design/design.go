package design

import . "github.com/goadesign/goa/design"
import . "github.com/goadesign/goa/design/apidsl"

var _ = API("pos", func() {
	Description("point of sale microservice")
	Host("localhost:5001")
})

var PurchasePayload = Type("PurchasePayload", func() {
	Description("Detailed information regarding a POS purchase operation")

	Attribute("Locator", String, "Operation reference code", func() {
		MinLength(1)
		MaxLength(30)
	})

	Attribute("PurchaseValue", Number, "Total amount paid", func() {
		Minimum(0.01)
	})

	Required("Locator", "PurchaseValue")
})

var PurchaseMedia = MediaType("vnd.application/pos.purchases", func() {
	TypeName("Purchase")
	Reference(PurchasePayload)

	Attributes(func() {
		Attribute("TransactionId", String, "Unique transaction identifier")
		Attribute("Locator")
		Attribute("PurchaseValue")
		Required("TransactionId", "Locator", "PurchaseValue")
	})

	View("default", func() {
		Attribute("TransactionId")
		Attribute("Locator")
		Attribute("PurchaseValue")
	})
})

var _ = Resource("Purchase", func() {
	Description("A pos purchase data")
	BasePath("/purchases")

	Action("create", func() {
		Description("creates a purchase")
		Routing(POST("/"))
		Payload(PurchasePayload)
		Response(Created)
	})

	Action("find", func() {
		Description("retrieve an specific purchase")
		Routing(GET("/:TransactionId"))
		Params(func() {
			Param("TransactionId", String)
		})
		Response(OK, PurchaseMedia)
	})
})
