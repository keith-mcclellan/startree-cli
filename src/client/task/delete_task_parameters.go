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

// NewDeleteTaskParams creates a new DeleteTaskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTaskParams() *DeleteTaskParams {
	return &DeleteTaskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTaskParamsWithTimeout creates a new DeleteTaskParams object
// with the ability to set a timeout on a request.
func NewDeleteTaskParamsWithTimeout(timeout time.Duration) *DeleteTaskParams {
	return &DeleteTaskParams{
		timeout: timeout,
	}
}

// NewDeleteTaskParamsWithContext creates a new DeleteTaskParams object
// with the ability to set a context for a request.
func NewDeleteTaskParamsWithContext(ctx context.Context) *DeleteTaskParams {
	return &DeleteTaskParams{
		Context: ctx,
	}
}

// NewDeleteTaskParamsWithHTTPClient creates a new DeleteTaskParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTaskParamsWithHTTPClient(client *http.Client) *DeleteTaskParams {
	return &DeleteTaskParams{
		HTTPClient: client,
	}
}

/*
DeleteTaskParams contains all the parameters to send to the API endpoint

	for the delete task operation.

	Typically these are written to a http.Request.
*/
type DeleteTaskParams struct {

	/* ForceDelete.

	   Whether to force deleting the task (expert only option, enable with cautious
	*/
	ForceDelete *bool

	/* TaskName.

	   Task name
	*/
	TaskName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTaskParams) WithDefaults() *DeleteTaskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTaskParams) SetDefaults() {
	var (
		forceDeleteDefault = bool(false)
	)

	val := DeleteTaskParams{
		ForceDelete: &forceDeleteDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete task params
func (o *DeleteTaskParams) WithTimeout(timeout time.Duration) *DeleteTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete task params
func (o *DeleteTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete task params
func (o *DeleteTaskParams) WithContext(ctx context.Context) *DeleteTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete task params
func (o *DeleteTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete task params
func (o *DeleteTaskParams) WithHTTPClient(client *http.Client) *DeleteTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete task params
func (o *DeleteTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceDelete adds the forceDelete to the delete task params
func (o *DeleteTaskParams) WithForceDelete(forceDelete *bool) *DeleteTaskParams {
	o.SetForceDelete(forceDelete)
	return o
}

// SetForceDelete adds the forceDelete to the delete task params
func (o *DeleteTaskParams) SetForceDelete(forceDelete *bool) {
	o.ForceDelete = forceDelete
}

// WithTaskName adds the taskName to the delete task params
func (o *DeleteTaskParams) WithTaskName(taskName string) *DeleteTaskParams {
	o.SetTaskName(taskName)
	return o
}

// SetTaskName adds the taskName to the delete task params
func (o *DeleteTaskParams) SetTaskName(taskName string) {
	o.TaskName = taskName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ForceDelete != nil {

		// query param forceDelete
		var qrForceDelete bool

		if o.ForceDelete != nil {
			qrForceDelete = *o.ForceDelete
		}
		qForceDelete := swag.FormatBool(qrForceDelete)
		if qForceDelete != "" {

			if err := r.SetQueryParam("forceDelete", qForceDelete); err != nil {
				return err
			}
		}
	}

	// path param taskName
	if err := r.SetPathParam("taskName", o.TaskName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
