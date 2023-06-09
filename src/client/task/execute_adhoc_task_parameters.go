// Code generated by go-swagger; DO NOT EDIT.

package task

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

	"startree.ai/cli/models"
)

// NewExecuteAdhocTaskParams creates a new ExecuteAdhocTaskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExecuteAdhocTaskParams() *ExecuteAdhocTaskParams {
	return &ExecuteAdhocTaskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExecuteAdhocTaskParamsWithTimeout creates a new ExecuteAdhocTaskParams object
// with the ability to set a timeout on a request.
func NewExecuteAdhocTaskParamsWithTimeout(timeout time.Duration) *ExecuteAdhocTaskParams {
	return &ExecuteAdhocTaskParams{
		timeout: timeout,
	}
}

// NewExecuteAdhocTaskParamsWithContext creates a new ExecuteAdhocTaskParams object
// with the ability to set a context for a request.
func NewExecuteAdhocTaskParamsWithContext(ctx context.Context) *ExecuteAdhocTaskParams {
	return &ExecuteAdhocTaskParams{
		Context: ctx,
	}
}

// NewExecuteAdhocTaskParamsWithHTTPClient creates a new ExecuteAdhocTaskParams object
// with the ability to set a custom HTTPClient for a request.
func NewExecuteAdhocTaskParamsWithHTTPClient(client *http.Client) *ExecuteAdhocTaskParams {
	return &ExecuteAdhocTaskParams{
		HTTPClient: client,
	}
}

/*
ExecuteAdhocTaskParams contains all the parameters to send to the API endpoint

	for the execute adhoc task operation.

	Typically these are written to a http.Request.
*/
type ExecuteAdhocTaskParams struct {

	// Body.
	Body *models.AdhocTaskConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the execute adhoc task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecuteAdhocTaskParams) WithDefaults() *ExecuteAdhocTaskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the execute adhoc task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecuteAdhocTaskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the execute adhoc task params
func (o *ExecuteAdhocTaskParams) WithTimeout(timeout time.Duration) *ExecuteAdhocTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the execute adhoc task params
func (o *ExecuteAdhocTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the execute adhoc task params
func (o *ExecuteAdhocTaskParams) WithContext(ctx context.Context) *ExecuteAdhocTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the execute adhoc task params
func (o *ExecuteAdhocTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the execute adhoc task params
func (o *ExecuteAdhocTaskParams) WithHTTPClient(client *http.Client) *ExecuteAdhocTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the execute adhoc task params
func (o *ExecuteAdhocTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the execute adhoc task params
func (o *ExecuteAdhocTaskParams) WithBody(body *models.AdhocTaskConfig) *ExecuteAdhocTaskParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the execute adhoc task params
func (o *ExecuteAdhocTaskParams) SetBody(body *models.AdhocTaskConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ExecuteAdhocTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
