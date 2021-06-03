// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// DeleteV1IntegrationsStatuspageConnectionsIDReader is a Reader for the DeleteV1IntegrationsStatuspageConnectionsID structure.
type DeleteV1IntegrationsStatuspageConnectionsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1IntegrationsStatuspageConnectionsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1IntegrationsStatuspageConnectionsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1IntegrationsStatuspageConnectionsIDOK creates a DeleteV1IntegrationsStatuspageConnectionsIDOK with default headers values
func NewDeleteV1IntegrationsStatuspageConnectionsIDOK() *DeleteV1IntegrationsStatuspageConnectionsIDOK {
	return &DeleteV1IntegrationsStatuspageConnectionsIDOK{}
}

/* DeleteV1IntegrationsStatuspageConnectionsIDOK describes a response with status code 200, with default header values.

Deletes the given Statuspage integration connection.
*/
type DeleteV1IntegrationsStatuspageConnectionsIDOK struct {
	Payload *models.ConnectionEntity
}

func (o *DeleteV1IntegrationsStatuspageConnectionsIDOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/integrations/statuspage/connections/{id}][%d] deleteV1IntegrationsStatuspageConnectionsIdOK  %+v", 200, o.Payload)
}
func (o *DeleteV1IntegrationsStatuspageConnectionsIDOK) GetPayload() *models.ConnectionEntity {
	return o.Payload
}

func (o *DeleteV1IntegrationsStatuspageConnectionsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConnectionEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}