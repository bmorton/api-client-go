// Code generated by go-swagger; DO NOT EDIT.

package post_mortems

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

// NewGetV1PostMortemsReportsReportIDActionItemsParams creates a new GetV1PostMortemsReportsReportIDActionItemsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1PostMortemsReportsReportIDActionItemsParams() *GetV1PostMortemsReportsReportIDActionItemsParams {
	return &GetV1PostMortemsReportsReportIDActionItemsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1PostMortemsReportsReportIDActionItemsParamsWithTimeout creates a new GetV1PostMortemsReportsReportIDActionItemsParams object
// with the ability to set a timeout on a request.
func NewGetV1PostMortemsReportsReportIDActionItemsParamsWithTimeout(timeout time.Duration) *GetV1PostMortemsReportsReportIDActionItemsParams {
	return &GetV1PostMortemsReportsReportIDActionItemsParams{
		timeout: timeout,
	}
}

// NewGetV1PostMortemsReportsReportIDActionItemsParamsWithContext creates a new GetV1PostMortemsReportsReportIDActionItemsParams object
// with the ability to set a context for a request.
func NewGetV1PostMortemsReportsReportIDActionItemsParamsWithContext(ctx context.Context) *GetV1PostMortemsReportsReportIDActionItemsParams {
	return &GetV1PostMortemsReportsReportIDActionItemsParams{
		Context: ctx,
	}
}

// NewGetV1PostMortemsReportsReportIDActionItemsParamsWithHTTPClient creates a new GetV1PostMortemsReportsReportIDActionItemsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1PostMortemsReportsReportIDActionItemsParamsWithHTTPClient(client *http.Client) *GetV1PostMortemsReportsReportIDActionItemsParams {
	return &GetV1PostMortemsReportsReportIDActionItemsParams{
		HTTPClient: client,
	}
}

/* GetV1PostMortemsReportsReportIDActionItemsParams contains all the parameters to send to the API endpoint
   for the get v1 post mortems reports report Id action items operation.

   Typically these are written to a http.Request.
*/
type GetV1PostMortemsReportsReportIDActionItemsParams struct {

	// Page.
	//
	// Format: int32
	Page *int32

	// PerPage.
	//
	// Format: int32
	PerPage *int32

	// ReportID.
	ReportID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 post mortems reports report Id action items params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1PostMortemsReportsReportIDActionItemsParams) WithDefaults() *GetV1PostMortemsReportsReportIDActionItemsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 post mortems reports report Id action items params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1PostMortemsReportsReportIDActionItemsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 post mortems reports report Id action items params
func (o *GetV1PostMortemsReportsReportIDActionItemsParams) WithTimeout(timeout time.Duration) *GetV1PostMortemsReportsReportIDActionItemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 post mortems reports report Id action items params
func (o *GetV1PostMortemsReportsReportIDActionItemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 post mortems reports report Id action items params
func (o *GetV1PostMortemsReportsReportIDActionItemsParams) WithContext(ctx context.Context) *GetV1PostMortemsReportsReportIDActionItemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 post mortems reports report Id action items params
func (o *GetV1PostMortemsReportsReportIDActionItemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 post mortems reports report Id action items params
func (o *GetV1PostMortemsReportsReportIDActionItemsParams) WithHTTPClient(client *http.Client) *GetV1PostMortemsReportsReportIDActionItemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 post mortems reports report Id action items params
func (o *GetV1PostMortemsReportsReportIDActionItemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get v1 post mortems reports report Id action items params
func (o *GetV1PostMortemsReportsReportIDActionItemsParams) WithPage(page *int32) *GetV1PostMortemsReportsReportIDActionItemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 post mortems reports report Id action items params
func (o *GetV1PostMortemsReportsReportIDActionItemsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 post mortems reports report Id action items params
func (o *GetV1PostMortemsReportsReportIDActionItemsParams) WithPerPage(perPage *int32) *GetV1PostMortemsReportsReportIDActionItemsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 post mortems reports report Id action items params
func (o *GetV1PostMortemsReportsReportIDActionItemsParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithReportID adds the reportID to the get v1 post mortems reports report Id action items params
func (o *GetV1PostMortemsReportsReportIDActionItemsParams) WithReportID(reportID string) *GetV1PostMortemsReportsReportIDActionItemsParams {
	o.SetReportID(reportID)
	return o
}

// SetReportID adds the reportId to the get v1 post mortems reports report Id action items params
func (o *GetV1PostMortemsReportsReportIDActionItemsParams) SetReportID(reportID string) {
	o.ReportID = reportID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1PostMortemsReportsReportIDActionItemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param report_id
	if err := r.SetPathParam("report_id", o.ReportID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
