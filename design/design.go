package design

import . "github.com/goadesign/goa/design"
import . "github.com/goadesign/goa/design/apidsl"

var _ = API("pos", func() {

	Title("Point Of Sale (POS)")
	Version("v1")
	License(func() {
		Name("GPL-3.0")
		URL("https://github.com/psavelis/goa-pos-poc/blob/master/LICENSE")
	})
	Description("go microservice")
	Host("psavelis.herokuapp.com")
	Scheme("https")
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

	Attribute("id", func() {
		Metadata("struct:tag:bson", "_id,omitempty")
		Metadata("struct:field:name", "ID")
		Metadata("struct:field:type", "bson.ObjectId", "gopkg.in/mgo.v2/bson")
		Example("")
	})

	Attribute("locator", String, "Operation reference code", func() {
		Metadata("struct:field:name", "Locator")
		Metadata("struct:tag:bson", "locator,omitempty")

		Example("MPOS00123820-UAT-A02")

		MinLength(1)
		MaxLength(30)
	})

	Attribute("purchase_value", Number, "Total amount paid", func() {
		Metadata("struct:field:name", "PurchaseValue")
		Metadata("struct:tag:bson", "purchase_value,omitempty")

		Example(119.99)

		Minimum(0.01)
	})

	Required("locator", "purchase_value")
})

var PurchaseMedia = MediaType("application/json", func() {
	TypeName("Purchase")
	Reference(PurchasePayload)

	Attributes(func() {

		// Inherited attributes from PurchasePayload
		Attribute("transaction_id", String, "Unique transaction identifier", func() {
			Metadata("struct:tag:json", "transaction_id")
			Metadata("struct:field:name", "TransactionID")
			Metadata("struct:tag:bson", "_id,omitempty")
			Pattern("^[0-9a-fA-F]{24}$")
		})
		Attribute("locator", String, "Operation reference code", func() {
			Metadata("struct:tag:json", "locator")
			Metadata("struct:field:name", "Locator")
		})

		Attribute("purchase_value", Number, "Total amount paid", func() {
			Metadata("struct:tag:json", "purchase_value")
			Metadata("struct:field:name", "PurchaseValue")
		})

		Attribute("href", String, "API href of Purchase", func() {
			Example("/pos/v1/purchases/5a06839d42e6552b004a7e03")
			Metadata("struct:tag:json", "href")
			Metadata("struct:field:name", "Href")
		})

		Required("transaction_id", "locator", "purchase_value", "href")
	})

	View("default", func() {
		Attribute("transaction_id")
		Attribute("locator")
		Attribute("purchase_value")
		Attribute("href")
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
		Response(BadRequest)
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
		Response(BadRequest)
		Response(NotFound)
	})

	// Action("list", func() {
	// 	Description("retrieves a purchase")
	// 	Routing(GET("/:transaction_id"))
	// 	Params(func() {
	// 		Param("_limit", String, "Unique transaction identifier", func() {
	// 			Pattern("^[0-9a-fA-F]{24}$")
	// 		})
	// 	})

	// 	Response(OK, PurchaseMedia)
	// 	Response(BadRequest)
	// 	Response(NotFound)
	// })
})

var _ = Resource("public", func() {
	Metadata("swagger:generate", "false")
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/ui", "public/html/index.html")
})

var _ = Resource("js", func() {
	Metadata("swagger:generate", "false")
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/js/*filepath", "public/js")
})

var _ = Resource("swagger", func() {
	Metadata("swagger:generate", "false")

	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/swagger.json", "public/swagger/swagger.json")
})
