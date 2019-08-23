// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/firehydrant/api-client-go/models"
)

// PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryReader is a Reader for the PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimary structure.
type PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK creates a PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK with default headers values
func NewPatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK() *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK {
	return &PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK{}
}

/*PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK handles this case with default header values.

Assign an alert a primary status
*/
type PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK struct {
	Payload *models.AlertEntity
}

func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/incidents/{incident_id}/alerts/{incident_alert_id}/primary][%d] patchV1IncidentsIncidentIdAlertsIncidentAlertIdPrimaryOK  %+v", 200, o.Payload)
}

func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}