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

// InstanceConstraintConfig instance constraint config
//
// swagger:model InstanceConstraintConfig
type InstanceConstraintConfig struct {

	// constraints
	// Required: true
	// Read Only: true
	Constraints []string `json:"constraints"`
}

// Validate validates this instance constraint config
func (m *InstanceConstraintConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstraints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceConstraintConfig) validateConstraints(formats strfmt.Registry) error {

	if err := validate.Required("constraints", "body", m.Constraints); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this instance constraint config based on the context it is used
func (m *InstanceConstraintConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceConstraintConfig) contextValidateConstraints(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "constraints", "body", []string(m.Constraints)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceConstraintConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceConstraintConfig) UnmarshalBinary(b []byte) error {
	var res InstanceConstraintConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}