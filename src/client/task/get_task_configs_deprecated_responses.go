// Code generated by go-swagger; DO NOT EDIT.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"startree.ai/cli/models"
)

// GetTaskConfigsDeprecatedReader is a Reader for the GetTaskConfigsDeprecated structure.
type GetTaskConfigsDeprecatedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskConfigsDeprecatedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTaskConfigsDeprecatedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTaskConfigsDeprecatedOK creates a GetTaskConfigsDeprecatedOK with default headers values
func NewGetTaskConfigsDeprecatedOK() *GetTaskConfigsDeprecatedOK {
	return &GetTaskConfigsDeprecatedOK{}
}

/*
GetTaskConfigsDeprecatedOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTaskConfigsDeprecatedOK struct {
	Payload []*models.PinotTaskConfig
}

// IsSuccess returns true when this get task configs deprecated o k response has a 2xx status code
func (o *GetTaskConfigsDeprecatedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get task configs deprecated o k response has a 3xx status code
func (o *GetTaskConfigsDeprecatedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task configs deprecated o k response has a 4xx status code
func (o *GetTaskConfigsDeprecatedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get task configs deprecated o k response has a 5xx status code
func (o *GetTaskConfigsDeprecatedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get task configs deprecated o k response a status code equal to that given
func (o *GetTaskConfigsDeprecatedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get task configs deprecated o k response
func (o *GetTaskConfigsDeprecatedOK) Code() int {
	return 200
}

func (o *GetTaskConfigsDeprecatedOK) Error() string {
	return fmt.Sprintf("[GET /tasks/taskconfig/{taskName}][%d] getTaskConfigsDeprecatedOK  %+v", 200, o.Payload)
}

func (o *GetTaskConfigsDeprecatedOK) String() string {
	return fmt.Sprintf("[GET /tasks/taskconfig/{taskName}][%d] getTaskConfigsDeprecatedOK  %+v", 200, o.Payload)
}

func (o *GetTaskConfigsDeprecatedOK) GetPayload() []*models.PinotTaskConfig {
	return o.Payload
}

func (o *GetTaskConfigsDeprecatedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}