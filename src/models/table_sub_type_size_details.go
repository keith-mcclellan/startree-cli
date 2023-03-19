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

// TableSubTypeSizeDetails table sub type size details
//
// swagger:model TableSubTypeSizeDetails
type TableSubTypeSizeDetails struct {

	// estimated size in bytes
	EstimatedSizeInBytes int64 `json:"estimatedSizeInBytes,omitempty"`

	// missing segments
	MissingSegments int32 `json:"missingSegments,omitempty"`

	// reported size in bytes
	ReportedSizeInBytes int64 `json:"reportedSizeInBytes,omitempty"`

	// segments
	Segments map[string]SegmentSizeDetails `json:"segments,omitempty"`
}

// Validate validates this table sub type size details
func (m *TableSubTypeSizeDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSegments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TableSubTypeSizeDetails) validateSegments(formats strfmt.Registry) error {
	if swag.IsZero(m.Segments) { // not required
		return nil
	}

	for k := range m.Segments {

		if err := validate.Required("segments"+"."+k, "body", m.Segments[k]); err != nil {
			return err
		}
		if val, ok := m.Segments[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("segments" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("segments" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this table sub type size details based on the context it is used
func (m *TableSubTypeSizeDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSegments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TableSubTypeSizeDetails) contextValidateSegments(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Segments {

		if val, ok := m.Segments[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TableSubTypeSizeDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TableSubTypeSizeDetails) UnmarshalBinary(b []byte) error {
	var res TableSubTypeSizeDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
