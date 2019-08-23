// Code generated by go-swagger; DO NOT EDIT.

package services

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

	models "github.com/firehydrant/api-client-go/models"
)

// NewPatchV1ServicesServiceIDParams creates a new PatchV1ServicesServiceIDParams object
// with the default values initialized.
func NewPatchV1ServicesServiceIDParams() *PatchV1ServicesServiceIDParams {
	var ()
	return &PatchV1ServicesServiceIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1ServicesServiceIDParamsWithTimeout creates a new PatchV1ServicesServiceIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchV1ServicesServiceIDParamsWithTimeout(timeout time.Duration) *PatchV1ServicesServiceIDParams {
	var ()
	return &PatchV1ServicesServiceIDParams{

		timeout: timeout,
	}
}

// NewPatchV1ServicesServiceIDParamsWithContext creates a new PatchV1ServicesServiceIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchV1ServicesServiceIDParamsWithContext(ctx context.Context) *PatchV1ServicesServiceIDParams {
	var ()
	return &PatchV1ServicesServiceIDParams{

		Context: ctx,
	}
}

// NewPatchV1ServicesServiceIDParamsWithHTTPClient creates a new PatchV1ServicesServiceIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchV1ServicesServiceIDParamsWithHTTPClient(client *http.Client) *PatchV1ServicesServiceIDParams {
	var ()
	return &PatchV1ServicesServiceIDParams{
		HTTPClient: client,
	}
}

/*PatchV1ServicesServiceIDParams contains all the parameters to send to the API endpoint
for the patch v1 services service Id operation typically these are written to a http.Request
*/
type PatchV1ServicesServiceIDParams struct {

	/*V1Services*/
	V1Services *models.PatchV1Services
	/*ServiceID*/
	ServiceID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch v1 services service Id params
func (o *PatchV1ServicesServiceIDParams) WithTimeout(timeout time.Duration) *PatchV1ServicesServiceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 services service Id params
func (o *PatchV1ServicesServiceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 services service Id params
func (o *PatchV1ServicesServiceIDParams) WithContext(ctx context.Context) *PatchV1ServicesServiceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 services service Id params
func (o *PatchV1ServicesServiceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 services service Id params
func (o *PatchV1ServicesServiceIDParams) WithHTTPClient(client *http.Client) *PatchV1ServicesServiceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 services service Id params
func (o *PatchV1ServicesServiceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1Services adds the v1Services to the patch v1 services service Id params
func (o *PatchV1ServicesServiceIDParams) WithV1Services(v1Services *models.PatchV1Services) *PatchV1ServicesServiceIDParams {
	o.SetV1Services(v1Services)
	return o
}

// SetV1Services adds the v1Services to the patch v1 services service Id params
func (o *PatchV1ServicesServiceIDParams) SetV1Services(v1Services *models.PatchV1Services) {
	o.V1Services = v1Services
}

// WithServiceID adds the serviceID to the patch v1 services service Id params
func (o *PatchV1ServicesServiceIDParams) WithServiceID(serviceID int32) *PatchV1ServicesServiceIDParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the patch v1 services service Id params
func (o *PatchV1ServicesServiceIDParams) SetServiceID(serviceID int32) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1ServicesServiceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.V1Services != nil {
		if err := r.SetBodyParam(o.V1Services); err != nil {
			return err
		}
	}

	// path param service_id
	if err := r.SetPathParam("service_id", swag.FormatInt32(o.ServiceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}