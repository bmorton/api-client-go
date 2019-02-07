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

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams creates a new DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams object
// with the default values initialized.
func NewDeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams() *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams {
	var ()
	return &DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParamsWithTimeout creates a new DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParamsWithTimeout(timeout time.Duration) *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams {
	var ()
	return &DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams{

		timeout: timeout,
	}
}

// NewDeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParamsWithContext creates a new DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParamsWithContext(ctx context.Context) *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams {
	var ()
	return &DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams{

		Context: ctx,
	}
}

// NewDeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParamsWithHTTPClient creates a new DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParamsWithHTTPClient(client *http.Client) *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams {
	var ()
	return &DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams{
		HTTPClient: client,
	}
}

/*DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams contains all the parameters to send to the API endpoint
for the delete v1 post mortems reports report Id affected components affected component Id operation typically these are written to a http.Request
*/
type DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams struct {

	/*AffectedComponentID
	  The affected component id. Note: This is not the component ID. You must retrieve the ID from the affected components endpoint first.

	*/
	AffectedComponentID string
	/*ReportID*/
	ReportID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete v1 post mortems reports report Id affected components affected component Id params
func (o *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams) WithTimeout(timeout time.Duration) *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 post mortems reports report Id affected components affected component Id params
func (o *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 post mortems reports report Id affected components affected component Id params
func (o *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams) WithContext(ctx context.Context) *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 post mortems reports report Id affected components affected component Id params
func (o *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 post mortems reports report Id affected components affected component Id params
func (o *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams) WithHTTPClient(client *http.Client) *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 post mortems reports report Id affected components affected component Id params
func (o *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAffectedComponentID adds the affectedComponentID to the delete v1 post mortems reports report Id affected components affected component Id params
func (o *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams) WithAffectedComponentID(affectedComponentID string) *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams {
	o.SetAffectedComponentID(affectedComponentID)
	return o
}

// SetAffectedComponentID adds the affectedComponentId to the delete v1 post mortems reports report Id affected components affected component Id params
func (o *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams) SetAffectedComponentID(affectedComponentID string) {
	o.AffectedComponentID = affectedComponentID
}

// WithReportID adds the reportID to the delete v1 post mortems reports report Id affected components affected component Id params
func (o *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams) WithReportID(reportID string) *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams {
	o.SetReportID(reportID)
	return o
}

// SetReportID adds the reportId to the delete v1 post mortems reports report Id affected components affected component Id params
func (o *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams) SetReportID(reportID string) {
	o.ReportID = reportID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param affected_component_id
	if err := r.SetPathParam("affected_component_id", o.AffectedComponentID); err != nil {
		return err
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