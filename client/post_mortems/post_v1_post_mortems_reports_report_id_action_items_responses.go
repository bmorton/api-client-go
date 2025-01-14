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

// PostV1PostMortemsReportsReportIDActionItemsReader is a Reader for the PostV1PostMortemsReportsReportIDActionItems structure.
type PostV1PostMortemsReportsReportIDActionItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1PostMortemsReportsReportIDActionItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1PostMortemsReportsReportIDActionItemsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1PostMortemsReportsReportIDActionItemsCreated creates a PostV1PostMortemsReportsReportIDActionItemsCreated with default headers values
func NewPostV1PostMortemsReportsReportIDActionItemsCreated() *PostV1PostMortemsReportsReportIDActionItemsCreated {
	return &PostV1PostMortemsReportsReportIDActionItemsCreated{}
}

/* PostV1PostMortemsReportsReportIDActionItemsCreated describes a response with status code 201, with default header values.

Create an action item on a report
*/
type PostV1PostMortemsReportsReportIDActionItemsCreated struct {
	Payload *models.ActionItemEntity
}

func (o *PostV1PostMortemsReportsReportIDActionItemsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/post_mortems/reports/{report_id}/action_items][%d] postV1PostMortemsReportsReportIdActionItemsCreated  %+v", 201, o.Payload)
}
func (o *PostV1PostMortemsReportsReportIDActionItemsCreated) GetPayload() *models.ActionItemEntity {
	return o.Payload
}

func (o *PostV1PostMortemsReportsReportIDActionItemsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionItemEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
