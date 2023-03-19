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

// NewGetTableTiersParams creates a new GetTableTiersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTableTiersParams() *GetTableTiersParams {
	return &GetTableTiersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTableTiersParamsWithTimeout creates a new GetTableTiersParams object
// with the ability to set a timeout on a request.
func NewGetTableTiersParamsWithTimeout(timeout time.Duration) *GetTableTiersParams {
	return &GetTableTiersParams{
		timeout: timeout,
	}
}

// NewGetTableTiersParamsWithContext creates a new GetTableTiersParams object
// with the ability to set a context for a request.
func NewGetTableTiersParamsWithContext(ctx context.Context) *GetTableTiersParams {
	return &GetTableTiersParams{
		Context: ctx,
	}
}

// NewGetTableTiersParamsWithHTTPClient creates a new GetTableTiersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTableTiersParamsWithHTTPClient(client *http.Client) *GetTableTiersParams {
	return &GetTableTiersParams{
		HTTPClient: client,
	}
}

/*
GetTableTiersParams contains all the parameters to send to the API endpoint

	for the get table tiers operation.

	Typically these are written to a http.Request.
*/
type GetTableTiersParams struct {

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

// WithDefaults hydrates default values in the get table tiers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTableTiersParams) WithDefaults() *GetTableTiersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get table tiers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTableTiersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get table tiers params
func (o *GetTableTiersParams) WithTimeout(timeout time.Duration) *GetTableTiersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get table tiers params
func (o *GetTableTiersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get table tiers params
func (o *GetTableTiersParams) WithContext(ctx context.Context) *GetTableTiersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get table tiers params
func (o *GetTableTiersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get table tiers params
func (o *GetTableTiersParams) WithHTTPClient(client *http.Client) *GetTableTiersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get table tiers params
func (o *GetTableTiersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTableName adds the tableName to the get table tiers params
func (o *GetTableTiersParams) WithTableName(tableName string) *GetTableTiersParams {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the get table tiers params
func (o *GetTableTiersParams) SetTableName(tableName string) {
	o.TableName = tableName
}

// WithType adds the typeVar to the get table tiers params
func (o *GetTableTiersParams) WithType(typeVar string) *GetTableTiersParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get table tiers params
func (o *GetTableTiersParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetTableTiersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
