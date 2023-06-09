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

// makeOperationTaskGetTaskStatesByTableCmd returns a cmd to handle operation getTaskStatesByTable
func makeOperationTaskGetTaskStatesByTableCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getTaskStatesByTable",
		Short: ``,
		RunE:  runOperationTaskGetTaskStatesByTable,
	}

	if err := registerOperationTaskGetTaskStatesByTableParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTaskGetTaskStatesByTable uses cmd flags to call endpoint api
func runOperationTaskGetTaskStatesByTable(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := task.NewGetTaskStatesByTableParams()
	if err, _ := retrieveOperationTaskGetTaskStatesByTableTableNameWithTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTaskGetTaskStatesByTableTaskTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationTaskGetTaskStatesByTableResult(appCli.Task.GetTaskStatesByTable(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationTaskGetTaskStatesByTableParamFlags registers all flags needed to fill params
func registerOperationTaskGetTaskStatesByTableParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTaskGetTaskStatesByTableTableNameWithTypeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTaskGetTaskStatesByTableTaskTypeParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTaskGetTaskStatesByTableTableNameWithTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationTaskGetTaskStatesByTableTaskTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationTaskGetTaskStatesByTableTableNameWithTypeFlag(m *task.GetTaskStatesByTableParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationTaskGetTaskStatesByTableTaskTypeFlag(m *task.GetTaskStatesByTableParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationTaskGetTaskStatesByTableResult parses request result and return the string content
func parseOperationTaskGetTaskStatesByTableResult(resp0 *task.GetTaskStatesByTableOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*task.GetTaskStatesByTableOK)
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
