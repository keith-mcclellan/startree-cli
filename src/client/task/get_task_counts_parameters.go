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

// NewGetTaskCountsParams creates a new GetTaskCountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTaskCountsParams() *GetTaskCountsParams {
	return &GetTaskCountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTaskCountsParamsWithTimeout creates a new GetTaskCountsParams object
// with the ability to set a timeout on a request.
func NewGetTaskCountsParamsWithTimeout(timeout time.Duration) *GetTaskCountsParams {
	return &GetTaskCountsParams{
		timeout: timeout,
	}
}

// NewGetTaskCountsParamsWithContext creates a new GetTaskCountsParams object
// with the ability to set a context for a request.
func NewGetTaskCountsParamsWithContext(ctx context.Context) *GetTaskCountsParams {
	return &GetTaskCountsParams{
		Context: ctx,
	}
}

// NewGetTaskCountsParamsWithHTTPClient creates a new GetTaskCountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTaskCountsParamsWithHTTPClient(client *http.Client) *GetTaskCountsParams {
	return &GetTaskCountsParams{
		HTTPClient: client,
	}
}

/*
GetTaskCountsParams contains all the parameters to send to the API endpoint

	for the get task counts operation.

	Typically these are written to a http.Request.
*/
type GetTaskCountsParams struct {

	/* TaskType.

	   Task type
	*/
	TaskType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get task counts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTaskCountsParams) WithDefaults() *GetTaskCountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get task counts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTaskCountsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get task counts params
func (o *GetTaskCountsParams) WithTimeout(timeout time.Duration) *GetTaskCountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get task counts params
func (o *GetTaskCountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get task counts params
func (o *GetTaskCountsParams) WithContext(ctx context.Context) *GetTaskCountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get task counts params
func (o *GetTaskCountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get task counts params
func (o *GetTaskCountsParams) WithHTTPClient(client *http.Client) *GetTaskCountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get task counts params
func (o *GetTaskCountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTaskType adds the taskType to the get task counts params
func (o *GetTaskCountsParams) WithTaskType(taskType string) *GetTaskCountsParams {
	o.SetTaskType(taskType)
	return o
}

// SetTaskType adds the taskType to the get task counts params
func (o *GetTaskCountsParams) SetTaskType(taskType string) {
	o.TaskType = taskType
}

// WriteToRequest writes these params to a swagger request
func (o *GetTaskCountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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