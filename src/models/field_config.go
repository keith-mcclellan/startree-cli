// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FieldConfig field config
//
// swagger:model FieldConfig
type FieldConfig struct {

	// compression codec
	// Read Only: true
	// Enum: [PASS_THROUGH SNAPPY ZSTANDARD LZ4]
	CompressionCodec string `json:"compressionCodec,omitempty"`

	// encoding type
	// Read Only: true
	// Enum: [RAW DICTIONARY]
	EncodingType string `json:"encodingType,omitempty"`

	// index type
	// Read Only: true
	// Enum: [INVERTED SORTED TEXT FST H3 JSON TIMESTAMP RANGE]
	IndexType string `json:"indexType,omitempty"`

	// index types
	// Read Only: true
	IndexTypes []string `json:"indexTypes"`

	// name
	// Required: true
	// Read Only: true
	Name string `json:"name"`

	// properties
	// Read Only: true
	Properties map[string]string `json:"properties,omitempty"`

	// timestamp config
	// Read Only: true
	TimestampConfig *TimestampConfig `json:"timestampConfig,omitempty"`
}

// Validate validates this field config
func (m *FieldConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompressionCodec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncodingType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndexType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndexTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestampConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var fieldConfigTypeCompressionCodecPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PASS_THROUGH","SNAPPY","ZSTANDARD","LZ4"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fieldConfigTypeCompressionCodecPropEnum = append(fieldConfigTypeCompressionCodecPropEnum, v)
	}
}

const (

	// FieldConfigCompressionCodecPASSTHROUGH captures enum value "PASS_THROUGH"
	FieldConfigCompressionCodecPASSTHROUGH string = "PASS_THROUGH"

	// FieldConfigCompressionCodecSNAPPY captures enum value "SNAPPY"
	FieldConfigCompressionCodecSNAPPY string = "SNAPPY"

	// FieldConfigCompressionCodecZSTANDARD captures enum value "ZSTANDARD"
	FieldConfigCompressionCodecZSTANDARD string = "ZSTANDARD"

	// FieldConfigCompressionCodecLZ4 captures enum value "LZ4"
	FieldConfigCompressionCodecLZ4 string = "LZ4"
)

// prop value enum
func (m *FieldConfig) validateCompressionCodecEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fieldConfigTypeCompressionCodecPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FieldConfig) validateCompressionCodec(formats strfmt.Registry) error {
	if swag.IsZero(m.CompressionCodec) { // not required
		return nil
	}

	// value enum
	if err := m.validateCompressionCodecEnum("compressionCodec", "body", m.CompressionCodec); err != nil {
		return err
	}

	return nil
}

var fieldConfigTypeEncodingTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RAW","DICTIONARY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fieldConfigTypeEncodingTypePropEnum = append(fieldConfigTypeEncodingTypePropEnum, v)
	}
}

const (

	// FieldConfigEncodingTypeRAW captures enum value "RAW"
	FieldConfigEncodingTypeRAW string = "RAW"

	// FieldConfigEncodingTypeDICTIONARY captures enum value "DICTIONARY"
	FieldConfigEncodingTypeDICTIONARY string = "DICTIONARY"
)

// prop value enum
func (m *FieldConfig) validateEncodingTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fieldConfigTypeEncodingTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FieldConfig) validateEncodingType(formats strfmt.Registry) error {
	if swag.IsZero(m.EncodingType) { // not required
		return nil
	}

	// value enum
	if err := m.validateEncodingTypeEnum("encodingType", "body", m.EncodingType); err != nil {
		return err
	}

	return nil
}

var fieldConfigTypeIndexTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INVERTED","SORTED","TEXT","FST","H3","JSON","TIMESTAMP","RANGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fieldConfigTypeIndexTypePropEnum = append(fieldConfigTypeIndexTypePropEnum, v)
	}
}

const (

	// FieldConfigIndexTypeINVERTED captures enum value "INVERTED"
	FieldConfigIndexTypeINVERTED string = "INVERTED"

	// FieldConfigIndexTypeSORTED captures enum value "SORTED"
	FieldConfigIndexTypeSORTED string = "SORTED"

	// FieldConfigIndexTypeTEXT captures enum value "TEXT"
	FieldConfigIndexTypeTEXT string = "TEXT"

	// FieldConfigIndexTypeFST captures enum value "FST"
	FieldConfigIndexTypeFST string = "FST"

	// FieldConfigIndexTypeH3 captures enum value "H3"
	FieldConfigIndexTypeH3 string = "H3"

	// FieldConfigIndexTypeJSON captures enum value "JSON"
	FieldConfigIndexTypeJSON string = "JSON"

	// FieldConfigIndexTypeTIMESTAMP captures enum value "TIMESTAMP"
	FieldConfigIndexTypeTIMESTAMP string = "TIMESTAMP"

	// FieldConfigIndexTypeRANGE captures enum value "RANGE"
	FieldConfigIndexTypeRANGE string = "RANGE"
)

// prop value enum
func (m *FieldConfig) validateIndexTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fieldConfigTypeIndexTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FieldConfig) validateIndexType(formats strfmt.Registry) error {
	if swag.IsZero(m.IndexType) { // not required
		return nil
	}

	// value enum
	if err := m.validateIndexTypeEnum("indexType", "body", m.IndexType); err != nil {
		return err
	}

	return nil
}

var fieldConfigIndexTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INVERTED","SORTED","TEXT","FST","H3","JSON","TIMESTAMP","RANGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fieldConfigIndexTypesItemsEnum = append(fieldConfigIndexTypesItemsEnum, v)
	}
}

func (m *FieldConfig) validateIndexTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fieldConfigIndexTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FieldConfig) validateIndexTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.IndexTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.IndexTypes); i++ {

		// value enum
		if err := m.validateIndexTypesItemsEnum("indexTypes"+"."+strconv.Itoa(i), "body", m.IndexTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *FieldConfig) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *FieldConfig) validateTimestampConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.TimestampConfig) { // not required
		return nil
	}

	if m.TimestampConfig != nil {
		if err := m.TimestampConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timestampConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timestampConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this field config based on the context it is used
func (m *FieldConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCompressionCodec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEncodingType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIndexType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIndexTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimestampConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FieldConfig) contextValidateCompressionCodec(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "compressionCodec", "body", string(m.CompressionCodec)); err != nil {
		return err
	}

	return nil
}

func (m *FieldConfig) contextValidateEncodingType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "encodingType", "body", string(m.EncodingType)); err != nil {
		return err
	}

	return nil
}

func (m *FieldConfig) contextValidateIndexType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "indexType", "body", string(m.IndexType)); err != nil {
		return err
	}

	return nil
}

func (m *FieldConfig) contextValidateIndexTypes(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "indexTypes", "body", []string(m.IndexTypes)); err != nil {
		return err
	}

	return nil
}

func (m *FieldConfig) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *FieldConfig) contextValidateProperties(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *FieldConfig) contextValidateTimestampConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.TimestampConfig != nil {
		if err := m.TimestampConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timestampConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timestampConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FieldConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FieldConfig) UnmarshalBinary(b []byte) error {
	var res FieldConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
