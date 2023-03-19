// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"startree.ai/cli/client/periodic_task"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationPeriodicTaskGetPeriodicTaskNamesCmd returns a cmd to handle operation getPeriodicTaskNames
func makeOperationPeriodicTaskGetPeriodicTaskNamesCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getPeriodicTaskNames",
		Short: ``,
		RunE:  runOperationPeriodicTaskGetPeriodicTaskNames,
	}

	if err := registerOperationPeriodicTaskGetPeriodicTaskNamesParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationPeriodicTaskGetPeriodicTaskNames uses cmd flags to call endpoint api
func runOperationPeriodicTaskGetPeriodicTaskNames(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := periodic_task.NewGetPeriodicTaskNamesParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationPeriodicTaskGetPeriodicTaskNamesResult(appCli.PeriodicTask.GetPeriodicTaskNames(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationPeriodicTaskGetPeriodicTaskNamesParamFlags registers all flags needed to fill params
func registerOperationPeriodicTaskGetPeriodicTaskNamesParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationPeriodicTaskGetPeriodicTaskNamesResult parses request result and return the string content
func parseOperationPeriodicTaskGetPeriodicTaskNamesResult(resp0 *periodic_task.GetPeriodicTaskNamesOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*periodic_task.GetPeriodicTaskNamesOK)
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
