// Code generated by go-swagger; DO NOT EDIT.

package version

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

// NewGetVersionNumberParams creates a new GetVersionNumberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVersionNumberParams() *GetVersionNumberParams {
	return &GetVersionNumberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVersionNumberParamsWithTimeout creates a new GetVersionNumberParams object
// with the ability to set a timeout on a request.
func NewGetVersionNumberParamsWithTimeout(timeout time.Duration) *GetVersionNumberParams {
	return &GetVersionNumberParams{
		timeout: timeout,
	}
}

// NewGetVersionNumberParamsWithContext creates a new GetVersionNumberParams object
// with the ability to set a context for a request.
func NewGetVersionNumberParamsWithContext(ctx context.Context) *GetVersionNumberParams {
	return &GetVersionNumberParams{
		Context: ctx,
	}
}

// NewGetVersionNumberParamsWithHTTPClient creates a new GetVersionNumberParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVersionNumberParamsWithHTTPClient(client *http.Client) *GetVersionNumberParams {
	return &GetVersionNumberParams{
		HTTPClient: client,
	}
}

/*
GetVersionNumberParams contains all the parameters to send to the API endpoint

	for the get version number operation.

	Typically these are written to a http.Request.
*/
type GetVersionNumberParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get version number params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVersionNumberParams) WithDefaults() *GetVersionNumberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get version number params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVersionNumberParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get version number params
func (o *GetVersionNumberParams) WithTimeout(timeout time.Duration) *GetVersionNumberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get version number params
func (o *GetVersionNumberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get version number params
func (o *GetVersionNumberParams) WithContext(ctx context.Context) *GetVersionNumberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get version number params
func (o *GetVersionNumberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get version number params
func (o *GetVersionNumberParams) WithHTTPClient(client *http.Client) *GetVersionNumberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get version number params
func (o *GetVersionNumberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetVersionNumberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
