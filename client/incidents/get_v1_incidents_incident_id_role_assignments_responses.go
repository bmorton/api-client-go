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

// GetV1IncidentsIncidentIDRoleAssignmentsReader is a Reader for the GetV1IncidentsIncidentIDRoleAssignments structure.
type GetV1IncidentsIncidentIDRoleAssignmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1IncidentsIncidentIDRoleAssignmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetV1IncidentsIncidentIDRoleAssignmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetV1IncidentsIncidentIDRoleAssignmentsOK creates a GetV1IncidentsIncidentIDRoleAssignmentsOK with default headers values
func NewGetV1IncidentsIncidentIDRoleAssignmentsOK() *GetV1IncidentsIncidentIDRoleAssignmentsOK {
	return &GetV1IncidentsIncidentIDRoleAssignmentsOK{}
}

/*GetV1IncidentsIncidentIDRoleAssignmentsOK handles this case with default header values.

List all role assignments for an incident
*/
type GetV1IncidentsIncidentIDRoleAssignmentsOK struct {
	Payload *models.RoleAssignmentEntityPaginated
}

func (o *GetV1IncidentsIncidentIDRoleAssignmentsOK) Error() string {
	return fmt.Sprintf("[GET /v1/incidents/{incident_id}/role_assignments][%d] getV1IncidentsIncidentIdRoleAssignmentsOK  %+v", 200, o.Payload)
}

func (o *GetV1IncidentsIncidentIDRoleAssignmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleAssignmentEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}