// Code generated by go-swagger; DO NOT EDIT.

package tokens

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

// NewGetV1TokensParams creates a new GetV1TokensParams object
// with the default values initialized.
func NewGetV1TokensParams() *GetV1TokensParams {

	return &GetV1TokensParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1TokensParamsWithTimeout creates a new GetV1TokensParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetV1TokensParamsWithTimeout(timeout time.Duration) *GetV1TokensParams {

	return &GetV1TokensParams{

		timeout: timeout,
	}
}

// NewGetV1TokensParamsWithContext creates a new GetV1TokensParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetV1TokensParamsWithContext(ctx context.Context) *GetV1TokensParams {

	return &GetV1TokensParams{

		Context: ctx,
	}
}

// NewGetV1TokensParamsWithHTTPClient creates a new GetV1TokensParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetV1TokensParamsWithHTTPClient(client *http.Client) *GetV1TokensParams {

	return &GetV1TokensParams{
		HTTPClient: client,
	}
}

/*GetV1TokensParams contains all the parameters to send to the API endpoint
for the get v1 tokens operation typically these are written to a http.Request
*/
type GetV1TokensParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v1 tokens params
func (o *GetV1TokensParams) WithTimeout(timeout time.Duration) *GetV1TokensParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 tokens params
func (o *GetV1TokensParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 tokens params
func (o *GetV1TokensParams) WithContext(ctx context.Context) *GetV1TokensParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 tokens params
func (o *GetV1TokensParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 tokens params
func (o *GetV1TokensParams) WithHTTPClient(client *http.Client) *GetV1TokensParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 tokens params
func (o *GetV1TokensParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1TokensParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
