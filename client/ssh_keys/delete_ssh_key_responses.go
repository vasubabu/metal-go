// Code generated by go-swagger; DO NOT EDIT.

package ssh_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteSSHKeyReader is a Reader for the DeleteSSHKey structure.
type DeleteSSHKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSSHKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteSSHKeyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteSSHKeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteSSHKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSSHKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSSHKeyNoContent creates a DeleteSSHKeyNoContent with default headers values
func NewDeleteSSHKeyNoContent() *DeleteSSHKeyNoContent {
	return &DeleteSSHKeyNoContent{}
}

/* DeleteSSHKeyNoContent describes a response with status code 204, with default header values.

no content
*/
type DeleteSSHKeyNoContent struct {
}

func (o *DeleteSSHKeyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ssh-keys/{id}][%d] deleteSshKeyNoContent ", 204)
}

func (o *DeleteSSHKeyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSSHKeyUnauthorized creates a DeleteSSHKeyUnauthorized with default headers values
func NewDeleteSSHKeyUnauthorized() *DeleteSSHKeyUnauthorized {
	return &DeleteSSHKeyUnauthorized{}
}

/* DeleteSSHKeyUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type DeleteSSHKeyUnauthorized struct {
}

func (o *DeleteSSHKeyUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /ssh-keys/{id}][%d] deleteSshKeyUnauthorized ", 401)
}

func (o *DeleteSSHKeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSSHKeyForbidden creates a DeleteSSHKeyForbidden with default headers values
func NewDeleteSSHKeyForbidden() *DeleteSSHKeyForbidden {
	return &DeleteSSHKeyForbidden{}
}

/* DeleteSSHKeyForbidden describes a response with status code 403, with default header values.

forbidden
*/
type DeleteSSHKeyForbidden struct {
}

func (o *DeleteSSHKeyForbidden) Error() string {
	return fmt.Sprintf("[DELETE /ssh-keys/{id}][%d] deleteSshKeyForbidden ", 403)
}

func (o *DeleteSSHKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSSHKeyNotFound creates a DeleteSSHKeyNotFound with default headers values
func NewDeleteSSHKeyNotFound() *DeleteSSHKeyNotFound {
	return &DeleteSSHKeyNotFound{}
}

/* DeleteSSHKeyNotFound describes a response with status code 404, with default header values.

not found
*/
type DeleteSSHKeyNotFound struct {
}

func (o *DeleteSSHKeyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /ssh-keys/{id}][%d] deleteSshKeyNotFound ", 404)
}

func (o *DeleteSSHKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
