// Code generated by go-swagger; DO NOT EDIT.

package reports

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

// NewGetV1ReportsMeanTimeParams creates a new GetV1ReportsMeanTimeParams object
// with the default values initialized.
func NewGetV1ReportsMeanTimeParams() *GetV1ReportsMeanTimeParams {
	var ()
	return &GetV1ReportsMeanTimeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ReportsMeanTimeParamsWithTimeout creates a new GetV1ReportsMeanTimeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetV1ReportsMeanTimeParamsWithTimeout(timeout time.Duration) *GetV1ReportsMeanTimeParams {
	var ()
	return &GetV1ReportsMeanTimeParams{

		timeout: timeout,
	}
}

// NewGetV1ReportsMeanTimeParamsWithContext creates a new GetV1ReportsMeanTimeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetV1ReportsMeanTimeParamsWithContext(ctx context.Context) *GetV1ReportsMeanTimeParams {
	var ()
	return &GetV1ReportsMeanTimeParams{

		Context: ctx,
	}
}

// NewGetV1ReportsMeanTimeParamsWithHTTPClient creates a new GetV1ReportsMeanTimeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetV1ReportsMeanTimeParamsWithHTTPClient(client *http.Client) *GetV1ReportsMeanTimeParams {
	var ()
	return &GetV1ReportsMeanTimeParams{
		HTTPClient: client,
	}
}

/*GetV1ReportsMeanTimeParams contains all the parameters to send to the API endpoint
for the get v1 reports mean time operation typically these are written to a http.Request
*/
type GetV1ReportsMeanTimeParams struct {

	/*CurrentMilestones
	  A comma separated list of current milestones

	*/
	CurrentMilestones *string
	/*EndDate
	  The end date to return incidents from

	*/
	EndDate *strfmt.Date
	/*Environments
	  A comma separated list of environment IDs

	*/
	Environments *string
	/*Query
	  A text query for an incident that searches on name, summary, and desciption

	*/
	Query *string
	/*SavedSearchID
	  The id of a previously saved search.

	*/
	SavedSearchID *string
	/*Services
	  A comma separated list of service IDs

	*/
	Services *string
	/*Severities
	  A comma separated list of severities

	*/
	Severities *string
	/*SeverityNotSet
	  Flag for including incidents where severity has not been set

	*/
	SeverityNotSet *bool
	/*StartDate
	  The start date to return incidents from

	*/
	StartDate *strfmt.Date
	/*Status
	  Incident status

	*/
	Status *string
	/*Teams
	  A comma separated list of team IDs

	*/
	Teams *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) WithTimeout(timeout time.Duration) *GetV1ReportsMeanTimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) WithContext(ctx context.Context) *GetV1ReportsMeanTimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) WithHTTPClient(client *http.Client) *GetV1ReportsMeanTimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCurrentMilestones adds the currentMilestones to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) WithCurrentMilestones(currentMilestones *string) *GetV1ReportsMeanTimeParams {
	o.SetCurrentMilestones(currentMilestones)
	return o
}

// SetCurrentMilestones adds the currentMilestones to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) SetCurrentMilestones(currentMilestones *string) {
	o.CurrentMilestones = currentMilestones
}

// WithEndDate adds the endDate to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) WithEndDate(endDate *strfmt.Date) *GetV1ReportsMeanTimeParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) SetEndDate(endDate *strfmt.Date) {
	o.EndDate = endDate
}

// WithEnvironments adds the environments to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) WithEnvironments(environments *string) *GetV1ReportsMeanTimeParams {
	o.SetEnvironments(environments)
	return o
}

// SetEnvironments adds the environments to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) SetEnvironments(environments *string) {
	o.Environments = environments
}

// WithQuery adds the query to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) WithQuery(query *string) *GetV1ReportsMeanTimeParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) SetQuery(query *string) {
	o.Query = query
}

// WithSavedSearchID adds the savedSearchID to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) WithSavedSearchID(savedSearchID *string) *GetV1ReportsMeanTimeParams {
	o.SetSavedSearchID(savedSearchID)
	return o
}

// SetSavedSearchID adds the savedSearchId to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) SetSavedSearchID(savedSearchID *string) {
	o.SavedSearchID = savedSearchID
}

// WithServices adds the services to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) WithServices(services *string) *GetV1ReportsMeanTimeParams {
	o.SetServices(services)
	return o
}

// SetServices adds the services to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) SetServices(services *string) {
	o.Services = services
}

// WithSeverities adds the severities to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) WithSeverities(severities *string) *GetV1ReportsMeanTimeParams {
	o.SetSeverities(severities)
	return o
}

// SetSeverities adds the severities to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) SetSeverities(severities *string) {
	o.Severities = severities
}

// WithSeverityNotSet adds the severityNotSet to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) WithSeverityNotSet(severityNotSet *bool) *GetV1ReportsMeanTimeParams {
	o.SetSeverityNotSet(severityNotSet)
	return o
}

// SetSeverityNotSet adds the severityNotSet to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) SetSeverityNotSet(severityNotSet *bool) {
	o.SeverityNotSet = severityNotSet
}

// WithStartDate adds the startDate to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) WithStartDate(startDate *strfmt.Date) *GetV1ReportsMeanTimeParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) SetStartDate(startDate *strfmt.Date) {
	o.StartDate = startDate
}

// WithStatus adds the status to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) WithStatus(status *string) *GetV1ReportsMeanTimeParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) SetStatus(status *string) {
	o.Status = status
}

// WithTeams adds the teams to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) WithTeams(teams *string) *GetV1ReportsMeanTimeParams {
	o.SetTeams(teams)
	return o
}

// SetTeams adds the teams to the get v1 reports mean time params
func (o *GetV1ReportsMeanTimeParams) SetTeams(teams *string) {
	o.Teams = teams
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ReportsMeanTimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CurrentMilestones != nil {

		// query param current_milestones
		var qrCurrentMilestones string
		if o.CurrentMilestones != nil {
			qrCurrentMilestones = *o.CurrentMilestones
		}
		qCurrentMilestones := qrCurrentMilestones
		if qCurrentMilestones != "" {
			if err := r.SetQueryParam("current_milestones", qCurrentMilestones); err != nil {
				return err
			}
		}

	}

	if o.EndDate != nil {

		// query param end_date
		var qrEndDate strfmt.Date
		if o.EndDate != nil {
			qrEndDate = *o.EndDate
		}
		qEndDate := qrEndDate.String()
		if qEndDate != "" {
			if err := r.SetQueryParam("end_date", qEndDate); err != nil {
				return err
			}
		}

	}

	if o.Environments != nil {

		// query param environments
		var qrEnvironments string
		if o.Environments != nil {
			qrEnvironments = *o.Environments
		}
		qEnvironments := qrEnvironments
		if qEnvironments != "" {
			if err := r.SetQueryParam("environments", qEnvironments); err != nil {
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

	if o.SavedSearchID != nil {

		// query param saved_search_id
		var qrSavedSearchID string
		if o.SavedSearchID != nil {
			qrSavedSearchID = *o.SavedSearchID
		}
		qSavedSearchID := qrSavedSearchID
		if qSavedSearchID != "" {
			if err := r.SetQueryParam("saved_search_id", qSavedSearchID); err != nil {
				return err
			}
		}

	}

	if o.Services != nil {

		// query param services
		var qrServices string
		if o.Services != nil {
			qrServices = *o.Services
		}
		qServices := qrServices
		if qServices != "" {
			if err := r.SetQueryParam("services", qServices); err != nil {
				return err
			}
		}

	}

	if o.Severities != nil {

		// query param severities
		var qrSeverities string
		if o.Severities != nil {
			qrSeverities = *o.Severities
		}
		qSeverities := qrSeverities
		if qSeverities != "" {
			if err := r.SetQueryParam("severities", qSeverities); err != nil {
				return err
			}
		}

	}

	if o.SeverityNotSet != nil {

		// query param severity_not_set
		var qrSeverityNotSet bool
		if o.SeverityNotSet != nil {
			qrSeverityNotSet = *o.SeverityNotSet
		}
		qSeverityNotSet := swag.FormatBool(qrSeverityNotSet)
		if qSeverityNotSet != "" {
			if err := r.SetQueryParam("severity_not_set", qSeverityNotSet); err != nil {
				return err
			}
		}

	}

	if o.StartDate != nil {

		// query param start_date
		var qrStartDate strfmt.Date
		if o.StartDate != nil {
			qrStartDate = *o.StartDate
		}
		qStartDate := qrStartDate.String()
		if qStartDate != "" {
			if err := r.SetQueryParam("start_date", qStartDate); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if o.Teams != nil {

		// query param teams
		var qrTeams string
		if o.Teams != nil {
			qrTeams = *o.Teams
		}
		qTeams := qrTeams
		if qTeams != "" {
			if err := r.SetQueryParam("teams", qTeams); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}