// Code generated by go-swagger; DO NOT EDIT.

package table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ForceCommitReader is a Reader for the ForceCommit structure.
type ForceCommitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ForceCommitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewForceCommitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewForceCommitOK creates a ForceCommitOK with default headers values
func NewForceCommitOK() *ForceCommitOK {
	return &ForceCommitOK{}
}

/*
ForceCommitOK describes a response with status code 200, with default header values.

successful operation
*/
type ForceCommitOK struct {
	Payload map[string]string
}

// IsSuccess returns true when this force commit o k response has a 2xx status code
func (o *ForceCommitOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this force commit o k response has a 3xx status code
func (o *ForceCommitOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this force commit o k response has a 4xx status code
func (o *ForceCommitOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this force commit o k response has a 5xx status code
func (o *ForceCommitOK) IsServerError() bool {
	return false
}

// IsCode returns true when this force commit o k response a status code equal to that given
func (o *ForceCommitOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the force commit o k response
func (o *ForceCommitOK) Code() int {
	return 200
}

func (o *ForceCommitOK) Error() string {
	return fmt.Sprintf("[POST /tables/{tableName}/forceCommit][%d] forceCommitOK  %+v", 200, o.Payload)
}

func (o *ForceCommitOK) String() string {
	return fmt.Sprintf("[POST /tables/{tableName}/forceCommit][%d] forceCommitOK  %+v", 200, o.Payload)
}

func (o *ForceCommitOK) GetPayload() map[string]string {
	return o.Payload
}

func (o *ForceCommitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
