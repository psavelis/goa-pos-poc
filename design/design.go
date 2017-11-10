package design

import . "github.com/goadesign/goa/design"
import . "github.com/goadesign/goa/design/apidsl"

var _ = API("pos", func() {
	Title("Point Of Sale API")
	Version("v1")
	Description("point of sale microservice")
	Host("localhost:5001")

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

		Status(422)
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

		MinLength(1)
		MaxLength(30)
	})

	Attribute("PurchaseValue", Number, "Total amount paid", func() {
		Metadata("struct:tag:json", "purchase_value")
		Metadata("struct:tag:bson", "purchase_value,omitempty")

		Minimum(0.01)
	})

	Required("Locator", "PurchaseValue")
})

var PurchaseMedia = MediaType("application/vnd.purchase+json", func() {
	TypeName("Purchase")
	Reference(PurchasePayload)

	Attributes(func() {

		Attribute("Href", String, "API href of Purchase", func() {
			Example("/purchases/1")
			Metadata("struct:tag:json", "href")
		})

		// Inherited attributes from PurchasePayload
		Attribute("TransactionId", String, "Unique transaction identifier", func() {
			Metadata("struct:tag:json", "transaction_id")
		})
		Attribute("Locator", String, "Operation reference code", func() {
			Metadata("struct:tag:json", "locator")
		})

		Attribute("PurchaseValue", Number, "Total amount paid", func() {
			Metadata("struct:tag:json", "purchase_value")
		})

		Required("TransactionId", "Locator", "PurchaseValue", "Href")
	})

	View("default", func() {
		Attribute("TransactionId")
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
		Description("retrieve an specific purchase")
		Routing(GET("/:TransactionId"))
		Params(func() {
			Param("TransactionId", String)
		})

		Response(OK, PurchaseMedia)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("Swagger", func() {
	Description("The API Swagger specification")

	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})

	Files("/swagger.json", "swagger/swagger.json")
	Files("swagger/*filepath", "swagger")
})
