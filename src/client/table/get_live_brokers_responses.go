// Code generated by go-swagger; DO NOT EDIT.

package table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetLiveBrokersReader is a Reader for the GetLiveBrokers structure.
type GetLiveBrokersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLiveBrokersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLiveBrokersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetLiveBrokersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLiveBrokersOK creates a GetLiveBrokersOK with default headers values
func NewGetLiveBrokersOK() *GetLiveBrokersOK {
	return &GetLiveBrokersOK{}
}

/*
GetLiveBrokersOK describes a response with status code 200, with default header values.

Success
*/
type GetLiveBrokersOK struct {
}

// IsSuccess returns true when this get live brokers o k response has a 2xx status code
func (o *GetLiveBrokersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get live brokers o k response has a 3xx status code
func (o *GetLiveBrokersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get live brokers o k response has a 4xx status code
func (o *GetLiveBrokersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get live brokers o k response has a 5xx status code
func (o *GetLiveBrokersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get live brokers o k response a status code equal to that given
func (o *GetLiveBrokersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get live brokers o k response
func (o *GetLiveBrokersOK) Code() int {
	return 200
}

func (o *GetLiveBrokersOK) Error() string {
	return fmt.Sprintf("[GET /tables/livebrokers][%d] getLiveBrokersOK ", 200)
}

func (o *GetLiveBrokersOK) String() string {
	return fmt.Sprintf("[GET /tables/livebrokers][%d] getLiveBrokersOK ", 200)
}

func (o *GetLiveBrokersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLiveBrokersInternalServerError creates a GetLiveBrokersInternalServerError with default headers values
func NewGetLiveBrokersInternalServerError() *GetLiveBrokersInternalServerError {
	return &GetLiveBrokersInternalServerError{}
}

/*
GetLiveBrokersInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetLiveBrokersInternalServerError struct {
}

// IsSuccess returns true when this get live brokers internal server error response has a 2xx status code
func (o *GetLiveBrokersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get live brokers internal server error response has a 3xx status code
func (o *GetLiveBrokersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get live brokers internal server error response has a 4xx status code
func (o *GetLiveBrokersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get live brokers internal server error response has a 5xx status code
func (o *GetLiveBrokersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get live brokers internal server error response a status code equal to that given
func (o *GetLiveBrokersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get live brokers internal server error response
func (o *GetLiveBrokersInternalServerError) Code() int {
	return 500
}

func (o *GetLiveBrokersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tables/livebrokers][%d] getLiveBrokersInternalServerError ", 500)
}

func (o *GetLiveBrokersInternalServerError) String() string {
	return fmt.Sprintf("[GET /tables/livebrokers][%d] getLiveBrokersInternalServerError ", 500)
}

func (o *GetLiveBrokersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
