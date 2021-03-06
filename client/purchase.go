// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "pos": Purchase Resource Client
//
// Command:
// $ goagen
// --design=github.com/psavelis/goa-pos-poc/design
// --out=$(GOPATH)src\github.com\psavelis\goa-pos-poc
// --version=v1.3.0

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CreatePurchasePath computes a request path to the create action of Purchase.
func CreatePurchasePath() string {

	return fmt.Sprintf("/pos/v1/purchases/")
}

// creates a purchase
func (c *Client) CreatePurchase(ctx context.Context, path string, payload *PurchasePayload) (*http.Response, error) {
	req, err := c.NewCreatePurchaseRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreatePurchaseRequest create the request corresponding to the create action endpoint of the Purchase resource.
func (c *Client) NewCreatePurchaseRequest(ctx context.Context, path string, payload *PurchasePayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return req, nil
}

// ShowPurchasePath computes a request path to the show action of Purchase.
func ShowPurchasePath(transactionID string) string {
	param0 := transactionID

	return fmt.Sprintf("/pos/v1/purchases/%s", param0)
}

// retrieves a purchase
func (c *Client) ShowPurchase(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowPurchaseRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowPurchaseRequest create the request corresponding to the show action endpoint of the Purchase resource.
func (c *Client) NewShowPurchaseRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
