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

// GetV1EventsReader is a Reader for the GetV1Events structure.
type GetV1EventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1EventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1EventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1EventsOK creates a GetV1EventsOK with default headers values
func NewGetV1EventsOK() *GetV1EventsOK {
	return &GetV1EventsOK{}
}

/* GetV1EventsOK describes a response with status code 200, with default header values.

Retrieve the timeline for an incident
*/
type GetV1EventsOK struct {
	Payload *models.IncidentEventEntityPaginated
}

func (o *GetV1EventsOK) Error() string {
	return fmt.Sprintf("[GET /v1/events][%d] getV1EventsOK  %+v", 200, o.Payload)
}
func (o *GetV1EventsOK) GetPayload() *models.IncidentEventEntityPaginated {
	return o.Payload
}

func (o *GetV1EventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IncidentEventEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}