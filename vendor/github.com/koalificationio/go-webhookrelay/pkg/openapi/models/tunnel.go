// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Tunnel Tunnel
// swagger:model Tunnel
type Tunnel struct {

	// auth
	Auth *TunnelAuth `json:"auth,omitempty"`

	// created at
	CreatedAt int64 `json:"created_at,omitempty"`

	// crypto
	Crypto string `json:"crypto,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// destination
	Destination string `json:"destination,omitempty"`

	// group
	Group string `json:"group,omitempty"`

	// host
	Host string `json:"host,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// protocol
	Protocol string `json:"protocol,omitempty"`

	// updated at
	UpdatedAt int64 `json:"updated_at,omitempty"`
}

// Validate validates this tunnel
func (m *Tunnel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tunnel) validateAuth(formats strfmt.Registry) error {

	if swag.IsZero(m.Auth) { // not required
		return nil
	}

	if m.Auth != nil {
		if err := m.Auth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Tunnel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tunnel) UnmarshalBinary(b []byte) error {
	var res Tunnel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
