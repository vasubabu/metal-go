// Code generated by go-swagger; DO NOT EDIT.

package ports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewBondPortParams creates a new BondPortParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBondPortParams() *BondPortParams {
	return &BondPortParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBondPortParamsWithTimeout creates a new BondPortParams object
// with the ability to set a timeout on a request.
func NewBondPortParamsWithTimeout(timeout time.Duration) *BondPortParams {
	return &BondPortParams{
		timeout: timeout,
	}
}

// NewBondPortParamsWithContext creates a new BondPortParams object
// with the ability to set a context for a request.
func NewBondPortParamsWithContext(ctx context.Context) *BondPortParams {
	return &BondPortParams{
		Context: ctx,
	}
}

// NewBondPortParamsWithHTTPClient creates a new BondPortParams object
// with the ability to set a custom HTTPClient for a request.
func NewBondPortParamsWithHTTPClient(client *http.Client) *BondPortParams {
	return &BondPortParams{
		HTTPClient: client,
	}
}

/* BondPortParams contains all the parameters to send to the API endpoint
   for the bond port operation.

   Typically these are written to a http.Request.
*/
type BondPortParams struct {

	/* BulkEnable.

	   enable both ports
	*/
	BulkEnable *bool

	/* ID.

	   Port UUID

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bond port params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BondPortParams) WithDefaults() *BondPortParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bond port params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BondPortParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bond port params
func (o *BondPortParams) WithTimeout(timeout time.Duration) *BondPortParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bond port params
func (o *BondPortParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bond port params
func (o *BondPortParams) WithContext(ctx context.Context) *BondPortParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bond port params
func (o *BondPortParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bond port params
func (o *BondPortParams) WithHTTPClient(client *http.Client) *BondPortParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bond port params
func (o *BondPortParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBulkEnable adds the bulkEnable to the bond port params
func (o *BondPortParams) WithBulkEnable(bulkEnable *bool) *BondPortParams {
	o.SetBulkEnable(bulkEnable)
	return o
}

// SetBulkEnable adds the bulkEnable to the bond port params
func (o *BondPortParams) SetBulkEnable(bulkEnable *bool) {
	o.BulkEnable = bulkEnable
}

// WithID adds the id to the bond port params
func (o *BondPortParams) WithID(id strfmt.UUID) *BondPortParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the bond port params
func (o *BondPortParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *BondPortParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BulkEnable != nil {

		// query param bulk_enable
		var qrBulkEnable bool

		if o.BulkEnable != nil {
			qrBulkEnable = *o.BulkEnable
		}
		qBulkEnable := swag.FormatBool(qrBulkEnable)
		if qBulkEnable != "" {

			if err := r.SetQueryParam("bulk_enable", qBulkEnable); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
