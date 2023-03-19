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

// NewGetTaskMetadataByTableParams creates a new GetTaskMetadataByTableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTaskMetadataByTableParams() *GetTaskMetadataByTableParams {
	return &GetTaskMetadataByTableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTaskMetadataByTableParamsWithTimeout creates a new GetTaskMetadataByTableParams object
// with the ability to set a timeout on a request.
func NewGetTaskMetadataByTableParamsWithTimeout(timeout time.Duration) *GetTaskMetadataByTableParams {
	return &GetTaskMetadataByTableParams{
		timeout: timeout,
	}
}

// NewGetTaskMetadataByTableParamsWithContext creates a new GetTaskMetadataByTableParams object
// with the ability to set a context for a request.
func NewGetTaskMetadataByTableParamsWithContext(ctx context.Context) *GetTaskMetadataByTableParams {
	return &GetTaskMetadataByTableParams{
		Context: ctx,
	}
}

// NewGetTaskMetadataByTableParamsWithHTTPClient creates a new GetTaskMetadataByTableParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTaskMetadataByTableParamsWithHTTPClient(client *http.Client) *GetTaskMetadataByTableParams {
	return &GetTaskMetadataByTableParams{
		HTTPClient: client,
	}
}

/*
GetTaskMetadataByTableParams contains all the parameters to send to the API endpoint

	for the get task metadata by table operation.

	Typically these are written to a http.Request.
*/
type GetTaskMetadataByTableParams struct {

	/* TableNameWithType.

	   Table name with type
	*/
	TableNameWithType string

	/* TaskType.

	   Task type
	*/
	TaskType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get task metadata by table params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTaskMetadataByTableParams) WithDefaults() *GetTaskMetadataByTableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get task metadata by table params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTaskMetadataByTableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get task metadata by table params
func (o *GetTaskMetadataByTableParams) WithTimeout(timeout time.Duration) *GetTaskMetadataByTableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get task metadata by table params
func (o *GetTaskMetadataByTableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get task metadata by table params
func (o *GetTaskMetadataByTableParams) WithContext(ctx context.Context) *GetTaskMetadataByTableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get task metadata by table params
func (o *GetTaskMetadataByTableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get task metadata by table params
func (o *GetTaskMetadataByTableParams) WithHTTPClient(client *http.Client) *GetTaskMetadataByTableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get task metadata by table params
func (o *GetTaskMetadataByTableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTableNameWithType adds the tableNameWithType to the get task metadata by table params
func (o *GetTaskMetadataByTableParams) WithTableNameWithType(tableNameWithType string) *GetTaskMetadataByTableParams {
	o.SetTableNameWithType(tableNameWithType)
	return o
}

// SetTableNameWithType adds the tableNameWithType to the get task metadata by table params
func (o *GetTaskMetadataByTableParams) SetTableNameWithType(tableNameWithType string) {
	o.TableNameWithType = tableNameWithType
}

// WithTaskType adds the taskType to the get task metadata by table params
func (o *GetTaskMetadataByTableParams) WithTaskType(taskType string) *GetTaskMetadataByTableParams {
	o.SetTaskType(taskType)
	return o
}

// SetTaskType adds the taskType to the get task metadata by table params
func (o *GetTaskMetadataByTableParams) SetTaskType(taskType string) {
	o.TaskType = taskType
}

// WriteToRequest writes these params to a swagger request
func (o *GetTaskMetadataByTableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tableNameWithType
	if err := r.SetPathParam("tableNameWithType", o.TableNameWithType); err != nil {
		return err
	}

	// path param taskType
	if err := r.SetPathParam("taskType", o.TaskType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}