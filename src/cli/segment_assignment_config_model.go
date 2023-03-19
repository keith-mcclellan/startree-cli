// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/spf13/cobra"
	"startree.ai/cli/models"
)

// Schema cli for SegmentAssignmentConfig

// register flags to command
func registerModelSegmentAssignmentConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSegmentAssignmentConfigAssignmentStrategy(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSegmentAssignmentConfigAssignmentStrategy(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	assignmentStrategyDescription := ``

	var assignmentStrategyFlagName string
	if cmdPrefix == "" {
		assignmentStrategyFlagName = "assignmentStrategy"
	} else {
		assignmentStrategyFlagName = fmt.Sprintf("%v.assignmentStrategy", cmdPrefix)
	}

	var assignmentStrategyFlagDefault string

	_ = cmd.PersistentFlags().String(assignmentStrategyFlagName, assignmentStrategyFlagDefault, assignmentStrategyDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSegmentAssignmentConfigFlags(depth int, m *models.SegmentAssignmentConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, assignmentStrategyAdded := retrieveSegmentAssignmentConfigAssignmentStrategyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || assignmentStrategyAdded

	return nil, retAdded
}

func retrieveSegmentAssignmentConfigAssignmentStrategyFlags(depth int, m *models.SegmentAssignmentConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	assignmentStrategyFlagName := fmt.Sprintf("%v.assignmentStrategy", cmdPrefix)
	if cmd.Flags().Changed(assignmentStrategyFlagName) {

		var assignmentStrategyFlagName string
		if cmdPrefix == "" {
			assignmentStrategyFlagName = "assignmentStrategy"
		} else {
			assignmentStrategyFlagName = fmt.Sprintf("%v.assignmentStrategy", cmdPrefix)
		}

		assignmentStrategyFlagValue, err := cmd.Flags().GetString(assignmentStrategyFlagName)
		if err != nil {
			return err, false
		}
		m.AssignmentStrategy = assignmentStrategyFlagValue

		retAdded = true
	}

	return nil, retAdded
}
