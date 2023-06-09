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

// InstanceAssignmentConfig instance assignment config
//
// swagger:model InstanceAssignmentConfig
type InstanceAssignmentConfig struct {

	// constraint config
	// Read Only: true
	ConstraintConfig *InstanceConstraintConfig `json:"constraintConfig,omitempty"`

	// partition selector
	// Read Only: true
	// Enum: [FD_AWARE_INSTANCE_PARTITION_SELECTOR INSTANCE_REPLICA_GROUP_PARTITION_SELECTOR]
	PartitionSelector string `json:"partitionSelector,omitempty"`

	// replica group partition config
	// Required: true
	// Read Only: true
	ReplicaGroupPartitionConfig *InstanceReplicaGroupPartitionConfig `json:"replicaGroupPartitionConfig"`

	// tag pool config
	// Required: true
	// Read Only: true
	TagPoolConfig *InstanceTagPoolConfig `json:"tagPoolConfig"`
}

// Validate validates this instance assignment config
func (m *InstanceAssignmentConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstraintConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartitionSelector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicaGroupPartitionConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagPoolConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceAssignmentConfig) validateConstraintConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.ConstraintConfig) { // not required
		return nil
	}

	if m.ConstraintConfig != nil {
		if err := m.ConstraintConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("constraintConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("constraintConfig")
			}
			return err
		}
	}

	return nil
}

var instanceAssignmentConfigTypePartitionSelectorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FD_AWARE_INSTANCE_PARTITION_SELECTOR","INSTANCE_REPLICA_GROUP_PARTITION_SELECTOR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceAssignmentConfigTypePartitionSelectorPropEnum = append(instanceAssignmentConfigTypePartitionSelectorPropEnum, v)
	}
}

const (

	// InstanceAssignmentConfigPartitionSelectorFDAWAREINSTANCEPARTITIONSELECTOR captures enum value "FD_AWARE_INSTANCE_PARTITION_SELECTOR"
	InstanceAssignmentConfigPartitionSelectorFDAWAREINSTANCEPARTITIONSELECTOR string = "FD_AWARE_INSTANCE_PARTITION_SELECTOR"

	// InstanceAssignmentConfigPartitionSelectorINSTANCEREPLICAGROUPPARTITIONSELECTOR captures enum value "INSTANCE_REPLICA_GROUP_PARTITION_SELECTOR"
	InstanceAssignmentConfigPartitionSelectorINSTANCEREPLICAGROUPPARTITIONSELECTOR string = "INSTANCE_REPLICA_GROUP_PARTITION_SELECTOR"
)

// prop value enum
func (m *InstanceAssignmentConfig) validatePartitionSelectorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, instanceAssignmentConfigTypePartitionSelectorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InstanceAssignmentConfig) validatePartitionSelector(formats strfmt.Registry) error {
	if swag.IsZero(m.PartitionSelector) { // not required
		return nil
	}

	// value enum
	if err := m.validatePartitionSelectorEnum("partitionSelector", "body", m.PartitionSelector); err != nil {
		return err
	}

	return nil
}

func (m *InstanceAssignmentConfig) validateReplicaGroupPartitionConfig(formats strfmt.Registry) error {

	if err := validate.Required("replicaGroupPartitionConfig", "body", m.ReplicaGroupPartitionConfig); err != nil {
		return err
	}

	if m.ReplicaGroupPartitionConfig != nil {
		if err := m.ReplicaGroupPartitionConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("replicaGroupPartitionConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("replicaGroupPartitionConfig")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceAssignmentConfig) validateTagPoolConfig(formats strfmt.Registry) error {

	if err := validate.Required("tagPoolConfig", "body", m.TagPoolConfig); err != nil {
		return err
	}

	if m.TagPoolConfig != nil {
		if err := m.TagPoolConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tagPoolConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tagPoolConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this instance assignment config based on the context it is used
func (m *InstanceAssignmentConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConstraintConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePartitionSelector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReplicaGroupPartitionConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTagPoolConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceAssignmentConfig) contextValidateConstraintConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.ConstraintConfig != nil {
		if err := m.ConstraintConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("constraintConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("constraintConfig")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceAssignmentConfig) contextValidatePartitionSelector(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "partitionSelector", "body", string(m.PartitionSelector)); err != nil {
		return err
	}

	return nil
}

func (m *InstanceAssignmentConfig) contextValidateReplicaGroupPartitionConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.ReplicaGroupPartitionConfig != nil {
		if err := m.ReplicaGroupPartitionConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("replicaGroupPartitionConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("replicaGroupPartitionConfig")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceAssignmentConfig) contextValidateTagPoolConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.TagPoolConfig != nil {
		if err := m.TagPoolConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tagPoolConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tagPoolConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceAssignmentConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceAssignmentConfig) UnmarshalBinary(b []byte) error {
	var res InstanceAssignmentConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
