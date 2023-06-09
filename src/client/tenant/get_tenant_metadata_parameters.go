// Code generated by go-swagger; DO NOT EDIT.

package tenant

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

// NewGetTenantMetadataParams creates a new GetTenantMetadataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTenantMetadataParams() *GetTenantMetadataParams {
	return &GetTenantMetadataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTenantMetadataParamsWithTimeout creates a new GetTenantMetadataParams object
// with the ability to set a timeout on a request.
func NewGetTenantMetadataParamsWithTimeout(timeout time.Duration) *GetTenantMetadataParams {
	return &GetTenantMetadataParams{
		timeout: timeout,
	}
}

// NewGetTenantMetadataParamsWithContext creates a new GetTenantMetadataParams object
// with the ability to set a context for a request.
func NewGetTenantMetadataParamsWithContext(ctx context.Context) *GetTenantMetadataParams {
	return &GetTenantMetadataParams{
		Context: ctx,
	}
}

// NewGetTenantMetadataParamsWithHTTPClient creates a new GetTenantMetadataParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTenantMetadataParamsWithHTTPClient(client *http.Client) *GetTenantMetadataParams {
	return &GetTenantMetadataParams{
		HTTPClient: client,
	}
}

/*
GetTenantMetadataParams contains all the parameters to send to the API endpoint

	for the get tenant metadata operation.

	Typically these are written to a http.Request.
*/
type GetTenantMetadataParams struct {

	/* TenantName.

	   Tenant name
	*/
	TenantName string

	/* Type.

	   tenant type
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get tenant metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTenantMetadataParams) WithDefaults() *GetTenantMetadataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get tenant metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTenantMetadataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get tenant metadata params
func (o *GetTenantMetadataParams) WithTimeout(timeout time.Duration) *GetTenantMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tenant metadata params
func (o *GetTenantMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tenant metadata params
func (o *GetTenantMetadataParams) WithContext(ctx context.Context) *GetTenantMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tenant metadata params
func (o *GetTenantMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tenant metadata params
func (o *GetTenantMetadataParams) WithHTTPClient(client *http.Client) *GetTenantMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tenant metadata params
func (o *GetTenantMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenantName adds the tenantName to the get tenant metadata params
func (o *GetTenantMetadataParams) WithTenantName(tenantName string) *GetTenantMetadataParams {
	o.SetTenantName(tenantName)
	return o
}

// SetTenantName adds the tenantName to the get tenant metadata params
func (o *GetTenantMetadataParams) SetTenantName(tenantName string) {
	o.TenantName = tenantName
}

// WithType adds the typeVar to the get tenant metadata params
func (o *GetTenantMetadataParams) WithType(typeVar *string) *GetTenantMetadataParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get tenant metadata params
func (o *GetTenantMetadataParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetTenantMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tenantName
	if err := r.SetPathParam("tenantName", o.TenantName); err != nil {
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
