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

// NewRecommendConfigParams creates a new RecommendConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRecommendConfigParams() *RecommendConfigParams {
	return &RecommendConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRecommendConfigParamsWithTimeout creates a new RecommendConfigParams object
// with the ability to set a timeout on a request.
func NewRecommendConfigParamsWithTimeout(timeout time.Duration) *RecommendConfigParams {
	return &RecommendConfigParams{
		timeout: timeout,
	}
}

// NewRecommendConfigParamsWithContext creates a new RecommendConfigParams object
// with the ability to set a context for a request.
func NewRecommendConfigParamsWithContext(ctx context.Context) *RecommendConfigParams {
	return &RecommendConfigParams{
		Context: ctx,
	}
}

// NewRecommendConfigParamsWithHTTPClient creates a new RecommendConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewRecommendConfigParamsWithHTTPClient(client *http.Client) *RecommendConfigParams {
	return &RecommendConfigParams{
		HTTPClient: client,
	}
}

/*
RecommendConfigParams contains all the parameters to send to the API endpoint

	for the recommend config operation.

	Typically these are written to a http.Request.
*/
type RecommendConfigParams struct {

	// Body.
	Body string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the recommend config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RecommendConfigParams) WithDefaults() *RecommendConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the recommend config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RecommendConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the recommend config params
func (o *RecommendConfigParams) WithTimeout(timeout time.Duration) *RecommendConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the recommend config params
func (o *RecommendConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the recommend config params
func (o *RecommendConfigParams) WithContext(ctx context.Context) *RecommendConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the recommend config params
func (o *RecommendConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the recommend config params
func (o *RecommendConfigParams) WithHTTPClient(client *http.Client) *RecommendConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the recommend config params
func (o *RecommendConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the recommend config params
func (o *RecommendConfigParams) WithBody(body string) *RecommendConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the recommend config params
func (o *RecommendConfigParams) SetBody(body string) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RecommendConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}