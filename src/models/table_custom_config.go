// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TableCustomConfig table custom config
//
// swagger:model TableCustomConfig
type TableCustomConfig struct {

	// custom configs
	// Read Only: true
	CustomConfigs map[string]string `json:"customConfigs,omitempty"`
}

// Validate validates this table custom config
func (m *TableCustomConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this table custom config based on the context it is used
func (m *TableCustomConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TableCustomConfig) contextValidateCustomConfigs(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *TableCustomConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TableCustomConfig) UnmarshalBinary(b []byte) error {
	var res TableCustomConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
