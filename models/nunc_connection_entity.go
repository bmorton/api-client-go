// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NuncConnectionEntity Retrieve the information displayed as part of your FireHydrant hosted status page.
//
// swagger:model NuncConnectionEntity
type NuncConnectionEntity struct {

	// cname
	Cname string `json:"cname,omitempty"`

	// company name
	CompanyName string `json:"company_name,omitempty"`

	// company tos url
	CompanyTosURL string `json:"company_tos_url,omitempty"`

	// company website
	CompanyWebsite string `json:"company_website,omitempty"`

	// components
	Components *NuncComponentEntity `json:"components,omitempty"`

	// conditions
	Conditions *NuncConditionEntity `json:"conditions,omitempty"`

	// domain
	Domain string `json:"domain,omitempty"`

	// greeting body
	GreetingBody string `json:"greeting_body,omitempty"`

	// greeting title
	GreetingTitle string `json:"greeting_title,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// logo
	Logo *NuncLogoEntity `json:"logo,omitempty"`

	// operational message
	OperationalMessage string `json:"operational_message,omitempty"`
}

// Validate validates this nunc connection entity
func (m *NuncConnectionEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NuncConnectionEntity) validateComponents(formats strfmt.Registry) error {
	if swag.IsZero(m.Components) { // not required
		return nil
	}

	if m.Components != nil {
		if err := m.Components.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("components")
			}
			return err
		}
	}

	return nil
}

func (m *NuncConnectionEntity) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	if m.Conditions != nil {
		if err := m.Conditions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conditions")
			}
			return err
		}
	}

	return nil
}

func (m *NuncConnectionEntity) validateLogo(formats strfmt.Registry) error {
	if swag.IsZero(m.Logo) { // not required
		return nil
	}

	if m.Logo != nil {
		if err := m.Logo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nunc connection entity based on the context it is used
func (m *NuncConnectionEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NuncConnectionEntity) contextValidateComponents(ctx context.Context, formats strfmt.Registry) error {

	if m.Components != nil {
		if err := m.Components.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("components")
			}
			return err
		}
	}

	return nil
}

func (m *NuncConnectionEntity) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	if m.Conditions != nil {
		if err := m.Conditions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conditions")
			}
			return err
		}
	}

	return nil
}

func (m *NuncConnectionEntity) contextValidateLogo(ctx context.Context, formats strfmt.Registry) error {

	if m.Logo != nil {
		if err := m.Logo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NuncConnectionEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NuncConnectionEntity) UnmarshalBinary(b []byte) error {
	var res NuncConnectionEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}