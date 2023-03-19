// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewDeleteClusterConfigParams creates a new DeleteClusterConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteClusterConfigParams() *DeleteClusterConfigParams {
	return &DeleteClusterConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClusterConfigParamsWithTimeout creates a new DeleteClusterConfigParams object
// with the ability to set a timeout on a request.
func NewDeleteClusterConfigParamsWithTimeout(timeout time.Duration) *DeleteClusterConfigParams {
	return &DeleteClusterConfigParams{
		timeout: timeout,
	}
}

// NewDeleteClusterConfigParamsWithContext creates a new DeleteClusterConfigParams object
// with the ability to set a context for a request.
func NewDeleteClusterConfigParamsWithContext(ctx context.Context) *DeleteClusterConfigParams {
	return &DeleteClusterConfigParams{
		Context: ctx,
	}
}

// NewDeleteClusterConfigParamsWithHTTPClient creates a new DeleteClusterConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteClusterConfigParamsWithHTTPClient(client *http.Client) *DeleteClusterConfigParams {
	return &DeleteClusterConfigParams{
		HTTPClient: client,
	}
}

/*
DeleteClusterConfigParams contains all the parameters to send to the API endpoint

	for the delete cluster config operation.

	Typically these are written to a http.Request.
*/
type DeleteClusterConfigParams struct {

	/* ConfigName.

	   Name of the config to delete
	*/
	ConfigName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete cluster config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteClusterConfigParams) WithDefaults() *DeleteClusterConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete cluster config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteClusterConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete cluster config params
func (o *DeleteClusterConfigParams) WithTimeout(timeout time.Duration) *DeleteClusterConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cluster config params
func (o *DeleteClusterConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cluster config params
func (o *DeleteClusterConfigParams) WithContext(ctx context.Context) *DeleteClusterConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cluster config params
func (o *DeleteClusterConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cluster config params
func (o *DeleteClusterConfigParams) WithHTTPClient(client *http.Client) *DeleteClusterConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cluster config params
func (o *DeleteClusterConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigName adds the configName to the delete cluster config params
func (o *DeleteClusterConfigParams) WithConfigName(configName string) *DeleteClusterConfigParams {
	o.SetConfigName(configName)
	return o
}

// SetConfigName adds the configName to the delete cluster config params
func (o *DeleteClusterConfigParams) SetConfigName(configName string) {
	o.ConfigName = configName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClusterConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param configName
	if err := r.SetPathParam("configName", o.ConfigName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}