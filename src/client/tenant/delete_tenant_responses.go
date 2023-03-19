// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteTenantReader is a Reader for the DeleteTenant structure.
type DeleteTenantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTenantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTenantOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTenantBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTenantNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTenantInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTenantOK creates a DeleteTenantOK with default headers values
func NewDeleteTenantOK() *DeleteTenantOK {
	return &DeleteTenantOK{}
}

/*
DeleteTenantOK describes a response with status code 200, with default header values.

Success
*/
type DeleteTenantOK struct {
}

// IsSuccess returns true when this delete tenant o k response has a 2xx status code
func (o *DeleteTenantOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete tenant o k response has a 3xx status code
func (o *DeleteTenantOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tenant o k response has a 4xx status code
func (o *DeleteTenantOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete tenant o k response has a 5xx status code
func (o *DeleteTenantOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete tenant o k response a status code equal to that given
func (o *DeleteTenantOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete tenant o k response
func (o *DeleteTenantOK) Code() int {
	return 200
}

func (o *DeleteTenantOK) Error() string {
	return fmt.Sprintf("[DELETE /tenants/{tenantName}][%d] deleteTenantOK ", 200)
}

func (o *DeleteTenantOK) String() string {
	return fmt.Sprintf("[DELETE /tenants/{tenantName}][%d] deleteTenantOK ", 200)
}

func (o *DeleteTenantOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTenantBadRequest creates a DeleteTenantBadRequest with default headers values
func NewDeleteTenantBadRequest() *DeleteTenantBadRequest {
	return &DeleteTenantBadRequest{}
}

/*
DeleteTenantBadRequest describes a response with status code 400, with default header values.

Tenant can not be deleted
*/
type DeleteTenantBadRequest struct {
}

// IsSuccess returns true when this delete tenant bad request response has a 2xx status code
func (o *DeleteTenantBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete tenant bad request response has a 3xx status code
func (o *DeleteTenantBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tenant bad request response has a 4xx status code
func (o *DeleteTenantBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete tenant bad request response has a 5xx status code
func (o *DeleteTenantBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete tenant bad request response a status code equal to that given
func (o *DeleteTenantBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete tenant bad request response
func (o *DeleteTenantBadRequest) Code() int {
	return 400
}

func (o *DeleteTenantBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /tenants/{tenantName}][%d] deleteTenantBadRequest ", 400)
}

func (o *DeleteTenantBadRequest) String() string {
	return fmt.Sprintf("[DELETE /tenants/{tenantName}][%d] deleteTenantBadRequest ", 400)
}

func (o *DeleteTenantBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTenantNotFound creates a DeleteTenantNotFound with default headers values
func NewDeleteTenantNotFound() *DeleteTenantNotFound {
	return &DeleteTenantNotFound{}
}

/*
DeleteTenantNotFound describes a response with status code 404, with default header values.

Tenant not found
*/
type DeleteTenantNotFound struct {
}

// IsSuccess returns true when this delete tenant not found response has a 2xx status code
func (o *DeleteTenantNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete tenant not found response has a 3xx status code
func (o *DeleteTenantNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tenant not found response has a 4xx status code
func (o *DeleteTenantNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete tenant not found response has a 5xx status code
func (o *DeleteTenantNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete tenant not found response a status code equal to that given
func (o *DeleteTenantNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete tenant not found response
func (o *DeleteTenantNotFound) Code() int {
	return 404
}

func (o *DeleteTenantNotFound) Error() string {
	return fmt.Sprintf("[DELETE /tenants/{tenantName}][%d] deleteTenantNotFound ", 404)
}

func (o *DeleteTenantNotFound) String() string {
	return fmt.Sprintf("[DELETE /tenants/{tenantName}][%d] deleteTenantNotFound ", 404)
}

func (o *DeleteTenantNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTenantInternalServerError creates a DeleteTenantInternalServerError with default headers values
func NewDeleteTenantInternalServerError() *DeleteTenantInternalServerError {
	return &DeleteTenantInternalServerError{}
}

/*
DeleteTenantInternalServerError describes a response with status code 500, with default header values.

Error deleting tenant
*/
type DeleteTenantInternalServerError struct {
}

// IsSuccess returns true when this delete tenant internal server error response has a 2xx status code
func (o *DeleteTenantInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete tenant internal server error response has a 3xx status code
func (o *DeleteTenantInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tenant internal server error response has a 4xx status code
func (o *DeleteTenantInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete tenant internal server error response has a 5xx status code
func (o *DeleteTenantInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete tenant internal server error response a status code equal to that given
func (o *DeleteTenantInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete tenant internal server error response
func (o *DeleteTenantInternalServerError) Code() int {
	return 500
}

func (o *DeleteTenantInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /tenants/{tenantName}][%d] deleteTenantInternalServerError ", 500)
}

func (o *DeleteTenantInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /tenants/{tenantName}][%d] deleteTenantInternalServerError ", 500)
}

func (o *DeleteTenantInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}