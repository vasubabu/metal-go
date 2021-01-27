// Code generated by go-swagger; DO NOT EDIT.

package facilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// FindFacilitiesByProjectReader is a Reader for the FindFacilitiesByProject structure.
type FindFacilitiesByProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindFacilitiesByProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindFacilitiesByProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindFacilitiesByProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFindFacilitiesByProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindFacilitiesByProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindFacilitiesByProjectOK creates a FindFacilitiesByProjectOK with default headers values
func NewFindFacilitiesByProjectOK() *FindFacilitiesByProjectOK {
	return &FindFacilitiesByProjectOK{}
}

/* FindFacilitiesByProjectOK describes a response with status code 200, with default header values.

ok
*/
type FindFacilitiesByProjectOK struct {
	Payload *types.FacilityList
}

func (o *FindFacilitiesByProjectOK) Error() string {
	return fmt.Sprintf("[GET /projects/{id}/facilities][%d] findFacilitiesByProjectOK  %+v", 200, o.Payload)
}
func (o *FindFacilitiesByProjectOK) GetPayload() *types.FacilityList {
	return o.Payload
}

func (o *FindFacilitiesByProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.FacilityList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindFacilitiesByProjectUnauthorized creates a FindFacilitiesByProjectUnauthorized with default headers values
func NewFindFacilitiesByProjectUnauthorized() *FindFacilitiesByProjectUnauthorized {
	return &FindFacilitiesByProjectUnauthorized{}
}

/* FindFacilitiesByProjectUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type FindFacilitiesByProjectUnauthorized struct {
}

func (o *FindFacilitiesByProjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{id}/facilities][%d] findFacilitiesByProjectUnauthorized ", 401)
}

func (o *FindFacilitiesByProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindFacilitiesByProjectForbidden creates a FindFacilitiesByProjectForbidden with default headers values
func NewFindFacilitiesByProjectForbidden() *FindFacilitiesByProjectForbidden {
	return &FindFacilitiesByProjectForbidden{}
}

/* FindFacilitiesByProjectForbidden describes a response with status code 403, with default header values.

forbidden
*/
type FindFacilitiesByProjectForbidden struct {
}

func (o *FindFacilitiesByProjectForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{id}/facilities][%d] findFacilitiesByProjectForbidden ", 403)
}

func (o *FindFacilitiesByProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindFacilitiesByProjectNotFound creates a FindFacilitiesByProjectNotFound with default headers values
func NewFindFacilitiesByProjectNotFound() *FindFacilitiesByProjectNotFound {
	return &FindFacilitiesByProjectNotFound{}
}

/* FindFacilitiesByProjectNotFound describes a response with status code 404, with default header values.

not found
*/
type FindFacilitiesByProjectNotFound struct {
}

func (o *FindFacilitiesByProjectNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{id}/facilities][%d] findFacilitiesByProjectNotFound ", 404)
}

func (o *FindFacilitiesByProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
