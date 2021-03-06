// Code generated by go-swagger; DO NOT EDIT.

package inputs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1BucketsBucketIDInputsInputIDReader is a Reader for the DeleteV1BucketsBucketIDInputsInputID structure.
type DeleteV1BucketsBucketIDInputsInputIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1BucketsBucketIDInputsInputIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1BucketsBucketIDInputsInputIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteV1BucketsBucketIDInputsInputIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1BucketsBucketIDInputsInputIDOK creates a DeleteV1BucketsBucketIDInputsInputIDOK with default headers values
func NewDeleteV1BucketsBucketIDInputsInputIDOK() *DeleteV1BucketsBucketIDInputsInputIDOK {
	return &DeleteV1BucketsBucketIDInputsInputIDOK{}
}

/*DeleteV1BucketsBucketIDInputsInputIDOK handles this case with default header values.

DeleteV1BucketsBucketIDInputsInputIDOK delete v1 buckets bucket Id inputs input Id o k
*/
type DeleteV1BucketsBucketIDInputsInputIDOK struct {
}

func (o *DeleteV1BucketsBucketIDInputsInputIDOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/buckets/{bucketID}/inputs/{inputID}][%d] deleteV1BucketsBucketIdInputsInputIdOK ", 200)
}

func (o *DeleteV1BucketsBucketIDInputsInputIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteV1BucketsBucketIDInputsInputIDNotFound creates a DeleteV1BucketsBucketIDInputsInputIDNotFound with default headers values
func NewDeleteV1BucketsBucketIDInputsInputIDNotFound() *DeleteV1BucketsBucketIDInputsInputIDNotFound {
	return &DeleteV1BucketsBucketIDInputsInputIDNotFound{}
}

/*DeleteV1BucketsBucketIDInputsInputIDNotFound handles this case with default header values.

Input or Bucket not found
*/
type DeleteV1BucketsBucketIDInputsInputIDNotFound struct {
}

func (o *DeleteV1BucketsBucketIDInputsInputIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/buckets/{bucketID}/inputs/{inputID}][%d] deleteV1BucketsBucketIdInputsInputIdNotFound ", 404)
}

func (o *DeleteV1BucketsBucketIDInputsInputIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
