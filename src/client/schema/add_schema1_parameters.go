// Code generated by go-swagger; DO NOT EDIT.

package schema

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

// NewAddSchema1Params creates a new AddSchema1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddSchema1Params() *AddSchema1Params {
	return &AddSchema1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddSchema1ParamsWithTimeout creates a new AddSchema1Params object
// with the ability to set a timeout on a request.
func NewAddSchema1ParamsWithTimeout(timeout time.Duration) *AddSchema1Params {
	return &AddSchema1Params{
		timeout: timeout,
	}
}

// NewAddSchema1ParamsWithContext creates a new AddSchema1Params object
// with the ability to set a context for a request.
func NewAddSchema1ParamsWithContext(ctx context.Context) *AddSchema1Params {
	return &AddSchema1Params{
		Context: ctx,
	}
}

// NewAddSchema1ParamsWithHTTPClient creates a new AddSchema1Params object
// with the ability to set a custom HTTPClient for a request.
func NewAddSchema1ParamsWithHTTPClient(client *http.Client) *AddSchema1Params {
	return &AddSchema1Params{
		HTTPClient: client,
	}
}

/*
AddSchema1Params contains all the parameters to send to the API endpoint

	for the add schema 1 operation.

	Typically these are written to a http.Request.
*/
type AddSchema1Params struct {

	// Body.
	Body string

	/* Force.

	   Whether to force overriding the schema if the schema exists
	*/
	Force *bool

	/* Override.

	   Whether to override the schema if the schema exists

	   Default: true
	*/
	Override *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add schema 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddSchema1Params) WithDefaults() *AddSchema1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add schema 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddSchema1Params) SetDefaults() {
	var (
		forceDefault = bool(false)

		overrideDefault = bool(true)
	)

	val := AddSchema1Params{
		Force:    &forceDefault,
		Override: &overrideDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the add schema 1 params
func (o *AddSchema1Params) WithTimeout(timeout time.Duration) *AddSchema1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add schema 1 params
func (o *AddSchema1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add schema 1 params
func (o *AddSchema1Params) WithContext(ctx context.Context) *AddSchema1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add schema 1 params
func (o *AddSchema1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add schema 1 params
func (o *AddSchema1Params) WithHTTPClient(client *http.Client) *AddSchema1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add schema 1 params
func (o *AddSchema1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add schema 1 params
func (o *AddSchema1Params) WithBody(body string) *AddSchema1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add schema 1 params
func (o *AddSchema1Params) SetBody(body string) {
	o.Body = body
}

// WithForce adds the force to the add schema 1 params
func (o *AddSchema1Params) WithForce(force *bool) *AddSchema1Params {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the add schema 1 params
func (o *AddSchema1Params) SetForce(force *bool) {
	o.Force = force
}

// WithOverride adds the override to the add schema 1 params
func (o *AddSchema1Params) WithOverride(override *bool) *AddSchema1Params {
	o.SetOverride(override)
	return o
}

// SetOverride adds the override to the add schema 1 params
func (o *AddSchema1Params) SetOverride(override *bool) {
	o.Override = override
}

// WriteToRequest writes these params to a swagger request
func (o *AddSchema1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if o.Force != nil {

		// query param force
		var qrForce bool

		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
	}

	if o.Override != nil {

		// query param override
		var qrOverride bool

		if o.Override != nil {
			qrOverride = *o.Override
		}
		qOverride := swag.FormatBool(qrOverride)
		if qOverride != "" {

			if err := r.SetQueryParam("override", qOverride); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}