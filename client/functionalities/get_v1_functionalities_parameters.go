// Code generated by go-swagger; DO NOT EDIT.

package functionalities

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

// NewGetV1FunctionalitiesParams creates a new GetV1FunctionalitiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1FunctionalitiesParams() *GetV1FunctionalitiesParams {
	return &GetV1FunctionalitiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1FunctionalitiesParamsWithTimeout creates a new GetV1FunctionalitiesParams object
// with the ability to set a timeout on a request.
func NewGetV1FunctionalitiesParamsWithTimeout(timeout time.Duration) *GetV1FunctionalitiesParams {
	return &GetV1FunctionalitiesParams{
		timeout: timeout,
	}
}

// NewGetV1FunctionalitiesParamsWithContext creates a new GetV1FunctionalitiesParams object
// with the ability to set a context for a request.
func NewGetV1FunctionalitiesParamsWithContext(ctx context.Context) *GetV1FunctionalitiesParams {
	return &GetV1FunctionalitiesParams{
		Context: ctx,
	}
}

// NewGetV1FunctionalitiesParamsWithHTTPClient creates a new GetV1FunctionalitiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1FunctionalitiesParamsWithHTTPClient(client *http.Client) *GetV1FunctionalitiesParams {
	return &GetV1FunctionalitiesParams{
		HTTPClient: client,
	}
}

/* GetV1FunctionalitiesParams contains all the parameters to send to the API endpoint
   for the get v1 functionalities operation.

   Typically these are written to a http.Request.
*/
type GetV1FunctionalitiesParams struct {

	// Page.
	//
	// Format: int32
	Page *int32

	// PerPage.
	//
	// Format: int32
	PerPage *int32

	/* Query.

	   A query to search functionalities by their name
	*/
	Query *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 functionalities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1FunctionalitiesParams) WithDefaults() *GetV1FunctionalitiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 functionalities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1FunctionalitiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 functionalities params
func (o *GetV1FunctionalitiesParams) WithTimeout(timeout time.Duration) *GetV1FunctionalitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 functionalities params
func (o *GetV1FunctionalitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 functionalities params
func (o *GetV1FunctionalitiesParams) WithContext(ctx context.Context) *GetV1FunctionalitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 functionalities params
func (o *GetV1FunctionalitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 functionalities params
func (o *GetV1FunctionalitiesParams) WithHTTPClient(client *http.Client) *GetV1FunctionalitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 functionalities params
func (o *GetV1FunctionalitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get v1 functionalities params
func (o *GetV1FunctionalitiesParams) WithPage(page *int32) *GetV1FunctionalitiesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 functionalities params
func (o *GetV1FunctionalitiesParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 functionalities params
func (o *GetV1FunctionalitiesParams) WithPerPage(perPage *int32) *GetV1FunctionalitiesParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 functionalities params
func (o *GetV1FunctionalitiesParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithQuery adds the query to the get v1 functionalities params
func (o *GetV1FunctionalitiesParams) WithQuery(query *string) *GetV1FunctionalitiesParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get v1 functionalities params
func (o *GetV1FunctionalitiesParams) SetQuery(query *string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1FunctionalitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int32

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
