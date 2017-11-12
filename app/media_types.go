// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "pos": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/psavelis/goa-pos-poc/design
// --out=$(GOPATH)src\github.com\psavelis\goa-pos-poc
// --version=v1.3.0

package app

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// Purchase media type (default view)
//
// Identifier: application/json; view=default
type Purchase struct {
	// API href of Purchase
	Href string `json:"href"`
	// Operation reference code
	Locator string `bson:"locator,omitempty" json:"locator"`
	// Total amount paid
	PurchaseValue float64 `bson:"purchase_value,omitempty" json:"purchase_value"`
	// Unique transaction identifier
	TransactionID string `bson:"_id,omitempty" json:"transaction_id"`
}

// Validate validates the Purchase media type instance.
func (mt *Purchase) Validate() (err error) {
	if mt.TransactionID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "transaction_id"))
	}
	if mt.Locator == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "locator"))
	}

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if utf8.RuneCountInString(mt.Locator) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.locator`, mt.Locator, utf8.RuneCountInString(mt.Locator), 1, true))
	}
	if utf8.RuneCountInString(mt.Locator) > 30 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.locator`, mt.Locator, utf8.RuneCountInString(mt.Locator), 30, false))
	}
	if mt.PurchaseValue < 0.010000 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.purchase_value`, mt.PurchaseValue, 0.010000, true))
	}
	if ok := goa.ValidatePattern(`^[0-9a-fA-F]{24}$`, mt.TransactionID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.transaction_id`, mt.TransactionID, `^[0-9a-fA-F]{24}$`))
	}
	return
}
