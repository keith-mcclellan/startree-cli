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

// NewGetTaskStatesParams creates a new GetTaskStatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTaskStatesParams() *GetTaskStatesParams {
	return &GetTaskStatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTaskStatesParamsWithTimeout creates a new GetTaskStatesParams object
// with the ability to set a timeout on a request.
func NewGetTaskStatesParamsWithTimeout(timeout time.Duration) *GetTaskStatesParams {
	return &GetTaskStatesParams{
		timeout: timeout,
	}
}

// NewGetTaskStatesParamsWithContext creates a new GetTaskStatesParams object
// with the ability to set a context for a request.
func NewGetTaskStatesParamsWithContext(ctx context.Context) *GetTaskStatesParams {
	return &GetTaskStatesParams{
		Context: ctx,
	}
}

// NewGetTaskStatesParamsWithHTTPClient creates a new GetTaskStatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTaskStatesParamsWithHTTPClient(client *http.Client) *GetTaskStatesParams {
	return &GetTaskStatesParams{
		HTTPClient: client,
	}
}

/*
GetTaskStatesParams contains all the parameters to send to the API endpoint

	for the get task states operation.

	Typically these are written to a http.Request.
*/
type GetTaskStatesParams struct {

	/* TaskType.

	   Task type
	*/
	TaskType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get task states params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTaskStatesParams) WithDefaults() *GetTaskStatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get task states params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTaskStatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get task states params
func (o *GetTaskStatesParams) WithTimeout(timeout time.Duration) *GetTaskStatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get task states params
func (o *GetTaskStatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get task states params
func (o *GetTaskStatesParams) WithContext(ctx context.Context) *GetTaskStatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get task states params
func (o *GetTaskStatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get task states params
func (o *GetTaskStatesParams) WithHTTPClient(client *http.Client) *GetTaskStatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get task states params
func (o *GetTaskStatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTaskType adds the taskType to the get task states params
func (o *GetTaskStatesParams) WithTaskType(taskType string) *GetTaskStatesParams {
	o.SetTaskType(taskType)
	return o
}

// SetTaskType adds the taskType to the get task states params
func (o *GetTaskStatesParams) SetTaskType(taskType string) {
	o.TaskType = taskType
}

// WriteToRequest writes these params to a swagger request
func (o *GetTaskStatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param taskType
	if err := r.SetPathParam("taskType", o.TaskType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
