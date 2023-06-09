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
	"github.com/go-openapi/swag"
)

// NewPutChildrenParams creates a new PutChildrenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutChildrenParams() *PutChildrenParams {
	return &PutChildrenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutChildrenParamsWithTimeout creates a new PutChildrenParams object
// with the ability to set a timeout on a request.
func NewPutChildrenParamsWithTimeout(timeout time.Duration) *PutChildrenParams {
	return &PutChildrenParams{
		timeout: timeout,
	}
}

// NewPutChildrenParamsWithContext creates a new PutChildrenParams object
// with the ability to set a context for a request.
func NewPutChildrenParamsWithContext(ctx context.Context) *PutChildrenParams {
	return &PutChildrenParams{
		Context: ctx,
	}
}

// NewPutChildrenParamsWithHTTPClient creates a new PutChildrenParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutChildrenParamsWithHTTPClient(client *http.Client) *PutChildrenParams {
	return &PutChildrenParams{
		HTTPClient: client,
	}
}

/*
PutChildrenParams contains all the parameters to send to the API endpoint

	for the put children operation.

	Typically these are written to a http.Request.
*/
type PutChildrenParams struct {

	/* AccessOption.

	   accessOption

	   Format: int32
	   Default: 1
	*/
	AccessOption *int32

	// Body.
	Body string

	/* Data.

	   Content
	*/
	Data *string

	/* ExpectedVersion.

	   expectedVersion

	   Format: int32
	   Default: -1
	*/
	ExpectedVersion *int32

	/* Path.

	   Zookeeper path of parent, must start with /
	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put children params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutChildrenParams) WithDefaults() *PutChildrenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put children params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutChildrenParams) SetDefaults() {
	var (
		accessOptionDefault = int32(1)

		expectedVersionDefault = int32(-1)
	)

	val := PutChildrenParams{
		AccessOption:    &accessOptionDefault,
		ExpectedVersion: &expectedVersionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the put children params
func (o *PutChildrenParams) WithTimeout(timeout time.Duration) *PutChildrenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put children params
func (o *PutChildrenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put children params
func (o *PutChildrenParams) WithContext(ctx context.Context) *PutChildrenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put children params
func (o *PutChildrenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put children params
func (o *PutChildrenParams) WithHTTPClient(client *http.Client) *PutChildrenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put children params
func (o *PutChildrenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessOption adds the accessOption to the put children params
func (o *PutChildrenParams) WithAccessOption(accessOption *int32) *PutChildrenParams {
	o.SetAccessOption(accessOption)
	return o
}

// SetAccessOption adds the accessOption to the put children params
func (o *PutChildrenParams) SetAccessOption(accessOption *int32) {
	o.AccessOption = accessOption
}

// WithBody adds the body to the put children params
func (o *PutChildrenParams) WithBody(body string) *PutChildrenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put children params
func (o *PutChildrenParams) SetBody(body string) {
	o.Body = body
}

// WithData adds the data to the put children params
func (o *PutChildrenParams) WithData(data *string) *PutChildrenParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the put children params
func (o *PutChildrenParams) SetData(data *string) {
	o.Data = data
}

// WithExpectedVersion adds the expectedVersion to the put children params
func (o *PutChildrenParams) WithExpectedVersion(expectedVersion *int32) *PutChildrenParams {
	o.SetExpectedVersion(expectedVersion)
	return o
}

// SetExpectedVersion adds the expectedVersion to the put children params
func (o *PutChildrenParams) SetExpectedVersion(expectedVersion *int32) {
	o.ExpectedVersion = expectedVersion
}

// WithPath adds the path to the put children params
func (o *PutChildrenParams) WithPath(path string) *PutChildrenParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the put children params
func (o *PutChildrenParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *PutChildrenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccessOption != nil {

		// query param accessOption
		var qrAccessOption int32

		if o.AccessOption != nil {
			qrAccessOption = *o.AccessOption
		}
		qAccessOption := swag.FormatInt32(qrAccessOption)
		if qAccessOption != "" {

			if err := r.SetQueryParam("accessOption", qAccessOption); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if o.Data != nil {

		// query param data
		var qrData string

		if o.Data != nil {
			qrData = *o.Data
		}
		qData := qrData
		if qData != "" {

			if err := r.SetQueryParam("data", qData); err != nil {
				return err
			}
		}
	}

	if o.ExpectedVersion != nil {

		// query param expectedVersion
		var qrExpectedVersion int32

		if o.ExpectedVersion != nil {
			qrExpectedVersion = *o.ExpectedVersion
		}
		qExpectedVersion := swag.FormatInt32(qrExpectedVersion)
		if qExpectedVersion != "" {

			if err := r.SetQueryParam("expectedVersion", qExpectedVersion); err != nil {
				return err
			}
		}
	}

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
