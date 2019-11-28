// Code generated by go-swagger; DO NOT EDIT.

package buckets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteV1BucketsBucketIDReader is a Reader for the DeleteV1BucketsBucketID structure.
type DeleteV1BucketsBucketIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1BucketsBucketIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1BucketsBucketIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteV1BucketsBucketIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteV1BucketsBucketIDOK creates a DeleteV1BucketsBucketIDOK with default headers values
func NewDeleteV1BucketsBucketIDOK() *DeleteV1BucketsBucketIDOK {
	return &DeleteV1BucketsBucketIDOK{}
}

/*DeleteV1BucketsBucketIDOK handles this case with default header values.

Successful Response
*/
type DeleteV1BucketsBucketIDOK struct {
}

func (o *DeleteV1BucketsBucketIDOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/buckets/{bucketID}][%d] deleteV1BucketsBucketIdOK ", 200)
}

func (o *DeleteV1BucketsBucketIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteV1BucketsBucketIDNotFound creates a DeleteV1BucketsBucketIDNotFound with default headers values
func NewDeleteV1BucketsBucketIDNotFound() *DeleteV1BucketsBucketIDNotFound {
	return &DeleteV1BucketsBucketIDNotFound{}
}

/*DeleteV1BucketsBucketIDNotFound handles this case with default header values.

Bucket not found
*/
type DeleteV1BucketsBucketIDNotFound struct {
}

func (o *DeleteV1BucketsBucketIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/buckets/{bucketID}][%d] deleteV1BucketsBucketIdNotFound ", 404)
}

func (o *DeleteV1BucketsBucketIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}