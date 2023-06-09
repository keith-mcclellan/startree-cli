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

// ComplexTypeConfig complex type config
//
// swagger:model ComplexTypeConfig
type ComplexTypeConfig struct {

	// collection not unnested to Json
	// Read Only: true
	// Enum: [NONE NON_PRIMITIVE ALL]
	CollectionNotUnnestedToJSON string `json:"collectionNotUnnestedToJson,omitempty"`

	// delimiter
	// Read Only: true
	Delimiter string `json:"delimiter,omitempty"`

	// fields to unnest
	// Read Only: true
	FieldsToUnnest []string `json:"fieldsToUnnest"`

	// prefixes to rename
	// Read Only: true
	PrefixesToRename map[string]string `json:"prefixesToRename,omitempty"`
}

// Validate validates this complex type config
func (m *ComplexTypeConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollectionNotUnnestedToJSON(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var complexTypeConfigTypeCollectionNotUnnestedToJSONPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","NON_PRIMITIVE","ALL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		complexTypeConfigTypeCollectionNotUnnestedToJSONPropEnum = append(complexTypeConfigTypeCollectionNotUnnestedToJSONPropEnum, v)
	}
}

const (

	// ComplexTypeConfigCollectionNotUnnestedToJSONNONE captures enum value "NONE"
	ComplexTypeConfigCollectionNotUnnestedToJSONNONE string = "NONE"

	// ComplexTypeConfigCollectionNotUnnestedToJSONNONPRIMITIVE captures enum value "NON_PRIMITIVE"
	ComplexTypeConfigCollectionNotUnnestedToJSONNONPRIMITIVE string = "NON_PRIMITIVE"

	// ComplexTypeConfigCollectionNotUnnestedToJSONALL captures enum value "ALL"
	ComplexTypeConfigCollectionNotUnnestedToJSONALL string = "ALL"
)

// prop value enum
func (m *ComplexTypeConfig) validateCollectionNotUnnestedToJSONEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, complexTypeConfigTypeCollectionNotUnnestedToJSONPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ComplexTypeConfig) validateCollectionNotUnnestedToJSON(formats strfmt.Registry) error {
	if swag.IsZero(m.CollectionNotUnnestedToJSON) { // not required
		return nil
	}

	// value enum
	if err := m.validateCollectionNotUnnestedToJSONEnum("collectionNotUnnestedToJson", "body", m.CollectionNotUnnestedToJSON); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this complex type config based on the context it is used
func (m *ComplexTypeConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCollectionNotUnnestedToJSON(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDelimiter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFieldsToUnnest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrefixesToRename(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComplexTypeConfig) contextValidateCollectionNotUnnestedToJSON(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "collectionNotUnnestedToJson", "body", string(m.CollectionNotUnnestedToJSON)); err != nil {
		return err
	}

	return nil
}

func (m *ComplexTypeConfig) contextValidateDelimiter(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "delimiter", "body", string(m.Delimiter)); err != nil {
		return err
	}

	return nil
}

func (m *ComplexTypeConfig) contextValidateFieldsToUnnest(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "fieldsToUnnest", "body", []string(m.FieldsToUnnest)); err != nil {
		return err
	}

	return nil
}

func (m *ComplexTypeConfig) contextValidatePrefixesToRename(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *ComplexTypeConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComplexTypeConfig) UnmarshalBinary(b []byte) error {
	var res ComplexTypeConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
