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

// NewGetTasksDeprecatedParams creates a new GetTasksDeprecatedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTasksDeprecatedParams() *GetTasksDeprecatedParams {
	return &GetTasksDeprecatedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTasksDeprecatedParamsWithTimeout creates a new GetTasksDeprecatedParams object
// with the ability to set a timeout on a request.
func NewGetTasksDeprecatedParamsWithTimeout(timeout time.Duration) *GetTasksDeprecatedParams {
	return &GetTasksDeprecatedParams{
		timeout: timeout,
	}
}

// NewGetTasksDeprecatedParamsWithContext creates a new GetTasksDeprecatedParams object
// with the ability to set a context for a request.
func NewGetTasksDeprecatedParamsWithContext(ctx context.Context) *GetTasksDeprecatedParams {
	return &GetTasksDeprecatedParams{
		Context: ctx,
	}
}

// NewGetTasksDeprecatedParamsWithHTTPClient creates a new GetTasksDeprecatedParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTasksDeprecatedParamsWithHTTPClient(client *http.Client) *GetTasksDeprecatedParams {
	return &GetTasksDeprecatedParams{
		HTTPClient: client,
	}
}

/*
GetTasksDeprecatedParams contains all the parameters to send to the API endpoint

	for the get tasks deprecated operation.

	Typically these are written to a http.Request.
*/
type GetTasksDeprecatedParams struct {

	/* TaskType.

	   Task type
	*/
	TaskType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get tasks deprecated params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTasksDeprecatedParams) WithDefaults() *GetTasksDeprecatedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get tasks deprecated params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTasksDeprecatedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get tasks deprecated params
func (o *GetTasksDeprecatedParams) WithTimeout(timeout time.Duration) *GetTasksDeprecatedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tasks deprecated params
func (o *GetTasksDeprecatedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tasks deprecated params
func (o *GetTasksDeprecatedParams) WithContext(ctx context.Context) *GetTasksDeprecatedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tasks deprecated params
func (o *GetTasksDeprecatedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tasks deprecated params
func (o *GetTasksDeprecatedParams) WithHTTPClient(client *http.Client) *GetTasksDeprecatedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tasks deprecated params
func (o *GetTasksDeprecatedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTaskType adds the taskType to the get tasks deprecated params
func (o *GetTasksDeprecatedParams) WithTaskType(taskType string) *GetTasksDeprecatedParams {
	o.SetTaskType(taskType)
	return o
}

// SetTaskType adds the taskType to the get tasks deprecated params
func (o *GetTasksDeprecatedParams) SetTaskType(taskType string) {
	o.TaskType = taskType
}

// WriteToRequest writes these params to a swagger request
func (o *GetTasksDeprecatedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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