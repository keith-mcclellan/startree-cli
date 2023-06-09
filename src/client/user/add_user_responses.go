// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"startree.ai/cli/models"
)

// AddUserReader is a Reader for the AddUser structure.
type AddUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddUserOK creates a AddUserOK with default headers values
func NewAddUserOK() *AddUserOK {
	return &AddUserOK{}
}

/*
AddUserOK describes a response with status code 200, with default header values.

successful operation
*/
type AddUserOK struct {
	Payload *models.SuccessResponse
}

// IsSuccess returns true when this add user o k response has a 2xx status code
func (o *AddUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add user o k response has a 3xx status code
func (o *AddUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user o k response has a 4xx status code
func (o *AddUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add user o k response has a 5xx status code
func (o *AddUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add user o k response a status code equal to that given
func (o *AddUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add user o k response
func (o *AddUserOK) Code() int {
	return 200
}

func (o *AddUserOK) Error() string {
	return fmt.Sprintf("[POST /users][%d] addUserOK  %+v", 200, o.Payload)
}

func (o *AddUserOK) String() string {
	return fmt.Sprintf("[POST /users][%d] addUserOK  %+v", 200, o.Payload)
}

func (o *AddUserOK) GetPayload() *models.SuccessResponse {
	return o.Payload
}

func (o *AddUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
