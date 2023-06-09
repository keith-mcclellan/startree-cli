// Code generated by go-swagger; DO NOT EDIT.

package query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CancelQueryReader is a Reader for the CancelQuery structure.
type CancelQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCancelQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCancelQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCancelQueryOK creates a CancelQueryOK with default headers values
func NewCancelQueryOK() *CancelQueryOK {
	return &CancelQueryOK{}
}

/*
CancelQueryOK describes a response with status code 200, with default header values.

Success
*/
type CancelQueryOK struct {
}

// IsSuccess returns true when this cancel query o k response has a 2xx status code
func (o *CancelQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cancel query o k response has a 3xx status code
func (o *CancelQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel query o k response has a 4xx status code
func (o *CancelQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel query o k response has a 5xx status code
func (o *CancelQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel query o k response a status code equal to that given
func (o *CancelQueryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cancel query o k response
func (o *CancelQueryOK) Code() int {
	return 200
}

func (o *CancelQueryOK) Error() string {
	return fmt.Sprintf("[DELETE /query/{brokerId}/{queryId}][%d] cancelQueryOK ", 200)
}

func (o *CancelQueryOK) String() string {
	return fmt.Sprintf("[DELETE /query/{brokerId}/{queryId}][%d] cancelQueryOK ", 200)
}

func (o *CancelQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancelQueryNotFound creates a CancelQueryNotFound with default headers values
func NewCancelQueryNotFound() *CancelQueryNotFound {
	return &CancelQueryNotFound{}
}

/*
CancelQueryNotFound describes a response with status code 404, with default header values.

Query not found on the requested broker
*/
type CancelQueryNotFound struct {
}

// IsSuccess returns true when this cancel query not found response has a 2xx status code
func (o *CancelQueryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel query not found response has a 3xx status code
func (o *CancelQueryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel query not found response has a 4xx status code
func (o *CancelQueryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel query not found response has a 5xx status code
func (o *CancelQueryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel query not found response a status code equal to that given
func (o *CancelQueryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the cancel query not found response
func (o *CancelQueryNotFound) Code() int {
	return 404
}

func (o *CancelQueryNotFound) Error() string {
	return fmt.Sprintf("[DELETE /query/{brokerId}/{queryId}][%d] cancelQueryNotFound ", 404)
}

func (o *CancelQueryNotFound) String() string {
	return fmt.Sprintf("[DELETE /query/{brokerId}/{queryId}][%d] cancelQueryNotFound ", 404)
}

func (o *CancelQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancelQueryInternalServerError creates a CancelQueryInternalServerError with default headers values
func NewCancelQueryInternalServerError() *CancelQueryInternalServerError {
	return &CancelQueryInternalServerError{}
}

/*
CancelQueryInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CancelQueryInternalServerError struct {
}

// IsSuccess returns true when this cancel query internal server error response has a 2xx status code
func (o *CancelQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel query internal server error response has a 3xx status code
func (o *CancelQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel query internal server error response has a 4xx status code
func (o *CancelQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel query internal server error response has a 5xx status code
func (o *CancelQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cancel query internal server error response a status code equal to that given
func (o *CancelQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the cancel query internal server error response
func (o *CancelQueryInternalServerError) Code() int {
	return 500
}

func (o *CancelQueryInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /query/{brokerId}/{queryId}][%d] cancelQueryInternalServerError ", 500)
}

func (o *CancelQueryInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /query/{brokerId}/{queryId}][%d] cancelQueryInternalServerError ", 500)
}

func (o *CancelQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
