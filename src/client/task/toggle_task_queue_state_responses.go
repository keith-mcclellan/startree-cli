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

// ToggleTaskQueueStateReader is a Reader for the ToggleTaskQueueState structure.
type ToggleTaskQueueStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ToggleTaskQueueStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewToggleTaskQueueStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewToggleTaskQueueStateOK creates a ToggleTaskQueueStateOK with default headers values
func NewToggleTaskQueueStateOK() *ToggleTaskQueueStateOK {
	return &ToggleTaskQueueStateOK{}
}

/*
ToggleTaskQueueStateOK describes a response with status code 200, with default header values.

successful operation
*/
type ToggleTaskQueueStateOK struct {
	Payload *models.SuccessResponse
}

// IsSuccess returns true when this toggle task queue state o k response has a 2xx status code
func (o *ToggleTaskQueueStateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this toggle task queue state o k response has a 3xx status code
func (o *ToggleTaskQueueStateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this toggle task queue state o k response has a 4xx status code
func (o *ToggleTaskQueueStateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this toggle task queue state o k response has a 5xx status code
func (o *ToggleTaskQueueStateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this toggle task queue state o k response a status code equal to that given
func (o *ToggleTaskQueueStateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the toggle task queue state o k response
func (o *ToggleTaskQueueStateOK) Code() int {
	return 200
}

func (o *ToggleTaskQueueStateOK) Error() string {
	return fmt.Sprintf("[PUT /tasks/taskqueue/{taskType}][%d] toggleTaskQueueStateOK  %+v", 200, o.Payload)
}

func (o *ToggleTaskQueueStateOK) String() string {
	return fmt.Sprintf("[PUT /tasks/taskqueue/{taskType}][%d] toggleTaskQueueStateOK  %+v", 200, o.Payload)
}

func (o *ToggleTaskQueueStateOK) GetPayload() *models.SuccessResponse {
	return o.Payload
}

func (o *ToggleTaskQueueStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
