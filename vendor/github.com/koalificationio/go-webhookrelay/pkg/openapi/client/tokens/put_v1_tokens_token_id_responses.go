// Code generated by go-swagger; DO NOT EDIT.

package tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/koalificationio/go-webhookrelay/pkg/openapi/models"
)

// PutV1TokensTokenIDReader is a Reader for the PutV1TokensTokenID structure.
type PutV1TokensTokenIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutV1TokensTokenIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutV1TokensTokenIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutV1TokensTokenIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutV1TokensTokenIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutV1TokensTokenIDOK creates a PutV1TokensTokenIDOK with default headers values
func NewPutV1TokensTokenIDOK() *PutV1TokensTokenIDOK {
	return &PutV1TokensTokenIDOK{}
}

/*PutV1TokensTokenIDOK handles this case with default header values.

Successful Response
*/
type PutV1TokensTokenIDOK struct {
	Payload *models.Token
}

func (o *PutV1TokensTokenIDOK) Error() string {
	return fmt.Sprintf("[PUT /v1/tokens/{tokenID}][%d] putV1TokensTokenIdOK  %+v", 200, o.Payload)
}

func (o *PutV1TokensTokenIDOK) GetPayload() *models.Token {
	return o.Payload
}

func (o *PutV1TokensTokenIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutV1TokensTokenIDBadRequest creates a PutV1TokensTokenIDBadRequest with default headers values
func NewPutV1TokensTokenIDBadRequest() *PutV1TokensTokenIDBadRequest {
	return &PutV1TokensTokenIDBadRequest{}
}

/*PutV1TokensTokenIDBadRequest handles this case with default header values.

Bad request (check response message)
*/
type PutV1TokensTokenIDBadRequest struct {
}

func (o *PutV1TokensTokenIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/tokens/{tokenID}][%d] putV1TokensTokenIdBadRequest ", 400)
}

func (o *PutV1TokensTokenIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutV1TokensTokenIDNotFound creates a PutV1TokensTokenIDNotFound with default headers values
func NewPutV1TokensTokenIDNotFound() *PutV1TokensTokenIDNotFound {
	return &PutV1TokensTokenIDNotFound{}
}

/*PutV1TokensTokenIDNotFound handles this case with default header values.

Tunnel not found
*/
type PutV1TokensTokenIDNotFound struct {
}

func (o *PutV1TokensTokenIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/tokens/{tokenID}][%d] putV1TokensTokenIdNotFound ", 404)
}

func (o *PutV1TokensTokenIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
