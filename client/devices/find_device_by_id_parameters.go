// Code generated by go-swagger; DO NOT EDIT.

package devices

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
)

// NewFindDeviceByIDParams creates a new FindDeviceByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindDeviceByIDParams() *FindDeviceByIDParams {
	return &FindDeviceByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindDeviceByIDParamsWithTimeout creates a new FindDeviceByIDParams object
// with the ability to set a timeout on a request.
func NewFindDeviceByIDParamsWithTimeout(timeout time.Duration) *FindDeviceByIDParams {
	return &FindDeviceByIDParams{
		timeout: timeout,
	}
}

// NewFindDeviceByIDParamsWithContext creates a new FindDeviceByIDParams object
// with the ability to set a context for a request.
func NewFindDeviceByIDParamsWithContext(ctx context.Context) *FindDeviceByIDParams {
	return &FindDeviceByIDParams{
		Context: ctx,
	}
}

// NewFindDeviceByIDParamsWithHTTPClient creates a new FindDeviceByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindDeviceByIDParamsWithHTTPClient(client *http.Client) *FindDeviceByIDParams {
	return &FindDeviceByIDParams{
		HTTPClient: client,
	}
}

/* FindDeviceByIDParams contains all the parameters to send to the API endpoint
   for the find device by Id operation.

   Typically these are written to a http.Request.
*/
type FindDeviceByIDParams struct {

	/* ID.

	   Device UUID

	   Format: uuid
	*/
	ID strfmt.UUID

	/* Include.

	   related attributes to include
	*/
	Include *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find device by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindDeviceByIDParams) WithDefaults() *FindDeviceByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find device by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindDeviceByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find device by Id params
func (o *FindDeviceByIDParams) WithTimeout(timeout time.Duration) *FindDeviceByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find device by Id params
func (o *FindDeviceByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find device by Id params
func (o *FindDeviceByIDParams) WithContext(ctx context.Context) *FindDeviceByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find device by Id params
func (o *FindDeviceByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find device by Id params
func (o *FindDeviceByIDParams) WithHTTPClient(client *http.Client) *FindDeviceByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find device by Id params
func (o *FindDeviceByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find device by Id params
func (o *FindDeviceByIDParams) WithID(id strfmt.UUID) *FindDeviceByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find device by Id params
func (o *FindDeviceByIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithInclude adds the include to the find device by Id params
func (o *FindDeviceByIDParams) WithInclude(include *string) *FindDeviceByIDParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the find device by Id params
func (o *FindDeviceByIDParams) SetInclude(include *string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *FindDeviceByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if o.Include != nil {

		// query param include
		var qrInclude string

		if o.Include != nil {
			qrInclude = *o.Include
		}
		qInclude := qrInclude
		if qInclude != "" {

			if err := r.SetQueryParam("include", qInclude); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
