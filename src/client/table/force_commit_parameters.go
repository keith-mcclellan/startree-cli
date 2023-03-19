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

// NewForceCommitParams creates a new ForceCommitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewForceCommitParams() *ForceCommitParams {
	return &ForceCommitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewForceCommitParamsWithTimeout creates a new ForceCommitParams object
// with the ability to set a timeout on a request.
func NewForceCommitParamsWithTimeout(timeout time.Duration) *ForceCommitParams {
	return &ForceCommitParams{
		timeout: timeout,
	}
}

// NewForceCommitParamsWithContext creates a new ForceCommitParams object
// with the ability to set a context for a request.
func NewForceCommitParamsWithContext(ctx context.Context) *ForceCommitParams {
	return &ForceCommitParams{
		Context: ctx,
	}
}

// NewForceCommitParamsWithHTTPClient creates a new ForceCommitParams object
// with the ability to set a custom HTTPClient for a request.
func NewForceCommitParamsWithHTTPClient(client *http.Client) *ForceCommitParams {
	return &ForceCommitParams{
		HTTPClient: client,
	}
}

/*
ForceCommitParams contains all the parameters to send to the API endpoint

	for the force commit operation.

	Typically these are written to a http.Request.
*/
type ForceCommitParams struct {

	/* TableName.

	   Name of the table
	*/
	TableName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the force commit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ForceCommitParams) WithDefaults() *ForceCommitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the force commit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ForceCommitParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the force commit params
func (o *ForceCommitParams) WithTimeout(timeout time.Duration) *ForceCommitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the force commit params
func (o *ForceCommitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the force commit params
func (o *ForceCommitParams) WithContext(ctx context.Context) *ForceCommitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the force commit params
func (o *ForceCommitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the force commit params
func (o *ForceCommitParams) WithHTTPClient(client *http.Client) *ForceCommitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the force commit params
func (o *ForceCommitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTableName adds the tableName to the force commit params
func (o *ForceCommitParams) WithTableName(tableName string) *ForceCommitParams {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the force commit params
func (o *ForceCommitParams) SetTableName(tableName string) {
	o.TableName = tableName
}

// WriteToRequest writes these params to a swagger request
func (o *ForceCommitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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