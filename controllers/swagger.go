package controllers

import (
	"github.com/goadesign/goa"
)

// SwaggerController implements the Swagger resource.
type SwaggerController struct {
	*goa.Controller
}

// NewSwaggerController creates a Swagger controller.
func NewSwaggerController(service *goa.Service) *SwaggerController {
	return &SwaggerController{Controller: service.NewController("SwaggerController")}
}
