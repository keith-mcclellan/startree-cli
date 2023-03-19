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

// GetTaskConfigsReader is a Reader for the GetTaskConfigs structure.
type GetTaskConfigsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskConfigsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTaskConfigsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTaskConfigsOK creates a GetTaskConfigsOK with default headers values
func NewGetTaskConfigsOK() *GetTaskConfigsOK {
	return &GetTaskConfigsOK{}
}

/*
GetTaskConfigsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTaskConfigsOK struct {
	Payload []*models.PinotTaskConfig
}

// IsSuccess returns true when this get task configs o k response has a 2xx status code
func (o *GetTaskConfigsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get task configs o k response has a 3xx status code
func (o *GetTaskConfigsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task configs o k response has a 4xx status code
func (o *GetTaskConfigsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get task configs o k response has a 5xx status code
func (o *GetTaskConfigsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get task configs o k response a status code equal to that given
func (o *GetTaskConfigsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get task configs o k response
func (o *GetTaskConfigsOK) Code() int {
	return 200
}

func (o *GetTaskConfigsOK) Error() string {
	return fmt.Sprintf("[GET /tasks/task/{taskName}/config][%d] getTaskConfigsOK  %+v", 200, o.Payload)
}

func (o *GetTaskConfigsOK) String() string {
	return fmt.Sprintf("[GET /tasks/task/{taskName}/config][%d] getTaskConfigsOK  %+v", 200, o.Payload)
}

func (o *GetTaskConfigsOK) GetPayload() []*models.PinotTaskConfig {
	return o.Payload
}

func (o *GetTaskConfigsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
