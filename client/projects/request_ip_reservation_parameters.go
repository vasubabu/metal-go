// Code generated by go-swagger; DO NOT EDIT.

package projects

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

	"github.com/t0mk/gometal/types"
)

// NewRequestIPReservationParams creates a new RequestIPReservationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRequestIPReservationParams() *RequestIPReservationParams {
	return &RequestIPReservationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRequestIPReservationParamsWithTimeout creates a new RequestIPReservationParams object
// with the ability to set a timeout on a request.
func NewRequestIPReservationParamsWithTimeout(timeout time.Duration) *RequestIPReservationParams {
	return &RequestIPReservationParams{
		timeout: timeout,
	}
}

// NewRequestIPReservationParamsWithContext creates a new RequestIPReservationParams object
// with the ability to set a context for a request.
func NewRequestIPReservationParamsWithContext(ctx context.Context) *RequestIPReservationParams {
	return &RequestIPReservationParams{
		Context: ctx,
	}
}

// NewRequestIPReservationParamsWithHTTPClient creates a new RequestIPReservationParams object
// with the ability to set a custom HTTPClient for a request.
func NewRequestIPReservationParamsWithHTTPClient(client *http.Client) *RequestIPReservationParams {
	return &RequestIPReservationParams{
		HTTPClient: client,
	}
}

/* RequestIPReservationParams contains all the parameters to send to the API endpoint
   for the request IP reservation operation.

   Typically these are written to a http.Request.
*/
type RequestIPReservationParams struct {

	/* ID.

	   Project UUID

	   Format: uuid
	*/
	ID strfmt.UUID

	/* IPReservationRequest.

	   IP Reservation Request to create
	*/
	IPReservationRequest *types.IPReservationRequestInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the request IP reservation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RequestIPReservationParams) WithDefaults() *RequestIPReservationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the request IP reservation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RequestIPReservationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the request IP reservation params
func (o *RequestIPReservationParams) WithTimeout(timeout time.Duration) *RequestIPReservationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the request IP reservation params
func (o *RequestIPReservationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the request IP reservation params
func (o *RequestIPReservationParams) WithContext(ctx context.Context) *RequestIPReservationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the request IP reservation params
func (o *RequestIPReservationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the request IP reservation params
func (o *RequestIPReservationParams) WithHTTPClient(client *http.Client) *RequestIPReservationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the request IP reservation params
func (o *RequestIPReservationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the request IP reservation params
func (o *RequestIPReservationParams) WithID(id strfmt.UUID) *RequestIPReservationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the request IP reservation params
func (o *RequestIPReservationParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithIPReservationRequest adds the iPReservationRequest to the request IP reservation params
func (o *RequestIPReservationParams) WithIPReservationRequest(iPReservationRequest *types.IPReservationRequestInput) *RequestIPReservationParams {
	o.SetIPReservationRequest(iPReservationRequest)
	return o
}

// SetIPReservationRequest adds the ipReservationRequest to the request IP reservation params
func (o *RequestIPReservationParams) SetIPReservationRequest(iPReservationRequest *types.IPReservationRequestInput) {
	o.IPReservationRequest = iPReservationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *RequestIPReservationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}
	if o.IPReservationRequest != nil {
		if err := r.SetBodyParam(o.IPReservationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
