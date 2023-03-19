// Code generated by go-swagger; DO NOT EDIT.

package table

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

// NewPauseConsumptionParams creates a new PauseConsumptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPauseConsumptionParams() *PauseConsumptionParams {
	return &PauseConsumptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPauseConsumptionParamsWithTimeout creates a new PauseConsumptionParams object
// with the ability to set a timeout on a request.
func NewPauseConsumptionParamsWithTimeout(timeout time.Duration) *PauseConsumptionParams {
	return &PauseConsumptionParams{
		timeout: timeout,
	}
}

// NewPauseConsumptionParamsWithContext creates a new PauseConsumptionParams object
// with the ability to set a context for a request.
func NewPauseConsumptionParamsWithContext(ctx context.Context) *PauseConsumptionParams {
	return &PauseConsumptionParams{
		Context: ctx,
	}
}

// NewPauseConsumptionParamsWithHTTPClient creates a new PauseConsumptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPauseConsumptionParamsWithHTTPClient(client *http.Client) *PauseConsumptionParams {
	return &PauseConsumptionParams{
		HTTPClient: client,
	}
}

/*
PauseConsumptionParams contains all the parameters to send to the API endpoint

	for the pause consumption operation.

	Typically these are written to a http.Request.
*/
type PauseConsumptionParams struct {

	/* TableName.

	   Name of the table
	*/
	TableName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pause consumption params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PauseConsumptionParams) WithDefaults() *PauseConsumptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pause consumption params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PauseConsumptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pause consumption params
func (o *PauseConsumptionParams) WithTimeout(timeout time.Duration) *PauseConsumptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pause consumption params
func (o *PauseConsumptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pause consumption params
func (o *PauseConsumptionParams) WithContext(ctx context.Context) *PauseConsumptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pause consumption params
func (o *PauseConsumptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pause consumption params
func (o *PauseConsumptionParams) WithHTTPClient(client *http.Client) *PauseConsumptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pause consumption params
func (o *PauseConsumptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTableName adds the tableName to the pause consumption params
func (o *PauseConsumptionParams) WithTableName(tableName string) *PauseConsumptionParams {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the pause consumption params
func (o *PauseConsumptionParams) SetTableName(tableName string) {
	o.TableName = tableName
}

// WriteToRequest writes these params to a swagger request
func (o *PauseConsumptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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