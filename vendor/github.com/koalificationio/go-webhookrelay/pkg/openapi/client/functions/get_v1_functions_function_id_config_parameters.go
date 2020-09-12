// Code generated by go-swagger; DO NOT EDIT.

package functions

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
)

// NewGetV1FunctionsFunctionIDConfigParams creates a new GetV1FunctionsFunctionIDConfigParams object
// with the default values initialized.
func NewGetV1FunctionsFunctionIDConfigParams() *GetV1FunctionsFunctionIDConfigParams {
	var ()
	return &GetV1FunctionsFunctionIDConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1FunctionsFunctionIDConfigParamsWithTimeout creates a new GetV1FunctionsFunctionIDConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetV1FunctionsFunctionIDConfigParamsWithTimeout(timeout time.Duration) *GetV1FunctionsFunctionIDConfigParams {
	var ()
	return &GetV1FunctionsFunctionIDConfigParams{

		timeout: timeout,
	}
}

// NewGetV1FunctionsFunctionIDConfigParamsWithContext creates a new GetV1FunctionsFunctionIDConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetV1FunctionsFunctionIDConfigParamsWithContext(ctx context.Context) *GetV1FunctionsFunctionIDConfigParams {
	var ()
	return &GetV1FunctionsFunctionIDConfigParams{

		Context: ctx,
	}
}

// NewGetV1FunctionsFunctionIDConfigParamsWithHTTPClient creates a new GetV1FunctionsFunctionIDConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetV1FunctionsFunctionIDConfigParamsWithHTTPClient(client *http.Client) *GetV1FunctionsFunctionIDConfigParams {
	var ()
	return &GetV1FunctionsFunctionIDConfigParams{
		HTTPClient: client,
	}
}

/*GetV1FunctionsFunctionIDConfigParams contains all the parameters to send to the API endpoint
for the get v1 functions function ID config operation typically these are written to a http.Request
*/
type GetV1FunctionsFunctionIDConfigParams struct {

	/*FunctionID
	  ID of a function to manage

	*/
	FunctionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v1 functions function ID config params
func (o *GetV1FunctionsFunctionIDConfigParams) WithTimeout(timeout time.Duration) *GetV1FunctionsFunctionIDConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 functions function ID config params
func (o *GetV1FunctionsFunctionIDConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 functions function ID config params
func (o *GetV1FunctionsFunctionIDConfigParams) WithContext(ctx context.Context) *GetV1FunctionsFunctionIDConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 functions function ID config params
func (o *GetV1FunctionsFunctionIDConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 functions function ID config params
func (o *GetV1FunctionsFunctionIDConfigParams) WithHTTPClient(client *http.Client) *GetV1FunctionsFunctionIDConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 functions function ID config params
func (o *GetV1FunctionsFunctionIDConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFunctionID adds the functionID to the get v1 functions function ID config params
func (o *GetV1FunctionsFunctionIDConfigParams) WithFunctionID(functionID string) *GetV1FunctionsFunctionIDConfigParams {
	o.SetFunctionID(functionID)
	return o
}

// SetFunctionID adds the functionId to the get v1 functions function ID config params
func (o *GetV1FunctionsFunctionIDConfigParams) SetFunctionID(functionID string) {
	o.FunctionID = functionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1FunctionsFunctionIDConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param functionID
	if err := r.SetPathParam("functionID", o.FunctionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}