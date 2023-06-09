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

// GetTaskDebugInfoReader is a Reader for the GetTaskDebugInfo structure.
type GetTaskDebugInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskDebugInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTaskDebugInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTaskDebugInfoOK creates a GetTaskDebugInfoOK with default headers values
func NewGetTaskDebugInfoOK() *GetTaskDebugInfoOK {
	return &GetTaskDebugInfoOK{}
}

/*
GetTaskDebugInfoOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTaskDebugInfoOK struct {
	Payload *models.TaskDebugInfo
}

// IsSuccess returns true when this get task debug info o k response has a 2xx status code
func (o *GetTaskDebugInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get task debug info o k response has a 3xx status code
func (o *GetTaskDebugInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task debug info o k response has a 4xx status code
func (o *GetTaskDebugInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get task debug info o k response has a 5xx status code
func (o *GetTaskDebugInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get task debug info o k response a status code equal to that given
func (o *GetTaskDebugInfoOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get task debug info o k response
func (o *GetTaskDebugInfoOK) Code() int {
	return 200
}

func (o *GetTaskDebugInfoOK) Error() string {
	return fmt.Sprintf("[GET /tasks/task/{taskName}/debug][%d] getTaskDebugInfoOK  %+v", 200, o.Payload)
}

func (o *GetTaskDebugInfoOK) String() string {
	return fmt.Sprintf("[GET /tasks/task/{taskName}/debug][%d] getTaskDebugInfoOK  %+v", 200, o.Payload)
}

func (o *GetTaskDebugInfoOK) GetPayload() *models.TaskDebugInfo {
	return o.Payload
}

func (o *GetTaskDebugInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskDebugInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
