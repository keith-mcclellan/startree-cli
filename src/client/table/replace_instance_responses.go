// Code generated by go-swagger; DO NOT EDIT.

package table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"startree.ai/cli/models"
)

// ReplaceInstanceReader is a Reader for the ReplaceInstance structure.
type ReplaceInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceInstanceOK creates a ReplaceInstanceOK with default headers values
func NewReplaceInstanceOK() *ReplaceInstanceOK {
	return &ReplaceInstanceOK{}
}

/*
ReplaceInstanceOK describes a response with status code 200, with default header values.

successful operation
*/
type ReplaceInstanceOK struct {
	Payload map[string]models.InstancePartitions
}

// IsSuccess returns true when this replace instance o k response has a 2xx status code
func (o *ReplaceInstanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this replace instance o k response has a 3xx status code
func (o *ReplaceInstanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this replace instance o k response has a 4xx status code
func (o *ReplaceInstanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this replace instance o k response has a 5xx status code
func (o *ReplaceInstanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this replace instance o k response a status code equal to that given
func (o *ReplaceInstanceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the replace instance o k response
func (o *ReplaceInstanceOK) Code() int {
	return 200
}

func (o *ReplaceInstanceOK) Error() string {
	return fmt.Sprintf("[POST /tables/{tableName}/replaceInstance][%d] replaceInstanceOK  %+v", 200, o.Payload)
}

func (o *ReplaceInstanceOK) String() string {
	return fmt.Sprintf("[POST /tables/{tableName}/replaceInstance][%d] replaceInstanceOK  %+v", 200, o.Payload)
}

func (o *ReplaceInstanceOK) GetPayload() map[string]models.InstancePartitions {
	return o.Payload
}

func (o *ReplaceInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}