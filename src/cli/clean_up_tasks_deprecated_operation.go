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

// makeOperationTaskCleanUpTasksDeprecatedCmd returns a cmd to handle operation cleanUpTasksDeprecated
func makeOperationTaskCleanUpTasksDeprecatedCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "cleanUpTasksDeprecated",
		Short: ``,
		RunE:  runOperationTaskCleanUpTasksDeprecated,
	}

	if err := registerOperationTaskCleanUpTasksDeprecatedParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTaskCleanUpTasksDeprecated uses cmd flags to call endpoint api
func runOperationTaskCleanUpTasksDeprecated(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := task.NewCleanUpTasksDeprecatedParams()
	if err, _ := retrieveOperationTaskCleanUpTasksDeprecatedTaskTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationTaskCleanUpTasksDeprecatedResult(appCli.Task.CleanUpTasksDeprecated(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationTaskCleanUpTasksDeprecatedParamFlags registers all flags needed to fill params
func registerOperationTaskCleanUpTasksDeprecatedParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTaskCleanUpTasksDeprecatedTaskTypeParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTaskCleanUpTasksDeprecatedTaskTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationTaskCleanUpTasksDeprecatedTaskTypeFlag(m *task.CleanUpTasksDeprecatedParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationTaskCleanUpTasksDeprecatedResult parses request result and return the string content
func parseOperationTaskCleanUpTasksDeprecatedResult(resp0 *task.CleanUpTasksDeprecatedOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*task.CleanUpTasksDeprecatedOK)
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
