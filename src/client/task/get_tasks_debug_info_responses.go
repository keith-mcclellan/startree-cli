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

// GetTasksDebugInfoReader is a Reader for the GetTasksDebugInfo structure.
type GetTasksDebugInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTasksDebugInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTasksDebugInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTasksDebugInfoOK creates a GetTasksDebugInfoOK with default headers values
func NewGetTasksDebugInfoOK() *GetTasksDebugInfoOK {
	return &GetTasksDebugInfoOK{}
}

/*
GetTasksDebugInfoOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTasksDebugInfoOK struct {
	Payload map[string]models.TaskDebugInfo
}

// IsSuccess returns true when this get tasks debug info o k response has a 2xx status code
func (o *GetTasksDebugInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tasks debug info o k response has a 3xx status code
func (o *GetTasksDebugInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tasks debug info o k response has a 4xx status code
func (o *GetTasksDebugInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tasks debug info o k response has a 5xx status code
func (o *GetTasksDebugInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tasks debug info o k response a status code equal to that given
func (o *GetTasksDebugInfoOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get tasks debug info o k response
func (o *GetTasksDebugInfoOK) Code() int {
	return 200
}

func (o *GetTasksDebugInfoOK) Error() string {
	return fmt.Sprintf("[GET /tasks/{taskType}/debug][%d] getTasksDebugInfoOK  %+v", 200, o.Payload)
}

func (o *GetTasksDebugInfoOK) String() string {
	return fmt.Sprintf("[GET /tasks/{taskType}/debug][%d] getTasksDebugInfoOK  %+v", 200, o.Payload)
}

func (o *GetTasksDebugInfoOK) GetPayload() map[string]models.TaskDebugInfo {
	return o.Payload
}

func (o *GetTasksDebugInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
