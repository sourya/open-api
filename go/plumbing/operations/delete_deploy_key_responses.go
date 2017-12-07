// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/go/models"
)

// DeleteDeployKeyReader is a Reader for the DeleteDeployKey structure.
type DeleteDeployKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDeployKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteDeployKeyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteDeployKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDeployKeyNoContent creates a DeleteDeployKeyNoContent with default headers values
func NewDeleteDeployKeyNoContent() *DeleteDeployKeyNoContent {
	return &DeleteDeployKeyNoContent{}
}

/*DeleteDeployKeyNoContent handles this case with default header values.

Not Content
*/
type DeleteDeployKeyNoContent struct {
}

func (o *DeleteDeployKeyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /deploy_keys/{key_id}][%d] deleteDeployKeyNoContent ", 204)
}

func (o *DeleteDeployKeyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeployKeyDefault creates a DeleteDeployKeyDefault with default headers values
func NewDeleteDeployKeyDefault(code int) *DeleteDeployKeyDefault {
	return &DeleteDeployKeyDefault{
		_statusCode: code,
	}
}

/*DeleteDeployKeyDefault handles this case with default header values.

error
*/
type DeleteDeployKeyDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete deploy key default response
func (o *DeleteDeployKeyDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDeployKeyDefault) Error() string {
	return fmt.Sprintf("[DELETE /deploy_keys/{key_id}][%d] deleteDeployKey default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDeployKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
