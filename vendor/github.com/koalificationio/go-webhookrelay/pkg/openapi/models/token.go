// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Token Token
// swagger:model Token
type Token struct {

	// Enable/disable API access. If disabled, it can only be used to open tunnel connections or forward webhooks but not change any existing configuration
	// Enum: [disabled enabled]
	APIAccess string `json:"api_access,omitempty"`

	// created at
	// Read Only: true
	CreatedAt int64 `json:"created_at,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// last login
	// Read Only: true
	LastLogin string `json:"last_login,omitempty"`

	// scopes
	Scopes *TokenScopes `json:"scopes,omitempty"`

	// updated at
	// Read Only: true
	UpdatedAt int64 `json:"updated_at,omitempty"`
}

// Validate validates this token
func (m *Token) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScopes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var tokenTypeAPIAccessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["disabled","enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tokenTypeAPIAccessPropEnum = append(tokenTypeAPIAccessPropEnum, v)
	}
}

const (

	// TokenAPIAccessDisabled captures enum value "disabled"
	TokenAPIAccessDisabled string = "disabled"

	// TokenAPIAccessEnabled captures enum value "enabled"
	TokenAPIAccessEnabled string = "enabled"
)

// prop value enum
func (m *Token) validateAPIAccessEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tokenTypeAPIAccessPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Token) validateAPIAccess(formats strfmt.Registry) error {

	if swag.IsZero(m.APIAccess) { // not required
		return nil
	}

	// value enum
	if err := m.validateAPIAccessEnum("api_access", "body", m.APIAccess); err != nil {
		return err
	}

	return nil
}

func (m *Token) validateScopes(formats strfmt.Registry) error {

	if swag.IsZero(m.Scopes) { // not required
		return nil
	}

	if m.Scopes != nil {
		if err := m.Scopes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scopes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Token) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Token) UnmarshalBinary(b []byte) error {
	var res Token
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
