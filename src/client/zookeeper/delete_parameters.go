// Code generated by go-swagger; DO NOT EDIT.

package zookeeper

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

// NewDeleteParams creates a new DeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteParams() *DeleteParams {
	return &DeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteParamsWithTimeout creates a new DeleteParams object
// with the ability to set a timeout on a request.
func NewDeleteParamsWithTimeout(timeout time.Duration) *DeleteParams {
	return &DeleteParams{
		timeout: timeout,
	}
}

// NewDeleteParamsWithContext creates a new DeleteParams object
// with the ability to set a context for a request.
func NewDeleteParamsWithContext(ctx context.Context) *DeleteParams {
	return &DeleteParams{
		Context: ctx,
	}
}

// NewDeleteParamsWithHTTPClient creates a new DeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteParamsWithHTTPClient(client *http.Client) *DeleteParams {
	return &DeleteParams{
		HTTPClient: client,
	}
}

/*
DeleteParams contains all the parameters to send to the API endpoint

	for the delete operation.

	Typically these are written to a http.Request.
*/
type DeleteParams struct {

	/* Path.

	   Zookeeper Path, must start with /
	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteParams) WithDefaults() *DeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete params
func (o *DeleteParams) WithTimeout(timeout time.Duration) *DeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete params
func (o *DeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete params
func (o *DeleteParams) WithContext(ctx context.Context) *DeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete params
func (o *DeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete params
func (o *DeleteParams) WithHTTPClient(client *http.Client) *DeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete params
func (o *DeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPath adds the path to the delete params
func (o *DeleteParams) WithPath(path string) *DeleteParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the delete params
func (o *DeleteParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param path
	qrPath := o.Path
	qPath := qrPath
	if qPath != "" {

		if err := r.SetQueryParam("path", qPath); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
