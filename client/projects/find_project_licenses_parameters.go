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
	"github.com/go-openapi/swag"
)

// NewFindProjectLicensesParams creates a new FindProjectLicensesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindProjectLicensesParams() *FindProjectLicensesParams {
	return &FindProjectLicensesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindProjectLicensesParamsWithTimeout creates a new FindProjectLicensesParams object
// with the ability to set a timeout on a request.
func NewFindProjectLicensesParamsWithTimeout(timeout time.Duration) *FindProjectLicensesParams {
	return &FindProjectLicensesParams{
		timeout: timeout,
	}
}

// NewFindProjectLicensesParamsWithContext creates a new FindProjectLicensesParams object
// with the ability to set a context for a request.
func NewFindProjectLicensesParamsWithContext(ctx context.Context) *FindProjectLicensesParams {
	return &FindProjectLicensesParams{
		Context: ctx,
	}
}

// NewFindProjectLicensesParamsWithHTTPClient creates a new FindProjectLicensesParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindProjectLicensesParamsWithHTTPClient(client *http.Client) *FindProjectLicensesParams {
	return &FindProjectLicensesParams{
		HTTPClient: client,
	}
}

/* FindProjectLicensesParams contains all the parameters to send to the API endpoint
   for the find project licenses operation.

   Typically these are written to a http.Request.
*/
type FindProjectLicensesParams struct {

	/* ID.

	   Project UUID

	   Format: uuid
	*/
	ID strfmt.UUID

	/* Include.

	   related attributes to include
	*/
	Include *string

	/* Page.

	   page to display, default to 1, max 100_000
	*/
	Page *int64

	/* PerPage.

	   items per page, default to 10, max 1_000
	*/
	PerPage *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find project licenses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindProjectLicensesParams) WithDefaults() *FindProjectLicensesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find project licenses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindProjectLicensesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find project licenses params
func (o *FindProjectLicensesParams) WithTimeout(timeout time.Duration) *FindProjectLicensesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find project licenses params
func (o *FindProjectLicensesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find project licenses params
func (o *FindProjectLicensesParams) WithContext(ctx context.Context) *FindProjectLicensesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find project licenses params
func (o *FindProjectLicensesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find project licenses params
func (o *FindProjectLicensesParams) WithHTTPClient(client *http.Client) *FindProjectLicensesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find project licenses params
func (o *FindProjectLicensesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find project licenses params
func (o *FindProjectLicensesParams) WithID(id strfmt.UUID) *FindProjectLicensesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find project licenses params
func (o *FindProjectLicensesParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithInclude adds the include to the find project licenses params
func (o *FindProjectLicensesParams) WithInclude(include *string) *FindProjectLicensesParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the find project licenses params
func (o *FindProjectLicensesParams) SetInclude(include *string) {
	o.Include = include
}

// WithPage adds the page to the find project licenses params
func (o *FindProjectLicensesParams) WithPage(page *int64) *FindProjectLicensesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the find project licenses params
func (o *FindProjectLicensesParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the find project licenses params
func (o *FindProjectLicensesParams) WithPerPage(perPage *int64) *FindProjectLicensesParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the find project licenses params
func (o *FindProjectLicensesParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *FindProjectLicensesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int64

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
