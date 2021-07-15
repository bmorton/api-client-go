// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ElementDynamicSelectEntity element dynamic select entity
//
// swagger:model ElementDynamicSelectEntity
type ElementDynamicSelectEntity struct {

	// async url
	AsyncURL string `json:"async_url,omitempty"`

	// clearable
	Clearable bool `json:"clearable,omitempty"`

	// default value
	DefaultValue interface{} `json:"default_value,omitempty"`

	// is multi
	IsMulti bool `json:"is_multi,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// options
	Options interface{} `json:"options,omitempty"`

	// placeholder
	Placeholder string `json:"placeholder,omitempty"`

	// required
	Required bool `json:"required,omitempty"`
}

// Validate validates this element dynamic select entity
func (m *ElementDynamicSelectEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this element dynamic select entity based on context it is used
func (m *ElementDynamicSelectEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ElementDynamicSelectEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElementDynamicSelectEntity) UnmarshalBinary(b []byte) error {
	var res ElementDynamicSelectEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
