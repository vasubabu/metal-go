// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// CreateSpotMarketRequestReader is a Reader for the CreateSpotMarketRequest structure.
type CreateSpotMarketRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSpotMarketRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateSpotMarketRequestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateSpotMarketRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateSpotMarketRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateSpotMarketRequestUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSpotMarketRequestCreated creates a CreateSpotMarketRequestCreated with default headers values
func NewCreateSpotMarketRequestCreated() *CreateSpotMarketRequestCreated {
	return &CreateSpotMarketRequestCreated{}
}

/* CreateSpotMarketRequestCreated describes a response with status code 201, with default header values.

created
*/
type CreateSpotMarketRequestCreated struct {
	Payload *types.SpotMarketRequest
}

func (o *CreateSpotMarketRequestCreated) Error() string {
	return fmt.Sprintf("[POST /projects/{id}/spot-market-requests][%d] createSpotMarketRequestCreated  %+v", 201, o.Payload)
}
func (o *CreateSpotMarketRequestCreated) GetPayload() *types.SpotMarketRequest {
	return o.Payload
}

func (o *CreateSpotMarketRequestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.SpotMarketRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpotMarketRequestUnauthorized creates a CreateSpotMarketRequestUnauthorized with default headers values
func NewCreateSpotMarketRequestUnauthorized() *CreateSpotMarketRequestUnauthorized {
	return &CreateSpotMarketRequestUnauthorized{}
}

/* CreateSpotMarketRequestUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type CreateSpotMarketRequestUnauthorized struct {
}

func (o *CreateSpotMarketRequestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{id}/spot-market-requests][%d] createSpotMarketRequestUnauthorized ", 401)
}

func (o *CreateSpotMarketRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSpotMarketRequestNotFound creates a CreateSpotMarketRequestNotFound with default headers values
func NewCreateSpotMarketRequestNotFound() *CreateSpotMarketRequestNotFound {
	return &CreateSpotMarketRequestNotFound{}
}

/* CreateSpotMarketRequestNotFound describes a response with status code 404, with default header values.

not found
*/
type CreateSpotMarketRequestNotFound struct {
}

func (o *CreateSpotMarketRequestNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{id}/spot-market-requests][%d] createSpotMarketRequestNotFound ", 404)
}

func (o *CreateSpotMarketRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSpotMarketRequestUnprocessableEntity creates a CreateSpotMarketRequestUnprocessableEntity with default headers values
func NewCreateSpotMarketRequestUnprocessableEntity() *CreateSpotMarketRequestUnprocessableEntity {
	return &CreateSpotMarketRequestUnprocessableEntity{}
}

/* CreateSpotMarketRequestUnprocessableEntity describes a response with status code 422, with default header values.

unprocessable entity
*/
type CreateSpotMarketRequestUnprocessableEntity struct {
}

func (o *CreateSpotMarketRequestUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /projects/{id}/spot-market-requests][%d] createSpotMarketRequestUnprocessableEntity ", 422)
}

func (o *CreateSpotMarketRequestUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
