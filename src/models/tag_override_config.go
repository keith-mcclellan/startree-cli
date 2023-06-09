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

// TagOverrideConfig tag override config
//
// swagger:model TagOverrideConfig
type TagOverrideConfig struct {

	// realtime completed
	// Read Only: true
	RealtimeCompleted string `json:"realtimeCompleted,omitempty"`

	// realtime consuming
	// Read Only: true
	RealtimeConsuming string `json:"realtimeConsuming,omitempty"`
}

// Validate validates this tag override config
func (m *TagOverrideConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this tag override config based on the context it is used
func (m *TagOverrideConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRealtimeCompleted(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRealtimeConsuming(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TagOverrideConfig) contextValidateRealtimeCompleted(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "realtimeCompleted", "body", string(m.RealtimeCompleted)); err != nil {
		return err
	}

	return nil
}

func (m *TagOverrideConfig) contextValidateRealtimeConsuming(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "realtimeConsuming", "body", string(m.RealtimeConsuming)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TagOverrideConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TagOverrideConfig) UnmarshalBinary(b []byte) error {
	var res TagOverrideConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
