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

// RemoveInstancePartitionsReader is a Reader for the RemoveInstancePartitions structure.
type RemoveInstancePartitionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveInstancePartitionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveInstancePartitionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveInstancePartitionsOK creates a RemoveInstancePartitionsOK with default headers values
func NewRemoveInstancePartitionsOK() *RemoveInstancePartitionsOK {
	return &RemoveInstancePartitionsOK{}
}

/*
RemoveInstancePartitionsOK describes a response with status code 200, with default header values.

successful operation
*/
type RemoveInstancePartitionsOK struct {
	Payload *models.SuccessResponse
}

// IsSuccess returns true when this remove instance partitions o k response has a 2xx status code
func (o *RemoveInstancePartitionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove instance partitions o k response has a 3xx status code
func (o *RemoveInstancePartitionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove instance partitions o k response has a 4xx status code
func (o *RemoveInstancePartitionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove instance partitions o k response has a 5xx status code
func (o *RemoveInstancePartitionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this remove instance partitions o k response a status code equal to that given
func (o *RemoveInstancePartitionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the remove instance partitions o k response
func (o *RemoveInstancePartitionsOK) Code() int {
	return 200
}

func (o *RemoveInstancePartitionsOK) Error() string {
	return fmt.Sprintf("[DELETE /tables/{tableName}/instancePartitions][%d] removeInstancePartitionsOK  %+v", 200, o.Payload)
}

func (o *RemoveInstancePartitionsOK) String() string {
	return fmt.Sprintf("[DELETE /tables/{tableName}/instancePartitions][%d] removeInstancePartitionsOK  %+v", 200, o.Payload)
}

func (o *RemoveInstancePartitionsOK) GetPayload() *models.SuccessResponse {
	return o.Payload
}

func (o *RemoveInstancePartitionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}