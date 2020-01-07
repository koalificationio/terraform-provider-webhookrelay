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

// RuleParameter RuleParameter
// swagger:model RuleParameter
type RuleParameter struct {

	// name
	Name string `json:"name,omitempty"`

	// source
	// Enum: [header payload query]
	Source string `json:"source,omitempty"`
}

// Validate validates this rule parameter
func (m *RuleParameter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ruleParameterTypeSourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["header","payload","query"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ruleParameterTypeSourcePropEnum = append(ruleParameterTypeSourcePropEnum, v)
	}
}

const (

	// RuleParameterSourceHeader captures enum value "header"
	RuleParameterSourceHeader string = "header"

	// RuleParameterSourcePayload captures enum value "payload"
	RuleParameterSourcePayload string = "payload"

	// RuleParameterSourceQuery captures enum value "query"
	RuleParameterSourceQuery string = "query"
)

// prop value enum
func (m *RuleParameter) validateSourceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ruleParameterTypeSourcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RuleParameter) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	// value enum
	if err := m.validateSourceEnum("source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RuleParameter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuleParameter) UnmarshalBinary(b []byte) error {
	var res RuleParameter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
