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
	"github.com/go-openapi/swag"

	"startree.ai/cli/models"
)

// NewStartReplaceSegmentsParams creates a new StartReplaceSegmentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartReplaceSegmentsParams() *StartReplaceSegmentsParams {
	return &StartReplaceSegmentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartReplaceSegmentsParamsWithTimeout creates a new StartReplaceSegmentsParams object
// with the ability to set a timeout on a request.
func NewStartReplaceSegmentsParamsWithTimeout(timeout time.Duration) *StartReplaceSegmentsParams {
	return &StartReplaceSegmentsParams{
		timeout: timeout,
	}
}

// NewStartReplaceSegmentsParamsWithContext creates a new StartReplaceSegmentsParams object
// with the ability to set a context for a request.
func NewStartReplaceSegmentsParamsWithContext(ctx context.Context) *StartReplaceSegmentsParams {
	return &StartReplaceSegmentsParams{
		Context: ctx,
	}
}

// NewStartReplaceSegmentsParamsWithHTTPClient creates a new StartReplaceSegmentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartReplaceSegmentsParamsWithHTTPClient(client *http.Client) *StartReplaceSegmentsParams {
	return &StartReplaceSegmentsParams{
		HTTPClient: client,
	}
}

/*
StartReplaceSegmentsParams contains all the parameters to send to the API endpoint

	for the start replace segments operation.

	Typically these are written to a http.Request.
*/
type StartReplaceSegmentsParams struct {

	/* Body.

	   Fields belonging to start replace segment request
	*/
	Body *models.StartReplaceSegmentsRequest

	/* ForceCleanup.

	   Force cleanup
	*/
	ForceCleanup *bool

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

// WithDefaults hydrates default values in the start replace segments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartReplaceSegmentsParams) WithDefaults() *StartReplaceSegmentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the start replace segments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartReplaceSegmentsParams) SetDefaults() {
	var (
		forceCleanupDefault = bool(false)
	)

	val := StartReplaceSegmentsParams{
		ForceCleanup: &forceCleanupDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the start replace segments params
func (o *StartReplaceSegmentsParams) WithTimeout(timeout time.Duration) *StartReplaceSegmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start replace segments params
func (o *StartReplaceSegmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start replace segments params
func (o *StartReplaceSegmentsParams) WithContext(ctx context.Context) *StartReplaceSegmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start replace segments params
func (o *StartReplaceSegmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start replace segments params
func (o *StartReplaceSegmentsParams) WithHTTPClient(client *http.Client) *StartReplaceSegmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start replace segments params
func (o *StartReplaceSegmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the start replace segments params
func (o *StartReplaceSegmentsParams) WithBody(body *models.StartReplaceSegmentsRequest) *StartReplaceSegmentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the start replace segments params
func (o *StartReplaceSegmentsParams) SetBody(body *models.StartReplaceSegmentsRequest) {
	o.Body = body
}

// WithForceCleanup adds the forceCleanup to the start replace segments params
func (o *StartReplaceSegmentsParams) WithForceCleanup(forceCleanup *bool) *StartReplaceSegmentsParams {
	o.SetForceCleanup(forceCleanup)
	return o
}

// SetForceCleanup adds the forceCleanup to the start replace segments params
func (o *StartReplaceSegmentsParams) SetForceCleanup(forceCleanup *bool) {
	o.ForceCleanup = forceCleanup
}

// WithTableName adds the tableName to the start replace segments params
func (o *StartReplaceSegmentsParams) WithTableName(tableName string) *StartReplaceSegmentsParams {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the start replace segments params
func (o *StartReplaceSegmentsParams) SetTableName(tableName string) {
	o.TableName = tableName
}

// WithType adds the typeVar to the start replace segments params
func (o *StartReplaceSegmentsParams) WithType(typeVar string) *StartReplaceSegmentsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the start replace segments params
func (o *StartReplaceSegmentsParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *StartReplaceSegmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.ForceCleanup != nil {

		// query param forceCleanup
		var qrForceCleanup bool

		if o.ForceCleanup != nil {
			qrForceCleanup = *o.ForceCleanup
		}
		qForceCleanup := swag.FormatBool(qrForceCleanup)
		if qForceCleanup != "" {

			if err := r.SetQueryParam("forceCleanup", qForceCleanup); err != nil {
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
