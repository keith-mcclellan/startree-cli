// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tenant API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tenant API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ChangeTenantState(params *ChangeTenantStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChangeTenantStateOK, error)

	CreateTenant(params *CreateTenantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateTenantOK, error)

	DeleteTenant(params *DeleteTenantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTenantOK, error)

	GetAllTenants(params *GetAllTenantsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllTenantsOK, error)

	GetTablesOnTenant(params *GetTablesOnTenantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTablesOnTenantOK, error)

	GetTenantMetadata(params *GetTenantMetadataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTenantMetadataOK, error)

	ListInstanceOrToggleTenantState(params *ListInstanceOrToggleTenantStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListInstanceOrToggleTenantStateOK, error)

	UpdateTenant(params *UpdateTenantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateTenantOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ChangeTenantState changes tenant state
*/
func (a *Client) ChangeTenantState(params *ChangeTenantStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChangeTenantStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeTenantStateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "changeTenantState",
		Method:             "POST",
		PathPattern:        "/tenants/{tenantName}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChangeTenantStateReader{formats: a.formats},
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
	success, ok := result.(*ChangeTenantStateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for changeTenantState: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateTenant creates a tenant
*/
func (a *Client) CreateTenant(params *CreateTenantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateTenantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTenantParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createTenant",
		Method:             "POST",
		PathPattern:        "/tenants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateTenantReader{formats: a.formats},
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
	success, ok := result.(*CreateTenantOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createTenant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteTenant deletes a tenant
*/
func (a *Client) DeleteTenant(params *DeleteTenantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTenantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTenantParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteTenant",
		Method:             "DELETE",
		PathPattern:        "/tenants/{tenantName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteTenantReader{formats: a.formats},
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
	success, ok := result.(*DeleteTenantOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteTenant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllTenants lists all tenants
*/
func (a *Client) GetAllTenants(params *GetAllTenantsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllTenantsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllTenantsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllTenants",
		Method:             "GET",
		PathPattern:        "/tenants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllTenantsReader{formats: a.formats},
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
	success, ok := result.(*GetAllTenantsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllTenants: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTablesOnTenant lists tables on a a server tenant
*/
func (a *Client) GetTablesOnTenant(params *GetTablesOnTenantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTablesOnTenantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTablesOnTenantParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTablesOnTenant",
		Method:             "GET",
		PathPattern:        "/tenants/{tenantName}/tables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTablesOnTenantReader{formats: a.formats},
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
	success, ok := result.(*GetTablesOnTenantOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTablesOnTenant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTenantMetadata gets tenant information
*/
func (a *Client) GetTenantMetadata(params *GetTenantMetadataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTenantMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTenantMetadataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTenantMetadata",
		Method:             "GET",
		PathPattern:        "/tenants/{tenantName}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTenantMetadataReader{formats: a.formats},
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
	success, ok := result.(*GetTenantMetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTenantMetadata: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListInstanceOrToggleTenantState lists instance for a tenant or enable disable drop a tenant
*/
func (a *Client) ListInstanceOrToggleTenantState(params *ListInstanceOrToggleTenantStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListInstanceOrToggleTenantStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListInstanceOrToggleTenantStateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listInstanceOrToggleTenantState",
		Method:             "GET",
		PathPattern:        "/tenants/{tenantName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListInstanceOrToggleTenantStateReader{formats: a.formats},
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
	success, ok := result.(*ListInstanceOrToggleTenantStateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listInstanceOrToggleTenantState: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateTenant updates a tenant
*/
func (a *Client) UpdateTenant(params *UpdateTenantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateTenantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTenantParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateTenant",
		Method:             "PUT",
		PathPattern:        "/tenants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateTenantReader{formats: a.formats},
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
	success, ok := result.(*UpdateTenantOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateTenant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
