// Code generated by go-swagger; DO NOT EDIT.

package atomic_ingestion

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EndDataIngestRequestReader is a Reader for the EndDataIngestRequest structure.
type EndDataIngestRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EndDataIngestRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewEndDataIngestRequestDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewEndDataIngestRequestDefault creates a EndDataIngestRequestDefault with default headers values
func NewEndDataIngestRequestDefault(code int) *EndDataIngestRequestDefault {
	return &EndDataIngestRequestDefault{
		_statusCode: code,
	}
}

/*
EndDataIngestRequestDefault describes a response with status code -1, with default header values.

successful operation
*/
type EndDataIngestRequestDefault struct {
	_statusCode int
}

// IsSuccess returns true when this end data ingest request default response has a 2xx status code
func (o *EndDataIngestRequestDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this end data ingest request default response has a 3xx status code
func (o *EndDataIngestRequestDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this end data ingest request default response has a 4xx status code
func (o *EndDataIngestRequestDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this end data ingest request default response has a 5xx status code
func (o *EndDataIngestRequestDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this end data ingest request default response a status code equal to that given
func (o *EndDataIngestRequestDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the end data ingest request default response
func (o *EndDataIngestRequestDefault) Code() int {
	return o._statusCode
}

func (o *EndDataIngestRequestDefault) Error() string {
	return fmt.Sprintf("[POST /segments/{tableName}/endDataIngestRequest][%d] endDataIngestRequest default ", o._statusCode)
}

func (o *EndDataIngestRequestDefault) String() string {
	return fmt.Sprintf("[POST /segments/{tableName}/endDataIngestRequest][%d] endDataIngestRequest default ", o._statusCode)
}

func (o *EndDataIngestRequestDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
