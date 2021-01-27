// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServerUnavailableInfo server unavailable info
//
// swagger:model ServerUnavailableInfo
type ServerUnavailableInfo struct {

	// available
	Available bool `json:"available,omitempty"`

	// facility
	Facility string `json:"facility,omitempty"`

	// plan
	Plan string `json:"plan,omitempty"`

	// quantity
	Quantity string `json:"quantity,omitempty"`
}

// Validate validates this server unavailable info
func (m *ServerUnavailableInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this server unavailable info based on context it is used
func (m *ServerUnavailableInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerUnavailableInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerUnavailableInfo) UnmarshalBinary(b []byte) error {
	var res ServerUnavailableInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
