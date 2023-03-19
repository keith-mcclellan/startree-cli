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

// makeOperationTaskGetTasksDeprecatedCmd returns a cmd to handle operation getTasksDeprecated
func makeOperationTaskGetTasksDeprecatedCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getTasksDeprecated",
		Short: ``,
		RunE:  runOperationTaskGetTasksDeprecated,
	}

	if err := registerOperationTaskGetTasksDeprecatedParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTaskGetTasksDeprecated uses cmd flags to call endpoint api
func runOperationTaskGetTasksDeprecated(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := task.NewGetTasksDeprecatedParams()
	if err, _ := retrieveOperationTaskGetTasksDeprecatedTaskTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationTaskGetTasksDeprecatedResult(appCli.Task.GetTasksDeprecated(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationTaskGetTasksDeprecatedParamFlags registers all flags needed to fill params
func registerOperationTaskGetTasksDeprecatedParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTaskGetTasksDeprecatedTaskTypeParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTaskGetTasksDeprecatedTaskTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationTaskGetTasksDeprecatedTaskTypeFlag(m *task.GetTasksDeprecatedParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationTaskGetTasksDeprecatedResult parses request result and return the string content
func parseOperationTaskGetTasksDeprecatedResult(resp0 *task.GetTasksDeprecatedOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*task.GetTasksDeprecatedOK)
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