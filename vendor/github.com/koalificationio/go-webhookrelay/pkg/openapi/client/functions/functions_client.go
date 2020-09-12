// Code generated by go-swagger; DO NOT EDIT.

package functions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new functions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for functions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteV1FunctionsFunctionID(params *DeleteV1FunctionsFunctionIDParams) (*DeleteV1FunctionsFunctionIDOK, error)

	GetV1Functions(params *GetV1FunctionsParams) (*GetV1FunctionsOK, error)

	GetV1FunctionsFunctionID(params *GetV1FunctionsFunctionIDParams) (*GetV1FunctionsFunctionIDOK, error)

	PostV1Functions(params *PostV1FunctionsParams) (*PostV1FunctionsCreated, error)

	PutV1FunctionsFunctionID(params *PutV1FunctionsFunctionIDParams) (*PutV1FunctionsFunctionIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteV1FunctionsFunctionID deletes function
*/
func (a *Client) DeleteV1FunctionsFunctionID(params *DeleteV1FunctionsFunctionIDParams) (*DeleteV1FunctionsFunctionIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1FunctionsFunctionIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteV1FunctionsFunctionID",
		Method:             "DELETE",
		PathPattern:        "/v1/functions/{functionID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1FunctionsFunctionIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteV1FunctionsFunctionIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteV1FunctionsFunctionID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1Functions lists all functions
*/
func (a *Client) GetV1Functions(params *GetV1FunctionsParams) (*GetV1FunctionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1FunctionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetV1Functions",
		Method:             "GET",
		PathPattern:        "/v1/functions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1FunctionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1FunctionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1Functions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1FunctionsFunctionID gets function details
*/
func (a *Client) GetV1FunctionsFunctionID(params *GetV1FunctionsFunctionIDParams) (*GetV1FunctionsFunctionIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1FunctionsFunctionIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetV1FunctionsFunctionID",
		Method:             "GET",
		PathPattern:        "/v1/functions/{functionID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1FunctionsFunctionIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1FunctionsFunctionIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1FunctionsFunctionID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostV1Functions creates a new function

  You may create your own functions using this action. It takes a JSON object containing a function request. Once function is created, you can  assign it based on its ID to bucket's input and/or output.
*/
func (a *Client) PostV1Functions(params *PostV1FunctionsParams) (*PostV1FunctionsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1FunctionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostV1Functions",
		Method:             "POST",
		PathPattern:        "/v1/functions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1FunctionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostV1FunctionsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostV1Functions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutV1FunctionsFunctionID updates function

  Update function.
*/
func (a *Client) PutV1FunctionsFunctionID(params *PutV1FunctionsFunctionIDParams) (*PutV1FunctionsFunctionIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutV1FunctionsFunctionIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutV1FunctionsFunctionID",
		Method:             "PUT",
		PathPattern:        "/v1/functions/{functionID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutV1FunctionsFunctionIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutV1FunctionsFunctionIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutV1FunctionsFunctionID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}