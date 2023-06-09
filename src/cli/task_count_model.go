// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/spf13/cobra"
	"startree.ai/cli/models"
)

// Schema cli for TaskCount

// register flags to command
func registerModelTaskCountFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTaskCountCompleted(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskCountError(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskCountRunning(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskCountTotal(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskCountUnknown(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskCountWaiting(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTaskCountCompleted(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	completedDescription := ``

	var completedFlagName string
	if cmdPrefix == "" {
		completedFlagName = "completed"
	} else {
		completedFlagName = fmt.Sprintf("%v.completed", cmdPrefix)
	}

	var completedFlagDefault int32

	_ = cmd.PersistentFlags().Int32(completedFlagName, completedFlagDefault, completedDescription)

	return nil
}

func registerTaskCountError(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	errorDescription := ``

	var errorFlagName string
	if cmdPrefix == "" {
		errorFlagName = "error"
	} else {
		errorFlagName = fmt.Sprintf("%v.error", cmdPrefix)
	}

	var errorFlagDefault int32

	_ = cmd.PersistentFlags().Int32(errorFlagName, errorFlagDefault, errorDescription)

	return nil
}

func registerTaskCountRunning(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	runningDescription := ``

	var runningFlagName string
	if cmdPrefix == "" {
		runningFlagName = "running"
	} else {
		runningFlagName = fmt.Sprintf("%v.running", cmdPrefix)
	}

	var runningFlagDefault int32

	_ = cmd.PersistentFlags().Int32(runningFlagName, runningFlagDefault, runningDescription)

	return nil
}

func registerTaskCountTotal(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalDescription := ``

	var totalFlagName string
	if cmdPrefix == "" {
		totalFlagName = "total"
	} else {
		totalFlagName = fmt.Sprintf("%v.total", cmdPrefix)
	}

	var totalFlagDefault int32

	_ = cmd.PersistentFlags().Int32(totalFlagName, totalFlagDefault, totalDescription)

	return nil
}

func registerTaskCountUnknown(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	unknownDescription := ``

	var unknownFlagName string
	if cmdPrefix == "" {
		unknownFlagName = "unknown"
	} else {
		unknownFlagName = fmt.Sprintf("%v.unknown", cmdPrefix)
	}

	var unknownFlagDefault int32

	_ = cmd.PersistentFlags().Int32(unknownFlagName, unknownFlagDefault, unknownDescription)

	return nil
}

func registerTaskCountWaiting(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	waitingDescription := ``

	var waitingFlagName string
	if cmdPrefix == "" {
		waitingFlagName = "waiting"
	} else {
		waitingFlagName = fmt.Sprintf("%v.waiting", cmdPrefix)
	}

	var waitingFlagDefault int32

	_ = cmd.PersistentFlags().Int32(waitingFlagName, waitingFlagDefault, waitingDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTaskCountFlags(depth int, m *models.TaskCount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, completedAdded := retrieveTaskCountCompletedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || completedAdded

	err, errorAdded := retrieveTaskCountErrorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorAdded

	err, runningAdded := retrieveTaskCountRunningFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || runningAdded

	err, totalAdded := retrieveTaskCountTotalFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalAdded

	err, unknownAdded := retrieveTaskCountUnknownFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || unknownAdded

	err, waitingAdded := retrieveTaskCountWaitingFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || waitingAdded

	return nil, retAdded
}

func retrieveTaskCountCompletedFlags(depth int, m *models.TaskCount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	completedFlagName := fmt.Sprintf("%v.completed", cmdPrefix)
	if cmd.Flags().Changed(completedFlagName) {

		var completedFlagName string
		if cmdPrefix == "" {
			completedFlagName = "completed"
		} else {
			completedFlagName = fmt.Sprintf("%v.completed", cmdPrefix)
		}

		completedFlagValue, err := cmd.Flags().GetInt32(completedFlagName)
		if err != nil {
			return err, false
		}
		m.Completed = completedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTaskCountErrorFlags(depth int, m *models.TaskCount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	errorFlagName := fmt.Sprintf("%v.error", cmdPrefix)
	if cmd.Flags().Changed(errorFlagName) {

		var errorFlagName string
		if cmdPrefix == "" {
			errorFlagName = "error"
		} else {
			errorFlagName = fmt.Sprintf("%v.error", cmdPrefix)
		}

		errorFlagValue, err := cmd.Flags().GetInt32(errorFlagName)
		if err != nil {
			return err, false
		}
		m.Error = errorFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTaskCountRunningFlags(depth int, m *models.TaskCount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	runningFlagName := fmt.Sprintf("%v.running", cmdPrefix)
	if cmd.Flags().Changed(runningFlagName) {

		var runningFlagName string
		if cmdPrefix == "" {
			runningFlagName = "running"
		} else {
			runningFlagName = fmt.Sprintf("%v.running", cmdPrefix)
		}

		runningFlagValue, err := cmd.Flags().GetInt32(runningFlagName)
		if err != nil {
			return err, false
		}
		m.Running = runningFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTaskCountTotalFlags(depth int, m *models.TaskCount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalFlagName := fmt.Sprintf("%v.total", cmdPrefix)
	if cmd.Flags().Changed(totalFlagName) {

		var totalFlagName string
		if cmdPrefix == "" {
			totalFlagName = "total"
		} else {
			totalFlagName = fmt.Sprintf("%v.total", cmdPrefix)
		}

		totalFlagValue, err := cmd.Flags().GetInt32(totalFlagName)
		if err != nil {
			return err, false
		}
		m.Total = totalFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTaskCountUnknownFlags(depth int, m *models.TaskCount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	unknownFlagName := fmt.Sprintf("%v.unknown", cmdPrefix)
	if cmd.Flags().Changed(unknownFlagName) {

		var unknownFlagName string
		if cmdPrefix == "" {
			unknownFlagName = "unknown"
		} else {
			unknownFlagName = fmt.Sprintf("%v.unknown", cmdPrefix)
		}

		unknownFlagValue, err := cmd.Flags().GetInt32(unknownFlagName)
		if err != nil {
			return err, false
		}
		m.Unknown = unknownFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTaskCountWaitingFlags(depth int, m *models.TaskCount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	waitingFlagName := fmt.Sprintf("%v.waiting", cmdPrefix)
	if cmd.Flags().Changed(waitingFlagName) {

		var waitingFlagName string
		if cmdPrefix == "" {
			waitingFlagName = "waiting"
		} else {
			waitingFlagName = fmt.Sprintf("%v.waiting", cmdPrefix)
		}

		waitingFlagValue, err := cmd.Flags().GetInt32(waitingFlagName)
		if err != nil {
			return err, false
		}
		m.Waiting = waitingFlagValue

		retAdded = true
	}

	return nil, retAdded
}
