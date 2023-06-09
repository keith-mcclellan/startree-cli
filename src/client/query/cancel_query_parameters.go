// Code generated by go-swagger; DO NOT EDIT.

package query

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
	"github.com/go-openapi/swag"
)

// NewCancelQueryParams creates a new CancelQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCancelQueryParams() *CancelQueryParams {
	return &CancelQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCancelQueryParamsWithTimeout creates a new CancelQueryParams object
// with the ability to set a timeout on a request.
func NewCancelQueryParamsWithTimeout(timeout time.Duration) *CancelQueryParams {
	return &CancelQueryParams{
		timeout: timeout,
	}
}

// NewCancelQueryParamsWithContext creates a new CancelQueryParams object
// with the ability to set a context for a request.
func NewCancelQueryParamsWithContext(ctx context.Context) *CancelQueryParams {
	return &CancelQueryParams{
		Context: ctx,
	}
}

// NewCancelQueryParamsWithHTTPClient creates a new CancelQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewCancelQueryParamsWithHTTPClient(client *http.Client) *CancelQueryParams {
	return &CancelQueryParams{
		HTTPClient: client,
	}
}

/*
CancelQueryParams contains all the parameters to send to the API endpoint

	for the cancel query operation.

	Typically these are written to a http.Request.
*/
type CancelQueryParams struct {

	/* BrokerID.

	   Broker that's running the query
	*/
	BrokerID string

	/* QueryID.

	   QueryId as assigned by the broker

	   Format: int64
	*/
	QueryID int64

	/* TimeoutMs.

	   Timeout for servers to respond the cancel request

	   Format: int32
	   Default: 3000
	*/
	TimeoutMs *int32

	/* Verbose.

	   Return verbose responses for troubleshooting
	*/
	Verbose *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cancel query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelQueryParams) WithDefaults() *CancelQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cancel query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelQueryParams) SetDefaults() {
	var (
		timeoutMsDefault = int32(3000)

		verboseDefault = bool(false)
	)

	val := CancelQueryParams{
		TimeoutMs: &timeoutMsDefault,
		Verbose:   &verboseDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cancel query params
func (o *CancelQueryParams) WithTimeout(timeout time.Duration) *CancelQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel query params
func (o *CancelQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel query params
func (o *CancelQueryParams) WithContext(ctx context.Context) *CancelQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel query params
func (o *CancelQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel query params
func (o *CancelQueryParams) WithHTTPClient(client *http.Client) *CancelQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel query params
func (o *CancelQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBrokerID adds the brokerID to the cancel query params
func (o *CancelQueryParams) WithBrokerID(brokerID string) *CancelQueryParams {
	o.SetBrokerID(brokerID)
	return o
}

// SetBrokerID adds the brokerId to the cancel query params
func (o *CancelQueryParams) SetBrokerID(brokerID string) {
	o.BrokerID = brokerID
}

// WithQueryID adds the queryID to the cancel query params
func (o *CancelQueryParams) WithQueryID(queryID int64) *CancelQueryParams {
	o.SetQueryID(queryID)
	return o
}

// SetQueryID adds the queryId to the cancel query params
func (o *CancelQueryParams) SetQueryID(queryID int64) {
	o.QueryID = queryID
}

// WithTimeoutMs adds the timeoutMs to the cancel query params
func (o *CancelQueryParams) WithTimeoutMs(timeoutMs *int32) *CancelQueryParams {
	o.SetTimeoutMs(timeoutMs)
	return o
}

// SetTimeoutMs adds the timeoutMs to the cancel query params
func (o *CancelQueryParams) SetTimeoutMs(timeoutMs *int32) {
	o.TimeoutMs = timeoutMs
}

// WithVerbose adds the verbose to the cancel query params
func (o *CancelQueryParams) WithVerbose(verbose *bool) *CancelQueryParams {
	o.SetVerbose(verbose)
	return o
}

// SetVerbose adds the verbose to the cancel query params
func (o *CancelQueryParams) SetVerbose(verbose *bool) {
	o.Verbose = verbose
}

// WriteToRequest writes these params to a swagger request
func (o *CancelQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param brokerId
	if err := r.SetPathParam("brokerId", o.BrokerID); err != nil {
		return err
	}

	// path param queryId
	if err := r.SetPathParam("queryId", swag.FormatInt64(o.QueryID)); err != nil {
		return err
	}

	if o.TimeoutMs != nil {

		// query param timeoutMs
		var qrTimeoutMs int32

		if o.TimeoutMs != nil {
			qrTimeoutMs = *o.TimeoutMs
		}
		qTimeoutMs := swag.FormatInt32(qrTimeoutMs)
		if qTimeoutMs != "" {

			if err := r.SetQueryParam("timeoutMs", qTimeoutMs); err != nil {
				return err
			}
		}
	}

	if o.Verbose != nil {

		// query param verbose
		var qrVerbose bool

		if o.Verbose != nil {
			qrVerbose = *o.Verbose
		}
		qVerbose := swag.FormatBool(qrVerbose)
		if qVerbose != "" {

			if err := r.SetQueryParam("verbose", qVerbose); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
