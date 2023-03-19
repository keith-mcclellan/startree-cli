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

// NewIngestFromFileParams creates a new IngestFromFileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIngestFromFileParams() *IngestFromFileParams {
	return &IngestFromFileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIngestFromFileParamsWithTimeout creates a new IngestFromFileParams object
// with the ability to set a timeout on a request.
func NewIngestFromFileParamsWithTimeout(timeout time.Duration) *IngestFromFileParams {
	return &IngestFromFileParams{
		timeout: timeout,
	}
}

// NewIngestFromFileParamsWithContext creates a new IngestFromFileParams object
// with the ability to set a context for a request.
func NewIngestFromFileParamsWithContext(ctx context.Context) *IngestFromFileParams {
	return &IngestFromFileParams{
		Context: ctx,
	}
}

// NewIngestFromFileParamsWithHTTPClient creates a new IngestFromFileParams object
// with the ability to set a custom HTTPClient for a request.
func NewIngestFromFileParamsWithHTTPClient(client *http.Client) *IngestFromFileParams {
	return &IngestFromFileParams{
		HTTPClient: client,
	}
}

/*
IngestFromFileParams contains all the parameters to send to the API endpoint

	for the ingest from file operation.

	Typically these are written to a http.Request.
*/
type IngestFromFileParams struct {

	/* BatchConfigMapStr.

	   Batch config Map as json string. Must pass inputFormat, and optionally record reader properties. e.g. {"inputFormat":"json"}
	*/
	BatchConfigMapStr string

	// Body.
	Body *models.FormDataMultiPart

	/* TableNameWithType.

	   Name of the table to upload the file to
	*/
	TableNameWithType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ingest from file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IngestFromFileParams) WithDefaults() *IngestFromFileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ingest from file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IngestFromFileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ingest from file params
func (o *IngestFromFileParams) WithTimeout(timeout time.Duration) *IngestFromFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ingest from file params
func (o *IngestFromFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ingest from file params
func (o *IngestFromFileParams) WithContext(ctx context.Context) *IngestFromFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ingest from file params
func (o *IngestFromFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ingest from file params
func (o *IngestFromFileParams) WithHTTPClient(client *http.Client) *IngestFromFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ingest from file params
func (o *IngestFromFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBatchConfigMapStr adds the batchConfigMapStr to the ingest from file params
func (o *IngestFromFileParams) WithBatchConfigMapStr(batchConfigMapStr string) *IngestFromFileParams {
	o.SetBatchConfigMapStr(batchConfigMapStr)
	return o
}

// SetBatchConfigMapStr adds the batchConfigMapStr to the ingest from file params
func (o *IngestFromFileParams) SetBatchConfigMapStr(batchConfigMapStr string) {
	o.BatchConfigMapStr = batchConfigMapStr
}

// WithBody adds the body to the ingest from file params
func (o *IngestFromFileParams) WithBody(body *models.FormDataMultiPart) *IngestFromFileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the ingest from file params
func (o *IngestFromFileParams) SetBody(body *models.FormDataMultiPart) {
	o.Body = body
}

// WithTableNameWithType adds the tableNameWithType to the ingest from file params
func (o *IngestFromFileParams) WithTableNameWithType(tableNameWithType string) *IngestFromFileParams {
	o.SetTableNameWithType(tableNameWithType)
	return o
}

// SetTableNameWithType adds the tableNameWithType to the ingest from file params
func (o *IngestFromFileParams) SetTableNameWithType(tableNameWithType string) {
	o.TableNameWithType = tableNameWithType
}

// WriteToRequest writes these params to a swagger request
func (o *IngestFromFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param batchConfigMapStr
	qrBatchConfigMapStr := o.BatchConfigMapStr
	qBatchConfigMapStr := qrBatchConfigMapStr
	if qBatchConfigMapStr != "" {

		if err := r.SetQueryParam("batchConfigMapStr", qBatchConfigMapStr); err != nil {
			return err
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param tableNameWithType
	qrTableNameWithType := o.TableNameWithType
	qTableNameWithType := qrTableNameWithType
	if qTableNameWithType != "" {

		if err := r.SetQueryParam("tableNameWithType", qTableNameWithType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}