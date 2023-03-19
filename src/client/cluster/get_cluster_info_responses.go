// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetClusterInfoReader is a Reader for the GetClusterInfo structure.
type GetClusterInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetClusterInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetClusterInfoOK creates a GetClusterInfoOK with default headers values
func NewGetClusterInfoOK() *GetClusterInfoOK {
	return &GetClusterInfoOK{}
}

/*
GetClusterInfoOK describes a response with status code 200, with default header values.

Success
*/
type GetClusterInfoOK struct {
}

// IsSuccess returns true when this get cluster info o k response has a 2xx status code
func (o *GetClusterInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster info o k response has a 3xx status code
func (o *GetClusterInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster info o k response has a 4xx status code
func (o *GetClusterInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster info o k response has a 5xx status code
func (o *GetClusterInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster info o k response a status code equal to that given
func (o *GetClusterInfoOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cluster info o k response
func (o *GetClusterInfoOK) Code() int {
	return 200
}

func (o *GetClusterInfoOK) Error() string {
	return fmt.Sprintf("[GET /cluster/info][%d] getClusterInfoOK ", 200)
}

func (o *GetClusterInfoOK) String() string {
	return fmt.Sprintf("[GET /cluster/info][%d] getClusterInfoOK ", 200)
}

func (o *GetClusterInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterInfoInternalServerError creates a GetClusterInfoInternalServerError with default headers values
func NewGetClusterInfoInternalServerError() *GetClusterInfoInternalServerError {
	return &GetClusterInfoInternalServerError{}
}

/*
GetClusterInfoInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetClusterInfoInternalServerError struct {
}

// IsSuccess returns true when this get cluster info internal server error response has a 2xx status code
func (o *GetClusterInfoInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster info internal server error response has a 3xx status code
func (o *GetClusterInfoInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster info internal server error response has a 4xx status code
func (o *GetClusterInfoInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster info internal server error response has a 5xx status code
func (o *GetClusterInfoInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get cluster info internal server error response a status code equal to that given
func (o *GetClusterInfoInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get cluster info internal server error response
func (o *GetClusterInfoInternalServerError) Code() int {
	return 500
}

func (o *GetClusterInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cluster/info][%d] getClusterInfoInternalServerError ", 500)
}

func (o *GetClusterInfoInternalServerError) String() string {
	return fmt.Sprintf("[GET /cluster/info][%d] getClusterInfoInternalServerError ", 500)
}

func (o *GetClusterInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
