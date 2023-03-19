// Code generated by go-swagger; DO NOT EDIT.

package tuner

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

// NewTuneTableParams creates a new TuneTableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTuneTableParams() *TuneTableParams {
	return &TuneTableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTuneTableParamsWithTimeout creates a new TuneTableParams object
// with the ability to set a timeout on a request.
func NewTuneTableParamsWithTimeout(timeout time.Duration) *TuneTableParams {
	return &TuneTableParams{
		timeout: timeout,
	}
}

// NewTuneTableParamsWithContext creates a new TuneTableParams object
// with the ability to set a context for a request.
func NewTuneTableParamsWithContext(ctx context.Context) *TuneTableParams {
	return &TuneTableParams{
		Context: ctx,
	}
}

// NewTuneTableParamsWithHTTPClient creates a new TuneTableParams object
// with the ability to set a custom HTTPClient for a request.
func NewTuneTableParamsWithHTTPClient(client *http.Client) *TuneTableParams {
	return &TuneTableParams{
		HTTPClient: client,
	}
}

/*
TuneTableParams contains all the parameters to send to the API endpoint

	for the tune table operation.

	Typically these are written to a http.Request.
*/
type TuneTableParams struct {

	/* TableName.

	   Name of the table
	*/
	TableName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tune table params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TuneTableParams) WithDefaults() *TuneTableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tune table params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TuneTableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tune table params
func (o *TuneTableParams) WithTimeout(timeout time.Duration) *TuneTableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tune table params
func (o *TuneTableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tune table params
func (o *TuneTableParams) WithContext(ctx context.Context) *TuneTableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tune table params
func (o *TuneTableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tune table params
func (o *TuneTableParams) WithHTTPClient(client *http.Client) *TuneTableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tune table params
func (o *TuneTableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTableName adds the tableName to the tune table params
func (o *TuneTableParams) WithTableName(tableName string) *TuneTableParams {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the tune table params
func (o *TuneTableParams) SetTableName(tableName string) {
	o.TableName = tableName
}

// WriteToRequest writes these params to a swagger request
func (o *TuneTableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tableName
	if err := r.SetPathParam("tableName", o.TableName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}