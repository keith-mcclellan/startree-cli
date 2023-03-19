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

// GetCronSchedulerJobKeysReader is a Reader for the GetCronSchedulerJobKeys structure.
type GetCronSchedulerJobKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCronSchedulerJobKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCronSchedulerJobKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCronSchedulerJobKeysOK creates a GetCronSchedulerJobKeysOK with default headers values
func NewGetCronSchedulerJobKeysOK() *GetCronSchedulerJobKeysOK {
	return &GetCronSchedulerJobKeysOK{}
}

/*
GetCronSchedulerJobKeysOK describes a response with status code 200, with default header values.

successful operation
*/
type GetCronSchedulerJobKeysOK struct {
	Payload []*models.JobKey
}

// IsSuccess returns true when this get cron scheduler job keys o k response has a 2xx status code
func (o *GetCronSchedulerJobKeysOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cron scheduler job keys o k response has a 3xx status code
func (o *GetCronSchedulerJobKeysOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cron scheduler job keys o k response has a 4xx status code
func (o *GetCronSchedulerJobKeysOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cron scheduler job keys o k response has a 5xx status code
func (o *GetCronSchedulerJobKeysOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cron scheduler job keys o k response a status code equal to that given
func (o *GetCronSchedulerJobKeysOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cron scheduler job keys o k response
func (o *GetCronSchedulerJobKeysOK) Code() int {
	return 200
}

func (o *GetCronSchedulerJobKeysOK) Error() string {
	return fmt.Sprintf("[GET /tasks/scheduler/jobKeys][%d] getCronSchedulerJobKeysOK  %+v", 200, o.Payload)
}

func (o *GetCronSchedulerJobKeysOK) String() string {
	return fmt.Sprintf("[GET /tasks/scheduler/jobKeys][%d] getCronSchedulerJobKeysOK  %+v", 200, o.Payload)
}

func (o *GetCronSchedulerJobKeysOK) GetPayload() []*models.JobKey {
	return o.Payload
}

func (o *GetCronSchedulerJobKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
