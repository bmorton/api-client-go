// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ElementInputEntity element input entity
//
// swagger:model ElementInputEntity
type ElementInputEntity struct {

	// default value
	DefaultValue string `json:"default_value,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// placeholder
	Placeholder string `json:"placeholder,omitempty"`
}

// Validate validates this element input entity
func (m *ElementInputEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this element input entity based on context it is used
func (m *ElementInputEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ElementInputEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElementInputEntity) UnmarshalBinary(b []byte) error {
	var res ElementInputEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
