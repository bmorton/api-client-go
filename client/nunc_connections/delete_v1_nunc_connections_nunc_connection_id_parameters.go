// Code generated by go-swagger; DO NOT EDIT.

package nunc_connections

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

// NewDeleteV1NuncConnectionsNuncConnectionIDParams creates a new DeleteV1NuncConnectionsNuncConnectionIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1NuncConnectionsNuncConnectionIDParams() *DeleteV1NuncConnectionsNuncConnectionIDParams {
	return &DeleteV1NuncConnectionsNuncConnectionIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1NuncConnectionsNuncConnectionIDParamsWithTimeout creates a new DeleteV1NuncConnectionsNuncConnectionIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1NuncConnectionsNuncConnectionIDParamsWithTimeout(timeout time.Duration) *DeleteV1NuncConnectionsNuncConnectionIDParams {
	return &DeleteV1NuncConnectionsNuncConnectionIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1NuncConnectionsNuncConnectionIDParamsWithContext creates a new DeleteV1NuncConnectionsNuncConnectionIDParams object
// with the ability to set a context for a request.
func NewDeleteV1NuncConnectionsNuncConnectionIDParamsWithContext(ctx context.Context) *DeleteV1NuncConnectionsNuncConnectionIDParams {
	return &DeleteV1NuncConnectionsNuncConnectionIDParams{
		Context: ctx,
	}
}

// NewDeleteV1NuncConnectionsNuncConnectionIDParamsWithHTTPClient creates a new DeleteV1NuncConnectionsNuncConnectionIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1NuncConnectionsNuncConnectionIDParamsWithHTTPClient(client *http.Client) *DeleteV1NuncConnectionsNuncConnectionIDParams {
	return &DeleteV1NuncConnectionsNuncConnectionIDParams{
		HTTPClient: client,
	}
}

/* DeleteV1NuncConnectionsNuncConnectionIDParams contains all the parameters to send to the API endpoint
   for the delete v1 nunc connections nunc connection Id operation.

   Typically these are written to a http.Request.
*/
type DeleteV1NuncConnectionsNuncConnectionIDParams struct {

	// NuncConnectionID.
	//
	// Format: int32
	NuncConnectionID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 nunc connections nunc connection Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1NuncConnectionsNuncConnectionIDParams) WithDefaults() *DeleteV1NuncConnectionsNuncConnectionIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 nunc connections nunc connection Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1NuncConnectionsNuncConnectionIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 nunc connections nunc connection Id params
func (o *DeleteV1NuncConnectionsNuncConnectionIDParams) WithTimeout(timeout time.Duration) *DeleteV1NuncConnectionsNuncConnectionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 nunc connections nunc connection Id params
func (o *DeleteV1NuncConnectionsNuncConnectionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 nunc connections nunc connection Id params
func (o *DeleteV1NuncConnectionsNuncConnectionIDParams) WithContext(ctx context.Context) *DeleteV1NuncConnectionsNuncConnectionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 nunc connections nunc connection Id params
func (o *DeleteV1NuncConnectionsNuncConnectionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 nunc connections nunc connection Id params
func (o *DeleteV1NuncConnectionsNuncConnectionIDParams) WithHTTPClient(client *http.Client) *DeleteV1NuncConnectionsNuncConnectionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 nunc connections nunc connection Id params
func (o *DeleteV1NuncConnectionsNuncConnectionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNuncConnectionID adds the nuncConnectionID to the delete v1 nunc connections nunc connection Id params
func (o *DeleteV1NuncConnectionsNuncConnectionIDParams) WithNuncConnectionID(nuncConnectionID int32) *DeleteV1NuncConnectionsNuncConnectionIDParams {
	o.SetNuncConnectionID(nuncConnectionID)
	return o
}

// SetNuncConnectionID adds the nuncConnectionId to the delete v1 nunc connections nunc connection Id params
func (o *DeleteV1NuncConnectionsNuncConnectionIDParams) SetNuncConnectionID(nuncConnectionID int32) {
	o.NuncConnectionID = nuncConnectionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1NuncConnectionsNuncConnectionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param nunc_connection_id
	if err := r.SetPathParam("nunc_connection_id", swag.FormatInt32(o.NuncConnectionID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}