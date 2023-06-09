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

// NewCleanUpTasksParams creates a new CleanUpTasksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCleanUpTasksParams() *CleanUpTasksParams {
	return &CleanUpTasksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCleanUpTasksParamsWithTimeout creates a new CleanUpTasksParams object
// with the ability to set a timeout on a request.
func NewCleanUpTasksParamsWithTimeout(timeout time.Duration) *CleanUpTasksParams {
	return &CleanUpTasksParams{
		timeout: timeout,
	}
}

// NewCleanUpTasksParamsWithContext creates a new CleanUpTasksParams object
// with the ability to set a context for a request.
func NewCleanUpTasksParamsWithContext(ctx context.Context) *CleanUpTasksParams {
	return &CleanUpTasksParams{
		Context: ctx,
	}
}

// NewCleanUpTasksParamsWithHTTPClient creates a new CleanUpTasksParams object
// with the ability to set a custom HTTPClient for a request.
func NewCleanUpTasksParamsWithHTTPClient(client *http.Client) *CleanUpTasksParams {
	return &CleanUpTasksParams{
		HTTPClient: client,
	}
}

/*
CleanUpTasksParams contains all the parameters to send to the API endpoint

	for the clean up tasks operation.

	Typically these are written to a http.Request.
*/
type CleanUpTasksParams struct {

	/* TaskType.

	   Task type
	*/
	TaskType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the clean up tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CleanUpTasksParams) WithDefaults() *CleanUpTasksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the clean up tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CleanUpTasksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the clean up tasks params
func (o *CleanUpTasksParams) WithTimeout(timeout time.Duration) *CleanUpTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the clean up tasks params
func (o *CleanUpTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the clean up tasks params
func (o *CleanUpTasksParams) WithContext(ctx context.Context) *CleanUpTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the clean up tasks params
func (o *CleanUpTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the clean up tasks params
func (o *CleanUpTasksParams) WithHTTPClient(client *http.Client) *CleanUpTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the clean up tasks params
func (o *CleanUpTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTaskType adds the taskType to the clean up tasks params
func (o *CleanUpTasksParams) WithTaskType(taskType string) *CleanUpTasksParams {
	o.SetTaskType(taskType)
	return o
}

// SetTaskType adds the taskType to the clean up tasks params
func (o *CleanUpTasksParams) SetTaskType(taskType string) {
	o.TaskType = taskType
}

// WriteToRequest writes these params to a swagger request
func (o *CleanUpTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
