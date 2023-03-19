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

// NewGetLocalLogFilesParams creates a new GetLocalLogFilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLocalLogFilesParams() *GetLocalLogFilesParams {
	return &GetLocalLogFilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLocalLogFilesParamsWithTimeout creates a new GetLocalLogFilesParams object
// with the ability to set a timeout on a request.
func NewGetLocalLogFilesParamsWithTimeout(timeout time.Duration) *GetLocalLogFilesParams {
	return &GetLocalLogFilesParams{
		timeout: timeout,
	}
}

// NewGetLocalLogFilesParamsWithContext creates a new GetLocalLogFilesParams object
// with the ability to set a context for a request.
func NewGetLocalLogFilesParamsWithContext(ctx context.Context) *GetLocalLogFilesParams {
	return &GetLocalLogFilesParams{
		Context: ctx,
	}
}

// NewGetLocalLogFilesParamsWithHTTPClient creates a new GetLocalLogFilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLocalLogFilesParamsWithHTTPClient(client *http.Client) *GetLocalLogFilesParams {
	return &GetLocalLogFilesParams{
		HTTPClient: client,
	}
}

/*
GetLocalLogFilesParams contains all the parameters to send to the API endpoint

	for the get local log files operation.

	Typically these are written to a http.Request.
*/
type GetLocalLogFilesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get local log files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLocalLogFilesParams) WithDefaults() *GetLocalLogFilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get local log files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLocalLogFilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get local log files params
func (o *GetLocalLogFilesParams) WithTimeout(timeout time.Duration) *GetLocalLogFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get local log files params
func (o *GetLocalLogFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get local log files params
func (o *GetLocalLogFilesParams) WithContext(ctx context.Context) *GetLocalLogFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get local log files params
func (o *GetLocalLogFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get local log files params
func (o *GetLocalLogFilesParams) WithHTTPClient(client *http.Client) *GetLocalLogFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get local log files params
func (o *GetLocalLogFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLocalLogFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}