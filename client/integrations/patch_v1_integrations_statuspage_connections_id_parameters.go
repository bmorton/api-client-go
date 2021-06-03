// Code generated by go-swagger; DO NOT EDIT.

package integrations

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

	"github.com/firehydrant/api-client-go/models"
)

// NewPatchV1IntegrationsStatuspageConnectionsIDParams creates a new PatchV1IntegrationsStatuspageConnectionsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1IntegrationsStatuspageConnectionsIDParams() *PatchV1IntegrationsStatuspageConnectionsIDParams {
	return &PatchV1IntegrationsStatuspageConnectionsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1IntegrationsStatuspageConnectionsIDParamsWithTimeout creates a new PatchV1IntegrationsStatuspageConnectionsIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1IntegrationsStatuspageConnectionsIDParamsWithTimeout(timeout time.Duration) *PatchV1IntegrationsStatuspageConnectionsIDParams {
	return &PatchV1IntegrationsStatuspageConnectionsIDParams{
		timeout: timeout,
	}
}

// NewPatchV1IntegrationsStatuspageConnectionsIDParamsWithContext creates a new PatchV1IntegrationsStatuspageConnectionsIDParams object
// with the ability to set a context for a request.
func NewPatchV1IntegrationsStatuspageConnectionsIDParamsWithContext(ctx context.Context) *PatchV1IntegrationsStatuspageConnectionsIDParams {
	return &PatchV1IntegrationsStatuspageConnectionsIDParams{
		Context: ctx,
	}
}

// NewPatchV1IntegrationsStatuspageConnectionsIDParamsWithHTTPClient creates a new PatchV1IntegrationsStatuspageConnectionsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1IntegrationsStatuspageConnectionsIDParamsWithHTTPClient(client *http.Client) *PatchV1IntegrationsStatuspageConnectionsIDParams {
	return &PatchV1IntegrationsStatuspageConnectionsIDParams{
		HTTPClient: client,
	}
}

/* PatchV1IntegrationsStatuspageConnectionsIDParams contains all the parameters to send to the API endpoint
   for the patch v1 integrations statuspage connections Id operation.

   Typically these are written to a http.Request.
*/
type PatchV1IntegrationsStatuspageConnectionsIDParams struct {

	// V1IntegrationsStatuspageConnections.
	V1IntegrationsStatuspageConnections *models.PatchV1IntegrationsStatuspageConnections

	/* ID.

	   Connection UUID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 integrations statuspage connections Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1IntegrationsStatuspageConnectionsIDParams) WithDefaults() *PatchV1IntegrationsStatuspageConnectionsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 integrations statuspage connections Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1IntegrationsStatuspageConnectionsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 integrations statuspage connections Id params
func (o *PatchV1IntegrationsStatuspageConnectionsIDParams) WithTimeout(timeout time.Duration) *PatchV1IntegrationsStatuspageConnectionsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 integrations statuspage connections Id params
func (o *PatchV1IntegrationsStatuspageConnectionsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 integrations statuspage connections Id params
func (o *PatchV1IntegrationsStatuspageConnectionsIDParams) WithContext(ctx context.Context) *PatchV1IntegrationsStatuspageConnectionsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 integrations statuspage connections Id params
func (o *PatchV1IntegrationsStatuspageConnectionsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 integrations statuspage connections Id params
func (o *PatchV1IntegrationsStatuspageConnectionsIDParams) WithHTTPClient(client *http.Client) *PatchV1IntegrationsStatuspageConnectionsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 integrations statuspage connections Id params
func (o *PatchV1IntegrationsStatuspageConnectionsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1IntegrationsStatuspageConnections adds the v1IntegrationsStatuspageConnections to the patch v1 integrations statuspage connections Id params
func (o *PatchV1IntegrationsStatuspageConnectionsIDParams) WithV1IntegrationsStatuspageConnections(v1IntegrationsStatuspageConnections *models.PatchV1IntegrationsStatuspageConnections) *PatchV1IntegrationsStatuspageConnectionsIDParams {
	o.SetV1IntegrationsStatuspageConnections(v1IntegrationsStatuspageConnections)
	return o
}

// SetV1IntegrationsStatuspageConnections adds the v1IntegrationsStatuspageConnections to the patch v1 integrations statuspage connections Id params
func (o *PatchV1IntegrationsStatuspageConnectionsIDParams) SetV1IntegrationsStatuspageConnections(v1IntegrationsStatuspageConnections *models.PatchV1IntegrationsStatuspageConnections) {
	o.V1IntegrationsStatuspageConnections = v1IntegrationsStatuspageConnections
}

// WithID adds the id to the patch v1 integrations statuspage connections Id params
func (o *PatchV1IntegrationsStatuspageConnectionsIDParams) WithID(id string) *PatchV1IntegrationsStatuspageConnectionsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch v1 integrations statuspage connections Id params
func (o *PatchV1IntegrationsStatuspageConnectionsIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1IntegrationsStatuspageConnectionsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1IntegrationsStatuspageConnections != nil {
		if err := r.SetBodyParam(o.V1IntegrationsStatuspageConnections); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}