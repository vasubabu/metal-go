// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// FindConnectionPortEventsReader is a Reader for the FindConnectionPortEvents structure.
type FindConnectionPortEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConnectionPortEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConnectionPortEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindConnectionPortEventsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFindConnectionPortEventsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindConnectionPortEventsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindConnectionPortEventsOK creates a FindConnectionPortEventsOK with default headers values
func NewFindConnectionPortEventsOK() *FindConnectionPortEventsOK {
	return &FindConnectionPortEventsOK{}
}

/* FindConnectionPortEventsOK describes a response with status code 200, with default header values.

ok
*/
type FindConnectionPortEventsOK struct {
	Payload *types.Event
}

func (o *FindConnectionPortEventsOK) Error() string {
	return fmt.Sprintf("[GET /connections/{connection_id}/ports/{id}/events][%d] findConnectionPortEventsOK  %+v", 200, o.Payload)
}
func (o *FindConnectionPortEventsOK) GetPayload() *types.Event {
	return o.Payload
}

func (o *FindConnectionPortEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.Event)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConnectionPortEventsUnauthorized creates a FindConnectionPortEventsUnauthorized with default headers values
func NewFindConnectionPortEventsUnauthorized() *FindConnectionPortEventsUnauthorized {
	return &FindConnectionPortEventsUnauthorized{}
}

/* FindConnectionPortEventsUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type FindConnectionPortEventsUnauthorized struct {
}

func (o *FindConnectionPortEventsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /connections/{connection_id}/ports/{id}/events][%d] findConnectionPortEventsUnauthorized ", 401)
}

func (o *FindConnectionPortEventsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindConnectionPortEventsForbidden creates a FindConnectionPortEventsForbidden with default headers values
func NewFindConnectionPortEventsForbidden() *FindConnectionPortEventsForbidden {
	return &FindConnectionPortEventsForbidden{}
}

/* FindConnectionPortEventsForbidden describes a response with status code 403, with default header values.

forbidden
*/
type FindConnectionPortEventsForbidden struct {
}

func (o *FindConnectionPortEventsForbidden) Error() string {
	return fmt.Sprintf("[GET /connections/{connection_id}/ports/{id}/events][%d] findConnectionPortEventsForbidden ", 403)
}

func (o *FindConnectionPortEventsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindConnectionPortEventsNotFound creates a FindConnectionPortEventsNotFound with default headers values
func NewFindConnectionPortEventsNotFound() *FindConnectionPortEventsNotFound {
	return &FindConnectionPortEventsNotFound{}
}

/* FindConnectionPortEventsNotFound describes a response with status code 404, with default header values.

not found
*/
type FindConnectionPortEventsNotFound struct {
}

func (o *FindConnectionPortEventsNotFound) Error() string {
	return fmt.Sprintf("[GET /connections/{connection_id}/ports/{id}/events][%d] findConnectionPortEventsNotFound ", 404)
}

func (o *FindConnectionPortEventsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
