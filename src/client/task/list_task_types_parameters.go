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

// NewListTaskTypesParams creates a new ListTaskTypesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListTaskTypesParams() *ListTaskTypesParams {
	return &ListTaskTypesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListTaskTypesParamsWithTimeout creates a new ListTaskTypesParams object
// with the ability to set a timeout on a request.
func NewListTaskTypesParamsWithTimeout(timeout time.Duration) *ListTaskTypesParams {
	return &ListTaskTypesParams{
		timeout: timeout,
	}
}

// NewListTaskTypesParamsWithContext creates a new ListTaskTypesParams object
// with the ability to set a context for a request.
func NewListTaskTypesParamsWithContext(ctx context.Context) *ListTaskTypesParams {
	return &ListTaskTypesParams{
		Context: ctx,
	}
}

// NewListTaskTypesParamsWithHTTPClient creates a new ListTaskTypesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListTaskTypesParamsWithHTTPClient(client *http.Client) *ListTaskTypesParams {
	return &ListTaskTypesParams{
		HTTPClient: client,
	}
}

/*
ListTaskTypesParams contains all the parameters to send to the API endpoint

	for the list task types operation.

	Typically these are written to a http.Request.
*/
type ListTaskTypesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list task types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTaskTypesParams) WithDefaults() *ListTaskTypesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list task types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTaskTypesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list task types params
func (o *ListTaskTypesParams) WithTimeout(timeout time.Duration) *ListTaskTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list task types params
func (o *ListTaskTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list task types params
func (o *ListTaskTypesParams) WithContext(ctx context.Context) *ListTaskTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list task types params
func (o *ListTaskTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list task types params
func (o *ListTaskTypesParams) WithHTTPClient(client *http.Client) *ListTaskTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list task types params
func (o *ListTaskTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListTaskTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
