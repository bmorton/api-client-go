// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/firehydrant/api-client-go/models"
)

// PatchV1IntegrationsAwsConnectionsIDReader is a Reader for the PatchV1IntegrationsAwsConnectionsID structure.
type PatchV1IntegrationsAwsConnectionsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1IntegrationsAwsConnectionsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchV1IntegrationsAwsConnectionsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchV1IntegrationsAwsConnectionsIDOK creates a PatchV1IntegrationsAwsConnectionsIDOK with default headers values
func NewPatchV1IntegrationsAwsConnectionsIDOK() *PatchV1IntegrationsAwsConnectionsIDOK {
	return &PatchV1IntegrationsAwsConnectionsIDOK{}
}

/*PatchV1IntegrationsAwsConnectionsIDOK handles this case with default header values.

translation missing: en.api.integrations.aws.connections.update.description
*/
type PatchV1IntegrationsAwsConnectionsIDOK struct {
	Payload *models.ConnectionEntity
}

func (o *PatchV1IntegrationsAwsConnectionsIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/integrations/aws/connections/{id}][%d] patchV1IntegrationsAwsConnectionsIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1IntegrationsAwsConnectionsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConnectionEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}