// Code generated by go-swagger; DO NOT EDIT.

package two_factor_auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EnableTfaAppReader is a Reader for the EnableTfaApp structure.
type EnableTfaAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableTfaAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnableTfaAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEnableTfaAppUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEnableTfaAppOK creates a EnableTfaAppOK with default headers values
func NewEnableTfaAppOK() *EnableTfaAppOK {
	return &EnableTfaAppOK{}
}

/* EnableTfaAppOK describes a response with status code 200, with default header values.

ok
*/
type EnableTfaAppOK struct {
}

func (o *EnableTfaAppOK) Error() string {
	return fmt.Sprintf("[POST /user/otp/app][%d] enableTfaAppOK ", 200)
}

func (o *EnableTfaAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableTfaAppUnauthorized creates a EnableTfaAppUnauthorized with default headers values
func NewEnableTfaAppUnauthorized() *EnableTfaAppUnauthorized {
	return &EnableTfaAppUnauthorized{}
}

/* EnableTfaAppUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type EnableTfaAppUnauthorized struct {
}

func (o *EnableTfaAppUnauthorized) Error() string {
	return fmt.Sprintf("[POST /user/otp/app][%d] enableTfaAppUnauthorized ", 401)
}

func (o *EnableTfaAppUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
