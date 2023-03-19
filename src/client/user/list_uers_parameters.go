// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewListUersParams creates a new ListUersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListUersParams() *ListUersParams {
	return &ListUersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListUersParamsWithTimeout creates a new ListUersParams object
// with the ability to set a timeout on a request.
func NewListUersParamsWithTimeout(timeout time.Duration) *ListUersParams {
	return &ListUersParams{
		timeout: timeout,
	}
}

// NewListUersParamsWithContext creates a new ListUersParams object
// with the ability to set a context for a request.
func NewListUersParamsWithContext(ctx context.Context) *ListUersParams {
	return &ListUersParams{
		Context: ctx,
	}
}

// NewListUersParamsWithHTTPClient creates a new ListUersParams object
// with the ability to set a custom HTTPClient for a request.
func NewListUersParamsWithHTTPClient(client *http.Client) *ListUersParams {
	return &ListUersParams{
		HTTPClient: client,
	}
}

/*
ListUersParams contains all the parameters to send to the API endpoint

	for the list uers operation.

	Typically these are written to a http.Request.
*/
type ListUersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list uers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListUersParams) WithDefaults() *ListUersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list uers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListUersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list uers params
func (o *ListUersParams) WithTimeout(timeout time.Duration) *ListUersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list uers params
func (o *ListUersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list uers params
func (o *ListUersParams) WithContext(ctx context.Context) *ListUersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list uers params
func (o *ListUersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list uers params
func (o *ListUersParams) WithHTTPClient(client *http.Client) *ListUersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list uers params
func (o *ListUersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListUersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
