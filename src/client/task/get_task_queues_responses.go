// Code generated by go-swagger; DO NOT EDIT.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetTaskQueuesReader is a Reader for the GetTaskQueues structure.
type GetTaskQueuesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskQueuesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTaskQueuesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTaskQueuesOK creates a GetTaskQueuesOK with default headers values
func NewGetTaskQueuesOK() *GetTaskQueuesOK {
	return &GetTaskQueuesOK{}
}

/*
GetTaskQueuesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTaskQueuesOK struct {
	Payload []string
}

// IsSuccess returns true when this get task queues o k response has a 2xx status code
func (o *GetTaskQueuesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get task queues o k response has a 3xx status code
func (o *GetTaskQueuesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task queues o k response has a 4xx status code
func (o *GetTaskQueuesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get task queues o k response has a 5xx status code
func (o *GetTaskQueuesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get task queues o k response a status code equal to that given
func (o *GetTaskQueuesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get task queues o k response
func (o *GetTaskQueuesOK) Code() int {
	return 200
}

func (o *GetTaskQueuesOK) Error() string {
	return fmt.Sprintf("[GET /tasks/taskqueues][%d] getTaskQueuesOK  %+v", 200, o.Payload)
}

func (o *GetTaskQueuesOK) String() string {
	return fmt.Sprintf("[GET /tasks/taskqueues][%d] getTaskQueuesOK  %+v", 200, o.Payload)
}

func (o *GetTaskQueuesOK) GetPayload() []string {
	return o.Payload
}

func (o *GetTaskQueuesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}