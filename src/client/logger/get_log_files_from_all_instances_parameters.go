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

// NewGetLogFilesFromAllInstancesParams creates a new GetLogFilesFromAllInstancesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLogFilesFromAllInstancesParams() *GetLogFilesFromAllInstancesParams {
	return &GetLogFilesFromAllInstancesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogFilesFromAllInstancesParamsWithTimeout creates a new GetLogFilesFromAllInstancesParams object
// with the ability to set a timeout on a request.
func NewGetLogFilesFromAllInstancesParamsWithTimeout(timeout time.Duration) *GetLogFilesFromAllInstancesParams {
	return &GetLogFilesFromAllInstancesParams{
		timeout: timeout,
	}
}

// NewGetLogFilesFromAllInstancesParamsWithContext creates a new GetLogFilesFromAllInstancesParams object
// with the ability to set a context for a request.
func NewGetLogFilesFromAllInstancesParamsWithContext(ctx context.Context) *GetLogFilesFromAllInstancesParams {
	return &GetLogFilesFromAllInstancesParams{
		Context: ctx,
	}
}

// NewGetLogFilesFromAllInstancesParamsWithHTTPClient creates a new GetLogFilesFromAllInstancesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLogFilesFromAllInstancesParamsWithHTTPClient(client *http.Client) *GetLogFilesFromAllInstancesParams {
	return &GetLogFilesFromAllInstancesParams{
		HTTPClient: client,
	}
}

/*
GetLogFilesFromAllInstancesParams contains all the parameters to send to the API endpoint

	for the get log files from all instances operation.

	Typically these are written to a http.Request.
*/
type GetLogFilesFromAllInstancesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get log files from all instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogFilesFromAllInstancesParams) WithDefaults() *GetLogFilesFromAllInstancesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get log files from all instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogFilesFromAllInstancesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get log files from all instances params
func (o *GetLogFilesFromAllInstancesParams) WithTimeout(timeout time.Duration) *GetLogFilesFromAllInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get log files from all instances params
func (o *GetLogFilesFromAllInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get log files from all instances params
func (o *GetLogFilesFromAllInstancesParams) WithContext(ctx context.Context) *GetLogFilesFromAllInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get log files from all instances params
func (o *GetLogFilesFromAllInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get log files from all instances params
func (o *GetLogFilesFromAllInstancesParams) WithHTTPClient(client *http.Client) *GetLogFilesFromAllInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get log files from all instances params
func (o *GetLogFilesFromAllInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogFilesFromAllInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
