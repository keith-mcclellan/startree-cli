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

// makeOperationTaskGetTasksCmd returns a cmd to handle operation getTasks
func makeOperationTaskGetTasksCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getTasks",
		Short: ``,
		RunE:  runOperationTaskGetTasks,
	}

	if err := registerOperationTaskGetTasksParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTaskGetTasks uses cmd flags to call endpoint api
func runOperationTaskGetTasks(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := task.NewGetTasksParams()
	if err, _ := retrieveOperationTaskGetTasksTaskTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationTaskGetTasksResult(appCli.Task.GetTasks(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationTaskGetTasksParamFlags registers all flags needed to fill params
func registerOperationTaskGetTasksParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTaskGetTasksTaskTypeParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTaskGetTasksTaskTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationTaskGetTasksTaskTypeFlag(m *task.GetTasksParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationTaskGetTasksResult parses request result and return the string content
func parseOperationTaskGetTasksResult(resp0 *task.GetTasksOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*task.GetTasksOK)
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
