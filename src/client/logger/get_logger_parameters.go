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

// NewGetLoggerParams creates a new GetLoggerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLoggerParams() *GetLoggerParams {
	return &GetLoggerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoggerParamsWithTimeout creates a new GetLoggerParams object
// with the ability to set a timeout on a request.
func NewGetLoggerParamsWithTimeout(timeout time.Duration) *GetLoggerParams {
	return &GetLoggerParams{
		timeout: timeout,
	}
}

// NewGetLoggerParamsWithContext creates a new GetLoggerParams object
// with the ability to set a context for a request.
func NewGetLoggerParamsWithContext(ctx context.Context) *GetLoggerParams {
	return &GetLoggerParams{
		Context: ctx,
	}
}

// NewGetLoggerParamsWithHTTPClient creates a new GetLoggerParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLoggerParamsWithHTTPClient(client *http.Client) *GetLoggerParams {
	return &GetLoggerParams{
		HTTPClient: client,
	}
}

/*
GetLoggerParams contains all the parameters to send to the API endpoint

	for the get logger operation.

	Typically these are written to a http.Request.
*/
type GetLoggerParams struct {

	/* LoggerName.

	   Logger name
	*/
	LoggerName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get logger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoggerParams) WithDefaults() *GetLoggerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get logger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoggerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get logger params
func (o *GetLoggerParams) WithTimeout(timeout time.Duration) *GetLoggerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get logger params
func (o *GetLoggerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get logger params
func (o *GetLoggerParams) WithContext(ctx context.Context) *GetLoggerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get logger params
func (o *GetLoggerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get logger params
func (o *GetLoggerParams) WithHTTPClient(client *http.Client) *GetLoggerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get logger params
func (o *GetLoggerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLoggerName adds the loggerName to the get logger params
func (o *GetLoggerParams) WithLoggerName(loggerName string) *GetLoggerParams {
	o.SetLoggerName(loggerName)
	return o
}

// SetLoggerName adds the loggerName to the get logger params
func (o *GetLoggerParams) SetLoggerName(loggerName string) {
	o.LoggerName = loggerName
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoggerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param loggerName
	if err := r.SetPathParam("loggerName", o.LoggerName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}