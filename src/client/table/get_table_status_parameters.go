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
)

// NewGetTableStatusParams creates a new GetTableStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTableStatusParams() *GetTableStatusParams {
	return &GetTableStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTableStatusParamsWithTimeout creates a new GetTableStatusParams object
// with the ability to set a timeout on a request.
func NewGetTableStatusParamsWithTimeout(timeout time.Duration) *GetTableStatusParams {
	return &GetTableStatusParams{
		timeout: timeout,
	}
}

// NewGetTableStatusParamsWithContext creates a new GetTableStatusParams object
// with the ability to set a context for a request.
func NewGetTableStatusParamsWithContext(ctx context.Context) *GetTableStatusParams {
	return &GetTableStatusParams{
		Context: ctx,
	}
}

// NewGetTableStatusParamsWithHTTPClient creates a new GetTableStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTableStatusParamsWithHTTPClient(client *http.Client) *GetTableStatusParams {
	return &GetTableStatusParams{
		HTTPClient: client,
	}
}

/*
GetTableStatusParams contains all the parameters to send to the API endpoint

	for the get table status operation.

	Typically these are written to a http.Request.
*/
type GetTableStatusParams struct {

	/* TableName.

	   Name of the table
	*/
	TableName string

	/* Type.

	   realtime|offline
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get table status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTableStatusParams) WithDefaults() *GetTableStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get table status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTableStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get table status params
func (o *GetTableStatusParams) WithTimeout(timeout time.Duration) *GetTableStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get table status params
func (o *GetTableStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get table status params
func (o *GetTableStatusParams) WithContext(ctx context.Context) *GetTableStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get table status params
func (o *GetTableStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get table status params
func (o *GetTableStatusParams) WithHTTPClient(client *http.Client) *GetTableStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get table status params
func (o *GetTableStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTableName adds the tableName to the get table status params
func (o *GetTableStatusParams) WithTableName(tableName string) *GetTableStatusParams {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the get table status params
func (o *GetTableStatusParams) SetTableName(tableName string) {
	o.TableName = tableName
}

// WithType adds the typeVar to the get table status params
func (o *GetTableStatusParams) WithType(typeVar *string) *GetTableStatusParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get table status params
func (o *GetTableStatusParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetTableStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
