// Code generated by go-swagger; DO NOT EDIT.

package ticketing

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

// NewPostV1TicketingTicketsParams creates a new PostV1TicketingTicketsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1TicketingTicketsParams() *PostV1TicketingTicketsParams {
	return &PostV1TicketingTicketsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1TicketingTicketsParamsWithTimeout creates a new PostV1TicketingTicketsParams object
// with the ability to set a timeout on a request.
func NewPostV1TicketingTicketsParamsWithTimeout(timeout time.Duration) *PostV1TicketingTicketsParams {
	return &PostV1TicketingTicketsParams{
		timeout: timeout,
	}
}

// NewPostV1TicketingTicketsParamsWithContext creates a new PostV1TicketingTicketsParams object
// with the ability to set a context for a request.
func NewPostV1TicketingTicketsParamsWithContext(ctx context.Context) *PostV1TicketingTicketsParams {
	return &PostV1TicketingTicketsParams{
		Context: ctx,
	}
}

// NewPostV1TicketingTicketsParamsWithHTTPClient creates a new PostV1TicketingTicketsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1TicketingTicketsParamsWithHTTPClient(client *http.Client) *PostV1TicketingTicketsParams {
	return &PostV1TicketingTicketsParams{
		HTTPClient: client,
	}
}

/* PostV1TicketingTicketsParams contains all the parameters to send to the API endpoint
   for the post v1 ticketing tickets operation.

   Typically these are written to a http.Request.
*/
type PostV1TicketingTicketsParams struct {

	// V1TicketingTickets.
	V1TicketingTickets *models.PostV1TicketingTickets

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 ticketing tickets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1TicketingTicketsParams) WithDefaults() *PostV1TicketingTicketsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 ticketing tickets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1TicketingTicketsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 ticketing tickets params
func (o *PostV1TicketingTicketsParams) WithTimeout(timeout time.Duration) *PostV1TicketingTicketsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 ticketing tickets params
func (o *PostV1TicketingTicketsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 ticketing tickets params
func (o *PostV1TicketingTicketsParams) WithContext(ctx context.Context) *PostV1TicketingTicketsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 ticketing tickets params
func (o *PostV1TicketingTicketsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 ticketing tickets params
func (o *PostV1TicketingTicketsParams) WithHTTPClient(client *http.Client) *PostV1TicketingTicketsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 ticketing tickets params
func (o *PostV1TicketingTicketsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1TicketingTickets adds the v1TicketingTickets to the post v1 ticketing tickets params
func (o *PostV1TicketingTicketsParams) WithV1TicketingTickets(v1TicketingTickets *models.PostV1TicketingTickets) *PostV1TicketingTicketsParams {
	o.SetV1TicketingTickets(v1TicketingTickets)
	return o
}

// SetV1TicketingTickets adds the v1TicketingTickets to the post v1 ticketing tickets params
func (o *PostV1TicketingTicketsParams) SetV1TicketingTickets(v1TicketingTickets *models.PostV1TicketingTickets) {
	o.V1TicketingTickets = v1TicketingTickets
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1TicketingTicketsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1TicketingTickets != nil {
		if err := r.SetBodyParam(o.V1TicketingTickets); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
