// Code generated by go-swagger; DO NOT EDIT.

package segment

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

// NewDeleteAllSegmentsParams creates a new DeleteAllSegmentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAllSegmentsParams() *DeleteAllSegmentsParams {
	return &DeleteAllSegmentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAllSegmentsParamsWithTimeout creates a new DeleteAllSegmentsParams object
// with the ability to set a timeout on a request.
func NewDeleteAllSegmentsParamsWithTimeout(timeout time.Duration) *DeleteAllSegmentsParams {
	return &DeleteAllSegmentsParams{
		timeout: timeout,
	}
}

// NewDeleteAllSegmentsParamsWithContext creates a new DeleteAllSegmentsParams object
// with the ability to set a context for a request.
func NewDeleteAllSegmentsParamsWithContext(ctx context.Context) *DeleteAllSegmentsParams {
	return &DeleteAllSegmentsParams{
		Context: ctx,
	}
}

// NewDeleteAllSegmentsParamsWithHTTPClient creates a new DeleteAllSegmentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAllSegmentsParamsWithHTTPClient(client *http.Client) *DeleteAllSegmentsParams {
	return &DeleteAllSegmentsParams{
		HTTPClient: client,
	}
}

/*
DeleteAllSegmentsParams contains all the parameters to send to the API endpoint

	for the delete all segments operation.

	Typically these are written to a http.Request.
*/
type DeleteAllSegmentsParams struct {

	/* Retention.

	   Retention period for the table segments (e.g. 12h, 3d); If not set, the retention period will default to the first config that's not null: the table config, then to cluster setting, then '7d'. Using 0d or -1d will instantly delete segments without retention
	*/
	Retention *string

	/* TableName.

	   Name of the table
	*/
	TableName string

	/* Type.

	   OFFLINE|REALTIME
	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete all segments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAllSegmentsParams) WithDefaults() *DeleteAllSegmentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete all segments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAllSegmentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete all segments params
func (o *DeleteAllSegmentsParams) WithTimeout(timeout time.Duration) *DeleteAllSegmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete all segments params
func (o *DeleteAllSegmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete all segments params
func (o *DeleteAllSegmentsParams) WithContext(ctx context.Context) *DeleteAllSegmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete all segments params
func (o *DeleteAllSegmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete all segments params
func (o *DeleteAllSegmentsParams) WithHTTPClient(client *http.Client) *DeleteAllSegmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete all segments params
func (o *DeleteAllSegmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRetention adds the retention to the delete all segments params
func (o *DeleteAllSegmentsParams) WithRetention(retention *string) *DeleteAllSegmentsParams {
	o.SetRetention(retention)
	return o
}

// SetRetention adds the retention to the delete all segments params
func (o *DeleteAllSegmentsParams) SetRetention(retention *string) {
	o.Retention = retention
}

// WithTableName adds the tableName to the delete all segments params
func (o *DeleteAllSegmentsParams) WithTableName(tableName string) *DeleteAllSegmentsParams {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the delete all segments params
func (o *DeleteAllSegmentsParams) SetTableName(tableName string) {
	o.TableName = tableName
}

// WithType adds the typeVar to the delete all segments params
func (o *DeleteAllSegmentsParams) WithType(typeVar string) *DeleteAllSegmentsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the delete all segments params
func (o *DeleteAllSegmentsParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAllSegmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Retention != nil {

		// query param retention
		var qrRetention string

		if o.Retention != nil {
			qrRetention = *o.Retention
		}
		qRetention := qrRetention
		if qRetention != "" {

			if err := r.SetQueryParam("retention", qRetention); err != nil {
				return err
			}
		}
	}

	// path param tableName
	if err := r.SetPathParam("tableName", o.TableName); err != nil {
		return err
	}

	// query param type
	qrType := o.Type
	qType := qrType
	if qType != "" {

		if err := r.SetQueryParam("type", qType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}