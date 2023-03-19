// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetAllTenantsReader is a Reader for the GetAllTenants structure.
type GetAllTenantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllTenantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllTenantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetAllTenantsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllTenantsOK creates a GetAllTenantsOK with default headers values
func NewGetAllTenantsOK() *GetAllTenantsOK {
	return &GetAllTenantsOK{}
}

/*
GetAllTenantsOK describes a response with status code 200, with default header values.

Success
*/
type GetAllTenantsOK struct {
}

// IsSuccess returns true when this get all tenants o k response has a 2xx status code
func (o *GetAllTenantsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all tenants o k response has a 3xx status code
func (o *GetAllTenantsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all tenants o k response has a 4xx status code
func (o *GetAllTenantsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all tenants o k response has a 5xx status code
func (o *GetAllTenantsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all tenants o k response a status code equal to that given
func (o *GetAllTenantsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all tenants o k response
func (o *GetAllTenantsOK) Code() int {
	return 200
}

func (o *GetAllTenantsOK) Error() string {
	return fmt.Sprintf("[GET /tenants][%d] getAllTenantsOK ", 200)
}

func (o *GetAllTenantsOK) String() string {
	return fmt.Sprintf("[GET /tenants][%d] getAllTenantsOK ", 200)
}

func (o *GetAllTenantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllTenantsInternalServerError creates a GetAllTenantsInternalServerError with default headers values
func NewGetAllTenantsInternalServerError() *GetAllTenantsInternalServerError {
	return &GetAllTenantsInternalServerError{}
}

/*
GetAllTenantsInternalServerError describes a response with status code 500, with default header values.

Error reading tenants list
*/
type GetAllTenantsInternalServerError struct {
}

// IsSuccess returns true when this get all tenants internal server error response has a 2xx status code
func (o *GetAllTenantsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all tenants internal server error response has a 3xx status code
func (o *GetAllTenantsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all tenants internal server error response has a 4xx status code
func (o *GetAllTenantsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all tenants internal server error response has a 5xx status code
func (o *GetAllTenantsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get all tenants internal server error response a status code equal to that given
func (o *GetAllTenantsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get all tenants internal server error response
func (o *GetAllTenantsInternalServerError) Code() int {
	return 500
}

func (o *GetAllTenantsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tenants][%d] getAllTenantsInternalServerError ", 500)
}

func (o *GetAllTenantsInternalServerError) String() string {
	return fmt.Sprintf("[GET /tenants][%d] getAllTenantsInternalServerError ", 500)
}

func (o *GetAllTenantsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
