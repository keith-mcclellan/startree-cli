// Code generated by go-swagger; DO NOT EDIT.

package query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetRunningQueriesReader is a Reader for the GetRunningQueries structure.
type GetRunningQueriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunningQueriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunningQueriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetRunningQueriesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRunningQueriesOK creates a GetRunningQueriesOK with default headers values
func NewGetRunningQueriesOK() *GetRunningQueriesOK {
	return &GetRunningQueriesOK{}
}

/*
GetRunningQueriesOK describes a response with status code 200, with default header values.

Success
*/
type GetRunningQueriesOK struct {
}

// IsSuccess returns true when this get running queries o k response has a 2xx status code
func (o *GetRunningQueriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get running queries o k response has a 3xx status code
func (o *GetRunningQueriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get running queries o k response has a 4xx status code
func (o *GetRunningQueriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get running queries o k response has a 5xx status code
func (o *GetRunningQueriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get running queries o k response a status code equal to that given
func (o *GetRunningQueriesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get running queries o k response
func (o *GetRunningQueriesOK) Code() int {
	return 200
}

func (o *GetRunningQueriesOK) Error() string {
	return fmt.Sprintf("[GET /queries][%d] getRunningQueriesOK ", 200)
}

func (o *GetRunningQueriesOK) String() string {
	return fmt.Sprintf("[GET /queries][%d] getRunningQueriesOK ", 200)
}

func (o *GetRunningQueriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRunningQueriesInternalServerError creates a GetRunningQueriesInternalServerError with default headers values
func NewGetRunningQueriesInternalServerError() *GetRunningQueriesInternalServerError {
	return &GetRunningQueriesInternalServerError{}
}

/*
GetRunningQueriesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetRunningQueriesInternalServerError struct {
}

// IsSuccess returns true when this get running queries internal server error response has a 2xx status code
func (o *GetRunningQueriesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get running queries internal server error response has a 3xx status code
func (o *GetRunningQueriesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get running queries internal server error response has a 4xx status code
func (o *GetRunningQueriesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get running queries internal server error response has a 5xx status code
func (o *GetRunningQueriesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get running queries internal server error response a status code equal to that given
func (o *GetRunningQueriesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get running queries internal server error response
func (o *GetRunningQueriesInternalServerError) Code() int {
	return 500
}

func (o *GetRunningQueriesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /queries][%d] getRunningQueriesInternalServerError ", 500)
}

func (o *GetRunningQueriesInternalServerError) String() string {
	return fmt.Sprintf("[GET /queries][%d] getRunningQueriesInternalServerError ", 500)
}

func (o *GetRunningQueriesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
