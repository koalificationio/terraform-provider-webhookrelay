// Code generated by go-swagger; DO NOT EDIT.

package buckets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetV1BucketsBucketIDParams creates a new GetV1BucketsBucketIDParams object
// with the default values initialized.
func NewGetV1BucketsBucketIDParams() *GetV1BucketsBucketIDParams {
	var ()
	return &GetV1BucketsBucketIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1BucketsBucketIDParamsWithTimeout creates a new GetV1BucketsBucketIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetV1BucketsBucketIDParamsWithTimeout(timeout time.Duration) *GetV1BucketsBucketIDParams {
	var ()
	return &GetV1BucketsBucketIDParams{

		timeout: timeout,
	}
}

// NewGetV1BucketsBucketIDParamsWithContext creates a new GetV1BucketsBucketIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetV1BucketsBucketIDParamsWithContext(ctx context.Context) *GetV1BucketsBucketIDParams {
	var ()
	return &GetV1BucketsBucketIDParams{

		Context: ctx,
	}
}

// NewGetV1BucketsBucketIDParamsWithHTTPClient creates a new GetV1BucketsBucketIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetV1BucketsBucketIDParamsWithHTTPClient(client *http.Client) *GetV1BucketsBucketIDParams {
	var ()
	return &GetV1BucketsBucketIDParams{
		HTTPClient: client,
	}
}

/*GetV1BucketsBucketIDParams contains all the parameters to send to the API endpoint
for the get v1 buckets bucket ID operation typically these are written to a http.Request
*/
type GetV1BucketsBucketIDParams struct {

	/*BucketID
	  ID of a bucket to return

	*/
	BucketID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v1 buckets bucket ID params
func (o *GetV1BucketsBucketIDParams) WithTimeout(timeout time.Duration) *GetV1BucketsBucketIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 buckets bucket ID params
func (o *GetV1BucketsBucketIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 buckets bucket ID params
func (o *GetV1BucketsBucketIDParams) WithContext(ctx context.Context) *GetV1BucketsBucketIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 buckets bucket ID params
func (o *GetV1BucketsBucketIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 buckets bucket ID params
func (o *GetV1BucketsBucketIDParams) WithHTTPClient(client *http.Client) *GetV1BucketsBucketIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 buckets bucket ID params
func (o *GetV1BucketsBucketIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketID adds the bucketID to the get v1 buckets bucket ID params
func (o *GetV1BucketsBucketIDParams) WithBucketID(bucketID string) *GetV1BucketsBucketIDParams {
	o.SetBucketID(bucketID)
	return o
}

// SetBucketID adds the bucketId to the get v1 buckets bucket ID params
func (o *GetV1BucketsBucketIDParams) SetBucketID(bucketID string) {
	o.BucketID = bucketID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1BucketsBucketIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bucketID
	if err := r.SetPathParam("bucketID", o.BucketID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}