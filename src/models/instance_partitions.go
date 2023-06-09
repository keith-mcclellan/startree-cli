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

// InstancePartitions instance partitions
//
// swagger:model InstancePartitions
type InstancePartitions struct {

	// instance partitions name
	// Read Only: true
	InstancePartitionsName string `json:"instancePartitionsName,omitempty"`

	// partition to instances map
	// Read Only: true
	PartitionToInstancesMap map[string][]string `json:"partitionToInstancesMap,omitempty"`
}

// Validate validates this instance partitions
func (m *InstancePartitions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this instance partitions based on the context it is used
func (m *InstancePartitions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInstancePartitionsName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePartitionToInstancesMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstancePartitions) contextValidateInstancePartitionsName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "instancePartitionsName", "body", string(m.InstancePartitionsName)); err != nil {
		return err
	}

	return nil
}

func (m *InstancePartitions) contextValidatePartitionToInstancesMap(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *InstancePartitions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstancePartitions) UnmarshalBinary(b []byte) error {
	var res InstancePartitions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
