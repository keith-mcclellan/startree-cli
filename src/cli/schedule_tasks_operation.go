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

// makeOperationTaskScheduleTasksCmd returns a cmd to handle operation scheduleTasks
func makeOperationTaskScheduleTasksCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "scheduleTasks",
		Short: ``,
		RunE:  runOperationTaskScheduleTasks,
	}

	if err := registerOperationTaskScheduleTasksParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTaskScheduleTasks uses cmd flags to call endpoint api
func runOperationTaskScheduleTasks(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := task.NewScheduleTasksParams()
	if err, _ := retrieveOperationTaskScheduleTasksTableNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTaskScheduleTasksTaskTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationTaskScheduleTasksResult(appCli.Task.ScheduleTasks(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationTaskScheduleTasksParamFlags registers all flags needed to fill params
func registerOperationTaskScheduleTasksParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTaskScheduleTasksTableNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTaskScheduleTasksTaskTypeParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTaskScheduleTasksTableNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	tableNameDescription := `Table name (with type suffix)`

	var tableNameFlagName string
	if cmdPrefix == "" {
		tableNameFlagName = "tableName"
	} else {
		tableNameFlagName = fmt.Sprintf("%v.tableName", cmdPrefix)
	}

	var tableNameFlagDefault string

	_ = cmd.PersistentFlags().String(tableNameFlagName, tableNameFlagDefault, tableNameDescription)

	return nil
}
func registerOperationTaskScheduleTasksTaskTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	taskTypeDescription := `Task type`

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

func retrieveOperationTaskScheduleTasksTableNameFlag(m *task.ScheduleTasksParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("tableName") {

		var tableNameFlagName string
		if cmdPrefix == "" {
			tableNameFlagName = "tableName"
		} else {
			tableNameFlagName = fmt.Sprintf("%v.tableName", cmdPrefix)
		}

		tableNameFlagValue, err := cmd.Flags().GetString(tableNameFlagName)
		if err != nil {
			return err, false
		}
		m.TableName = &tableNameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationTaskScheduleTasksTaskTypeFlag(m *task.ScheduleTasksParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.TaskType = &taskTypeFlagValue

	}
	return nil, retAdded
}

// parseOperationTaskScheduleTasksResult parses request result and return the string content
func parseOperationTaskScheduleTasksResult(resp0 *task.ScheduleTasksOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*task.ScheduleTasksOK)
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