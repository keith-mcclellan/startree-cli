// Code generated by go-swagger; DO NOT EDIT.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"startree.ai/cli/models"
)

// DeleteTaskReader is a Reader for the DeleteTask structure.
type DeleteTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTaskOK creates a DeleteTaskOK with default headers values
func NewDeleteTaskOK() *DeleteTaskOK {
	return &DeleteTaskOK{}
}

/*
DeleteTaskOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteTaskOK struct {
	Payload *models.SuccessResponse
}

// IsSuccess returns true when this delete task o k response has a 2xx status code
func (o *DeleteTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete task o k response has a 3xx status code
func (o *DeleteTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete task o k response has a 4xx status code
func (o *DeleteTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete task o k response has a 5xx status code
func (o *DeleteTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete task o k response a status code equal to that given
func (o *DeleteTaskOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete task o k response
func (o *DeleteTaskOK) Code() int {
	return 200
}

func (o *DeleteTaskOK) Error() string {
	return fmt.Sprintf("[DELETE /tasks/task/{taskName}][%d] deleteTaskOK  %+v", 200, o.Payload)
}

func (o *DeleteTaskOK) String() string {
	return fmt.Sprintf("[DELETE /tasks/task/{taskName}][%d] deleteTaskOK  %+v", 200, o.Payload)
}

func (o *DeleteTaskOK) GetPayload() *models.SuccessResponse {
	return o.Payload
}

func (o *DeleteTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
