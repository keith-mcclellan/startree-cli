// Code generated by go-swagger; DO NOT EDIT.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExecuteAdhocTaskReader is a Reader for the ExecuteAdhocTask structure.
type ExecuteAdhocTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecuteAdhocTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewExecuteAdhocTaskDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewExecuteAdhocTaskDefault creates a ExecuteAdhocTaskDefault with default headers values
func NewExecuteAdhocTaskDefault(code int) *ExecuteAdhocTaskDefault {
	return &ExecuteAdhocTaskDefault{
		_statusCode: code,
	}
}

/*
ExecuteAdhocTaskDefault describes a response with status code -1, with default header values.

successful operation
*/
type ExecuteAdhocTaskDefault struct {
	_statusCode int
}

// IsSuccess returns true when this execute adhoc task default response has a 2xx status code
func (o *ExecuteAdhocTaskDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this execute adhoc task default response has a 3xx status code
func (o *ExecuteAdhocTaskDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this execute adhoc task default response has a 4xx status code
func (o *ExecuteAdhocTaskDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this execute adhoc task default response has a 5xx status code
func (o *ExecuteAdhocTaskDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this execute adhoc task default response a status code equal to that given
func (o *ExecuteAdhocTaskDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the execute adhoc task default response
func (o *ExecuteAdhocTaskDefault) Code() int {
	return o._statusCode
}

func (o *ExecuteAdhocTaskDefault) Error() string {
	return fmt.Sprintf("[POST /tasks/execute][%d] executeAdhocTask default ", o._statusCode)
}

func (o *ExecuteAdhocTaskDefault) String() string {
	return fmt.Sprintf("[POST /tasks/execute][%d] executeAdhocTask default ", o._statusCode)
}

func (o *ExecuteAdhocTaskDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
