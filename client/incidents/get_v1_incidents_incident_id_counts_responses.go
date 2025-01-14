// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetV1IncidentsIncidentIDCountsReader is a Reader for the GetV1IncidentsIncidentIDCounts structure.
type GetV1IncidentsIncidentIDCountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1IncidentsIncidentIDCountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1IncidentsIncidentIDCountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1IncidentsIncidentIDCountsOK creates a GetV1IncidentsIncidentIDCountsOK with default headers values
func NewGetV1IncidentsIncidentIDCountsOK() *GetV1IncidentsIncidentIDCountsOK {
	return &GetV1IncidentsIncidentIDCountsOK{}
}

/* GetV1IncidentsIncidentIDCountsOK describes a response with status code 200, with default header values.

get Count(s)
*/
type GetV1IncidentsIncidentIDCountsOK struct {
}

func (o *GetV1IncidentsIncidentIDCountsOK) Error() string {
	return fmt.Sprintf("[GET /v1/incidents/{incident_id}/counts][%d] getV1IncidentsIncidentIdCountsOK ", 200)
}

func (o *GetV1IncidentsIncidentIDCountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
