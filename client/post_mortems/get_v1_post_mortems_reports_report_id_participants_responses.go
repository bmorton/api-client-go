// Code generated by go-swagger; DO NOT EDIT.

package post_mortems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1PostMortemsReportsReportIDParticipantsReader is a Reader for the GetV1PostMortemsReportsReportIDParticipants structure.
type GetV1PostMortemsReportsReportIDParticipantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1PostMortemsReportsReportIDParticipantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1PostMortemsReportsReportIDParticipantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1PostMortemsReportsReportIDParticipantsOK creates a GetV1PostMortemsReportsReportIDParticipantsOK with default headers values
func NewGetV1PostMortemsReportsReportIDParticipantsOK() *GetV1PostMortemsReportsReportIDParticipantsOK {
	return &GetV1PostMortemsReportsReportIDParticipantsOK{}
}

/* GetV1PostMortemsReportsReportIDParticipantsOK describes a response with status code 200, with default header values.

List participants for a post mortem report
*/
type GetV1PostMortemsReportsReportIDParticipantsOK struct {
	Payload *models.ParticipantEntityPaginated
}

func (o *GetV1PostMortemsReportsReportIDParticipantsOK) Error() string {
	return fmt.Sprintf("[GET /v1/post_mortems/reports/{report_id}/participants][%d] getV1PostMortemsReportsReportIdParticipantsOK  %+v", 200, o.Payload)
}
func (o *GetV1PostMortemsReportsReportIDParticipantsOK) GetPayload() *models.ParticipantEntityPaginated {
	return o.Payload
}

func (o *GetV1PostMortemsReportsReportIDParticipantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ParticipantEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
