// Code generated by go-swagger; DO NOT EDIT.

package broker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new broker API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for broker API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetBrokersForTable(params *GetBrokersForTableParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBrokersForTableOK, error)

	GetBrokersForTableV2(params *GetBrokersForTableV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBrokersForTableV2OK, error)

	GetBrokersForTenant(params *GetBrokersForTenantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBrokersForTenantOK, error)

	GetBrokersForTenantV2(params *GetBrokersForTenantV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBrokersForTenantV2OK, error)

	GetTablesToBrokersMapping(params *GetTablesToBrokersMappingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTablesToBrokersMappingOK, error)

	GetTablesToBrokersMappingV2(params *GetTablesToBrokersMappingV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTablesToBrokersMappingV2OK, error)

	GetTenantsToBrokersMapping(params *GetTenantsToBrokersMappingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTenantsToBrokersMappingOK, error)

	GetTenantsToBrokersMappingV2(params *GetTenantsToBrokersMappingV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTenantsToBrokersMappingV2OK, error)

	ListBrokersMapping(params *ListBrokersMappingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListBrokersMappingOK, error)

	ListBrokersMappingV2(params *ListBrokersMappingV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListBrokersMappingV2OK, error)

	ToggleQueryRateLimiting(params *ToggleQueryRateLimitingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ToggleQueryRateLimitingOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetBrokersForTable lists brokers for a given table

List brokers for a given table
*/
func (a *Client) GetBrokersForTable(params *GetBrokersForTableParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBrokersForTableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBrokersForTableParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBrokersForTable",
		Method:             "GET",
		PathPattern:        "/brokers/tables/{tableName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBrokersForTableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBrokersForTableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBrokersForTable: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBrokersForTableV2 lists brokers for a given table

List brokers for a given table
*/
func (a *Client) GetBrokersForTableV2(params *GetBrokersForTableV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBrokersForTableV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBrokersForTableV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBrokersForTableV2",
		Method:             "GET",
		PathPattern:        "/v2/brokers/tables/{tableName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBrokersForTableV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBrokersForTableV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBrokersForTableV2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBrokersForTenant lists brokers for a given tenant

List brokers for a given tenant
*/
func (a *Client) GetBrokersForTenant(params *GetBrokersForTenantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBrokersForTenantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBrokersForTenantParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBrokersForTenant",
		Method:             "GET",
		PathPattern:        "/brokers/tenants/{tenantName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBrokersForTenantReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBrokersForTenantOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBrokersForTenant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBrokersForTenantV2 lists brokers for a given tenant

List brokers for a given tenant
*/
func (a *Client) GetBrokersForTenantV2(params *GetBrokersForTenantV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBrokersForTenantV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBrokersForTenantV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBrokersForTenantV2",
		Method:             "GET",
		PathPattern:        "/v2/brokers/tenants/{tenantName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBrokersForTenantV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBrokersForTenantV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBrokersForTenantV2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTablesToBrokersMapping lists tables to brokers mappings

List tables to brokers mappings
*/
func (a *Client) GetTablesToBrokersMapping(params *GetTablesToBrokersMappingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTablesToBrokersMappingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTablesToBrokersMappingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTablesToBrokersMapping",
		Method:             "GET",
		PathPattern:        "/brokers/tables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTablesToBrokersMappingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTablesToBrokersMappingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTablesToBrokersMapping: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTablesToBrokersMappingV2 lists tables to brokers mappings

List tables to brokers mappings
*/
func (a *Client) GetTablesToBrokersMappingV2(params *GetTablesToBrokersMappingV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTablesToBrokersMappingV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTablesToBrokersMappingV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTablesToBrokersMappingV2",
		Method:             "GET",
		PathPattern:        "/v2/brokers/tables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTablesToBrokersMappingV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTablesToBrokersMappingV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTablesToBrokersMappingV2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTenantsToBrokersMapping lists tenants to brokers mappings

List tenants to brokers mappings
*/
func (a *Client) GetTenantsToBrokersMapping(params *GetTenantsToBrokersMappingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTenantsToBrokersMappingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTenantsToBrokersMappingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTenantsToBrokersMapping",
		Method:             "GET",
		PathPattern:        "/brokers/tenants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTenantsToBrokersMappingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTenantsToBrokersMappingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTenantsToBrokersMapping: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTenantsToBrokersMappingV2 lists tenants to brokers mappings

List tenants to brokers mappings
*/
func (a *Client) GetTenantsToBrokersMappingV2(params *GetTenantsToBrokersMappingV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTenantsToBrokersMappingV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTenantsToBrokersMappingV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTenantsToBrokersMappingV2",
		Method:             "GET",
		PathPattern:        "/v2/brokers/tenants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTenantsToBrokersMappingV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTenantsToBrokersMappingV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTenantsToBrokersMappingV2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListBrokersMapping lists tenants and tables to brokers mappings

List tenants and tables to brokers mappings
*/
func (a *Client) ListBrokersMapping(params *ListBrokersMappingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListBrokersMappingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListBrokersMappingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listBrokersMapping",
		Method:             "GET",
		PathPattern:        "/brokers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListBrokersMappingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListBrokersMappingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listBrokersMapping: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListBrokersMappingV2 lists tenants and tables to brokers mappings

List tenants and tables to brokers mappings
*/
func (a *Client) ListBrokersMappingV2(params *ListBrokersMappingV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListBrokersMappingV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListBrokersMappingV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listBrokersMappingV2",
		Method:             "GET",
		PathPattern:        "/v2/brokers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListBrokersMappingV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListBrokersMappingV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listBrokersMappingV2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ToggleQueryRateLimiting enables disable the query rate limiting for a broker instance

Enable/disable the query rate limiting for a broker instance
*/
func (a *Client) ToggleQueryRateLimiting(params *ToggleQueryRateLimitingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ToggleQueryRateLimitingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewToggleQueryRateLimitingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "toggleQueryRateLimiting",
		Method:             "POST",
		PathPattern:        "/brokers/instances/{instanceName}/qps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ToggleQueryRateLimitingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ToggleQueryRateLimitingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for toggleQueryRateLimiting: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
