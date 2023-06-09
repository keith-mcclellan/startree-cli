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

// AggregationConfig aggregation config
//
// swagger:model AggregationConfig
type AggregationConfig struct {

	// aggregation function
	// Read Only: true
	AggregationFunction string `json:"aggregationFunction,omitempty"`

	// column name
	// Read Only: true
	ColumnName string `json:"columnName,omitempty"`
}

// Validate validates this aggregation config
func (m *AggregationConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this aggregation config based on the context it is used
func (m *AggregationConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggregationFunction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateColumnName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AggregationConfig) contextValidateAggregationFunction(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "aggregationFunction", "body", string(m.AggregationFunction)); err != nil {
		return err
	}

	return nil
}

func (m *AggregationConfig) contextValidateColumnName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "columnName", "body", string(m.ColumnName)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AggregationConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AggregationConfig) UnmarshalBinary(b []byte) error {
	var res AggregationConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
