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

// NewPutV1IncidentsIncidentIDResolveParams creates a new PutV1IncidentsIncidentIDResolveParams object
// with the default values initialized.
func NewPutV1IncidentsIncidentIDResolveParams() *PutV1IncidentsIncidentIDResolveParams {
	var ()
	return &PutV1IncidentsIncidentIDResolveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutV1IncidentsIncidentIDResolveParamsWithTimeout creates a new PutV1IncidentsIncidentIDResolveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutV1IncidentsIncidentIDResolveParamsWithTimeout(timeout time.Duration) *PutV1IncidentsIncidentIDResolveParams {
	var ()
	return &PutV1IncidentsIncidentIDResolveParams{

		timeout: timeout,
	}
}

// NewPutV1IncidentsIncidentIDResolveParamsWithContext creates a new PutV1IncidentsIncidentIDResolveParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutV1IncidentsIncidentIDResolveParamsWithContext(ctx context.Context) *PutV1IncidentsIncidentIDResolveParams {
	var ()
	return &PutV1IncidentsIncidentIDResolveParams{

		Context: ctx,
	}
}

// NewPutV1IncidentsIncidentIDResolveParamsWithHTTPClient creates a new PutV1IncidentsIncidentIDResolveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutV1IncidentsIncidentIDResolveParamsWithHTTPClient(client *http.Client) *PutV1IncidentsIncidentIDResolveParams {
	var ()
	return &PutV1IncidentsIncidentIDResolveParams{
		HTTPClient: client,
	}
}

/*PutV1IncidentsIncidentIDResolveParams contains all the parameters to send to the API endpoint
for the put v1 incidents incident Id resolve operation typically these are written to a http.Request
*/
type PutV1IncidentsIncidentIDResolveParams struct {

	/*IncidentID*/
	IncidentID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put v1 incidents incident Id resolve params
func (o *PutV1IncidentsIncidentIDResolveParams) WithTimeout(timeout time.Duration) *PutV1IncidentsIncidentIDResolveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put v1 incidents incident Id resolve params
func (o *PutV1IncidentsIncidentIDResolveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put v1 incidents incident Id resolve params
func (o *PutV1IncidentsIncidentIDResolveParams) WithContext(ctx context.Context) *PutV1IncidentsIncidentIDResolveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put v1 incidents incident Id resolve params
func (o *PutV1IncidentsIncidentIDResolveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put v1 incidents incident Id resolve params
func (o *PutV1IncidentsIncidentIDResolveParams) WithHTTPClient(client *http.Client) *PutV1IncidentsIncidentIDResolveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put v1 incidents incident Id resolve params
func (o *PutV1IncidentsIncidentIDResolveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentID adds the incidentID to the put v1 incidents incident Id resolve params
func (o *PutV1IncidentsIncidentIDResolveParams) WithIncidentID(incidentID int32) *PutV1IncidentsIncidentIDResolveParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the put v1 incidents incident Id resolve params
func (o *PutV1IncidentsIncidentIDResolveParams) SetIncidentID(incidentID int32) {
	o.IncidentID = incidentID
}

// WriteToRequest writes these params to a swagger request
func (o *PutV1IncidentsIncidentIDResolveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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