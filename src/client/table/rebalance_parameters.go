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
	"github.com/go-openapi/swag"
)

// NewRebalanceParams creates a new RebalanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRebalanceParams() *RebalanceParams {
	return &RebalanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRebalanceParamsWithTimeout creates a new RebalanceParams object
// with the ability to set a timeout on a request.
func NewRebalanceParamsWithTimeout(timeout time.Duration) *RebalanceParams {
	return &RebalanceParams{
		timeout: timeout,
	}
}

// NewRebalanceParamsWithContext creates a new RebalanceParams object
// with the ability to set a context for a request.
func NewRebalanceParamsWithContext(ctx context.Context) *RebalanceParams {
	return &RebalanceParams{
		Context: ctx,
	}
}

// NewRebalanceParamsWithHTTPClient creates a new RebalanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewRebalanceParamsWithHTTPClient(client *http.Client) *RebalanceParams {
	return &RebalanceParams{
		HTTPClient: client,
	}
}

/*
RebalanceParams contains all the parameters to send to the API endpoint

	for the rebalance operation.

	Typically these are written to a http.Request.
*/
type RebalanceParams struct {

	/* BestEfforts.

	   Whether to use best-efforts to rebalance (not fail the rebalance when the no-downtime contract cannot be achieved)
	*/
	BestEfforts *bool

	/* Bootstrap.

	   Whether to rebalance table in bootstrap mode (regardless of minimum segment movement, reassign all segments in a round-robin fashion as if adding new segments to an empty table)
	*/
	Bootstrap *bool

	/* Downtime.

	   Whether to allow downtime for the rebalance
	*/
	Downtime *bool

	/* DryRun.

	   Whether to rebalance table in dry-run mode
	*/
	DryRun *bool

	/* ExternalViewCheckIntervalInMs.

	   How often to check if external view converges with ideal states

	   Format: int64
	   Default: 1000
	*/
	ExternalViewCheckIntervalInMs *int64

	/* ExternalViewStabilizationTimeoutInMs.

	   How long to wait till external view converges with ideal states

	   Format: int64
	   Default: 3600000
	*/
	ExternalViewStabilizationTimeoutInMs *int64

	/* IncludeConsuming.

	   Whether to reassign CONSUMING segments for real-time table
	*/
	IncludeConsuming *bool

	/* MinAvailableReplicas.

	   For no-downtime rebalance, minimum number of replicas to keep alive during rebalance, or maximum number of replicas allowed to be unavailable if value is negative

	   Format: int32
	   Default: 1
	*/
	MinAvailableReplicas *int32

	/* ReassignInstances.

	   Whether to reassign instances before reassigning segments
	*/
	ReassignInstances *bool

	/* TableName.

	   Name of the table to rebalance
	*/
	TableName string

	/* Type.

	   OFFLINE|REALTIME
	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rebalance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RebalanceParams) WithDefaults() *RebalanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rebalance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RebalanceParams) SetDefaults() {
	var (
		bestEffortsDefault = bool(false)

		bootstrapDefault = bool(false)

		downtimeDefault = bool(false)

		dryRunDefault = bool(false)

		externalViewCheckIntervalInMsDefault = int64(1000)

		externalViewStabilizationTimeoutInMsDefault = int64(3.6e+06)

		includeConsumingDefault = bool(false)

		minAvailableReplicasDefault = int32(1)

		reassignInstancesDefault = bool(false)
	)

	val := RebalanceParams{
		BestEfforts:                          &bestEffortsDefault,
		Bootstrap:                            &bootstrapDefault,
		Downtime:                             &downtimeDefault,
		DryRun:                               &dryRunDefault,
		ExternalViewCheckIntervalInMs:        &externalViewCheckIntervalInMsDefault,
		ExternalViewStabilizationTimeoutInMs: &externalViewStabilizationTimeoutInMsDefault,
		IncludeConsuming:                     &includeConsumingDefault,
		MinAvailableReplicas:                 &minAvailableReplicasDefault,
		ReassignInstances:                    &reassignInstancesDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the rebalance params
func (o *RebalanceParams) WithTimeout(timeout time.Duration) *RebalanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rebalance params
func (o *RebalanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rebalance params
func (o *RebalanceParams) WithContext(ctx context.Context) *RebalanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rebalance params
func (o *RebalanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rebalance params
func (o *RebalanceParams) WithHTTPClient(client *http.Client) *RebalanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rebalance params
func (o *RebalanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBestEfforts adds the bestEfforts to the rebalance params
func (o *RebalanceParams) WithBestEfforts(bestEfforts *bool) *RebalanceParams {
	o.SetBestEfforts(bestEfforts)
	return o
}

// SetBestEfforts adds the bestEfforts to the rebalance params
func (o *RebalanceParams) SetBestEfforts(bestEfforts *bool) {
	o.BestEfforts = bestEfforts
}

// WithBootstrap adds the bootstrap to the rebalance params
func (o *RebalanceParams) WithBootstrap(bootstrap *bool) *RebalanceParams {
	o.SetBootstrap(bootstrap)
	return o
}

// SetBootstrap adds the bootstrap to the rebalance params
func (o *RebalanceParams) SetBootstrap(bootstrap *bool) {
	o.Bootstrap = bootstrap
}

// WithDowntime adds the downtime to the rebalance params
func (o *RebalanceParams) WithDowntime(downtime *bool) *RebalanceParams {
	o.SetDowntime(downtime)
	return o
}

// SetDowntime adds the downtime to the rebalance params
func (o *RebalanceParams) SetDowntime(downtime *bool) {
	o.Downtime = downtime
}

// WithDryRun adds the dryRun to the rebalance params
func (o *RebalanceParams) WithDryRun(dryRun *bool) *RebalanceParams {
	o.SetDryRun(dryRun)
	return o
}

// SetDryRun adds the dryRun to the rebalance params
func (o *RebalanceParams) SetDryRun(dryRun *bool) {
	o.DryRun = dryRun
}

// WithExternalViewCheckIntervalInMs adds the externalViewCheckIntervalInMs to the rebalance params
func (o *RebalanceParams) WithExternalViewCheckIntervalInMs(externalViewCheckIntervalInMs *int64) *RebalanceParams {
	o.SetExternalViewCheckIntervalInMs(externalViewCheckIntervalInMs)
	return o
}

// SetExternalViewCheckIntervalInMs adds the externalViewCheckIntervalInMs to the rebalance params
func (o *RebalanceParams) SetExternalViewCheckIntervalInMs(externalViewCheckIntervalInMs *int64) {
	o.ExternalViewCheckIntervalInMs = externalViewCheckIntervalInMs
}

// WithExternalViewStabilizationTimeoutInMs adds the externalViewStabilizationTimeoutInMs to the rebalance params
func (o *RebalanceParams) WithExternalViewStabilizationTimeoutInMs(externalViewStabilizationTimeoutInMs *int64) *RebalanceParams {
	o.SetExternalViewStabilizationTimeoutInMs(externalViewStabilizationTimeoutInMs)
	return o
}

// SetExternalViewStabilizationTimeoutInMs adds the externalViewStabilizationTimeoutInMs to the rebalance params
func (o *RebalanceParams) SetExternalViewStabilizationTimeoutInMs(externalViewStabilizationTimeoutInMs *int64) {
	o.ExternalViewStabilizationTimeoutInMs = externalViewStabilizationTimeoutInMs
}

// WithIncludeConsuming adds the includeConsuming to the rebalance params
func (o *RebalanceParams) WithIncludeConsuming(includeConsuming *bool) *RebalanceParams {
	o.SetIncludeConsuming(includeConsuming)
	return o
}

// SetIncludeConsuming adds the includeConsuming to the rebalance params
func (o *RebalanceParams) SetIncludeConsuming(includeConsuming *bool) {
	o.IncludeConsuming = includeConsuming
}

// WithMinAvailableReplicas adds the minAvailableReplicas to the rebalance params
func (o *RebalanceParams) WithMinAvailableReplicas(minAvailableReplicas *int32) *RebalanceParams {
	o.SetMinAvailableReplicas(minAvailableReplicas)
	return o
}

// SetMinAvailableReplicas adds the minAvailableReplicas to the rebalance params
func (o *RebalanceParams) SetMinAvailableReplicas(minAvailableReplicas *int32) {
	o.MinAvailableReplicas = minAvailableReplicas
}

// WithReassignInstances adds the reassignInstances to the rebalance params
func (o *RebalanceParams) WithReassignInstances(reassignInstances *bool) *RebalanceParams {
	o.SetReassignInstances(reassignInstances)
	return o
}

// SetReassignInstances adds the reassignInstances to the rebalance params
func (o *RebalanceParams) SetReassignInstances(reassignInstances *bool) {
	o.ReassignInstances = reassignInstances
}

// WithTableName adds the tableName to the rebalance params
func (o *RebalanceParams) WithTableName(tableName string) *RebalanceParams {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the rebalance params
func (o *RebalanceParams) SetTableName(tableName string) {
	o.TableName = tableName
}

// WithType adds the typeVar to the rebalance params
func (o *RebalanceParams) WithType(typeVar string) *RebalanceParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the rebalance params
func (o *RebalanceParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *RebalanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BestEfforts != nil {

		// query param bestEfforts
		var qrBestEfforts bool

		if o.BestEfforts != nil {
			qrBestEfforts = *o.BestEfforts
		}
		qBestEfforts := swag.FormatBool(qrBestEfforts)
		if qBestEfforts != "" {

			if err := r.SetQueryParam("bestEfforts", qBestEfforts); err != nil {
				return err
			}
		}
	}

	if o.Bootstrap != nil {

		// query param bootstrap
		var qrBootstrap bool

		if o.Bootstrap != nil {
			qrBootstrap = *o.Bootstrap
		}
		qBootstrap := swag.FormatBool(qrBootstrap)
		if qBootstrap != "" {

			if err := r.SetQueryParam("bootstrap", qBootstrap); err != nil {
				return err
			}
		}
	}

	if o.Downtime != nil {

		// query param downtime
		var qrDowntime bool

		if o.Downtime != nil {
			qrDowntime = *o.Downtime
		}
		qDowntime := swag.FormatBool(qrDowntime)
		if qDowntime != "" {

			if err := r.SetQueryParam("downtime", qDowntime); err != nil {
				return err
			}
		}
	}

	if o.DryRun != nil {

		// query param dryRun
		var qrDryRun bool

		if o.DryRun != nil {
			qrDryRun = *o.DryRun
		}
		qDryRun := swag.FormatBool(qrDryRun)
		if qDryRun != "" {

			if err := r.SetQueryParam("dryRun", qDryRun); err != nil {
				return err
			}
		}
	}

	if o.ExternalViewCheckIntervalInMs != nil {

		// query param externalViewCheckIntervalInMs
		var qrExternalViewCheckIntervalInMs int64

		if o.ExternalViewCheckIntervalInMs != nil {
			qrExternalViewCheckIntervalInMs = *o.ExternalViewCheckIntervalInMs
		}
		qExternalViewCheckIntervalInMs := swag.FormatInt64(qrExternalViewCheckIntervalInMs)
		if qExternalViewCheckIntervalInMs != "" {

			if err := r.SetQueryParam("externalViewCheckIntervalInMs", qExternalViewCheckIntervalInMs); err != nil {
				return err
			}
		}
	}

	if o.ExternalViewStabilizationTimeoutInMs != nil {

		// query param externalViewStabilizationTimeoutInMs
		var qrExternalViewStabilizationTimeoutInMs int64

		if o.ExternalViewStabilizationTimeoutInMs != nil {
			qrExternalViewStabilizationTimeoutInMs = *o.ExternalViewStabilizationTimeoutInMs
		}
		qExternalViewStabilizationTimeoutInMs := swag.FormatInt64(qrExternalViewStabilizationTimeoutInMs)
		if qExternalViewStabilizationTimeoutInMs != "" {

			if err := r.SetQueryParam("externalViewStabilizationTimeoutInMs", qExternalViewStabilizationTimeoutInMs); err != nil {
				return err
			}
		}
	}

	if o.IncludeConsuming != nil {

		// query param includeConsuming
		var qrIncludeConsuming bool

		if o.IncludeConsuming != nil {
			qrIncludeConsuming = *o.IncludeConsuming
		}
		qIncludeConsuming := swag.FormatBool(qrIncludeConsuming)
		if qIncludeConsuming != "" {

			if err := r.SetQueryParam("includeConsuming", qIncludeConsuming); err != nil {
				return err
			}
		}
	}

	if o.MinAvailableReplicas != nil {

		// query param minAvailableReplicas
		var qrMinAvailableReplicas int32

		if o.MinAvailableReplicas != nil {
			qrMinAvailableReplicas = *o.MinAvailableReplicas
		}
		qMinAvailableReplicas := swag.FormatInt32(qrMinAvailableReplicas)
		if qMinAvailableReplicas != "" {

			if err := r.SetQueryParam("minAvailableReplicas", qMinAvailableReplicas); err != nil {
				return err
			}
		}
	}

	if o.ReassignInstances != nil {

		// query param reassignInstances
		var qrReassignInstances bool

		if o.ReassignInstances != nil {
			qrReassignInstances = *o.ReassignInstances
		}
		qReassignInstances := swag.FormatBool(qrReassignInstances)
		if qReassignInstances != "" {

			if err := r.SetQueryParam("reassignInstances", qReassignInstances); err != nil {
				return err
			}
		}
	}

	// path param tableName
	if err := r.SetPathParam("tableName", o.TableName); err != nil {
		return err
	}

	// query param type
	qrType := o.Type
	qType := qrType
	if qType != "" {

		if err := r.SetQueryParam("type", qType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
