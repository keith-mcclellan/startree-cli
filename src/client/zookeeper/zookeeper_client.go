// Code generated by go-swagger; DO NOT EDIT.

package zookeeper

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new zookeeper API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for zookeeper API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Delete(params *DeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOK, *DeleteNoContent, error)

	GetChildren(params *GetChildrenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetChildrenOK, *GetChildrenNoContent, error)

	GetData(params *GetDataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDataOK, *GetDataNoContent, error)

	Ls(params *LsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LsOK, error)

	Lsl(params *LslParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LslOK, error)

	PutChildren(params *PutChildrenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutChildrenOK, *PutChildrenNoContent, error)

	PutData(params *PutDataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutDataOK, *PutDataNoContent, error)

	Stat(params *StatParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Delete deletes the znode at this path
*/
func (a *Client) Delete(params *DeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOK, *DeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete",
		Method:             "DELETE",
		PathPattern:        "/zk/delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteOK:
		return value, nil, nil
	case *DeleteNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zookeeper: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetChildren gets all child znodes
*/
func (a *Client) GetChildren(params *GetChildrenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetChildrenOK, *GetChildrenNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetChildrenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getChildren",
		Method:             "GET",
		PathPattern:        "/zk/getChildren",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetChildrenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetChildrenOK:
		return value, nil, nil
	case *GetChildrenNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zookeeper: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetData gets content of the znode
*/
func (a *Client) GetData(params *GetDataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDataOK, *GetDataNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getData",
		Method:             "GET",
		PathPattern:        "/zk/get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetDataOK:
		return value, nil, nil
	case *GetDataNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zookeeper: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Ls lists the child znodes
*/
func (a *Client) Ls(params *LsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ls",
		Method:             "GET",
		PathPattern:        "/zk/ls",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LsReader{formats: a.formats},
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
	success, ok := result.(*LsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ls: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Lsl lists the child znodes along with stats
*/
func (a *Client) Lsl(params *LslParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LslOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLslParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "lsl",
		Method:             "GET",
		PathPattern:        "/zk/lsl",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LslReader{formats: a.formats},
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
	success, ok := result.(*LslOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for lsl: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutChildren updates the content of multiple zn record node under the same path
*/
func (a *Client) PutChildren(params *PutChildrenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutChildrenOK, *PutChildrenNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutChildrenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "putChildren",
		Method:             "PUT",
		PathPattern:        "/zk/putChildren",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutChildrenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PutChildrenOK:
		return value, nil, nil
	case *PutChildrenNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zookeeper: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutData updates the content of the node
*/
func (a *Client) PutData(params *PutDataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutDataOK, *PutDataNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutDataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "putData",
		Method:             "PUT",
		PathPattern:        "/zk/put",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutDataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PutDataOK:
		return value, nil, nil
	case *PutDataNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zookeeper: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Stat gets the stat

	Use this api to fetch additional details of a znode such as creation time, modified time, numChildren etc
*/
func (a *Client) Stat(params *StatParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "stat",
		Method:             "GET",
		PathPattern:        "/zk/stat",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StatReader{formats: a.formats},
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
	success, ok := result.(*StatOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for stat: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}