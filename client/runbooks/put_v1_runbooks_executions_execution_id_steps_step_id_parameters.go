// Code generated by go-swagger; DO NOT EDIT.

package runbooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/firehydrant/api-client-go/models"
)

// NewPutV1RunbooksExecutionsExecutionIDStepsStepIDParams creates a new PutV1RunbooksExecutionsExecutionIDStepsStepIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutV1RunbooksExecutionsExecutionIDStepsStepIDParams() *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams {
	return &PutV1RunbooksExecutionsExecutionIDStepsStepIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutV1RunbooksExecutionsExecutionIDStepsStepIDParamsWithTimeout creates a new PutV1RunbooksExecutionsExecutionIDStepsStepIDParams object
// with the ability to set a timeout on a request.
func NewPutV1RunbooksExecutionsExecutionIDStepsStepIDParamsWithTimeout(timeout time.Duration) *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams {
	return &PutV1RunbooksExecutionsExecutionIDStepsStepIDParams{
		timeout: timeout,
	}
}

// NewPutV1RunbooksExecutionsExecutionIDStepsStepIDParamsWithContext creates a new PutV1RunbooksExecutionsExecutionIDStepsStepIDParams object
// with the ability to set a context for a request.
func NewPutV1RunbooksExecutionsExecutionIDStepsStepIDParamsWithContext(ctx context.Context) *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams {
	return &PutV1RunbooksExecutionsExecutionIDStepsStepIDParams{
		Context: ctx,
	}
}

// NewPutV1RunbooksExecutionsExecutionIDStepsStepIDParamsWithHTTPClient creates a new PutV1RunbooksExecutionsExecutionIDStepsStepIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutV1RunbooksExecutionsExecutionIDStepsStepIDParamsWithHTTPClient(client *http.Client) *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams {
	return &PutV1RunbooksExecutionsExecutionIDStepsStepIDParams{
		HTTPClient: client,
	}
}

/* PutV1RunbooksExecutionsExecutionIDStepsStepIDParams contains all the parameters to send to the API endpoint
   for the put v1 runbooks executions execution Id steps step Id operation.

   Typically these are written to a http.Request.
*/
type PutV1RunbooksExecutionsExecutionIDStepsStepIDParams struct {

	// V1RunbooksExecutionsExecutionIDSteps.
	V1RunbooksExecutionsExecutionIDSteps *models.PutV1RunbooksExecutionsExecutionIDSteps

	// ExecutionID.
	//
	// Format: int32
	ExecutionID int32

	// StepID.
	//
	// Format: int32
	StepID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put v1 runbooks executions execution Id steps step Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams) WithDefaults() *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put v1 runbooks executions execution Id steps step Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put v1 runbooks executions execution Id steps step Id params
func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams) WithTimeout(timeout time.Duration) *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put v1 runbooks executions execution Id steps step Id params
func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put v1 runbooks executions execution Id steps step Id params
func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams) WithContext(ctx context.Context) *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put v1 runbooks executions execution Id steps step Id params
func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put v1 runbooks executions execution Id steps step Id params
func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams) WithHTTPClient(client *http.Client) *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put v1 runbooks executions execution Id steps step Id params
func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1RunbooksExecutionsExecutionIDSteps adds the v1RunbooksExecutionsExecutionIDSteps to the put v1 runbooks executions execution Id steps step Id params
func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams) WithV1RunbooksExecutionsExecutionIDSteps(v1RunbooksExecutionsExecutionIDSteps *models.PutV1RunbooksExecutionsExecutionIDSteps) *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams {
	o.SetV1RunbooksExecutionsExecutionIDSteps(v1RunbooksExecutionsExecutionIDSteps)
	return o
}

// SetV1RunbooksExecutionsExecutionIDSteps adds the v1RunbooksExecutionsExecutionIdSteps to the put v1 runbooks executions execution Id steps step Id params
func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams) SetV1RunbooksExecutionsExecutionIDSteps(v1RunbooksExecutionsExecutionIDSteps *models.PutV1RunbooksExecutionsExecutionIDSteps) {
	o.V1RunbooksExecutionsExecutionIDSteps = v1RunbooksExecutionsExecutionIDSteps
}

// WithExecutionID adds the executionID to the put v1 runbooks executions execution Id steps step Id params
func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams) WithExecutionID(executionID int32) *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the put v1 runbooks executions execution Id steps step Id params
func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams) SetExecutionID(executionID int32) {
	o.ExecutionID = executionID
}

// WithStepID adds the stepID to the put v1 runbooks executions execution Id steps step Id params
func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams) WithStepID(stepID int32) *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams {
	o.SetStepID(stepID)
	return o
}

// SetStepID adds the stepId to the put v1 runbooks executions execution Id steps step Id params
func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams) SetStepID(stepID int32) {
	o.StepID = stepID
}

// WriteToRequest writes these params to a swagger request
func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1RunbooksExecutionsExecutionIDSteps != nil {
		if err := r.SetBodyParam(o.V1RunbooksExecutionsExecutionIDSteps); err != nil {
			return err
		}
	}

	// path param execution_id
	if err := r.SetPathParam("execution_id", swag.FormatInt32(o.ExecutionID)); err != nil {
		return err
	}

	// path param step_id
	if err := r.SetPathParam("step_id", swag.FormatInt32(o.StepID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
