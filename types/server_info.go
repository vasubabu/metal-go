// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServerInfo server info
//
// swagger:model ServerInfo
type ServerInfo struct {

	// facility
	Facility string `json:"facility,omitempty"`

	// plan
	Plan string `json:"plan,omitempty"`

	// quantity
	Quantity string `json:"quantity,omitempty"`
}

// Validate validates this server info
func (m *ServerInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this server info based on context it is used
func (m *ServerInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerInfo) UnmarshalBinary(b []byte) error {
	var res ServerInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
