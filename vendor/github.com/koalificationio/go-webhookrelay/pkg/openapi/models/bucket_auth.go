// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// BucketAuth BucketAuth
// swagger:model BucketAuth
type BucketAuth struct {

	// created at
	CreatedAt int64 `json:"created_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated at
	UpdatedAt int64 `json:"updated_at,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this bucket auth
func (m *BucketAuth) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BucketAuth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BucketAuth) UnmarshalBinary(b []byte) error {
	var res BucketAuth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
