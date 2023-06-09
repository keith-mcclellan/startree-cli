// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/spf13/cobra"
	"startree.ai/cli/models"
)

// Schema cli for AuthWorkflowInfo

// register flags to command
func registerModelAuthWorkflowInfoFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerAuthWorkflowInfoWorkflow(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerAuthWorkflowInfoWorkflow(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	workflowDescription := ``

	var workflowFlagName string
	if cmdPrefix == "" {
		workflowFlagName = "workflow"
	} else {
		workflowFlagName = fmt.Sprintf("%v.workflow", cmdPrefix)
	}

	var workflowFlagDefault string

	_ = cmd.PersistentFlags().String(workflowFlagName, workflowFlagDefault, workflowDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelAuthWorkflowInfoFlags(depth int, m *models.AuthWorkflowInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, workflowAdded := retrieveAuthWorkflowInfoWorkflowFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || workflowAdded

	return nil, retAdded
}

func retrieveAuthWorkflowInfoWorkflowFlags(depth int, m *models.AuthWorkflowInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	workflowFlagName := fmt.Sprintf("%v.workflow", cmdPrefix)
	if cmd.Flags().Changed(workflowFlagName) {

		var workflowFlagName string
		if cmdPrefix == "" {
			workflowFlagName = "workflow"
		} else {
			workflowFlagName = fmt.Sprintf("%v.workflow", cmdPrefix)
		}

		workflowFlagValue, err := cmd.Flags().GetString(workflowFlagName)
		if err != nil {
			return err, false
		}
		m.Workflow = workflowFlagValue

		retAdded = true
	}

	return nil, retAdded
}
