// Code generated by go-swagger; DO NOT EDIT.

package broker

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

// NewGetBrokersForTableParams creates a new GetBrokersForTableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBrokersForTableParams() *GetBrokersForTableParams {
	return &GetBrokersForTableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBrokersForTableParamsWithTimeout creates a new GetBrokersForTableParams object
// with the ability to set a timeout on a request.
func NewGetBrokersForTableParamsWithTimeout(timeout time.Duration) *GetBrokersForTableParams {
	return &GetBrokersForTableParams{
		timeout: timeout,
	}
}

// NewGetBrokersForTableParamsWithContext creates a new GetBrokersForTableParams object
// with the ability to set a context for a request.
func NewGetBrokersForTableParamsWithContext(ctx context.Context) *GetBrokersForTableParams {
	return &GetBrokersForTableParams{
		Context: ctx,
	}
}

// NewGetBrokersForTableParamsWithHTTPClient creates a new GetBrokersForTableParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBrokersForTableParamsWithHTTPClient(client *http.Client) *GetBrokersForTableParams {
	return &GetBrokersForTableParams{
		HTTPClient: client,
	}
}

/*
GetBrokersForTableParams contains all the parameters to send to the API endpoint

	for the get brokers for table operation.

	Typically these are written to a http.Request.
*/
type GetBrokersForTableParams struct {

	/* State.

	   ONLINE|OFFLINE
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

// WithDefaults hydrates default values in the get brokers for table params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBrokersForTableParams) WithDefaults() *GetBrokersForTableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get brokers for table params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBrokersForTableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get brokers for table params
func (o *GetBrokersForTableParams) WithTimeout(timeout time.Duration) *GetBrokersForTableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get brokers for table params
func (o *GetBrokersForTableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get brokers for table params
func (o *GetBrokersForTableParams) WithContext(ctx context.Context) *GetBrokersForTableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get brokers for table params
func (o *GetBrokersForTableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get brokers for table params
func (o *GetBrokersForTableParams) WithHTTPClient(client *http.Client) *GetBrokersForTableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get brokers for table params
func (o *GetBrokersForTableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithState adds the state to the get brokers for table params
func (o *GetBrokersForTableParams) WithState(state *string) *GetBrokersForTableParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the get brokers for table params
func (o *GetBrokersForTableParams) SetState(state *string) {
	o.State = state
}

// WithTableName adds the tableName to the get brokers for table params
func (o *GetBrokersForTableParams) WithTableName(tableName string) *GetBrokersForTableParams {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the get brokers for table params
func (o *GetBrokersForTableParams) SetTableName(tableName string) {
	o.TableName = tableName
}

// WithType adds the typeVar to the get brokers for table params
func (o *GetBrokersForTableParams) WithType(typeVar *string) *GetBrokersForTableParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get brokers for table params
func (o *GetBrokersForTableParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetBrokersForTableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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