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

// PostV1TunnelsReader is a Reader for the PostV1Tunnels structure.
type PostV1TunnelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1TunnelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1TunnelsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostV1TunnelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostV1TunnelsCreated creates a PostV1TunnelsCreated with default headers values
func NewPostV1TunnelsCreated() *PostV1TunnelsCreated {
	return &PostV1TunnelsCreated{}
}

/*PostV1TunnelsCreated handles this case with default header values.

PostV1TunnelsCreated post v1 tunnels created
*/
type PostV1TunnelsCreated struct {
	Payload *models.Tunnel
}

func (o *PostV1TunnelsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/tunnels][%d] postV1TunnelsCreated  %+v", 201, o.Payload)
}

func (o *PostV1TunnelsCreated) GetPayload() *models.Tunnel {
	return o.Payload
}

func (o *PostV1TunnelsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tunnel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1TunnelsBadRequest creates a PostV1TunnelsBadRequest with default headers values
func NewPostV1TunnelsBadRequest() *PostV1TunnelsBadRequest {
	return &PostV1TunnelsBadRequest{}
}

/*PostV1TunnelsBadRequest handles this case with default header values.

Invalid tunnel request supplied
*/
type PostV1TunnelsBadRequest struct {
}

func (o *PostV1TunnelsBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/tunnels][%d] postV1TunnelsBadRequest ", 400)
}

func (o *PostV1TunnelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}