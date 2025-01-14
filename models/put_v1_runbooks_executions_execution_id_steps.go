// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PutV1RunbooksExecutionsExecutionIDSteps Updates a runbook step execution, especially for changing the state of a step execution.
//
// swagger:model putV1RunbooksExecutionsExecutionIdSteps
type PutV1RunbooksExecutionsExecutionIDSteps struct {

	// Data for execution of this step
	Data interface{} `json:"data,omitempty"`

	// repeats at
	// Format: date-time
	RepeatsAt strfmt.DateTime `json:"repeats_at,omitempty"`

	// schedule for
	// Format: date-time
	ScheduleFor strfmt.DateTime `json:"schedule_for,omitempty"`

	// state
	// Required: true
	State *string `json:"state"`
}

// Validate validates this put v1 runbooks executions execution Id steps
func (m *PutV1RunbooksExecutionsExecutionIDSteps) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepeatsAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduleFor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutV1RunbooksExecutionsExecutionIDSteps) validateRepeatsAt(formats strfmt.Registry) error {
	if swag.IsZero(m.RepeatsAt) { // not required
		return nil
	}

	if err := validate.FormatOf("repeats_at", "body", "date-time", m.RepeatsAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PutV1RunbooksExecutionsExecutionIDSteps) validateScheduleFor(formats strfmt.Registry) error {
	if swag.IsZero(m.ScheduleFor) { // not required
		return nil
	}

	if err := validate.FormatOf("schedule_for", "body", "date-time", m.ScheduleFor.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PutV1RunbooksExecutionsExecutionIDSteps) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this put v1 runbooks executions execution Id steps based on context it is used
func (m *PutV1RunbooksExecutionsExecutionIDSteps) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PutV1RunbooksExecutionsExecutionIDSteps) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutV1RunbooksExecutionsExecutionIDSteps) UnmarshalBinary(b []byte) error {
	var res PutV1RunbooksExecutionsExecutionIDSteps
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
