// Code generated by go-swagger; DO NOT EDIT.

package instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateInstanceTagsReader is a Reader for the UpdateInstanceTags structure.
type UpdateInstanceTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateInstanceTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateInstanceTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateInstanceTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateInstanceTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateInstanceTagsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateInstanceTagsOK creates a UpdateInstanceTagsOK with default headers values
func NewUpdateInstanceTagsOK() *UpdateInstanceTagsOK {
	return &UpdateInstanceTagsOK{}
}

/*
UpdateInstanceTagsOK describes a response with status code 200, with default header values.

Success
*/
type UpdateInstanceTagsOK struct {
}

// IsSuccess returns true when this update instance tags o k response has a 2xx status code
func (o *UpdateInstanceTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update instance tags o k response has a 3xx status code
func (o *UpdateInstanceTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update instance tags o k response has a 4xx status code
func (o *UpdateInstanceTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update instance tags o k response has a 5xx status code
func (o *UpdateInstanceTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update instance tags o k response a status code equal to that given
func (o *UpdateInstanceTagsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update instance tags o k response
func (o *UpdateInstanceTagsOK) Code() int {
	return 200
}

func (o *UpdateInstanceTagsOK) Error() string {
	return fmt.Sprintf("[PUT /instances/{instanceName}/updateTags][%d] updateInstanceTagsOK ", 200)
}

func (o *UpdateInstanceTagsOK) String() string {
	return fmt.Sprintf("[PUT /instances/{instanceName}/updateTags][%d] updateInstanceTagsOK ", 200)
}

func (o *UpdateInstanceTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstanceTagsBadRequest creates a UpdateInstanceTagsBadRequest with default headers values
func NewUpdateInstanceTagsBadRequest() *UpdateInstanceTagsBadRequest {
	return &UpdateInstanceTagsBadRequest{}
}

/*
UpdateInstanceTagsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateInstanceTagsBadRequest struct {
}

// IsSuccess returns true when this update instance tags bad request response has a 2xx status code
func (o *UpdateInstanceTagsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update instance tags bad request response has a 3xx status code
func (o *UpdateInstanceTagsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update instance tags bad request response has a 4xx status code
func (o *UpdateInstanceTagsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update instance tags bad request response has a 5xx status code
func (o *UpdateInstanceTagsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update instance tags bad request response a status code equal to that given
func (o *UpdateInstanceTagsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update instance tags bad request response
func (o *UpdateInstanceTagsBadRequest) Code() int {
	return 400
}

func (o *UpdateInstanceTagsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /instances/{instanceName}/updateTags][%d] updateInstanceTagsBadRequest ", 400)
}

func (o *UpdateInstanceTagsBadRequest) String() string {
	return fmt.Sprintf("[PUT /instances/{instanceName}/updateTags][%d] updateInstanceTagsBadRequest ", 400)
}

func (o *UpdateInstanceTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstanceTagsNotFound creates a UpdateInstanceTagsNotFound with default headers values
func NewUpdateInstanceTagsNotFound() *UpdateInstanceTagsNotFound {
	return &UpdateInstanceTagsNotFound{}
}

/*
UpdateInstanceTagsNotFound describes a response with status code 404, with default header values.

Instance not found
*/
type UpdateInstanceTagsNotFound struct {
}

// IsSuccess returns true when this update instance tags not found response has a 2xx status code
func (o *UpdateInstanceTagsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update instance tags not found response has a 3xx status code
func (o *UpdateInstanceTagsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update instance tags not found response has a 4xx status code
func (o *UpdateInstanceTagsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update instance tags not found response has a 5xx status code
func (o *UpdateInstanceTagsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update instance tags not found response a status code equal to that given
func (o *UpdateInstanceTagsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update instance tags not found response
func (o *UpdateInstanceTagsNotFound) Code() int {
	return 404
}

func (o *UpdateInstanceTagsNotFound) Error() string {
	return fmt.Sprintf("[PUT /instances/{instanceName}/updateTags][%d] updateInstanceTagsNotFound ", 404)
}

func (o *UpdateInstanceTagsNotFound) String() string {
	return fmt.Sprintf("[PUT /instances/{instanceName}/updateTags][%d] updateInstanceTagsNotFound ", 404)
}

func (o *UpdateInstanceTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInstanceTagsInternalServerError creates a UpdateInstanceTagsInternalServerError with default headers values
func NewUpdateInstanceTagsInternalServerError() *UpdateInstanceTagsInternalServerError {
	return &UpdateInstanceTagsInternalServerError{}
}

/*
UpdateInstanceTagsInternalServerError describes a response with status code 500, with default header values.

Internal error
*/
type UpdateInstanceTagsInternalServerError struct {
}

// IsSuccess returns true when this update instance tags internal server error response has a 2xx status code
func (o *UpdateInstanceTagsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update instance tags internal server error response has a 3xx status code
func (o *UpdateInstanceTagsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update instance tags internal server error response has a 4xx status code
func (o *UpdateInstanceTagsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update instance tags internal server error response has a 5xx status code
func (o *UpdateInstanceTagsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update instance tags internal server error response a status code equal to that given
func (o *UpdateInstanceTagsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update instance tags internal server error response
func (o *UpdateInstanceTagsInternalServerError) Code() int {
	return 500
}

func (o *UpdateInstanceTagsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /instances/{instanceName}/updateTags][%d] updateInstanceTagsInternalServerError ", 500)
}

func (o *UpdateInstanceTagsInternalServerError) String() string {
	return fmt.Sprintf("[PUT /instances/{instanceName}/updateTags][%d] updateInstanceTagsInternalServerError ", 500)
}

func (o *UpdateInstanceTagsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
