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

	"startree.ai/cli/models"
)

// NewValidateTableAndSchemaParams creates a new ValidateTableAndSchemaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateTableAndSchemaParams() *ValidateTableAndSchemaParams {
	return &ValidateTableAndSchemaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateTableAndSchemaParamsWithTimeout creates a new ValidateTableAndSchemaParams object
// with the ability to set a timeout on a request.
func NewValidateTableAndSchemaParamsWithTimeout(timeout time.Duration) *ValidateTableAndSchemaParams {
	return &ValidateTableAndSchemaParams{
		timeout: timeout,
	}
}

// NewValidateTableAndSchemaParamsWithContext creates a new ValidateTableAndSchemaParams object
// with the ability to set a context for a request.
func NewValidateTableAndSchemaParamsWithContext(ctx context.Context) *ValidateTableAndSchemaParams {
	return &ValidateTableAndSchemaParams{
		Context: ctx,
	}
}

// NewValidateTableAndSchemaParamsWithHTTPClient creates a new ValidateTableAndSchemaParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateTableAndSchemaParamsWithHTTPClient(client *http.Client) *ValidateTableAndSchemaParams {
	return &ValidateTableAndSchemaParams{
		HTTPClient: client,
	}
}

/*
ValidateTableAndSchemaParams contains all the parameters to send to the API endpoint

	for the validate table and schema operation.

	Typically these are written to a http.Request.
*/
type ValidateTableAndSchemaParams struct {

	// Body.
	Body *models.TableAndSchemaConfig

	/* ValidationTypesToSkip.

	   comma separated list of validation type(s) to skip. supported types: (ALL|TASK|UPSERT)
	*/
	ValidationTypesToSkip *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate table and schema params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateTableAndSchemaParams) WithDefaults() *ValidateTableAndSchemaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate table and schema params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateTableAndSchemaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate table and schema params
func (o *ValidateTableAndSchemaParams) WithTimeout(timeout time.Duration) *ValidateTableAndSchemaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate table and schema params
func (o *ValidateTableAndSchemaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate table and schema params
func (o *ValidateTableAndSchemaParams) WithContext(ctx context.Context) *ValidateTableAndSchemaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate table and schema params
func (o *ValidateTableAndSchemaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate table and schema params
func (o *ValidateTableAndSchemaParams) WithHTTPClient(client *http.Client) *ValidateTableAndSchemaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate table and schema params
func (o *ValidateTableAndSchemaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate table and schema params
func (o *ValidateTableAndSchemaParams) WithBody(body *models.TableAndSchemaConfig) *ValidateTableAndSchemaParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate table and schema params
func (o *ValidateTableAndSchemaParams) SetBody(body *models.TableAndSchemaConfig) {
	o.Body = body
}

// WithValidationTypesToSkip adds the validationTypesToSkip to the validate table and schema params
func (o *ValidateTableAndSchemaParams) WithValidationTypesToSkip(validationTypesToSkip *string) *ValidateTableAndSchemaParams {
	o.SetValidationTypesToSkip(validationTypesToSkip)
	return o
}

// SetValidationTypesToSkip adds the validationTypesToSkip to the validate table and schema params
func (o *ValidateTableAndSchemaParams) SetValidationTypesToSkip(validationTypesToSkip *string) {
	o.ValidationTypesToSkip = validationTypesToSkip
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateTableAndSchemaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.ValidationTypesToSkip != nil {

		// query param validationTypesToSkip
		var qrValidationTypesToSkip string

		if o.ValidationTypesToSkip != nil {
			qrValidationTypesToSkip = *o.ValidationTypesToSkip
		}
		qValidationTypesToSkip := qrValidationTypesToSkip
		if qValidationTypesToSkip != "" {

			if err := r.SetQueryParam("validationTypesToSkip", qValidationTypesToSkip); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
