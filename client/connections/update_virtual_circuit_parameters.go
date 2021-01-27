// Code generated by go-swagger; DO NOT EDIT.

package connections

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

// NewUpdateVirtualCircuitParams creates a new UpdateVirtualCircuitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVirtualCircuitParams() *UpdateVirtualCircuitParams {
	return &UpdateVirtualCircuitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVirtualCircuitParamsWithTimeout creates a new UpdateVirtualCircuitParams object
// with the ability to set a timeout on a request.
func NewUpdateVirtualCircuitParamsWithTimeout(timeout time.Duration) *UpdateVirtualCircuitParams {
	return &UpdateVirtualCircuitParams{
		timeout: timeout,
	}
}

// NewUpdateVirtualCircuitParamsWithContext creates a new UpdateVirtualCircuitParams object
// with the ability to set a context for a request.
func NewUpdateVirtualCircuitParamsWithContext(ctx context.Context) *UpdateVirtualCircuitParams {
	return &UpdateVirtualCircuitParams{
		Context: ctx,
	}
}

// NewUpdateVirtualCircuitParamsWithHTTPClient creates a new UpdateVirtualCircuitParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVirtualCircuitParamsWithHTTPClient(client *http.Client) *UpdateVirtualCircuitParams {
	return &UpdateVirtualCircuitParams{
		HTTPClient: client,
	}
}

/* UpdateVirtualCircuitParams contains all the parameters to send to the API endpoint
   for the update virtual circuit operation.

   Typically these are written to a http.Request.
*/
type UpdateVirtualCircuitParams struct {

	/* ID.

	   Virtual Circuit UUID

	   Format: uuid
	*/
	ID strfmt.UUID

	/* VirtualCircuit.

	   Updated Virtual Circuit details
	*/
	VirtualCircuit *types.VirtualCircuitUpdateInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update virtual circuit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVirtualCircuitParams) WithDefaults() *UpdateVirtualCircuitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update virtual circuit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVirtualCircuitParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update virtual circuit params
func (o *UpdateVirtualCircuitParams) WithTimeout(timeout time.Duration) *UpdateVirtualCircuitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update virtual circuit params
func (o *UpdateVirtualCircuitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update virtual circuit params
func (o *UpdateVirtualCircuitParams) WithContext(ctx context.Context) *UpdateVirtualCircuitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update virtual circuit params
func (o *UpdateVirtualCircuitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update virtual circuit params
func (o *UpdateVirtualCircuitParams) WithHTTPClient(client *http.Client) *UpdateVirtualCircuitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update virtual circuit params
func (o *UpdateVirtualCircuitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update virtual circuit params
func (o *UpdateVirtualCircuitParams) WithID(id strfmt.UUID) *UpdateVirtualCircuitParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update virtual circuit params
func (o *UpdateVirtualCircuitParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithVirtualCircuit adds the virtualCircuit to the update virtual circuit params
func (o *UpdateVirtualCircuitParams) WithVirtualCircuit(virtualCircuit *types.VirtualCircuitUpdateInput) *UpdateVirtualCircuitParams {
	o.SetVirtualCircuit(virtualCircuit)
	return o
}

// SetVirtualCircuit adds the virtualCircuit to the update virtual circuit params
func (o *UpdateVirtualCircuitParams) SetVirtualCircuit(virtualCircuit *types.VirtualCircuitUpdateInput) {
	o.VirtualCircuit = virtualCircuit
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVirtualCircuitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}
	if o.VirtualCircuit != nil {
		if err := r.SetBodyParam(o.VirtualCircuit); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
