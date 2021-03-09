// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteDeviceReader is a Reader for the DeleteDevice structure.
type DeleteDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteDeviceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteDeviceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDeviceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDeviceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteDeviceUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDeviceNoContent creates a DeleteDeviceNoContent with default headers values
func NewDeleteDeviceNoContent() *DeleteDeviceNoContent {
	return &DeleteDeviceNoContent{}
}

/* DeleteDeviceNoContent describes a response with status code 204, with default header values.

no content
*/
type DeleteDeviceNoContent struct {
}

func (o *DeleteDeviceNoContent) Error() string {
	return fmt.Sprintf("[DELETE /devices/{id}][%d] deleteDeviceNoContent ", 204)
}

func (o *DeleteDeviceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceUnauthorized creates a DeleteDeviceUnauthorized with default headers values
func NewDeleteDeviceUnauthorized() *DeleteDeviceUnauthorized {
	return &DeleteDeviceUnauthorized{}
}

/* DeleteDeviceUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type DeleteDeviceUnauthorized struct {
}

func (o *DeleteDeviceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /devices/{id}][%d] deleteDeviceUnauthorized ", 401)
}

func (o *DeleteDeviceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceForbidden creates a DeleteDeviceForbidden with default headers values
func NewDeleteDeviceForbidden() *DeleteDeviceForbidden {
	return &DeleteDeviceForbidden{}
}

/* DeleteDeviceForbidden describes a response with status code 403, with default header values.

forbidden
*/
type DeleteDeviceForbidden struct {
}

func (o *DeleteDeviceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /devices/{id}][%d] deleteDeviceForbidden ", 403)
}

func (o *DeleteDeviceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceNotFound creates a DeleteDeviceNotFound with default headers values
func NewDeleteDeviceNotFound() *DeleteDeviceNotFound {
	return &DeleteDeviceNotFound{}
}

/* DeleteDeviceNotFound describes a response with status code 404, with default header values.

not found
*/
type DeleteDeviceNotFound struct {
}

func (o *DeleteDeviceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /devices/{id}][%d] deleteDeviceNotFound ", 404)
}

func (o *DeleteDeviceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceUnprocessableEntity creates a DeleteDeviceUnprocessableEntity with default headers values
func NewDeleteDeviceUnprocessableEntity() *DeleteDeviceUnprocessableEntity {
	return &DeleteDeviceUnprocessableEntity{}
}

/* DeleteDeviceUnprocessableEntity describes a response with status code 422, with default header values.

unprocessable entity
*/
type DeleteDeviceUnprocessableEntity struct {
}

func (o *DeleteDeviceUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /devices/{id}][%d] deleteDeviceUnprocessableEntity ", 422)
}

func (o *DeleteDeviceUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
