// Code generated by go-swagger; DO NOT EDIT.

package table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IngestFromFileReader is a Reader for the IngestFromFile structure.
type IngestFromFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IngestFromFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewIngestFromFileDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewIngestFromFileDefault creates a IngestFromFileDefault with default headers values
func NewIngestFromFileDefault(code int) *IngestFromFileDefault {
	return &IngestFromFileDefault{
		_statusCode: code,
	}
}

/*
IngestFromFileDefault describes a response with status code -1, with default header values.

successful operation
*/
type IngestFromFileDefault struct {
	_statusCode int
}

// IsSuccess returns true when this ingest from file default response has a 2xx status code
func (o *IngestFromFileDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ingest from file default response has a 3xx status code
func (o *IngestFromFileDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ingest from file default response has a 4xx status code
func (o *IngestFromFileDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ingest from file default response has a 5xx status code
func (o *IngestFromFileDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ingest from file default response a status code equal to that given
func (o *IngestFromFileDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ingest from file default response
func (o *IngestFromFileDefault) Code() int {
	return o._statusCode
}

func (o *IngestFromFileDefault) Error() string {
	return fmt.Sprintf("[POST /ingestFromFile][%d] ingestFromFile default ", o._statusCode)
}

func (o *IngestFromFileDefault) String() string {
	return fmt.Sprintf("[POST /ingestFromFile][%d] ingestFromFile default ", o._statusCode)
}

func (o *IngestFromFileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}