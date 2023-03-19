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
	"github.com/go-openapi/swag"
)

// NewGetTasksDebugInfoParams creates a new GetTasksDebugInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTasksDebugInfoParams() *GetTasksDebugInfoParams {
	return &GetTasksDebugInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTasksDebugInfoParamsWithTimeout creates a new GetTasksDebugInfoParams object
// with the ability to set a timeout on a request.
func NewGetTasksDebugInfoParamsWithTimeout(timeout time.Duration) *GetTasksDebugInfoParams {
	return &GetTasksDebugInfoParams{
		timeout: timeout,
	}
}

// NewGetTasksDebugInfoParamsWithContext creates a new GetTasksDebugInfoParams object
// with the ability to set a context for a request.
func NewGetTasksDebugInfoParamsWithContext(ctx context.Context) *GetTasksDebugInfoParams {
	return &GetTasksDebugInfoParams{
		Context: ctx,
	}
}

// NewGetTasksDebugInfoParamsWithHTTPClient creates a new GetTasksDebugInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTasksDebugInfoParamsWithHTTPClient(client *http.Client) *GetTasksDebugInfoParams {
	return &GetTasksDebugInfoParams{
		HTTPClient: client,
	}
}

/*
GetTasksDebugInfoParams contains all the parameters to send to the API endpoint

	for the get tasks debug info operation.

	Typically these are written to a http.Request.
*/
type GetTasksDebugInfoParams struct {

	/* TaskType.

	   Task type
	*/
	TaskType string

	/* Verbosity.

	   verbosity (Prints information for all the tasks for the given task type.By default, only prints subtask details for running and error tasks. Value of > 0 prints subtask details for all tasks)

	   Format: int32
	*/
	Verbosity *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get tasks debug info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTasksDebugInfoParams) WithDefaults() *GetTasksDebugInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get tasks debug info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTasksDebugInfoParams) SetDefaults() {
	var (
		verbosityDefault = int32(0)
	)

	val := GetTasksDebugInfoParams{
		Verbosity: &verbosityDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get tasks debug info params
func (o *GetTasksDebugInfoParams) WithTimeout(timeout time.Duration) *GetTasksDebugInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tasks debug info params
func (o *GetTasksDebugInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tasks debug info params
func (o *GetTasksDebugInfoParams) WithContext(ctx context.Context) *GetTasksDebugInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tasks debug info params
func (o *GetTasksDebugInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tasks debug info params
func (o *GetTasksDebugInfoParams) WithHTTPClient(client *http.Client) *GetTasksDebugInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tasks debug info params
func (o *GetTasksDebugInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTaskType adds the taskType to the get tasks debug info params
func (o *GetTasksDebugInfoParams) WithTaskType(taskType string) *GetTasksDebugInfoParams {
	o.SetTaskType(taskType)
	return o
}

// SetTaskType adds the taskType to the get tasks debug info params
func (o *GetTasksDebugInfoParams) SetTaskType(taskType string) {
	o.TaskType = taskType
}

// WithVerbosity adds the verbosity to the get tasks debug info params
func (o *GetTasksDebugInfoParams) WithVerbosity(verbosity *int32) *GetTasksDebugInfoParams {
	o.SetVerbosity(verbosity)
	return o
}

// SetVerbosity adds the verbosity to the get tasks debug info params
func (o *GetTasksDebugInfoParams) SetVerbosity(verbosity *int32) {
	o.Verbosity = verbosity
}

// WriteToRequest writes these params to a swagger request
func (o *GetTasksDebugInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param taskType
	if err := r.SetPathParam("taskType", o.TaskType); err != nil {
		return err
	}

	if o.Verbosity != nil {

		// query param verbosity
		var qrVerbosity int32

		if o.Verbosity != nil {
			qrVerbosity = *o.Verbosity
		}
		qVerbosity := swag.FormatInt32(qrVerbosity)
		if qVerbosity != "" {

			if err := r.SetQueryParam("verbosity", qVerbosity); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
