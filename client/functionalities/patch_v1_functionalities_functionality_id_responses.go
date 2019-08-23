// Code generated by go-swagger; DO NOT EDIT.

package functionalities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/firehydrant/api-client-go/models"
)

// PatchV1FunctionalitiesFunctionalityIDReader is a Reader for the PatchV1FunctionalitiesFunctionalityID structure.
type PatchV1FunctionalitiesFunctionalityIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1FunctionalitiesFunctionalityIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchV1FunctionalitiesFunctionalityIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchV1FunctionalitiesFunctionalityIDOK creates a PatchV1FunctionalitiesFunctionalityIDOK with default headers values
func NewPatchV1FunctionalitiesFunctionalityIDOK() *PatchV1FunctionalitiesFunctionalityIDOK {
	return &PatchV1FunctionalitiesFunctionalityIDOK{}
}

/*PatchV1FunctionalitiesFunctionalityIDOK handles this case with default header values.

Update a functionality's attributes

*/
type PatchV1FunctionalitiesFunctionalityIDOK struct {
	Payload *models.FunctionalityEntity
}

func (o *PatchV1FunctionalitiesFunctionalityIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/functionalities/{functionality_id}][%d] patchV1FunctionalitiesFunctionalityIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1FunctionalitiesFunctionalityIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FunctionalityEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}