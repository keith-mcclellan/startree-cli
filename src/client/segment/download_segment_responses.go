// Code generated by go-swagger; DO NOT EDIT.

package segment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DownloadSegmentReader is a Reader for the DownloadSegment structure.
type DownloadSegmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DownloadSegmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDownloadSegmentDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDownloadSegmentDefault creates a DownloadSegmentDefault with default headers values
func NewDownloadSegmentDefault(code int) *DownloadSegmentDefault {
	return &DownloadSegmentDefault{
		_statusCode: code,
	}
}

/*
DownloadSegmentDefault describes a response with status code -1, with default header values.

successful operation
*/
type DownloadSegmentDefault struct {
	_statusCode int
}

// IsSuccess returns true when this download segment default response has a 2xx status code
func (o *DownloadSegmentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this download segment default response has a 3xx status code
func (o *DownloadSegmentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this download segment default response has a 4xx status code
func (o *DownloadSegmentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this download segment default response has a 5xx status code
func (o *DownloadSegmentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this download segment default response a status code equal to that given
func (o *DownloadSegmentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the download segment default response
func (o *DownloadSegmentDefault) Code() int {
	return o._statusCode
}

func (o *DownloadSegmentDefault) Error() string {
	return fmt.Sprintf("[GET /segments/{tableName}/{segmentName}][%d] downloadSegment default ", o._statusCode)
}

func (o *DownloadSegmentDefault) String() string {
	return fmt.Sprintf("[GET /segments/{tableName}/{segmentName}][%d] downloadSegment default ", o._statusCode)
}

func (o *DownloadSegmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
