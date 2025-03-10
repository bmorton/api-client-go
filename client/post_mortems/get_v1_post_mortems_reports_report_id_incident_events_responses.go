// Code generated by go-swagger; DO NOT EDIT.

package post_mortems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetV1PostMortemsReportsReportIDIncidentEventsReader is a Reader for the GetV1PostMortemsReportsReportIDIncidentEvents structure.
type GetV1PostMortemsReportsReportIDIncidentEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1PostMortemsReportsReportIDIncidentEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1PostMortemsReportsReportIDIncidentEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1PostMortemsReportsReportIDIncidentEventsOK creates a GetV1PostMortemsReportsReportIDIncidentEventsOK with default headers values
func NewGetV1PostMortemsReportsReportIDIncidentEventsOK() *GetV1PostMortemsReportsReportIDIncidentEventsOK {
	return &GetV1PostMortemsReportsReportIDIncidentEventsOK{}
}

/* GetV1PostMortemsReportsReportIDIncidentEventsOK describes a response with status code 200, with default header values.

get IncidentEvent(s)
*/
type GetV1PostMortemsReportsReportIDIncidentEventsOK struct {
}

func (o *GetV1PostMortemsReportsReportIDIncidentEventsOK) Error() string {
	return fmt.Sprintf("[GET /v1/post_mortems/reports/{report_id}/incident_events][%d] getV1PostMortemsReportsReportIdIncidentEventsOK ", 200)
}

func (o *GetV1PostMortemsReportsReportIDIncidentEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
