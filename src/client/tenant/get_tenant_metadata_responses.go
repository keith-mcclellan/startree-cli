// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"startree.ai/cli/models"
)

// GetTenantMetadataReader is a Reader for the GetTenantMetadata structure.
type GetTenantMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTenantMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTenantMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTenantMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTenantMetadataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTenantMetadataOK creates a GetTenantMetadataOK with default headers values
func NewGetTenantMetadataOK() *GetTenantMetadataOK {
	return &GetTenantMetadataOK{}
}

/*
GetTenantMetadataOK describes a response with status code 200, with default header values.

Success
*/
type GetTenantMetadataOK struct {
	Payload *models.TenantMetadata
}

// IsSuccess returns true when this get tenant metadata o k response has a 2xx status code
func (o *GetTenantMetadataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tenant metadata o k response has a 3xx status code
func (o *GetTenantMetadataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tenant metadata o k response has a 4xx status code
func (o *GetTenantMetadataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tenant metadata o k response has a 5xx status code
func (o *GetTenantMetadataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tenant metadata o k response a status code equal to that given
func (o *GetTenantMetadataOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get tenant metadata o k response
func (o *GetTenantMetadataOK) Code() int {
	return 200
}

func (o *GetTenantMetadataOK) Error() string {
	return fmt.Sprintf("[GET /tenants/{tenantName}/metadata][%d] getTenantMetadataOK  %+v", 200, o.Payload)
}

func (o *GetTenantMetadataOK) String() string {
	return fmt.Sprintf("[GET /tenants/{tenantName}/metadata][%d] getTenantMetadataOK  %+v", 200, o.Payload)
}

func (o *GetTenantMetadataOK) GetPayload() *models.TenantMetadata {
	return o.Payload
}

func (o *GetTenantMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTenantMetadataNotFound creates a GetTenantMetadataNotFound with default headers values
func NewGetTenantMetadataNotFound() *GetTenantMetadataNotFound {
	return &GetTenantMetadataNotFound{}
}

/*
GetTenantMetadataNotFound describes a response with status code 404, with default header values.

Tenant not found
*/
type GetTenantMetadataNotFound struct {
}

// IsSuccess returns true when this get tenant metadata not found response has a 2xx status code
func (o *GetTenantMetadataNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tenant metadata not found response has a 3xx status code
func (o *GetTenantMetadataNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tenant metadata not found response has a 4xx status code
func (o *GetTenantMetadataNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tenant metadata not found response has a 5xx status code
func (o *GetTenantMetadataNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get tenant metadata not found response a status code equal to that given
func (o *GetTenantMetadataNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get tenant metadata not found response
func (o *GetTenantMetadataNotFound) Code() int {
	return 404
}

func (o *GetTenantMetadataNotFound) Error() string {
	return fmt.Sprintf("[GET /tenants/{tenantName}/metadata][%d] getTenantMetadataNotFound ", 404)
}

func (o *GetTenantMetadataNotFound) String() string {
	return fmt.Sprintf("[GET /tenants/{tenantName}/metadata][%d] getTenantMetadataNotFound ", 404)
}

func (o *GetTenantMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTenantMetadataInternalServerError creates a GetTenantMetadataInternalServerError with default headers values
func NewGetTenantMetadataInternalServerError() *GetTenantMetadataInternalServerError {
	return &GetTenantMetadataInternalServerError{}
}

/*
GetTenantMetadataInternalServerError describes a response with status code 500, with default header values.

Server error reading tenant information
*/
type GetTenantMetadataInternalServerError struct {
}

// IsSuccess returns true when this get tenant metadata internal server error response has a 2xx status code
func (o *GetTenantMetadataInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tenant metadata internal server error response has a 3xx status code
func (o *GetTenantMetadataInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tenant metadata internal server error response has a 4xx status code
func (o *GetTenantMetadataInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tenant metadata internal server error response has a 5xx status code
func (o *GetTenantMetadataInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get tenant metadata internal server error response a status code equal to that given
func (o *GetTenantMetadataInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get tenant metadata internal server error response
func (o *GetTenantMetadataInternalServerError) Code() int {
	return 500
}

func (o *GetTenantMetadataInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tenants/{tenantName}/metadata][%d] getTenantMetadataInternalServerError ", 500)
}

func (o *GetTenantMetadataInternalServerError) String() string {
	return fmt.Sprintf("[GET /tenants/{tenantName}/metadata][%d] getTenantMetadataInternalServerError ", 500)
}

func (o *GetTenantMetadataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}