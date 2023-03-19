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

// TableTierDetails table tier details
//
// swagger:model TableTierDetails
type TableTierDetails struct {

	// Storage tiers of segments for the given table
	// Read Only: true
	SegmentTiers map[string]map[string]string `json:"segmentTiers,omitempty"`

	// Name of table to look for segment storage tiers
	// Read Only: true
	TableName string `json:"tableName,omitempty"`
}

// Validate validates this table tier details
func (m *TableTierDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this table tier details based on the context it is used
func (m *TableTierDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSegmentTiers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTableName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TableTierDetails) contextValidateSegmentTiers(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *TableTierDetails) contextValidateTableName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "tableName", "body", string(m.TableName)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TableTierDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TableTierDetails) UnmarshalBinary(b []byte) error {
	var res TableTierDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}