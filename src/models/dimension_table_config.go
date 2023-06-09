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

// DimensionTableConfig dimension table config
//
// swagger:model DimensionTableConfig
type DimensionTableConfig struct {

	// disable preload
	// Required: true
	// Read Only: true
	DisablePreload bool `json:"disablePreload"`
}

// Validate validates this dimension table config
func (m *DimensionTableConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisablePreload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DimensionTableConfig) validateDisablePreload(formats strfmt.Registry) error {

	if err := validate.Required("disablePreload", "body", bool(m.DisablePreload)); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this dimension table config based on the context it is used
func (m *DimensionTableConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisablePreload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DimensionTableConfig) contextValidateDisablePreload(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "disablePreload", "body", bool(m.DisablePreload)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DimensionTableConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DimensionTableConfig) UnmarshalBinary(b []byte) error {
	var res DimensionTableConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
