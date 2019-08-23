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

// NewGetV1IncidentsIncidentIDChannelParams creates a new GetV1IncidentsIncidentIDChannelParams object
// with the default values initialized.
func NewGetV1IncidentsIncidentIDChannelParams() *GetV1IncidentsIncidentIDChannelParams {
	var ()
	return &GetV1IncidentsIncidentIDChannelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IncidentsIncidentIDChannelParamsWithTimeout creates a new GetV1IncidentsIncidentIDChannelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetV1IncidentsIncidentIDChannelParamsWithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDChannelParams {
	var ()
	return &GetV1IncidentsIncidentIDChannelParams{

		timeout: timeout,
	}
}

// NewGetV1IncidentsIncidentIDChannelParamsWithContext creates a new GetV1IncidentsIncidentIDChannelParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetV1IncidentsIncidentIDChannelParamsWithContext(ctx context.Context) *GetV1IncidentsIncidentIDChannelParams {
	var ()
	return &GetV1IncidentsIncidentIDChannelParams{

		Context: ctx,
	}
}

// NewGetV1IncidentsIncidentIDChannelParamsWithHTTPClient creates a new GetV1IncidentsIncidentIDChannelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetV1IncidentsIncidentIDChannelParamsWithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDChannelParams {
	var ()
	return &GetV1IncidentsIncidentIDChannelParams{
		HTTPClient: client,
	}
}

/*GetV1IncidentsIncidentIDChannelParams contains all the parameters to send to the API endpoint
for the get v1 incidents incident Id channel operation typically these are written to a http.Request
*/
type GetV1IncidentsIncidentIDChannelParams struct {

	/*IncidentID*/
	IncidentID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v1 incidents incident Id channel params
func (o *GetV1IncidentsIncidentIDChannelParams) WithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDChannelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 incidents incident Id channel params
func (o *GetV1IncidentsIncidentIDChannelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 incidents incident Id channel params
func (o *GetV1IncidentsIncidentIDChannelParams) WithContext(ctx context.Context) *GetV1IncidentsIncidentIDChannelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 incidents incident Id channel params
func (o *GetV1IncidentsIncidentIDChannelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 incidents incident Id channel params
func (o *GetV1IncidentsIncidentIDChannelParams) WithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDChannelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 incidents incident Id channel params
func (o *GetV1IncidentsIncidentIDChannelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentID adds the incidentID to the get v1 incidents incident Id channel params
func (o *GetV1IncidentsIncidentIDChannelParams) WithIncidentID(incidentID int32) *GetV1IncidentsIncidentIDChannelParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the get v1 incidents incident Id channel params
func (o *GetV1IncidentsIncidentIDChannelParams) SetIncidentID(incidentID int32) {
	o.IncidentID = incidentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IncidentsIncidentIDChannelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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