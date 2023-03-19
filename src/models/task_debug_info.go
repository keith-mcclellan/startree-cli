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

// TaskDebugInfo task debug info
//
// swagger:model TaskDebugInfo
type TaskDebugInfo struct {

	// execution start time
	ExecutionStartTime string `json:"executionStartTime,omitempty"`

	// finish time
	FinishTime string `json:"finishTime,omitempty"`

	// start time
	StartTime string `json:"startTime,omitempty"`

	// subtask count
	SubtaskCount *TaskCount `json:"subtaskCount,omitempty"`

	// subtask infos
	SubtaskInfos []*SubtaskDebugInfo `json:"subtaskInfos"`

	// task state
	// Enum: [NOT_STARTED IN_PROGRESS STOPPED STOPPING FAILED COMPLETED ABORTED TIMED_OUT TIMING_OUT FAILING]
	TaskState string `json:"taskState,omitempty"`
}

// Validate validates this task debug info
func (m *TaskDebugInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubtaskCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubtaskInfos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskDebugInfo) validateSubtaskCount(formats strfmt.Registry) error {
	if swag.IsZero(m.SubtaskCount) { // not required
		return nil
	}

	if m.SubtaskCount != nil {
		if err := m.SubtaskCount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subtaskCount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subtaskCount")
			}
			return err
		}
	}

	return nil
}

func (m *TaskDebugInfo) validateSubtaskInfos(formats strfmt.Registry) error {
	if swag.IsZero(m.SubtaskInfos) { // not required
		return nil
	}

	for i := 0; i < len(m.SubtaskInfos); i++ {
		if swag.IsZero(m.SubtaskInfos[i]) { // not required
			continue
		}

		if m.SubtaskInfos[i] != nil {
			if err := m.SubtaskInfos[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subtaskInfos" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subtaskInfos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var taskDebugInfoTypeTaskStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NOT_STARTED","IN_PROGRESS","STOPPED","STOPPING","FAILED","COMPLETED","ABORTED","TIMED_OUT","TIMING_OUT","FAILING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		taskDebugInfoTypeTaskStatePropEnum = append(taskDebugInfoTypeTaskStatePropEnum, v)
	}
}

const (

	// TaskDebugInfoTaskStateNOTSTARTED captures enum value "NOT_STARTED"
	TaskDebugInfoTaskStateNOTSTARTED string = "NOT_STARTED"

	// TaskDebugInfoTaskStateINPROGRESS captures enum value "IN_PROGRESS"
	TaskDebugInfoTaskStateINPROGRESS string = "IN_PROGRESS"

	// TaskDebugInfoTaskStateSTOPPED captures enum value "STOPPED"
	TaskDebugInfoTaskStateSTOPPED string = "STOPPED"

	// TaskDebugInfoTaskStateSTOPPING captures enum value "STOPPING"
	TaskDebugInfoTaskStateSTOPPING string = "STOPPING"

	// TaskDebugInfoTaskStateFAILED captures enum value "FAILED"
	TaskDebugInfoTaskStateFAILED string = "FAILED"

	// TaskDebugInfoTaskStateCOMPLETED captures enum value "COMPLETED"
	TaskDebugInfoTaskStateCOMPLETED string = "COMPLETED"

	// TaskDebugInfoTaskStateABORTED captures enum value "ABORTED"
	TaskDebugInfoTaskStateABORTED string = "ABORTED"

	// TaskDebugInfoTaskStateTIMEDOUT captures enum value "TIMED_OUT"
	TaskDebugInfoTaskStateTIMEDOUT string = "TIMED_OUT"

	// TaskDebugInfoTaskStateTIMINGOUT captures enum value "TIMING_OUT"
	TaskDebugInfoTaskStateTIMINGOUT string = "TIMING_OUT"

	// TaskDebugInfoTaskStateFAILING captures enum value "FAILING"
	TaskDebugInfoTaskStateFAILING string = "FAILING"
)

// prop value enum
func (m *TaskDebugInfo) validateTaskStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, taskDebugInfoTypeTaskStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TaskDebugInfo) validateTaskState(formats strfmt.Registry) error {
	if swag.IsZero(m.TaskState) { // not required
		return nil
	}

	// value enum
	if err := m.validateTaskStateEnum("taskState", "body", m.TaskState); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this task debug info based on the context it is used
func (m *TaskDebugInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSubtaskCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubtaskInfos(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskDebugInfo) contextValidateSubtaskCount(ctx context.Context, formats strfmt.Registry) error {

	if m.SubtaskCount != nil {
		if err := m.SubtaskCount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subtaskCount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subtaskCount")
			}
			return err
		}
	}

	return nil
}

func (m *TaskDebugInfo) contextValidateSubtaskInfos(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SubtaskInfos); i++ {

		if m.SubtaskInfos[i] != nil {
			if err := m.SubtaskInfos[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subtaskInfos" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subtaskInfos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskDebugInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskDebugInfo) UnmarshalBinary(b []byte) error {
	var res TaskDebugInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}