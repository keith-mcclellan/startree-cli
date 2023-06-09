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

// NewGetSegmentMetadataDeprecated2Params creates a new GetSegmentMetadataDeprecated2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSegmentMetadataDeprecated2Params() *GetSegmentMetadataDeprecated2Params {
	return &GetSegmentMetadataDeprecated2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSegmentMetadataDeprecated2ParamsWithTimeout creates a new GetSegmentMetadataDeprecated2Params object
// with the ability to set a timeout on a request.
func NewGetSegmentMetadataDeprecated2ParamsWithTimeout(timeout time.Duration) *GetSegmentMetadataDeprecated2Params {
	return &GetSegmentMetadataDeprecated2Params{
		timeout: timeout,
	}
}

// NewGetSegmentMetadataDeprecated2ParamsWithContext creates a new GetSegmentMetadataDeprecated2Params object
// with the ability to set a context for a request.
func NewGetSegmentMetadataDeprecated2ParamsWithContext(ctx context.Context) *GetSegmentMetadataDeprecated2Params {
	return &GetSegmentMetadataDeprecated2Params{
		Context: ctx,
	}
}

// NewGetSegmentMetadataDeprecated2ParamsWithHTTPClient creates a new GetSegmentMetadataDeprecated2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetSegmentMetadataDeprecated2ParamsWithHTTPClient(client *http.Client) *GetSegmentMetadataDeprecated2Params {
	return &GetSegmentMetadataDeprecated2Params{
		HTTPClient: client,
	}
}

/*
GetSegmentMetadataDeprecated2Params contains all the parameters to send to the API endpoint

	for the get segment metadata deprecated2 operation.

	Typically these are written to a http.Request.
*/
type GetSegmentMetadataDeprecated2Params struct {

	/* SegmentName.

	   Name of the segment
	*/
	SegmentName string

	/* State.

	   MUST be null
	*/
	State *string

	/* TableName.

	   Name of the table
	*/
	TableName string

	/* Type.

	   OFFLINE|REALTIME
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get segment metadata deprecated2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSegmentMetadataDeprecated2Params) WithDefaults() *GetSegmentMetadataDeprecated2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get segment metadata deprecated2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSegmentMetadataDeprecated2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get segment metadata deprecated2 params
func (o *GetSegmentMetadataDeprecated2Params) WithTimeout(timeout time.Duration) *GetSegmentMetadataDeprecated2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get segment metadata deprecated2 params
func (o *GetSegmentMetadataDeprecated2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get segment metadata deprecated2 params
func (o *GetSegmentMetadataDeprecated2Params) WithContext(ctx context.Context) *GetSegmentMetadataDeprecated2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get segment metadata deprecated2 params
func (o *GetSegmentMetadataDeprecated2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get segment metadata deprecated2 params
func (o *GetSegmentMetadataDeprecated2Params) WithHTTPClient(client *http.Client) *GetSegmentMetadataDeprecated2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get segment metadata deprecated2 params
func (o *GetSegmentMetadataDeprecated2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSegmentName adds the segmentName to the get segment metadata deprecated2 params
func (o *GetSegmentMetadataDeprecated2Params) WithSegmentName(segmentName string) *GetSegmentMetadataDeprecated2Params {
	o.SetSegmentName(segmentName)
	return o
}

// SetSegmentName adds the segmentName to the get segment metadata deprecated2 params
func (o *GetSegmentMetadataDeprecated2Params) SetSegmentName(segmentName string) {
	o.SegmentName = segmentName
}

// WithState adds the state to the get segment metadata deprecated2 params
func (o *GetSegmentMetadataDeprecated2Params) WithState(state *string) *GetSegmentMetadataDeprecated2Params {
	o.SetState(state)
	return o
}

// SetState adds the state to the get segment metadata deprecated2 params
func (o *GetSegmentMetadataDeprecated2Params) SetState(state *string) {
	o.State = state
}

// WithTableName adds the tableName to the get segment metadata deprecated2 params
func (o *GetSegmentMetadataDeprecated2Params) WithTableName(tableName string) *GetSegmentMetadataDeprecated2Params {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the get segment metadata deprecated2 params
func (o *GetSegmentMetadataDeprecated2Params) SetTableName(tableName string) {
	o.TableName = tableName
}

// WithType adds the typeVar to the get segment metadata deprecated2 params
func (o *GetSegmentMetadataDeprecated2Params) WithType(typeVar *string) *GetSegmentMetadataDeprecated2Params {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get segment metadata deprecated2 params
func (o *GetSegmentMetadataDeprecated2Params) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetSegmentMetadataDeprecated2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param segmentName
	if err := r.SetPathParam("segmentName", o.SegmentName); err != nil {
		return err
	}

	if o.State != nil {

		// query param state
		var qrState string

		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}
	}

	// path param tableName
	if err := r.SetPathParam("tableName", o.TableName); err != nil {
		return err
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
