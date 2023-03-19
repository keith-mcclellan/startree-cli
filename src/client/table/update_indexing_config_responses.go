// Code generated by go-swagger; DO NOT EDIT.

package table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateIndexingConfigReader is a Reader for the UpdateIndexingConfig structure.
type UpdateIndexingConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateIndexingConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateIndexingConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateIndexingConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateIndexingConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateIndexingConfigOK creates a UpdateIndexingConfigOK with default headers values
func NewUpdateIndexingConfigOK() *UpdateIndexingConfigOK {
	return &UpdateIndexingConfigOK{}
}

/*
UpdateIndexingConfigOK describes a response with status code 200, with default header values.

Success
*/
type UpdateIndexingConfigOK struct {
}

// IsSuccess returns true when this update indexing config o k response has a 2xx status code
func (o *UpdateIndexingConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update indexing config o k response has a 3xx status code
func (o *UpdateIndexingConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update indexing config o k response has a 4xx status code
func (o *UpdateIndexingConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update indexing config o k response has a 5xx status code
func (o *UpdateIndexingConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update indexing config o k response a status code equal to that given
func (o *UpdateIndexingConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update indexing config o k response
func (o *UpdateIndexingConfigOK) Code() int {
	return 200
}

func (o *UpdateIndexingConfigOK) Error() string {
	return fmt.Sprintf("[PUT /tables/{tableName}/indexingConfigs][%d] updateIndexingConfigOK ", 200)
}

func (o *UpdateIndexingConfigOK) String() string {
	return fmt.Sprintf("[PUT /tables/{tableName}/indexingConfigs][%d] updateIndexingConfigOK ", 200)
}

func (o *UpdateIndexingConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateIndexingConfigNotFound creates a UpdateIndexingConfigNotFound with default headers values
func NewUpdateIndexingConfigNotFound() *UpdateIndexingConfigNotFound {
	return &UpdateIndexingConfigNotFound{}
}

/*
UpdateIndexingConfigNotFound describes a response with status code 404, with default header values.

Table not found
*/
type UpdateIndexingConfigNotFound struct {
}

// IsSuccess returns true when this update indexing config not found response has a 2xx status code
func (o *UpdateIndexingConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update indexing config not found response has a 3xx status code
func (o *UpdateIndexingConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update indexing config not found response has a 4xx status code
func (o *UpdateIndexingConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update indexing config not found response has a 5xx status code
func (o *UpdateIndexingConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update indexing config not found response a status code equal to that given
func (o *UpdateIndexingConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update indexing config not found response
func (o *UpdateIndexingConfigNotFound) Code() int {
	return 404
}

func (o *UpdateIndexingConfigNotFound) Error() string {
	return fmt.Sprintf("[PUT /tables/{tableName}/indexingConfigs][%d] updateIndexingConfigNotFound ", 404)
}

func (o *UpdateIndexingConfigNotFound) String() string {
	return fmt.Sprintf("[PUT /tables/{tableName}/indexingConfigs][%d] updateIndexingConfigNotFound ", 404)
}

func (o *UpdateIndexingConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateIndexingConfigInternalServerError creates a UpdateIndexingConfigInternalServerError with default headers values
func NewUpdateIndexingConfigInternalServerError() *UpdateIndexingConfigInternalServerError {
	return &UpdateIndexingConfigInternalServerError{}
}

/*
UpdateIndexingConfigInternalServerError describes a response with status code 500, with default header values.

Server error updating configuration
*/
type UpdateIndexingConfigInternalServerError struct {
}

// IsSuccess returns true when this update indexing config internal server error response has a 2xx status code
func (o *UpdateIndexingConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update indexing config internal server error response has a 3xx status code
func (o *UpdateIndexingConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update indexing config internal server error response has a 4xx status code
func (o *UpdateIndexingConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update indexing config internal server error response has a 5xx status code
func (o *UpdateIndexingConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update indexing config internal server error response a status code equal to that given
func (o *UpdateIndexingConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update indexing config internal server error response
func (o *UpdateIndexingConfigInternalServerError) Code() int {
	return 500
}

func (o *UpdateIndexingConfigInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /tables/{tableName}/indexingConfigs][%d] updateIndexingConfigInternalServerError ", 500)
}

func (o *UpdateIndexingConfigInternalServerError) String() string {
	return fmt.Sprintf("[PUT /tables/{tableName}/indexingConfigs][%d] updateIndexingConfigInternalServerError ", 500)
}

func (o *UpdateIndexingConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
