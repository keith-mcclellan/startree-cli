// Code generated by go-swagger; DO NOT EDIT.

package segment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"startree.ai/cli/models"
)

// DeleteSegmentReader is a Reader for the DeleteSegment structure.
type DeleteSegmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSegmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSegmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSegmentOK creates a DeleteSegmentOK with default headers values
func NewDeleteSegmentOK() *DeleteSegmentOK {
	return &DeleteSegmentOK{}
}

/*
DeleteSegmentOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteSegmentOK struct {
	Payload *models.SuccessResponse
}

// IsSuccess returns true when this delete segment o k response has a 2xx status code
func (o *DeleteSegmentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete segment o k response has a 3xx status code
func (o *DeleteSegmentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete segment o k response has a 4xx status code
func (o *DeleteSegmentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete segment o k response has a 5xx status code
func (o *DeleteSegmentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete segment o k response a status code equal to that given
func (o *DeleteSegmentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete segment o k response
func (o *DeleteSegmentOK) Code() int {
	return 200
}

func (o *DeleteSegmentOK) Error() string {
	return fmt.Sprintf("[DELETE /segments/{tableName}/{segmentName}][%d] deleteSegmentOK  %+v", 200, o.Payload)
}

func (o *DeleteSegmentOK) String() string {
	return fmt.Sprintf("[DELETE /segments/{tableName}/{segmentName}][%d] deleteSegmentOK  %+v", 200, o.Payload)
}

func (o *DeleteSegmentOK) GetPayload() *models.SuccessResponse {
	return o.Payload
}

func (o *DeleteSegmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
