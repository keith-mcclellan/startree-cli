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

// ReloadSegmentDeprecated1Reader is a Reader for the ReloadSegmentDeprecated1 structure.
type ReloadSegmentDeprecated1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReloadSegmentDeprecated1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReloadSegmentDeprecated1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReloadSegmentDeprecated1OK creates a ReloadSegmentDeprecated1OK with default headers values
func NewReloadSegmentDeprecated1OK() *ReloadSegmentDeprecated1OK {
	return &ReloadSegmentDeprecated1OK{}
}

/*
ReloadSegmentDeprecated1OK describes a response with status code 200, with default header values.

successful operation
*/
type ReloadSegmentDeprecated1OK struct {
	Payload *models.SuccessResponse
}

// IsSuccess returns true when this reload segment deprecated1 o k response has a 2xx status code
func (o *ReloadSegmentDeprecated1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this reload segment deprecated1 o k response has a 3xx status code
func (o *ReloadSegmentDeprecated1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reload segment deprecated1 o k response has a 4xx status code
func (o *ReloadSegmentDeprecated1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this reload segment deprecated1 o k response has a 5xx status code
func (o *ReloadSegmentDeprecated1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this reload segment deprecated1 o k response a status code equal to that given
func (o *ReloadSegmentDeprecated1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the reload segment deprecated1 o k response
func (o *ReloadSegmentDeprecated1OK) Code() int {
	return 200
}

func (o *ReloadSegmentDeprecated1OK) Error() string {
	return fmt.Sprintf("[POST /tables/{tableName}/segments/{segmentName}/reload][%d] reloadSegmentDeprecated1OK  %+v", 200, o.Payload)
}

func (o *ReloadSegmentDeprecated1OK) String() string {
	return fmt.Sprintf("[POST /tables/{tableName}/segments/{segmentName}/reload][%d] reloadSegmentDeprecated1OK  %+v", 200, o.Payload)
}

func (o *ReloadSegmentDeprecated1OK) GetPayload() *models.SuccessResponse {
	return o.Payload
}

func (o *ReloadSegmentDeprecated1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
