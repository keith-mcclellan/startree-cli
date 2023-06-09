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

// NewGetSubtaskStatesParams creates a new GetSubtaskStatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSubtaskStatesParams() *GetSubtaskStatesParams {
	return &GetSubtaskStatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSubtaskStatesParamsWithTimeout creates a new GetSubtaskStatesParams object
// with the ability to set a timeout on a request.
func NewGetSubtaskStatesParamsWithTimeout(timeout time.Duration) *GetSubtaskStatesParams {
	return &GetSubtaskStatesParams{
		timeout: timeout,
	}
}

// NewGetSubtaskStatesParamsWithContext creates a new GetSubtaskStatesParams object
// with the ability to set a context for a request.
func NewGetSubtaskStatesParamsWithContext(ctx context.Context) *GetSubtaskStatesParams {
	return &GetSubtaskStatesParams{
		Context: ctx,
	}
}

// NewGetSubtaskStatesParamsWithHTTPClient creates a new GetSubtaskStatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSubtaskStatesParamsWithHTTPClient(client *http.Client) *GetSubtaskStatesParams {
	return &GetSubtaskStatesParams{
		HTTPClient: client,
	}
}

/*
GetSubtaskStatesParams contains all the parameters to send to the API endpoint

	for the get subtask states operation.

	Typically these are written to a http.Request.
*/
type GetSubtaskStatesParams struct {

	/* TaskName.

	   Task name
	*/
	TaskName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get subtask states params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSubtaskStatesParams) WithDefaults() *GetSubtaskStatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get subtask states params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSubtaskStatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get subtask states params
func (o *GetSubtaskStatesParams) WithTimeout(timeout time.Duration) *GetSubtaskStatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get subtask states params
func (o *GetSubtaskStatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get subtask states params
func (o *GetSubtaskStatesParams) WithContext(ctx context.Context) *GetSubtaskStatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get subtask states params
func (o *GetSubtaskStatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get subtask states params
func (o *GetSubtaskStatesParams) WithHTTPClient(client *http.Client) *GetSubtaskStatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get subtask states params
func (o *GetSubtaskStatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTaskName adds the taskName to the get subtask states params
func (o *GetSubtaskStatesParams) WithTaskName(taskName string) *GetSubtaskStatesParams {
	o.SetTaskName(taskName)
	return o
}

// SetTaskName adds the taskName to the get subtask states params
func (o *GetSubtaskStatesParams) SetTaskName(taskName string) {
	o.TaskName = taskName
}

// WriteToRequest writes these params to a swagger request
func (o *GetSubtaskStatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
