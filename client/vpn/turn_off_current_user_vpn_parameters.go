// Code generated by go-swagger; DO NOT EDIT.

package vpn

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

// NewTurnOffCurrentUserVPNParams creates a new TurnOffCurrentUserVPNParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTurnOffCurrentUserVPNParams() *TurnOffCurrentUserVPNParams {
	return &TurnOffCurrentUserVPNParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTurnOffCurrentUserVPNParamsWithTimeout creates a new TurnOffCurrentUserVPNParams object
// with the ability to set a timeout on a request.
func NewTurnOffCurrentUserVPNParamsWithTimeout(timeout time.Duration) *TurnOffCurrentUserVPNParams {
	return &TurnOffCurrentUserVPNParams{
		timeout: timeout,
	}
}

// NewTurnOffCurrentUserVPNParamsWithContext creates a new TurnOffCurrentUserVPNParams object
// with the ability to set a context for a request.
func NewTurnOffCurrentUserVPNParamsWithContext(ctx context.Context) *TurnOffCurrentUserVPNParams {
	return &TurnOffCurrentUserVPNParams{
		Context: ctx,
	}
}

// NewTurnOffCurrentUserVPNParamsWithHTTPClient creates a new TurnOffCurrentUserVPNParams object
// with the ability to set a custom HTTPClient for a request.
func NewTurnOffCurrentUserVPNParamsWithHTTPClient(client *http.Client) *TurnOffCurrentUserVPNParams {
	return &TurnOffCurrentUserVPNParams{
		HTTPClient: client,
	}
}

/* TurnOffCurrentUserVPNParams contains all the parameters to send to the API endpoint
   for the turn off current user Vpn operation.

   Typically these are written to a http.Request.
*/
type TurnOffCurrentUserVPNParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the turn off current user Vpn params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TurnOffCurrentUserVPNParams) WithDefaults() *TurnOffCurrentUserVPNParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the turn off current user Vpn params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TurnOffCurrentUserVPNParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the turn off current user Vpn params
func (o *TurnOffCurrentUserVPNParams) WithTimeout(timeout time.Duration) *TurnOffCurrentUserVPNParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the turn off current user Vpn params
func (o *TurnOffCurrentUserVPNParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the turn off current user Vpn params
func (o *TurnOffCurrentUserVPNParams) WithContext(ctx context.Context) *TurnOffCurrentUserVPNParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the turn off current user Vpn params
func (o *TurnOffCurrentUserVPNParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the turn off current user Vpn params
func (o *TurnOffCurrentUserVPNParams) WithHTTPClient(client *http.Client) *TurnOffCurrentUserVPNParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the turn off current user Vpn params
func (o *TurnOffCurrentUserVPNParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *TurnOffCurrentUserVPNParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
