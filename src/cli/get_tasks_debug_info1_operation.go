// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"startree.ai/cli/client/task"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationTaskGetTasksDebugInfo1Cmd returns a cmd to handle operation getTasksDebugInfo1
func makeOperationTaskGetTasksDebugInfo1Cmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getTasksDebugInfo_1",
		Short: ``,
		RunE:  runOperationTaskGetTasksDebugInfo1,
	}

	if err := registerOperationTaskGetTasksDebugInfo1ParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTaskGetTasksDebugInfo1 uses cmd flags to call endpoint api
func runOperationTaskGetTasksDebugInfo1(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := task.NewGetTasksDebugInfo1Params()
	if err, _ := retrieveOperationTaskGetTasksDebugInfo1TableNameWithTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTaskGetTasksDebugInfo1TaskTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTaskGetTasksDebugInfo1VerbosityFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationTaskGetTasksDebugInfo1Result(appCli.Task.GetTasksDebugInfo1(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationTaskGetTasksDebugInfo1ParamFlags registers all flags needed to fill params
func registerOperationTaskGetTasksDebugInfo1ParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTaskGetTasksDebugInfo1TableNameWithTypeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTaskGetTasksDebugInfo1TaskTypeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTaskGetTasksDebugInfo1VerbosityParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTaskGetTasksDebugInfo1TableNameWithTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	tableNameWithTypeDescription := `Required. Table name with type`

	var tableNameWithTypeFlagName string
	if cmdPrefix == "" {
		tableNameWithTypeFlagName = "tableNameWithType"
	} else {
		tableNameWithTypeFlagName = fmt.Sprintf("%v.tableNameWithType", cmdPrefix)
	}

	var tableNameWithTypeFlagDefault string

	_ = cmd.PersistentFlags().String(tableNameWithTypeFlagName, tableNameWithTypeFlagDefault, tableNameWithTypeDescription)

	return nil
}
func registerOperationTaskGetTasksDebugInfo1TaskTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	taskTypeDescription := `Required. Task type`

	var taskTypeFlagName string
	if cmdPrefix == "" {
		taskTypeFlagName = "taskType"
	} else {
		taskTypeFlagName = fmt.Sprintf("%v.taskType", cmdPrefix)
	}

	var taskTypeFlagDefault string

	_ = cmd.PersistentFlags().String(taskTypeFlagName, taskTypeFlagDefault, taskTypeDescription)

	return nil
}
func registerOperationTaskGetTasksDebugInfo1VerbosityParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	verbosityDescription := `verbosity (Prints information for all the tasks for the given task type and table.By default, only prints subtask details for running and error tasks. Value of > 0 prints subtask details for all tasks)`

	var verbosityFlagName string
	if cmdPrefix == "" {
		verbosityFlagName = "verbosity"
	} else {
		verbosityFlagName = fmt.Sprintf("%v.verbosity", cmdPrefix)
	}

	var verbosityFlagDefault int32

	_ = cmd.PersistentFlags().Int32(verbosityFlagName, verbosityFlagDefault, verbosityDescription)

	return nil
}

func retrieveOperationTaskGetTasksDebugInfo1TableNameWithTypeFlag(m *task.GetTasksDebugInfo1Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("tableNameWithType") {

		var tableNameWithTypeFlagName string
		if cmdPrefix == "" {
			tableNameWithTypeFlagName = "tableNameWithType"
		} else {
			tableNameWithTypeFlagName = fmt.Sprintf("%v.tableNameWithType", cmdPrefix)
		}

		tableNameWithTypeFlagValue, err := cmd.Flags().GetString(tableNameWithTypeFlagName)
		if err != nil {
			return err, false
		}
		m.TableNameWithType = tableNameWithTypeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationTaskGetTasksDebugInfo1TaskTypeFlag(m *task.GetTasksDebugInfo1Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("taskType") {

		var taskTypeFlagName string
		if cmdPrefix == "" {
			taskTypeFlagName = "taskType"
		} else {
			taskTypeFlagName = fmt.Sprintf("%v.taskType", cmdPrefix)
		}

		taskTypeFlagValue, err := cmd.Flags().GetString(taskTypeFlagName)
		if err != nil {
			return err, false
		}
		m.TaskType = taskTypeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationTaskGetTasksDebugInfo1VerbosityFlag(m *task.GetTasksDebugInfo1Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("verbosity") {

		var verbosityFlagName string
		if cmdPrefix == "" {
			verbosityFlagName = "verbosity"
		} else {
			verbosityFlagName = fmt.Sprintf("%v.verbosity", cmdPrefix)
		}

		verbosityFlagValue, err := cmd.Flags().GetInt32(verbosityFlagName)
		if err != nil {
			return err, false
		}
		m.Verbosity = &verbosityFlagValue

	}
	return nil, retAdded
}

// parseOperationTaskGetTasksDebugInfo1Result parses request result and return the string content
func parseOperationTaskGetTasksDebugInfo1Result(resp0 *task.GetTasksDebugInfo1OK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*task.GetTasksDebugInfo1OK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}