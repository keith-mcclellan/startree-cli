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

// NewToggleQueryRateLimitingParams creates a new ToggleQueryRateLimitingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewToggleQueryRateLimitingParams() *ToggleQueryRateLimitingParams {
	return &ToggleQueryRateLimitingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewToggleQueryRateLimitingParamsWithTimeout creates a new ToggleQueryRateLimitingParams object
// with the ability to set a timeout on a request.
func NewToggleQueryRateLimitingParamsWithTimeout(timeout time.Duration) *ToggleQueryRateLimitingParams {
	return &ToggleQueryRateLimitingParams{
		timeout: timeout,
	}
}

// NewToggleQueryRateLimitingParamsWithContext creates a new ToggleQueryRateLimitingParams object
// with the ability to set a context for a request.
func NewToggleQueryRateLimitingParamsWithContext(ctx context.Context) *ToggleQueryRateLimitingParams {
	return &ToggleQueryRateLimitingParams{
		Context: ctx,
	}
}

// NewToggleQueryRateLimitingParamsWithHTTPClient creates a new ToggleQueryRateLimitingParams object
// with the ability to set a custom HTTPClient for a request.
func NewToggleQueryRateLimitingParamsWithHTTPClient(client *http.Client) *ToggleQueryRateLimitingParams {
	return &ToggleQueryRateLimitingParams{
		HTTPClient: client,
	}
}

/*
ToggleQueryRateLimitingParams contains all the parameters to send to the API endpoint

	for the toggle query rate limiting operation.

	Typically these are written to a http.Request.
*/
type ToggleQueryRateLimitingParams struct {

	/* InstanceName.

	   Broker instance name
	*/
	InstanceName string

	/* State.

	   ENABLE|DISABLE
	*/
	State string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the toggle query rate limiting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ToggleQueryRateLimitingParams) WithDefaults() *ToggleQueryRateLimitingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the toggle query rate limiting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ToggleQueryRateLimitingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the toggle query rate limiting params
func (o *ToggleQueryRateLimitingParams) WithTimeout(timeout time.Duration) *ToggleQueryRateLimitingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the toggle query rate limiting params
func (o *ToggleQueryRateLimitingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the toggle query rate limiting params
func (o *ToggleQueryRateLimitingParams) WithContext(ctx context.Context) *ToggleQueryRateLimitingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the toggle query rate limiting params
func (o *ToggleQueryRateLimitingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the toggle query rate limiting params
func (o *ToggleQueryRateLimitingParams) WithHTTPClient(client *http.Client) *ToggleQueryRateLimitingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the toggle query rate limiting params
func (o *ToggleQueryRateLimitingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstanceName adds the instanceName to the toggle query rate limiting params
func (o *ToggleQueryRateLimitingParams) WithInstanceName(instanceName string) *ToggleQueryRateLimitingParams {
	o.SetInstanceName(instanceName)
	return o
}

// SetInstanceName adds the instanceName to the toggle query rate limiting params
func (o *ToggleQueryRateLimitingParams) SetInstanceName(instanceName string) {
	o.InstanceName = instanceName
}

// WithState adds the state to the toggle query rate limiting params
func (o *ToggleQueryRateLimitingParams) WithState(state string) *ToggleQueryRateLimitingParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the toggle query rate limiting params
func (o *ToggleQueryRateLimitingParams) SetState(state string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *ToggleQueryRateLimitingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param instanceName
	if err := r.SetPathParam("instanceName", o.InstanceName); err != nil {
		return err
	}

	// query param state
	qrState := o.State
	qState := qrState
	if qState != "" {

		if err := r.SetQueryParam("state", qState); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
