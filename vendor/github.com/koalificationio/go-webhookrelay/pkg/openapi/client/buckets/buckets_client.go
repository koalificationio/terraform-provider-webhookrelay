// Code generated by go-swagger; DO NOT EDIT.

package buckets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new buckets API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for buckets API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteV1BucketsBucketID deletes bucket
*/
func (a *Client) DeleteV1BucketsBucketID(params *DeleteV1BucketsBucketIDParams) (*DeleteV1BucketsBucketIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1BucketsBucketIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteV1BucketsBucketID",
		Method:             "DELETE",
		PathPattern:        "/v1/buckets/{bucketID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1BucketsBucketIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteV1BucketsBucketIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteV1BucketsBucketID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1Buckets lists all buckets
*/
func (a *Client) GetV1Buckets(params *GetV1BucketsParams) (*GetV1BucketsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1BucketsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetV1Buckets",
		Method:             "GET",
		PathPattern:        "/v1/buckets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1BucketsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1BucketsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1Buckets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1BucketsBucketID gets bucket details
*/
func (a *Client) GetV1BucketsBucketID(params *GetV1BucketsBucketIDParams) (*GetV1BucketsBucketIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1BucketsBucketIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetV1BucketsBucketID",
		Method:             "GET",
		PathPattern:        "/v1/buckets/{bucketID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1BucketsBucketIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1BucketsBucketIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1BucketsBucketID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1Buckets creates a new bucket

You may create your own bucket using this action. It takes a JSON object containing a bucket request. Once bucket is created, it gets  assigned a default input to accept webhooks but you will still have to  create a new output to give it a destination.
*/
func (a *Client) PostV1Buckets(params *PostV1BucketsParams) (*PostV1BucketsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1BucketsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostV1Buckets",
		Method:             "POST",
		PathPattern:        "/v1/buckets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1BucketsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostV1BucketsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostV1Buckets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutV1BucketsBucketID updates bucket

Update bucket.
*/
func (a *Client) PutV1BucketsBucketID(params *PutV1BucketsBucketIDParams) (*PutV1BucketsBucketIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutV1BucketsBucketIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutV1BucketsBucketID",
		Method:             "PUT",
		PathPattern:        "/v1/buckets/{bucketID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutV1BucketsBucketIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutV1BucketsBucketIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutV1BucketsBucketID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}