// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1EventsEventIDReader is a Reader for the GetV1EventsEventID structure.
type GetV1EventsEventIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1EventsEventIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1EventsEventIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1EventsEventIDOK creates a GetV1EventsEventIDOK with default headers values
func NewGetV1EventsEventIDOK() *GetV1EventsEventIDOK {
	return &GetV1EventsEventIDOK{}
}

/* GetV1EventsEventIDOK describes a response with status code 200, with default header values.

Get an individual incident event
*/
type GetV1EventsEventIDOK struct {
	Payload *models.IncidentEventEntity
}

func (o *GetV1EventsEventIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/events/{event_id}][%d] getV1EventsEventIdOK  %+v", 200, o.Payload)
}
func (o *GetV1EventsEventIDOK) GetPayload() *models.IncidentEventEntity {
	return o.Payload
}

func (o *GetV1EventsEventIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IncidentEventEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}