// Code generated by go-swagger; DO NOT EDIT.

package table

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

// NewListTablesParams creates a new ListTablesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListTablesParams() *ListTablesParams {
	return &ListTablesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListTablesParamsWithTimeout creates a new ListTablesParams object
// with the ability to set a timeout on a request.
func NewListTablesParamsWithTimeout(timeout time.Duration) *ListTablesParams {
	return &ListTablesParams{
		timeout: timeout,
	}
}

// NewListTablesParamsWithContext creates a new ListTablesParams object
// with the ability to set a context for a request.
func NewListTablesParamsWithContext(ctx context.Context) *ListTablesParams {
	return &ListTablesParams{
		Context: ctx,
	}
}

// NewListTablesParamsWithHTTPClient creates a new ListTablesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListTablesParamsWithHTTPClient(client *http.Client) *ListTablesParams {
	return &ListTablesParams{
		HTTPClient: client,
	}
}

/*
ListTablesParams contains all the parameters to send to the API endpoint

	for the list tables operation.

	Typically these are written to a http.Request.
*/
type ListTablesParams struct {

	/* SortAsc.

	   true|false

	   Default: true
	*/
	SortAsc *bool

	/* SortType.

	   name|creationTime|lastModifiedTime
	*/
	SortType *string

	/* TaskType.

	   Task type
	*/
	TaskType *string

	/* Type.

	   realtime|offline
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list tables params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTablesParams) WithDefaults() *ListTablesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list tables params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTablesParams) SetDefaults() {
	var (
		sortAscDefault = bool(true)
	)

	val := ListTablesParams{
		SortAsc: &sortAscDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list tables params
func (o *ListTablesParams) WithTimeout(timeout time.Duration) *ListTablesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list tables params
func (o *ListTablesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list tables params
func (o *ListTablesParams) WithContext(ctx context.Context) *ListTablesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list tables params
func (o *ListTablesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list tables params
func (o *ListTablesParams) WithHTTPClient(client *http.Client) *ListTablesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list tables params
func (o *ListTablesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSortAsc adds the sortAsc to the list tables params
func (o *ListTablesParams) WithSortAsc(sortAsc *bool) *ListTablesParams {
	o.SetSortAsc(sortAsc)
	return o
}

// SetSortAsc adds the sortAsc to the list tables params
func (o *ListTablesParams) SetSortAsc(sortAsc *bool) {
	o.SortAsc = sortAsc
}

// WithSortType adds the sortType to the list tables params
func (o *ListTablesParams) WithSortType(sortType *string) *ListTablesParams {
	o.SetSortType(sortType)
	return o
}

// SetSortType adds the sortType to the list tables params
func (o *ListTablesParams) SetSortType(sortType *string) {
	o.SortType = sortType
}

// WithTaskType adds the taskType to the list tables params
func (o *ListTablesParams) WithTaskType(taskType *string) *ListTablesParams {
	o.SetTaskType(taskType)
	return o
}

// SetTaskType adds the taskType to the list tables params
func (o *ListTablesParams) SetTaskType(taskType *string) {
	o.TaskType = taskType
}

// WithType adds the typeVar to the list tables params
func (o *ListTablesParams) WithType(typeVar *string) *ListTablesParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the list tables params
func (o *ListTablesParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *ListTablesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SortAsc != nil {

		// query param sortAsc
		var qrSortAsc bool

		if o.SortAsc != nil {
			qrSortAsc = *o.SortAsc
		}
		qSortAsc := swag.FormatBool(qrSortAsc)
		if qSortAsc != "" {

			if err := r.SetQueryParam("sortAsc", qSortAsc); err != nil {
				return err
			}
		}
	}

	if o.SortType != nil {

		// query param sortType
		var qrSortType string

		if o.SortType != nil {
			qrSortType = *o.SortType
		}
		qSortType := qrSortType
		if qSortType != "" {

			if err := r.SetQueryParam("sortType", qSortType); err != nil {
				return err
			}
		}
	}

	if o.TaskType != nil {

		// query param taskType
		var qrTaskType string

		if o.TaskType != nil {
			qrTaskType = *o.TaskType
		}
		qTaskType := qrTaskType
		if qTaskType != "" {

			if err := r.SetQueryParam("taskType", qTaskType); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
