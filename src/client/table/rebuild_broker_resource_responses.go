// Code generated by go-swagger; DO NOT EDIT.

package table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RebuildBrokerResourceReader is a Reader for the RebuildBrokerResource structure.
type RebuildBrokerResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RebuildBrokerResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRebuildBrokerResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRebuildBrokerResourceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRebuildBrokerResourceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRebuildBrokerResourceOK creates a RebuildBrokerResourceOK with default headers values
func NewRebuildBrokerResourceOK() *RebuildBrokerResourceOK {
	return &RebuildBrokerResourceOK{}
}

/*
RebuildBrokerResourceOK describes a response with status code 200, with default header values.

Success
*/
type RebuildBrokerResourceOK struct {
}

// IsSuccess returns true when this rebuild broker resource o k response has a 2xx status code
func (o *RebuildBrokerResourceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rebuild broker resource o k response has a 3xx status code
func (o *RebuildBrokerResourceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rebuild broker resource o k response has a 4xx status code
func (o *RebuildBrokerResourceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this rebuild broker resource o k response has a 5xx status code
func (o *RebuildBrokerResourceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this rebuild broker resource o k response a status code equal to that given
func (o *RebuildBrokerResourceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the rebuild broker resource o k response
func (o *RebuildBrokerResourceOK) Code() int {
	return 200
}

func (o *RebuildBrokerResourceOK) Error() string {
	return fmt.Sprintf("[POST /tables/{tableName}/rebuildBrokerResourceFromHelixTags][%d] rebuildBrokerResourceOK ", 200)
}

func (o *RebuildBrokerResourceOK) String() string {
	return fmt.Sprintf("[POST /tables/{tableName}/rebuildBrokerResourceFromHelixTags][%d] rebuildBrokerResourceOK ", 200)
}

func (o *RebuildBrokerResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRebuildBrokerResourceBadRequest creates a RebuildBrokerResourceBadRequest with default headers values
func NewRebuildBrokerResourceBadRequest() *RebuildBrokerResourceBadRequest {
	return &RebuildBrokerResourceBadRequest{}
}

/*
RebuildBrokerResourceBadRequest describes a response with status code 400, with default header values.

Bad request: table name has to be with table type
*/
type RebuildBrokerResourceBadRequest struct {
}

// IsSuccess returns true when this rebuild broker resource bad request response has a 2xx status code
func (o *RebuildBrokerResourceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rebuild broker resource bad request response has a 3xx status code
func (o *RebuildBrokerResourceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rebuild broker resource bad request response has a 4xx status code
func (o *RebuildBrokerResourceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this rebuild broker resource bad request response has a 5xx status code
func (o *RebuildBrokerResourceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this rebuild broker resource bad request response a status code equal to that given
func (o *RebuildBrokerResourceBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the rebuild broker resource bad request response
func (o *RebuildBrokerResourceBadRequest) Code() int {
	return 400
}

func (o *RebuildBrokerResourceBadRequest) Error() string {
	return fmt.Sprintf("[POST /tables/{tableName}/rebuildBrokerResourceFromHelixTags][%d] rebuildBrokerResourceBadRequest ", 400)
}

func (o *RebuildBrokerResourceBadRequest) String() string {
	return fmt.Sprintf("[POST /tables/{tableName}/rebuildBrokerResourceFromHelixTags][%d] rebuildBrokerResourceBadRequest ", 400)
}

func (o *RebuildBrokerResourceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRebuildBrokerResourceInternalServerError creates a RebuildBrokerResourceInternalServerError with default headers values
func NewRebuildBrokerResourceInternalServerError() *RebuildBrokerResourceInternalServerError {
	return &RebuildBrokerResourceInternalServerError{}
}

/*
RebuildBrokerResourceInternalServerError describes a response with status code 500, with default header values.

Internal error rebuilding broker resource or serializing response
*/
type RebuildBrokerResourceInternalServerError struct {
}

// IsSuccess returns true when this rebuild broker resource internal server error response has a 2xx status code
func (o *RebuildBrokerResourceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rebuild broker resource internal server error response has a 3xx status code
func (o *RebuildBrokerResourceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rebuild broker resource internal server error response has a 4xx status code
func (o *RebuildBrokerResourceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this rebuild broker resource internal server error response has a 5xx status code
func (o *RebuildBrokerResourceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this rebuild broker resource internal server error response a status code equal to that given
func (o *RebuildBrokerResourceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the rebuild broker resource internal server error response
func (o *RebuildBrokerResourceInternalServerError) Code() int {
	return 500
}

func (o *RebuildBrokerResourceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /tables/{tableName}/rebuildBrokerResourceFromHelixTags][%d] rebuildBrokerResourceInternalServerError ", 500)
}

func (o *RebuildBrokerResourceInternalServerError) String() string {
	return fmt.Sprintf("[POST /tables/{tableName}/rebuildBrokerResourceFromHelixTags][%d] rebuildBrokerResourceInternalServerError ", 500)
}

func (o *RebuildBrokerResourceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
