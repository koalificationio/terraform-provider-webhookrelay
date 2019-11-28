// Code generated by go-swagger; DO NOT EDIT.

package tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/koalificationio/go-webhookrelay/pkg/openapi/models"
)

// PostV1TokensReader is a Reader for the PostV1Tokens structure.
type PostV1TokensReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1TokensReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1TokensCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostV1TokensBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostV1TokensCreated creates a PostV1TokensCreated with default headers values
func NewPostV1TokensCreated() *PostV1TokensCreated {
	return &PostV1TokensCreated{}
}

/*PostV1TokensCreated handles this case with default header values.

Create a new token
*/
type PostV1TokensCreated struct {
	Payload *models.TokenCreateResponse
}

func (o *PostV1TokensCreated) Error() string {
	return fmt.Sprintf("[POST /v1/tokens][%d] postV1TokensCreated  %+v", 201, o.Payload)
}

func (o *PostV1TokensCreated) GetPayload() *models.TokenCreateResponse {
	return o.Payload
}

func (o *PostV1TokensCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TokenCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1TokensBadRequest creates a PostV1TokensBadRequest with default headers values
func NewPostV1TokensBadRequest() *PostV1TokensBadRequest {
	return &PostV1TokensBadRequest{}
}

/*PostV1TokensBadRequest handles this case with default header values.

PostV1TokensBadRequest post v1 tokens bad request
*/
type PostV1TokensBadRequest struct {
}

func (o *PostV1TokensBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/tokens][%d] postV1TokensBadRequest ", 400)
}

func (o *PostV1TokensBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}