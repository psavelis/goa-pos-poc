// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "pos": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/psavelis/goa-pos-poc/design
// --out=c:\code\src\github.com\psavelis\goa-pos-poc
// --version=v1.3.0

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// CreatePurchaseContext provides the Purchase create action context.
type CreatePurchaseContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *PurchasePayload
}

// NewCreatePurchaseContext parses the incoming request URL and body, performs validations and creates the
// context used by the Purchase controller create action.
func NewCreatePurchaseContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreatePurchaseContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreatePurchaseContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreatePurchaseContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// FindPurchaseContext provides the Purchase find action context.
type FindPurchaseContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	TransactionID string
}

// NewFindPurchaseContext parses the incoming request URL and body, performs validations and creates the
// context used by the Purchase controller find action.
func NewFindPurchaseContext(ctx context.Context, r *http.Request, service *goa.Service) (*FindPurchaseContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := FindPurchaseContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramTransactionID := req.Params["TransactionId"]
	if len(paramTransactionID) > 0 {
		rawTransactionID := paramTransactionID[0]
		rctx.TransactionID = rawTransactionID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *FindPurchaseContext) OK(r *Purchase) error {
	ctx.ResponseData.Header().Set("Content-Type", "vnd.application/pos.purchases")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}
