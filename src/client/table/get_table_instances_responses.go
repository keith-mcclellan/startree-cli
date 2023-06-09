// Code generated by go-swagger; DO NOT EDIT.

package table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetTableInstancesReader is a Reader for the GetTableInstances structure.
type GetTableInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTableInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTableInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTableInstancesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTableInstancesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTableInstancesOK creates a GetTableInstancesOK with default headers values
func NewGetTableInstancesOK() *GetTableInstancesOK {
	return &GetTableInstancesOK{}
}

/*
GetTableInstancesOK describes a response with status code 200, with default header values.

Success
*/
type GetTableInstancesOK struct {
}

// IsSuccess returns true when this get table instances o k response has a 2xx status code
func (o *GetTableInstancesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get table instances o k response has a 3xx status code
func (o *GetTableInstancesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get table instances o k response has a 4xx status code
func (o *GetTableInstancesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get table instances o k response has a 5xx status code
func (o *GetTableInstancesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get table instances o k response a status code equal to that given
func (o *GetTableInstancesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get table instances o k response
func (o *GetTableInstancesOK) Code() int {
	return 200
}

func (o *GetTableInstancesOK) Error() string {
	return fmt.Sprintf("[GET /tables/{tableName}/instances][%d] getTableInstancesOK ", 200)
}

func (o *GetTableInstancesOK) String() string {
	return fmt.Sprintf("[GET /tables/{tableName}/instances][%d] getTableInstancesOK ", 200)
}

func (o *GetTableInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTableInstancesNotFound creates a GetTableInstancesNotFound with default headers values
func NewGetTableInstancesNotFound() *GetTableInstancesNotFound {
	return &GetTableInstancesNotFound{}
}

/*
GetTableInstancesNotFound describes a response with status code 404, with default header values.

Table not found
*/
type GetTableInstancesNotFound struct {
}

// IsSuccess returns true when this get table instances not found response has a 2xx status code
func (o *GetTableInstancesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get table instances not found response has a 3xx status code
func (o *GetTableInstancesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get table instances not found response has a 4xx status code
func (o *GetTableInstancesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get table instances not found response has a 5xx status code
func (o *GetTableInstancesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get table instances not found response a status code equal to that given
func (o *GetTableInstancesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get table instances not found response
func (o *GetTableInstancesNotFound) Code() int {
	return 404
}

func (o *GetTableInstancesNotFound) Error() string {
	return fmt.Sprintf("[GET /tables/{tableName}/instances][%d] getTableInstancesNotFound ", 404)
}

func (o *GetTableInstancesNotFound) String() string {
	return fmt.Sprintf("[GET /tables/{tableName}/instances][%d] getTableInstancesNotFound ", 404)
}

func (o *GetTableInstancesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTableInstancesInternalServerError creates a GetTableInstancesInternalServerError with default headers values
func NewGetTableInstancesInternalServerError() *GetTableInstancesInternalServerError {
	return &GetTableInstancesInternalServerError{}
}

/*
GetTableInstancesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetTableInstancesInternalServerError struct {
}

// IsSuccess returns true when this get table instances internal server error response has a 2xx status code
func (o *GetTableInstancesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get table instances internal server error response has a 3xx status code
func (o *GetTableInstancesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get table instances internal server error response has a 4xx status code
func (o *GetTableInstancesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get table instances internal server error response has a 5xx status code
func (o *GetTableInstancesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get table instances internal server error response a status code equal to that given
func (o *GetTableInstancesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get table instances internal server error response
func (o *GetTableInstancesInternalServerError) Code() int {
	return 500
}

func (o *GetTableInstancesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tables/{tableName}/instances][%d] getTableInstancesInternalServerError ", 500)
}

func (o *GetTableInstancesInternalServerError) String() string {
	return fmt.Sprintf("[GET /tables/{tableName}/instances][%d] getTableInstancesInternalServerError ", 500)
}

func (o *GetTableInstancesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
