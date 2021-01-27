// Code generated by go-swagger; DO NOT EDIT.

package emails

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// CreateEmailReader is a Reader for the CreateEmail structure.
type CreateEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateEmailCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateEmailUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateEmailUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateEmailCreated creates a CreateEmailCreated with default headers values
func NewCreateEmailCreated() *CreateEmailCreated {
	return &CreateEmailCreated{}
}

/* CreateEmailCreated describes a response with status code 201, with default header values.

created
*/
type CreateEmailCreated struct {
	Payload *types.Email
}

func (o *CreateEmailCreated) Error() string {
	return fmt.Sprintf("[POST /emails][%d] createEmailCreated  %+v", 201, o.Payload)
}
func (o *CreateEmailCreated) GetPayload() *types.Email {
	return o.Payload
}

func (o *CreateEmailCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.Email)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEmailUnauthorized creates a CreateEmailUnauthorized with default headers values
func NewCreateEmailUnauthorized() *CreateEmailUnauthorized {
	return &CreateEmailUnauthorized{}
}

/* CreateEmailUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type CreateEmailUnauthorized struct {
}

func (o *CreateEmailUnauthorized) Error() string {
	return fmt.Sprintf("[POST /emails][%d] createEmailUnauthorized ", 401)
}

func (o *CreateEmailUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateEmailUnprocessableEntity creates a CreateEmailUnprocessableEntity with default headers values
func NewCreateEmailUnprocessableEntity() *CreateEmailUnprocessableEntity {
	return &CreateEmailUnprocessableEntity{}
}

/* CreateEmailUnprocessableEntity describes a response with status code 422, with default header values.

unprocessable entity
*/
type CreateEmailUnprocessableEntity struct {
}

func (o *CreateEmailUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /emails][%d] createEmailUnprocessableEntity ", 422)
}

func (o *CreateEmailUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
