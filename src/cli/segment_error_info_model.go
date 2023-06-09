// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/spf13/cobra"
	"startree.ai/cli/models"
)

// Schema cli for SegmentErrorInfo

// register flags to command
func registerModelSegmentErrorInfoFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSegmentErrorInfoErrorMessage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSegmentErrorInfoStackTrace(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSegmentErrorInfoTimestamp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSegmentErrorInfoErrorMessage(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	errorMessageDescription := ``

	var errorMessageFlagName string
	if cmdPrefix == "" {
		errorMessageFlagName = "errorMessage"
	} else {
		errorMessageFlagName = fmt.Sprintf("%v.errorMessage", cmdPrefix)
	}

	var errorMessageFlagDefault string

	_ = cmd.PersistentFlags().String(errorMessageFlagName, errorMessageFlagDefault, errorMessageDescription)

	return nil
}

func registerSegmentErrorInfoStackTrace(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	stackTraceDescription := ``

	var stackTraceFlagName string
	if cmdPrefix == "" {
		stackTraceFlagName = "stackTrace"
	} else {
		stackTraceFlagName = fmt.Sprintf("%v.stackTrace", cmdPrefix)
	}

	var stackTraceFlagDefault string

	_ = cmd.PersistentFlags().String(stackTraceFlagName, stackTraceFlagDefault, stackTraceDescription)

	return nil
}

func registerSegmentErrorInfoTimestamp(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	timestampDescription := ``

	var timestampFlagName string
	if cmdPrefix == "" {
		timestampFlagName = "timestamp"
	} else {
		timestampFlagName = fmt.Sprintf("%v.timestamp", cmdPrefix)
	}

	var timestampFlagDefault string

	_ = cmd.PersistentFlags().String(timestampFlagName, timestampFlagDefault, timestampDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSegmentErrorInfoFlags(depth int, m *models.SegmentErrorInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, errorMessageAdded := retrieveSegmentErrorInfoErrorMessageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorMessageAdded

	err, stackTraceAdded := retrieveSegmentErrorInfoStackTraceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || stackTraceAdded

	err, timestampAdded := retrieveSegmentErrorInfoTimestampFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || timestampAdded

	return nil, retAdded
}

func retrieveSegmentErrorInfoErrorMessageFlags(depth int, m *models.SegmentErrorInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	errorMessageFlagName := fmt.Sprintf("%v.errorMessage", cmdPrefix)
	if cmd.Flags().Changed(errorMessageFlagName) {

		var errorMessageFlagName string
		if cmdPrefix == "" {
			errorMessageFlagName = "errorMessage"
		} else {
			errorMessageFlagName = fmt.Sprintf("%v.errorMessage", cmdPrefix)
		}

		errorMessageFlagValue, err := cmd.Flags().GetString(errorMessageFlagName)
		if err != nil {
			return err, false
		}
		m.ErrorMessage = errorMessageFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSegmentErrorInfoStackTraceFlags(depth int, m *models.SegmentErrorInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	stackTraceFlagName := fmt.Sprintf("%v.stackTrace", cmdPrefix)
	if cmd.Flags().Changed(stackTraceFlagName) {

		var stackTraceFlagName string
		if cmdPrefix == "" {
			stackTraceFlagName = "stackTrace"
		} else {
			stackTraceFlagName = fmt.Sprintf("%v.stackTrace", cmdPrefix)
		}

		stackTraceFlagValue, err := cmd.Flags().GetString(stackTraceFlagName)
		if err != nil {
			return err, false
		}
		m.StackTrace = stackTraceFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSegmentErrorInfoTimestampFlags(depth int, m *models.SegmentErrorInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	timestampFlagName := fmt.Sprintf("%v.timestamp", cmdPrefix)
	if cmd.Flags().Changed(timestampFlagName) {

		var timestampFlagName string
		if cmdPrefix == "" {
			timestampFlagName = "timestamp"
		} else {
			timestampFlagName = fmt.Sprintf("%v.timestamp", cmdPrefix)
		}

		timestampFlagValue, err := cmd.Flags().GetString(timestampFlagName)
		if err != nil {
			return err, false
		}
		m.Timestamp = timestampFlagValue

		retAdded = true
	}

	return nil, retAdded
}
