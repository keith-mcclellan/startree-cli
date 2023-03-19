// Code generated by go-swagger; DO NOT EDIT.

package segment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EndReplaceSegmentsReader is a Reader for the EndReplaceSegments structure.
type EndReplaceSegmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EndReplaceSegmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewEndReplaceSegmentsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewEndReplaceSegmentsDefault creates a EndReplaceSegmentsDefault with default headers values
func NewEndReplaceSegmentsDefault(code int) *EndReplaceSegmentsDefault {
	return &EndReplaceSegmentsDefault{
		_statusCode: code,
	}
}

/*
EndReplaceSegmentsDefault describes a response with status code -1, with default header values.

successful operation
*/
type EndReplaceSegmentsDefault struct {
	_statusCode int
}

// IsSuccess returns true when this end replace segments default response has a 2xx status code
func (o *EndReplaceSegmentsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this end replace segments default response has a 3xx status code
func (o *EndReplaceSegmentsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this end replace segments default response has a 4xx status code
func (o *EndReplaceSegmentsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this end replace segments default response has a 5xx status code
func (o *EndReplaceSegmentsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this end replace segments default response a status code equal to that given
func (o *EndReplaceSegmentsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the end replace segments default response
func (o *EndReplaceSegmentsDefault) Code() int {
	return o._statusCode
}

func (o *EndReplaceSegmentsDefault) Error() string {
	return fmt.Sprintf("[POST /segments/{tableName}/endReplaceSegments][%d] endReplaceSegments default ", o._statusCode)
}

func (o *EndReplaceSegmentsDefault) String() string {
	return fmt.Sprintf("[POST /segments/{tableName}/endReplaceSegments][%d] endReplaceSegments default ", o._statusCode)
}

func (o *EndReplaceSegmentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
