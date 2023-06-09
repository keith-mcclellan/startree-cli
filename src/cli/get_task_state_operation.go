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

// makeOperationTaskGetTaskStateCmd returns a cmd to handle operation getTaskState
func makeOperationTaskGetTaskStateCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getTaskState",
		Short: ``,
		RunE:  runOperationTaskGetTaskState,
	}

	if err := registerOperationTaskGetTaskStateParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTaskGetTaskState uses cmd flags to call endpoint api
func runOperationTaskGetTaskState(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := task.NewGetTaskStateParams()
	if err, _ := retrieveOperationTaskGetTaskStateTaskNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationTaskGetTaskStateResult(appCli.Task.GetTaskState(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationTaskGetTaskStateParamFlags registers all flags needed to fill params
func registerOperationTaskGetTaskStateParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTaskGetTaskStateTaskNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTaskGetTaskStateTaskNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	taskNameDescription := `Required. Task name`

	var taskNameFlagName string
	if cmdPrefix == "" {
		taskNameFlagName = "taskName"
	} else {
		taskNameFlagName = fmt.Sprintf("%v.taskName", cmdPrefix)
	}

	var taskNameFlagDefault string

	_ = cmd.PersistentFlags().String(taskNameFlagName, taskNameFlagDefault, taskNameDescription)

	return nil
}

func retrieveOperationTaskGetTaskStateTaskNameFlag(m *task.GetTaskStateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("taskName") {

		var taskNameFlagName string
		if cmdPrefix == "" {
			taskNameFlagName = "taskName"
		} else {
			taskNameFlagName = fmt.Sprintf("%v.taskName", cmdPrefix)
		}

		taskNameFlagValue, err := cmd.Flags().GetString(taskNameFlagName)
		if err != nil {
			return err, false
		}
		m.TaskName = taskNameFlagValue

	}
	return nil, retAdded
}

// parseOperationTaskGetTaskStateResult parses request result and return the string content
func parseOperationTaskGetTaskStateResult(resp0 *task.GetTaskStateOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*task.GetTaskStateOK)
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
		msgStr := fmt.Sprintf("%v", resp0.Payload)
		return string(msgStr), nil
	}

	return "", nil
}
