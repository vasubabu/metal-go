// Code generated by go-swagger; DO NOT EDIT.

package incidents

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

// NewGetIncidentsParams creates a new GetIncidentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIncidentsParams() *GetIncidentsParams {
	return &GetIncidentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIncidentsParamsWithTimeout creates a new GetIncidentsParams object
// with the ability to set a timeout on a request.
func NewGetIncidentsParamsWithTimeout(timeout time.Duration) *GetIncidentsParams {
	return &GetIncidentsParams{
		timeout: timeout,
	}
}

// NewGetIncidentsParamsWithContext creates a new GetIncidentsParams object
// with the ability to set a context for a request.
func NewGetIncidentsParamsWithContext(ctx context.Context) *GetIncidentsParams {
	return &GetIncidentsParams{
		Context: ctx,
	}
}

// NewGetIncidentsParamsWithHTTPClient creates a new GetIncidentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIncidentsParamsWithHTTPClient(client *http.Client) *GetIncidentsParams {
	return &GetIncidentsParams{
		HTTPClient: client,
	}
}

/* GetIncidentsParams contains all the parameters to send to the API endpoint
   for the get incidents operation.

   Typically these are written to a http.Request.
*/
type GetIncidentsParams struct {

	/* Include.

	   related attributes to include
	*/
	Include *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get incidents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIncidentsParams) WithDefaults() *GetIncidentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get incidents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIncidentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get incidents params
func (o *GetIncidentsParams) WithTimeout(timeout time.Duration) *GetIncidentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get incidents params
func (o *GetIncidentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get incidents params
func (o *GetIncidentsParams) WithContext(ctx context.Context) *GetIncidentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get incidents params
func (o *GetIncidentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get incidents params
func (o *GetIncidentsParams) WithHTTPClient(client *http.Client) *GetIncidentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get incidents params
func (o *GetIncidentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInclude adds the include to the get incidents params
func (o *GetIncidentsParams) WithInclude(include *string) *GetIncidentsParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the get incidents params
func (o *GetIncidentsParams) SetInclude(include *string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *GetIncidentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
