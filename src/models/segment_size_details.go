// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SegmentSizeDetails segment size details
//
// swagger:model SegmentSizeDetails
type SegmentSizeDetails struct {

	// estimated size in bytes
	EstimatedSizeInBytes int64 `json:"estimatedSizeInBytes,omitempty"`

	// reported size in bytes
	ReportedSizeInBytes int64 `json:"reportedSizeInBytes,omitempty"`

	// server info
	ServerInfo map[string]SegmentSizeInfo `json:"serverInfo,omitempty"`
}

// Validate validates this segment size details
func (m *SegmentSizeDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServerInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SegmentSizeDetails) validateServerInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerInfo) { // not required
		return nil
	}

	for k := range m.ServerInfo {

		if err := validate.Required("serverInfo"+"."+k, "body", m.ServerInfo[k]); err != nil {
			return err
		}
		if val, ok := m.ServerInfo[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serverInfo" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serverInfo" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this segment size details based on the context it is used
func (m *SegmentSizeDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServerInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SegmentSizeDetails) contextValidateServerInfo(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.ServerInfo {

		if val, ok := m.ServerInfo[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SegmentSizeDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SegmentSizeDetails) UnmarshalBinary(b []byte) error {
	var res SegmentSizeDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
