// Code generated by go-swagger; DO NOT EDIT.

package logger

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

// NewGetLoggersParams creates a new GetLoggersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLoggersParams() *GetLoggersParams {
	return &GetLoggersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoggersParamsWithTimeout creates a new GetLoggersParams object
// with the ability to set a timeout on a request.
func NewGetLoggersParamsWithTimeout(timeout time.Duration) *GetLoggersParams {
	return &GetLoggersParams{
		timeout: timeout,
	}
}

// NewGetLoggersParamsWithContext creates a new GetLoggersParams object
// with the ability to set a context for a request.
func NewGetLoggersParamsWithContext(ctx context.Context) *GetLoggersParams {
	return &GetLoggersParams{
		Context: ctx,
	}
}

// NewGetLoggersParamsWithHTTPClient creates a new GetLoggersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLoggersParamsWithHTTPClient(client *http.Client) *GetLoggersParams {
	return &GetLoggersParams{
		HTTPClient: client,
	}
}

/*
GetLoggersParams contains all the parameters to send to the API endpoint

	for the get loggers operation.

	Typically these are written to a http.Request.
*/
type GetLoggersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get loggers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoggersParams) WithDefaults() *GetLoggersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get loggers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoggersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get loggers params
func (o *GetLoggersParams) WithTimeout(timeout time.Duration) *GetLoggersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get loggers params
func (o *GetLoggersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get loggers params
func (o *GetLoggersParams) WithContext(ctx context.Context) *GetLoggersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get loggers params
func (o *GetLoggersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get loggers params
func (o *GetLoggersParams) WithHTTPClient(client *http.Client) *GetLoggersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get loggers params
func (o *GetLoggersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoggersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
