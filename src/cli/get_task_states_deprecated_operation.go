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

// makeOperationTaskGetTaskStatesDeprecatedCmd returns a cmd to handle operation getTaskStatesDeprecated
func makeOperationTaskGetTaskStatesDeprecatedCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getTaskStatesDeprecated",
		Short: ``,
		RunE:  runOperationTaskGetTaskStatesDeprecated,
	}

	if err := registerOperationTaskGetTaskStatesDeprecatedParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTaskGetTaskStatesDeprecated uses cmd flags to call endpoint api
func runOperationTaskGetTaskStatesDeprecated(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := task.NewGetTaskStatesDeprecatedParams()
	if err, _ := retrieveOperationTaskGetTaskStatesDeprecatedTaskTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationTaskGetTaskStatesDeprecatedResult(appCli.Task.GetTaskStatesDeprecated(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationTaskGetTaskStatesDeprecatedParamFlags registers all flags needed to fill params
func registerOperationTaskGetTaskStatesDeprecatedParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTaskGetTaskStatesDeprecatedTaskTypeParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTaskGetTaskStatesDeprecatedTaskTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationTaskGetTaskStatesDeprecatedTaskTypeFlag(m *task.GetTaskStatesDeprecatedParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationTaskGetTaskStatesDeprecatedResult parses request result and return the string content
func parseOperationTaskGetTaskStatesDeprecatedResult(resp0 *task.GetTaskStatesDeprecatedOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*task.GetTaskStatesDeprecatedOK)
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
