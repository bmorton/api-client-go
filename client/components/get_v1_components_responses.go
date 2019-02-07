// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/firehydrant/api-client-go/models"
)

// GetV1ComponentsReader is a Reader for the GetV1Components structure.
type GetV1ComponentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ComponentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetV1ComponentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetV1ComponentsOK creates a GetV1ComponentsOK with default headers values
func NewGetV1ComponentsOK() *GetV1ComponentsOK {
	return &GetV1ComponentsOK{}
}

/*GetV1ComponentsOK handles this case with default header values.

Retrieve all components
*/
type GetV1ComponentsOK struct {
	Payload *models.ComponentEntityPaginated
}

func (o *GetV1ComponentsOK) Error() string {
	return fmt.Sprintf("[GET /v1/components][%d] getV1ComponentsOK  %+v", 200, o.Payload)
}

func (o *GetV1ComponentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComponentEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}