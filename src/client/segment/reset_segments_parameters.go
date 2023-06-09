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
)

// NewResetSegmentsParams creates a new ResetSegmentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResetSegmentsParams() *ResetSegmentsParams {
	return &ResetSegmentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResetSegmentsParamsWithTimeout creates a new ResetSegmentsParams object
// with the ability to set a timeout on a request.
func NewResetSegmentsParamsWithTimeout(timeout time.Duration) *ResetSegmentsParams {
	return &ResetSegmentsParams{
		timeout: timeout,
	}
}

// NewResetSegmentsParamsWithContext creates a new ResetSegmentsParams object
// with the ability to set a context for a request.
func NewResetSegmentsParamsWithContext(ctx context.Context) *ResetSegmentsParams {
	return &ResetSegmentsParams{
		Context: ctx,
	}
}

// NewResetSegmentsParamsWithHTTPClient creates a new ResetSegmentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewResetSegmentsParamsWithHTTPClient(client *http.Client) *ResetSegmentsParams {
	return &ResetSegmentsParams{
		HTTPClient: client,
	}
}

/*
ResetSegmentsParams contains all the parameters to send to the API endpoint

	for the reset segments operation.

	Typically these are written to a http.Request.
*/
type ResetSegmentsParams struct {

	/* ErrorSegmentsOnly.

	   Whether to reset only segments with error state
	*/
	ErrorSegmentsOnly *bool

	/* TableNameWithType.

	   Name of the table with type
	*/
	TableNameWithType string

	/* TargetInstance.

	   Name of the target instance to reset
	*/
	TargetInstance *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reset segments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetSegmentsParams) WithDefaults() *ResetSegmentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reset segments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetSegmentsParams) SetDefaults() {
	var (
		errorSegmentsOnlyDefault = bool(false)
	)

	val := ResetSegmentsParams{
		ErrorSegmentsOnly: &errorSegmentsOnlyDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the reset segments params
func (o *ResetSegmentsParams) WithTimeout(timeout time.Duration) *ResetSegmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reset segments params
func (o *ResetSegmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reset segments params
func (o *ResetSegmentsParams) WithContext(ctx context.Context) *ResetSegmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reset segments params
func (o *ResetSegmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reset segments params
func (o *ResetSegmentsParams) WithHTTPClient(client *http.Client) *ResetSegmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reset segments params
func (o *ResetSegmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithErrorSegmentsOnly adds the errorSegmentsOnly to the reset segments params
func (o *ResetSegmentsParams) WithErrorSegmentsOnly(errorSegmentsOnly *bool) *ResetSegmentsParams {
	o.SetErrorSegmentsOnly(errorSegmentsOnly)
	return o
}

// SetErrorSegmentsOnly adds the errorSegmentsOnly to the reset segments params
func (o *ResetSegmentsParams) SetErrorSegmentsOnly(errorSegmentsOnly *bool) {
	o.ErrorSegmentsOnly = errorSegmentsOnly
}

// WithTableNameWithType adds the tableNameWithType to the reset segments params
func (o *ResetSegmentsParams) WithTableNameWithType(tableNameWithType string) *ResetSegmentsParams {
	o.SetTableNameWithType(tableNameWithType)
	return o
}

// SetTableNameWithType adds the tableNameWithType to the reset segments params
func (o *ResetSegmentsParams) SetTableNameWithType(tableNameWithType string) {
	o.TableNameWithType = tableNameWithType
}

// WithTargetInstance adds the targetInstance to the reset segments params
func (o *ResetSegmentsParams) WithTargetInstance(targetInstance *string) *ResetSegmentsParams {
	o.SetTargetInstance(targetInstance)
	return o
}

// SetTargetInstance adds the targetInstance to the reset segments params
func (o *ResetSegmentsParams) SetTargetInstance(targetInstance *string) {
	o.TargetInstance = targetInstance
}

// WriteToRequest writes these params to a swagger request
func (o *ResetSegmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ErrorSegmentsOnly != nil {

		// query param errorSegmentsOnly
		var qrErrorSegmentsOnly bool

		if o.ErrorSegmentsOnly != nil {
			qrErrorSegmentsOnly = *o.ErrorSegmentsOnly
		}
		qErrorSegmentsOnly := swag.FormatBool(qrErrorSegmentsOnly)
		if qErrorSegmentsOnly != "" {

			if err := r.SetQueryParam("errorSegmentsOnly", qErrorSegmentsOnly); err != nil {
				return err
			}
		}
	}

	// path param tableNameWithType
	if err := r.SetPathParam("tableNameWithType", o.TableNameWithType); err != nil {
		return err
	}

	if o.TargetInstance != nil {

		// query param targetInstance
		var qrTargetInstance string

		if o.TargetInstance != nil {
			qrTargetInstance = *o.TargetInstance
		}
		qTargetInstance := qrTargetInstance
		if qTargetInstance != "" {

			if err := r.SetQueryParam("targetInstance", qTargetInstance); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
