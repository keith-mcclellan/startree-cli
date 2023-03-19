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
)

// NewGetTaskConfigParams creates a new GetTaskConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTaskConfigParams() *GetTaskConfigParams {
	return &GetTaskConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTaskConfigParamsWithTimeout creates a new GetTaskConfigParams object
// with the ability to set a timeout on a request.
func NewGetTaskConfigParamsWithTimeout(timeout time.Duration) *GetTaskConfigParams {
	return &GetTaskConfigParams{
		timeout: timeout,
	}
}

// NewGetTaskConfigParamsWithContext creates a new GetTaskConfigParams object
// with the ability to set a context for a request.
func NewGetTaskConfigParamsWithContext(ctx context.Context) *GetTaskConfigParams {
	return &GetTaskConfigParams{
		Context: ctx,
	}
}

// NewGetTaskConfigParamsWithHTTPClient creates a new GetTaskConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTaskConfigParamsWithHTTPClient(client *http.Client) *GetTaskConfigParams {
	return &GetTaskConfigParams{
		HTTPClient: client,
	}
}

/*
GetTaskConfigParams contains all the parameters to send to the API endpoint

	for the get task config operation.

	Typically these are written to a http.Request.
*/
type GetTaskConfigParams struct {

	/* TaskName.

	   Task name
	*/
	TaskName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get task config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTaskConfigParams) WithDefaults() *GetTaskConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get task config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTaskConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get task config params
func (o *GetTaskConfigParams) WithTimeout(timeout time.Duration) *GetTaskConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get task config params
func (o *GetTaskConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get task config params
func (o *GetTaskConfigParams) WithContext(ctx context.Context) *GetTaskConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get task config params
func (o *GetTaskConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get task config params
func (o *GetTaskConfigParams) WithHTTPClient(client *http.Client) *GetTaskConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get task config params
func (o *GetTaskConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTaskName adds the taskName to the get task config params
func (o *GetTaskConfigParams) WithTaskName(taskName string) *GetTaskConfigParams {
	o.SetTaskName(taskName)
	return o
}

// SetTaskName adds the taskName to the get task config params
func (o *GetTaskConfigParams) SetTaskName(taskName string) {
	o.TaskName = taskName
}

// WriteToRequest writes these params to a swagger request
func (o *GetTaskConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param taskName
	if err := r.SetPathParam("taskName", o.TaskName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
