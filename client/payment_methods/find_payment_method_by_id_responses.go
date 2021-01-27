// Code generated by go-swagger; DO NOT EDIT.

package payment_methods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// FindPaymentMethodByIDReader is a Reader for the FindPaymentMethodByID structure.
type FindPaymentMethodByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindPaymentMethodByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindPaymentMethodByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindPaymentMethodByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindPaymentMethodByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindPaymentMethodByIDOK creates a FindPaymentMethodByIDOK with default headers values
func NewFindPaymentMethodByIDOK() *FindPaymentMethodByIDOK {
	return &FindPaymentMethodByIDOK{}
}

/* FindPaymentMethodByIDOK describes a response with status code 200, with default header values.

ok
*/
type FindPaymentMethodByIDOK struct {
	Payload *types.PaymentMethod
}

func (o *FindPaymentMethodByIDOK) Error() string {
	return fmt.Sprintf("[GET /payment-methods/{id}][%d] findPaymentMethodByIdOK  %+v", 200, o.Payload)
}
func (o *FindPaymentMethodByIDOK) GetPayload() *types.PaymentMethod {
	return o.Payload
}

func (o *FindPaymentMethodByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.PaymentMethod)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindPaymentMethodByIDUnauthorized creates a FindPaymentMethodByIDUnauthorized with default headers values
func NewFindPaymentMethodByIDUnauthorized() *FindPaymentMethodByIDUnauthorized {
	return &FindPaymentMethodByIDUnauthorized{}
}

/* FindPaymentMethodByIDUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type FindPaymentMethodByIDUnauthorized struct {
}

func (o *FindPaymentMethodByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /payment-methods/{id}][%d] findPaymentMethodByIdUnauthorized ", 401)
}

func (o *FindPaymentMethodByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindPaymentMethodByIDNotFound creates a FindPaymentMethodByIDNotFound with default headers values
func NewFindPaymentMethodByIDNotFound() *FindPaymentMethodByIDNotFound {
	return &FindPaymentMethodByIDNotFound{}
}

/* FindPaymentMethodByIDNotFound describes a response with status code 404, with default header values.

not found
*/
type FindPaymentMethodByIDNotFound struct {
}

func (o *FindPaymentMethodByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /payment-methods/{id}][%d] findPaymentMethodByIdNotFound ", 404)
}

func (o *FindPaymentMethodByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
