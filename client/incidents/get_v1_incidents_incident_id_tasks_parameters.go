// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetV1IncidentsIncidentIDTasksParams creates a new GetV1IncidentsIncidentIDTasksParams object
// with the default values initialized.
func NewGetV1IncidentsIncidentIDTasksParams() *GetV1IncidentsIncidentIDTasksParams {
	var ()
	return &GetV1IncidentsIncidentIDTasksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IncidentsIncidentIDTasksParamsWithTimeout creates a new GetV1IncidentsIncidentIDTasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetV1IncidentsIncidentIDTasksParamsWithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDTasksParams {
	var ()
	return &GetV1IncidentsIncidentIDTasksParams{

		timeout: timeout,
	}
}

// NewGetV1IncidentsIncidentIDTasksParamsWithContext creates a new GetV1IncidentsIncidentIDTasksParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetV1IncidentsIncidentIDTasksParamsWithContext(ctx context.Context) *GetV1IncidentsIncidentIDTasksParams {
	var ()
	return &GetV1IncidentsIncidentIDTasksParams{

		Context: ctx,
	}
}

// NewGetV1IncidentsIncidentIDTasksParamsWithHTTPClient creates a new GetV1IncidentsIncidentIDTasksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetV1IncidentsIncidentIDTasksParamsWithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDTasksParams {
	var ()
	return &GetV1IncidentsIncidentIDTasksParams{
		HTTPClient: client,
	}
}

/*GetV1IncidentsIncidentIDTasksParams contains all the parameters to send to the API endpoint
for the get v1 incidents incident Id tasks operation typically these are written to a http.Request
*/
type GetV1IncidentsIncidentIDTasksParams struct {

	/*IncidentID*/
	IncidentID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v1 incidents incident Id tasks params
func (o *GetV1IncidentsIncidentIDTasksParams) WithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 incidents incident Id tasks params
func (o *GetV1IncidentsIncidentIDTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 incidents incident Id tasks params
func (o *GetV1IncidentsIncidentIDTasksParams) WithContext(ctx context.Context) *GetV1IncidentsIncidentIDTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 incidents incident Id tasks params
func (o *GetV1IncidentsIncidentIDTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 incidents incident Id tasks params
func (o *GetV1IncidentsIncidentIDTasksParams) WithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 incidents incident Id tasks params
func (o *GetV1IncidentsIncidentIDTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentID adds the incidentID to the get v1 incidents incident Id tasks params
func (o *GetV1IncidentsIncidentIDTasksParams) WithIncidentID(incidentID int32) *GetV1IncidentsIncidentIDTasksParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the get v1 incidents incident Id tasks params
func (o *GetV1IncidentsIncidentIDTasksParams) SetIncidentID(incidentID int32) {
	o.IncidentID = incidentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IncidentsIncidentIDTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param incident_id
	if err := r.SetPathParam("incident_id", swag.FormatInt32(o.IncidentID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}