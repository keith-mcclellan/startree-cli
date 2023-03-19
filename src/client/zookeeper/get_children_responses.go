// Code generated by go-swagger; DO NOT EDIT.

package zookeeper

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetChildrenReader is a Reader for the GetChildren structure.
type GetChildrenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChildrenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetChildrenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetChildrenNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetChildrenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetChildrenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetChildrenOK creates a GetChildrenOK with default headers values
func NewGetChildrenOK() *GetChildrenOK {
	return &GetChildrenOK{}
}

/*
GetChildrenOK describes a response with status code 200, with default header values.

Success
*/
type GetChildrenOK struct {
}

// IsSuccess returns true when this get children o k response has a 2xx status code
func (o *GetChildrenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get children o k response has a 3xx status code
func (o *GetChildrenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get children o k response has a 4xx status code
func (o *GetChildrenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get children o k response has a 5xx status code
func (o *GetChildrenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get children o k response a status code equal to that given
func (o *GetChildrenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get children o k response
func (o *GetChildrenOK) Code() int {
	return 200
}

func (o *GetChildrenOK) Error() string {
	return fmt.Sprintf("[GET /zk/getChildren][%d] getChildrenOK ", 200)
}

func (o *GetChildrenOK) String() string {
	return fmt.Sprintf("[GET /zk/getChildren][%d] getChildrenOK ", 200)
}

func (o *GetChildrenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetChildrenNoContent creates a GetChildrenNoContent with default headers values
func NewGetChildrenNoContent() *GetChildrenNoContent {
	return &GetChildrenNoContent{}
}

/*
GetChildrenNoContent describes a response with status code 204, with default header values.

No Content
*/
type GetChildrenNoContent struct {
}

// IsSuccess returns true when this get children no content response has a 2xx status code
func (o *GetChildrenNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get children no content response has a 3xx status code
func (o *GetChildrenNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get children no content response has a 4xx status code
func (o *GetChildrenNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this get children no content response has a 5xx status code
func (o *GetChildrenNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this get children no content response a status code equal to that given
func (o *GetChildrenNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the get children no content response
func (o *GetChildrenNoContent) Code() int {
	return 204
}

func (o *GetChildrenNoContent) Error() string {
	return fmt.Sprintf("[GET /zk/getChildren][%d] getChildrenNoContent ", 204)
}

func (o *GetChildrenNoContent) String() string {
	return fmt.Sprintf("[GET /zk/getChildren][%d] getChildrenNoContent ", 204)
}

func (o *GetChildrenNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetChildrenNotFound creates a GetChildrenNotFound with default headers values
func NewGetChildrenNotFound() *GetChildrenNotFound {
	return &GetChildrenNotFound{}
}

/*
GetChildrenNotFound describes a response with status code 404, with default header values.

ZK Path not found
*/
type GetChildrenNotFound struct {
}

// IsSuccess returns true when this get children not found response has a 2xx status code
func (o *GetChildrenNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get children not found response has a 3xx status code
func (o *GetChildrenNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get children not found response has a 4xx status code
func (o *GetChildrenNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get children not found response has a 5xx status code
func (o *GetChildrenNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get children not found response a status code equal to that given
func (o *GetChildrenNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get children not found response
func (o *GetChildrenNotFound) Code() int {
	return 404
}

func (o *GetChildrenNotFound) Error() string {
	return fmt.Sprintf("[GET /zk/getChildren][%d] getChildrenNotFound ", 404)
}

func (o *GetChildrenNotFound) String() string {
	return fmt.Sprintf("[GET /zk/getChildren][%d] getChildrenNotFound ", 404)
}

func (o *GetChildrenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetChildrenInternalServerError creates a GetChildrenInternalServerError with default headers values
func NewGetChildrenInternalServerError() *GetChildrenInternalServerError {
	return &GetChildrenInternalServerError{}
}

/*
GetChildrenInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetChildrenInternalServerError struct {
}

// IsSuccess returns true when this get children internal server error response has a 2xx status code
func (o *GetChildrenInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get children internal server error response has a 3xx status code
func (o *GetChildrenInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get children internal server error response has a 4xx status code
func (o *GetChildrenInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get children internal server error response has a 5xx status code
func (o *GetChildrenInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get children internal server error response a status code equal to that given
func (o *GetChildrenInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get children internal server error response
func (o *GetChildrenInternalServerError) Code() int {
	return 500
}

func (o *GetChildrenInternalServerError) Error() string {
	return fmt.Sprintf("[GET /zk/getChildren][%d] getChildrenInternalServerError ", 500)
}

func (o *GetChildrenInternalServerError) String() string {
	return fmt.Sprintf("[GET /zk/getChildren][%d] getChildrenInternalServerError ", 500)
}

func (o *GetChildrenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
