// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DedupConfig dedup config
//
// swagger:model DedupConfig
type DedupConfig struct {

	// dedup enabled
	// Required: true
	// Read Only: true
	DedupEnabled bool `json:"dedupEnabled"`

	// hash function
	// Read Only: true
	// Enum: [NONE MD5 MURMUR3]
	HashFunction string `json:"hashFunction,omitempty"`
}

// Validate validates this dedup config
func (m *DedupConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDedupEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHashFunction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DedupConfig) validateDedupEnabled(formats strfmt.Registry) error {

	if err := validate.Required("dedupEnabled", "body", bool(m.DedupEnabled)); err != nil {
		return err
	}

	return nil
}

var dedupConfigTypeHashFunctionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","MD5","MURMUR3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dedupConfigTypeHashFunctionPropEnum = append(dedupConfigTypeHashFunctionPropEnum, v)
	}
}

const (

	// DedupConfigHashFunctionNONE captures enum value "NONE"
	DedupConfigHashFunctionNONE string = "NONE"

	// DedupConfigHashFunctionMD5 captures enum value "MD5"
	DedupConfigHashFunctionMD5 string = "MD5"

	// DedupConfigHashFunctionMURMUR3 captures enum value "MURMUR3"
	DedupConfigHashFunctionMURMUR3 string = "MURMUR3"
)

// prop value enum
func (m *DedupConfig) validateHashFunctionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dedupConfigTypeHashFunctionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DedupConfig) validateHashFunction(formats strfmt.Registry) error {
	if swag.IsZero(m.HashFunction) { // not required
		return nil
	}

	// value enum
	if err := m.validateHashFunctionEnum("hashFunction", "body", m.HashFunction); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this dedup config based on the context it is used
func (m *DedupConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDedupEnabled(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHashFunction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DedupConfig) contextValidateDedupEnabled(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dedupEnabled", "body", bool(m.DedupEnabled)); err != nil {
		return err
	}

	return nil
}

func (m *DedupConfig) contextValidateHashFunction(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "hashFunction", "body", string(m.HashFunction)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DedupConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DedupConfig) UnmarshalBinary(b []byte) error {
	var res DedupConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}