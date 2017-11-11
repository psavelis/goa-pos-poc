package design

import . "github.com/goadesign/goa/design"
import . "github.com/goadesign/goa/design/apidsl"

var _ = API("pos", func() {
	Title("Point Of Sale API")
	Version("v1")
	Description("point of sale microservice")
	Host("localhost:5001")
	Scheme("http")
	BasePath("/pos/v1")

	Consumes("application/json")
	Produces("application/json")

	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})

	ResponseTemplate(Conflict, func() {
		Description("A conflict arose from your request. (e.g. resource already exists)")
		Headers(func() {
			Header("X-Request-Id", func() {
				Pattern("[a-f0-9]+")
			})
		})

		Status(409)
	})
})

var PurchasePayload = Type("PurchasePayload", func() {
	Description("Detailed information regarding a POS purchase operation")

	Attribute("ID", func() {
		Metadata("struct:tag:json", "id")
		Metadata("struct:tag:bson", "_id,omitempty")
		Metadata("struct:field:type", "bson.ObjectId", "gopkg.in/mgo.v2/bson")
		Metadata("swagger:generate", "false")
	})

	Attribute("Locator", String, "Operation reference code", func() {
		Metadata("struct:tag:json", "locator")
		Metadata("struct:tag:bson", "locator,omitempty")
		Metadata("swagger:tag:json", "locator")

		MinLength(1)
		MaxLength(30)
	})

	Attribute("PurchaseValue", Number, "Total amount paid", func() {
		Metadata("struct:tag:json", "purchase_value")
		Metadata("struct:tag:bson", "purchase_value,omitempty")
		Metadata("swagger:tag:json", "purchase_value")
		Minimum(0.01)
	})

	Required("Locator", "PurchaseValue")
})

var PurchaseMedia = MediaType("application/vnd.pos.purchase+json", func() {
	TypeName("Purchase")
	Reference(PurchasePayload)

	Attributes(func() {

		// Inherited attributes from PurchasePayload
		Attribute("TransactionID", String, "Unique transaction identifier", func() {
			Metadata("struct:tag:json", "transaction_id")
			Metadata("swagger:tag:json", "transaction_id")
			Metadata("struct:tag:bson", "_id,omitempty")
			Pattern("^[0-9a-fA-F]{24}$")
		})
		Attribute("Locator", String, "Operation reference code", func() {
			Metadata("struct:tag:json", "locator")
			Metadata("swagger:tag:json", "locator")
		})

		Attribute("PurchaseValue", Number, "Total amount paid", func() {
			Metadata("swagger:tag:json", "purchase_value")
		})

		Attribute("Href", String, "API href of Purchase", func() {
			Example("/pos/v1/purchases/5a06839d42e6552b004a7e03")
			Metadata("struct:tag:json", "href")
			Metadata("swagger:tag:json", "href")
		})

		Required("TransactionID", "Locator", "PurchaseValue", "Href")
	})

	View("default", func() {
		Attribute("TransactionID")
		Attribute("Locator")
		Attribute("PurchaseValue")
		Attribute("Href")
	})
})

var _ = Resource("Purchase", func() {
	DefaultMedia(PurchaseMedia)
	Description("A pos purchase data")
	BasePath("/purchases")

	Action("create", func() {
		Description("creates a purchase")
		Routing(POST("/"))
		Payload(PurchasePayload)
		Response(Created, "^/purchases/[A-Za-z0-9_.]+$")
		Response(BadRequest, ErrorMedia)
		Response(Conflict)
	})

	Action("show", func() {
		Description("retrieves a purchase")
		Routing(GET("/:transaction_id"))
		Params(func() {
			Param("transaction_id", String, "Unique transaction identifier", func() {
				Pattern("^[0-9a-fA-F]{24}$")
			})
		})

		Response(OK, PurchaseMedia)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})
