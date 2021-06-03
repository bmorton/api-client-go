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

// NewGetV1RunbooksExecutionsExecutionIDVotesStatusParams creates a new GetV1RunbooksExecutionsExecutionIDVotesStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1RunbooksExecutionsExecutionIDVotesStatusParams() *GetV1RunbooksExecutionsExecutionIDVotesStatusParams {
	return &GetV1RunbooksExecutionsExecutionIDVotesStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1RunbooksExecutionsExecutionIDVotesStatusParamsWithTimeout creates a new GetV1RunbooksExecutionsExecutionIDVotesStatusParams object
// with the ability to set a timeout on a request.
func NewGetV1RunbooksExecutionsExecutionIDVotesStatusParamsWithTimeout(timeout time.Duration) *GetV1RunbooksExecutionsExecutionIDVotesStatusParams {
	return &GetV1RunbooksExecutionsExecutionIDVotesStatusParams{
		timeout: timeout,
	}
}

// NewGetV1RunbooksExecutionsExecutionIDVotesStatusParamsWithContext creates a new GetV1RunbooksExecutionsExecutionIDVotesStatusParams object
// with the ability to set a context for a request.
func NewGetV1RunbooksExecutionsExecutionIDVotesStatusParamsWithContext(ctx context.Context) *GetV1RunbooksExecutionsExecutionIDVotesStatusParams {
	return &GetV1RunbooksExecutionsExecutionIDVotesStatusParams{
		Context: ctx,
	}
}

// NewGetV1RunbooksExecutionsExecutionIDVotesStatusParamsWithHTTPClient creates a new GetV1RunbooksExecutionsExecutionIDVotesStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1RunbooksExecutionsExecutionIDVotesStatusParamsWithHTTPClient(client *http.Client) *GetV1RunbooksExecutionsExecutionIDVotesStatusParams {
	return &GetV1RunbooksExecutionsExecutionIDVotesStatusParams{
		HTTPClient: client,
	}
}

/* GetV1RunbooksExecutionsExecutionIDVotesStatusParams contains all the parameters to send to the API endpoint
   for the get v1 runbooks executions execution Id votes status operation.

   Typically these are written to a http.Request.
*/
type GetV1RunbooksExecutionsExecutionIDVotesStatusParams struct {

	// ExecutionID.
	//
	// Format: int32
	ExecutionID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 runbooks executions execution Id votes status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusParams) WithDefaults() *GetV1RunbooksExecutionsExecutionIDVotesStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 runbooks executions execution Id votes status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 runbooks executions execution Id votes status params
func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusParams) WithTimeout(timeout time.Duration) *GetV1RunbooksExecutionsExecutionIDVotesStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 runbooks executions execution Id votes status params
func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 runbooks executions execution Id votes status params
func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusParams) WithContext(ctx context.Context) *GetV1RunbooksExecutionsExecutionIDVotesStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 runbooks executions execution Id votes status params
func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 runbooks executions execution Id votes status params
func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusParams) WithHTTPClient(client *http.Client) *GetV1RunbooksExecutionsExecutionIDVotesStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 runbooks executions execution Id votes status params
func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExecutionID adds the executionID to the get v1 runbooks executions execution Id votes status params
func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusParams) WithExecutionID(executionID int32) *GetV1RunbooksExecutionsExecutionIDVotesStatusParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the get v1 runbooks executions execution Id votes status params
func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusParams) SetExecutionID(executionID int32) {
	o.ExecutionID = executionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param execution_id
	if err := r.SetPathParam("execution_id", swag.FormatInt32(o.ExecutionID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}