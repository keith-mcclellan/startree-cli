// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/spf13/cobra"
	"startree.ai/cli/models"
)

// Schema cli for SuccessResponse

// register flags to command
func registerModelSuccessResponseFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSuccessResponseStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSuccessResponseStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statusDescription := ``

	var statusFlagName string
	if cmdPrefix == "" {
		statusFlagName = "status"
	} else {
		statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
	}

	var statusFlagDefault string

	_ = cmd.PersistentFlags().String(statusFlagName, statusFlagDefault, statusDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSuccessResponseFlags(depth int, m *models.SuccessResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, statusAdded := retrieveSuccessResponseStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded

	return nil, retAdded
}

func retrieveSuccessResponseStatusFlags(depth int, m *models.SuccessResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statusFlagName := fmt.Sprintf("%v.status", cmdPrefix)
	if cmd.Flags().Changed(statusFlagName) {

		var statusFlagName string
		if cmdPrefix == "" {
			statusFlagName = "status"
		} else {
			statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
		}

		statusFlagValue, err := cmd.Flags().GetString(statusFlagName)
		if err != nil {
			return err, false
		}
		m.Status = statusFlagValue

		retAdded = true
	}

	return nil, retAdded
}
