// Code generated by go-swagger; DO NOT EDIT.

package tunnels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/koalificationio/go-webhookrelay/pkg/openapi/models"
)

// GetV1TunnelsTunnelIDReader is a Reader for the GetV1TunnelsTunnelID structure.
type GetV1TunnelsTunnelIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1TunnelsTunnelIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1TunnelsTunnelIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetV1TunnelsTunnelIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetV1TunnelsTunnelIDOK creates a GetV1TunnelsTunnelIDOK with default headers values
func NewGetV1TunnelsTunnelIDOK() *GetV1TunnelsTunnelIDOK {
	return &GetV1TunnelsTunnelIDOK{}
}

/*GetV1TunnelsTunnelIDOK handles this case with default header values.

Successful Response
*/
type GetV1TunnelsTunnelIDOK struct {
	Payload *models.Tunnel
}

func (o *GetV1TunnelsTunnelIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/tunnels/{tunnelID}][%d] getV1TunnelsTunnelIdOK  %+v", 200, o.Payload)
}

func (o *GetV1TunnelsTunnelIDOK) GetPayload() *models.Tunnel {
	return o.Payload
}

func (o *GetV1TunnelsTunnelIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tunnel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1TunnelsTunnelIDNotFound creates a GetV1TunnelsTunnelIDNotFound with default headers values
func NewGetV1TunnelsTunnelIDNotFound() *GetV1TunnelsTunnelIDNotFound {
	return &GetV1TunnelsTunnelIDNotFound{}
}

/*GetV1TunnelsTunnelIDNotFound handles this case with default header values.

Tunnel not found
*/
type GetV1TunnelsTunnelIDNotFound struct {
}

func (o *GetV1TunnelsTunnelIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/tunnels/{tunnelID}][%d] getV1TunnelsTunnelIdNotFound ", 404)
}

func (o *GetV1TunnelsTunnelIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}