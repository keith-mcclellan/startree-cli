// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterHealthResponse cluster health response
//
// swagger:model ClusterHealthResponse
type ClusterHealthResponse struct {

	// table to error segments count
	TableToErrorSegmentsCount map[string]int32 `json:"tableToErrorSegmentsCount,omitempty"`

	// table to misconfigured segment count
	TableToMisconfiguredSegmentCount map[string]int32 `json:"tableToMisconfiguredSegmentCount,omitempty"`

	// table to segments wit h missing columns count
	TableToSegmentsWitHMissingColumnsCount map[string]int32 `json:"tableToSegmentsWitHMissingColumnsCount,omitempty"`

	// unhealthy server count
	UnhealthyServerCount int64 `json:"unhealthyServerCount,omitempty"`
}

// Validate validates this cluster health response
func (m *ClusterHealthResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster health response based on context it is used
func (m *ClusterHealthResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterHealthResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterHealthResponse) UnmarshalBinary(b []byte) error {
	var res ClusterHealthResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}