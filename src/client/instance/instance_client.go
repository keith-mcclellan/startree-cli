// Code generated by go-swagger; DO NOT EDIT.

package instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new instance API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for instance API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddInstance(params *AddInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddInstanceOK, error)

	DropInstance(params *DropInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DropInstanceOK, error)

	GetAllInstances(params *GetAllInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllInstancesOK, error)

	GetInstance(params *GetInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInstanceOK, error)

	ToggleInstanceState(params *ToggleInstanceStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ToggleInstanceStateOK, error)

	UpdateBrokerResource(params *UpdateBrokerResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBrokerResourceOK, error)

	UpdateInstance(params *UpdateInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateInstanceOK, error)

	UpdateInstanceTags(params *UpdateInstanceTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateInstanceTagsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddInstance creates a new instance

Creates a new instance with given instance config
*/
func (a *Client) AddInstance(params *AddInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addInstance",
		Method:             "POST",
		PathPattern:        "/instances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddInstanceReader{formats: a.formats},
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
	success, ok := result.(*AddInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addInstance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DropInstance drops an instance

Drop an instance
*/
func (a *Client) DropInstance(params *DropInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DropInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDropInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dropInstance",
		Method:             "DELETE",
		PathPattern:        "/instances/{instanceName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DropInstanceReader{formats: a.formats},
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
	success, ok := result.(*DropInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dropInstance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllInstances lists all instances
*/
func (a *Client) GetAllInstances(params *GetAllInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllInstancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllInstancesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllInstances",
		Method:             "GET",
		PathPattern:        "/instances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllInstancesReader{formats: a.formats},
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
	success, ok := result.(*GetAllInstancesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllInstances: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInstance gets instance information
*/
func (a *Client) GetInstance(params *GetInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getInstance",
		Method:             "GET",
		PathPattern:        "/instances/{instanceName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInstanceReader{formats: a.formats},
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
	success, ok := result.(*GetInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInstance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ToggleInstanceState enables disable drop an instance

Enable/disable/drop an instance
*/
func (a *Client) ToggleInstanceState(params *ToggleInstanceStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ToggleInstanceStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewToggleInstanceStateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "toggleInstanceState",
		Method:             "POST",
		PathPattern:        "/instances/{instanceName}/state",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ToggleInstanceStateReader{formats: a.formats},
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
	success, ok := result.(*ToggleInstanceStateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for toggleInstanceState: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateBrokerResource updates the tables served by the specified broker instance in the broker resource

Broker resource should be updated when a new broker instance is added, or the tags for an existing broker are changed. Updating broker resource requires reading all the table configs, which can be costly for large cluster. Consider updating broker resource for each table individually.
*/
func (a *Client) UpdateBrokerResource(params *UpdateBrokerResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBrokerResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBrokerResourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateBrokerResource",
		Method:             "POST",
		PathPattern:        "/instances/{instanceName}/updateBrokerResource",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateBrokerResourceReader{formats: a.formats},
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
	success, ok := result.(*UpdateBrokerResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateBrokerResource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateInstance updates the specified instance

Update specified instance with given instance config
*/
func (a *Client) UpdateInstance(params *UpdateInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateInstance",
		Method:             "PUT",
		PathPattern:        "/instances/{instanceName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateInstanceReader{formats: a.formats},
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
	success, ok := result.(*UpdateInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateInstance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateInstanceTags updates the tags of the specified instance

Update the tags of the specified instance
*/
func (a *Client) UpdateInstanceTags(params *UpdateInstanceTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateInstanceTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateInstanceTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateInstanceTags",
		Method:             "PUT",
		PathPattern:        "/instances/{instanceName}/updateTags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateInstanceTagsReader{formats: a.formats},
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
	success, ok := result.(*UpdateInstanceTagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateInstanceTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
