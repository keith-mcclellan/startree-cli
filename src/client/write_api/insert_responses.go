// Code generated by go-swagger; DO NOT EDIT.

package write_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InsertReader is a Reader for the Insert structure.
type InsertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InsertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewInsertDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewInsertDefault creates a InsertDefault with default headers values
func NewInsertDefault(code int) *InsertDefault {
	return &InsertDefault{
		_statusCode: code,
	}
}

/*
InsertDefault describes a response with status code -1, with default header values.

successful operation
*/
type InsertDefault struct {
	_statusCode int
}

// IsSuccess returns true when this insert default response has a 2xx status code
func (o *InsertDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this insert default response has a 3xx status code
func (o *InsertDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this insert default response has a 4xx status code
func (o *InsertDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this insert default response has a 5xx status code
func (o *InsertDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this insert default response a status code equal to that given
func (o *InsertDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the insert default response
func (o *InsertDefault) Code() int {
	return o._statusCode
}

func (o *InsertDefault) Error() string {
	return fmt.Sprintf("[POST /v1/write/{table}][%d] insert default ", o._statusCode)
}

func (o *InsertDefault) String() string {
	return fmt.Sprintf("[POST /v1/write/{table}][%d] insert default ", o._statusCode)
}

func (o *InsertDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}