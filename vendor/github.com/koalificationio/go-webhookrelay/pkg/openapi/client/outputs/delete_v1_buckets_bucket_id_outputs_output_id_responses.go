// Code generated by go-swagger; DO NOT EDIT.

package outputs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1BucketsBucketIDOutputsOutputIDReader is a Reader for the DeleteV1BucketsBucketIDOutputsOutputID structure.
type DeleteV1BucketsBucketIDOutputsOutputIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1BucketsBucketIDOutputsOutputIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1BucketsBucketIDOutputsOutputIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteV1BucketsBucketIDOutputsOutputIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1BucketsBucketIDOutputsOutputIDOK creates a DeleteV1BucketsBucketIDOutputsOutputIDOK with default headers values
func NewDeleteV1BucketsBucketIDOutputsOutputIDOK() *DeleteV1BucketsBucketIDOutputsOutputIDOK {
	return &DeleteV1BucketsBucketIDOutputsOutputIDOK{}
}

/*DeleteV1BucketsBucketIDOutputsOutputIDOK handles this case with default header values.

DeleteV1BucketsBucketIDOutputsOutputIDOK delete v1 buckets bucket Id outputs output Id o k
*/
type DeleteV1BucketsBucketIDOutputsOutputIDOK struct {
}

func (o *DeleteV1BucketsBucketIDOutputsOutputIDOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/buckets/{bucketID}/outputs/{outputID}][%d] deleteV1BucketsBucketIdOutputsOutputIdOK ", 200)
}

func (o *DeleteV1BucketsBucketIDOutputsOutputIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteV1BucketsBucketIDOutputsOutputIDNotFound creates a DeleteV1BucketsBucketIDOutputsOutputIDNotFound with default headers values
func NewDeleteV1BucketsBucketIDOutputsOutputIDNotFound() *DeleteV1BucketsBucketIDOutputsOutputIDNotFound {
	return &DeleteV1BucketsBucketIDOutputsOutputIDNotFound{}
}

/*DeleteV1BucketsBucketIDOutputsOutputIDNotFound handles this case with default header values.

Output or Bucket not found
*/
type DeleteV1BucketsBucketIDOutputsOutputIDNotFound struct {
}

func (o *DeleteV1BucketsBucketIDOutputsOutputIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/buckets/{bucketID}/outputs/{outputID}][%d] deleteV1BucketsBucketIdOutputsOutputIdNotFound ", 404)
}

func (o *DeleteV1BucketsBucketIDOutputsOutputIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
