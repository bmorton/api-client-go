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

// ExecutionStepExecutionEntity execution step execution entity
//
// swagger:model ExecutionStepExecutionEntity
type ExecutionStepExecutionEntity struct {

	// data
	Data string `json:"data,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// performed at
	PerformedAt string `json:"performed_at,omitempty"`

	// performed by
	PerformedBy *ActorEntity `json:"performed_by,omitempty"`

	// scheduled for
	ScheduledFor string `json:"scheduled_for,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this execution step execution entity
func (m *ExecutionStepExecutionEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePerformedBy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExecutionStepExecutionEntity) validatePerformedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.PerformedBy) { // not required
		return nil
	}

	if m.PerformedBy != nil {
		if err := m.PerformedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("performed_by")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this execution step execution entity based on the context it is used
func (m *ExecutionStepExecutionEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePerformedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExecutionStepExecutionEntity) contextValidatePerformedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.PerformedBy != nil {
		if err := m.PerformedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("performed_by")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExecutionStepExecutionEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExecutionStepExecutionEntity) UnmarshalBinary(b []byte) error {
	var res ExecutionStepExecutionEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}