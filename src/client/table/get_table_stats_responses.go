// Code generated by go-swagger; DO NOT EDIT.

package table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetTableStatsReader is a Reader for the GetTableStats structure.
type GetTableStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTableStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTableStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTableStatsOK creates a GetTableStatsOK with default headers values
func NewGetTableStatsOK() *GetTableStatsOK {
	return &GetTableStatsOK{}
}

/*
GetTableStatsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTableStatsOK struct {
	Payload string
}

// IsSuccess returns true when this get table stats o k response has a 2xx status code
func (o *GetTableStatsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get table stats o k response has a 3xx status code
func (o *GetTableStatsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get table stats o k response has a 4xx status code
func (o *GetTableStatsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get table stats o k response has a 5xx status code
func (o *GetTableStatsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get table stats o k response a status code equal to that given
func (o *GetTableStatsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get table stats o k response
func (o *GetTableStatsOK) Code() int {
	return 200
}

func (o *GetTableStatsOK) Error() string {
	return fmt.Sprintf("[GET /tables/{tableName}/stats][%d] getTableStatsOK  %+v", 200, o.Payload)
}

func (o *GetTableStatsOK) String() string {
	return fmt.Sprintf("[GET /tables/{tableName}/stats][%d] getTableStatsOK  %+v", 200, o.Payload)
}

func (o *GetTableStatsOK) GetPayload() string {
	return o.Payload
}

func (o *GetTableStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
