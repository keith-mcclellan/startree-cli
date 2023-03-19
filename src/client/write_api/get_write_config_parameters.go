// Code generated by go-swagger; DO NOT EDIT.

package write_api

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

// NewGetWriteConfigParams creates a new GetWriteConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWriteConfigParams() *GetWriteConfigParams {
	return &GetWriteConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWriteConfigParamsWithTimeout creates a new GetWriteConfigParams object
// with the ability to set a timeout on a request.
func NewGetWriteConfigParamsWithTimeout(timeout time.Duration) *GetWriteConfigParams {
	return &GetWriteConfigParams{
		timeout: timeout,
	}
}

// NewGetWriteConfigParamsWithContext creates a new GetWriteConfigParams object
// with the ability to set a context for a request.
func NewGetWriteConfigParamsWithContext(ctx context.Context) *GetWriteConfigParams {
	return &GetWriteConfigParams{
		Context: ctx,
	}
}

// NewGetWriteConfigParamsWithHTTPClient creates a new GetWriteConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWriteConfigParamsWithHTTPClient(client *http.Client) *GetWriteConfigParams {
	return &GetWriteConfigParams{
		HTTPClient: client,
	}
}

/*
GetWriteConfigParams contains all the parameters to send to the API endpoint

	for the get write config operation.

	Typically these are written to a http.Request.
*/
type GetWriteConfigParams struct {

	// Table.
	Table string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get write config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWriteConfigParams) WithDefaults() *GetWriteConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get write config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWriteConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get write config params
func (o *GetWriteConfigParams) WithTimeout(timeout time.Duration) *GetWriteConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get write config params
func (o *GetWriteConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get write config params
func (o *GetWriteConfigParams) WithContext(ctx context.Context) *GetWriteConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get write config params
func (o *GetWriteConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get write config params
func (o *GetWriteConfigParams) WithHTTPClient(client *http.Client) *GetWriteConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get write config params
func (o *GetWriteConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTable adds the table to the get write config params
func (o *GetWriteConfigParams) WithTable(table string) *GetWriteConfigParams {
	o.SetTable(table)
	return o
}

// SetTable adds the table to the get write config params
func (o *GetWriteConfigParams) SetTable(table string) {
	o.Table = table
}

// WriteToRequest writes these params to a swagger request
func (o *GetWriteConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param table
	if err := r.SetPathParam("table", o.Table); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
