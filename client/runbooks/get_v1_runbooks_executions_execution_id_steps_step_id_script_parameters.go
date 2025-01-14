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
)

// NewGetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams creates a new GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams() *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams {
	return &GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParamsWithTimeout creates a new GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams object
// with the ability to set a timeout on a request.
func NewGetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParamsWithTimeout(timeout time.Duration) *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams {
	return &GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams{
		timeout: timeout,
	}
}

// NewGetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParamsWithContext creates a new GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams object
// with the ability to set a context for a request.
func NewGetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParamsWithContext(ctx context.Context) *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams {
	return &GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams{
		Context: ctx,
	}
}

// NewGetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParamsWithHTTPClient creates a new GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParamsWithHTTPClient(client *http.Client) *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams {
	return &GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams{
		HTTPClient: client,
	}
}

/* GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams contains all the parameters to send to the API endpoint
   for the get v1 runbooks executions execution Id steps step Id script operation.

   Typically these are written to a http.Request.
*/
type GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams struct {

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

// WithDefaults hydrates default values in the get v1 runbooks executions execution Id steps step Id script params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams) WithDefaults() *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 runbooks executions execution Id steps step Id script params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 runbooks executions execution Id steps step Id script params
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams) WithTimeout(timeout time.Duration) *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 runbooks executions execution Id steps step Id script params
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 runbooks executions execution Id steps step Id script params
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams) WithContext(ctx context.Context) *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 runbooks executions execution Id steps step Id script params
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 runbooks executions execution Id steps step Id script params
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams) WithHTTPClient(client *http.Client) *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 runbooks executions execution Id steps step Id script params
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExecutionID adds the executionID to the get v1 runbooks executions execution Id steps step Id script params
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams) WithExecutionID(executionID int32) *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the get v1 runbooks executions execution Id steps step Id script params
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams) SetExecutionID(executionID int32) {
	o.ExecutionID = executionID
}

// WithStepID adds the stepID to the get v1 runbooks executions execution Id steps step Id script params
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams) WithStepID(stepID int32) *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams {
	o.SetStepID(stepID)
	return o
}

// SetStepID adds the stepId to the get v1 runbooks executions execution Id steps step Id script params
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams) SetStepID(stepID int32) {
	o.StepID = stepID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDScriptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
