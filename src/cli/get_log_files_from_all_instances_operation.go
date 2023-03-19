// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"startree.ai/cli/client/logger"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationLoggerGetLogFilesFromAllInstancesCmd returns a cmd to handle operation getLogFilesFromAllInstances
func makeOperationLoggerGetLogFilesFromAllInstancesCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getLogFilesFromAllInstances",
		Short: ``,
		RunE:  runOperationLoggerGetLogFilesFromAllInstances,
	}

	if err := registerOperationLoggerGetLogFilesFromAllInstancesParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationLoggerGetLogFilesFromAllInstances uses cmd flags to call endpoint api
func runOperationLoggerGetLogFilesFromAllInstances(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := logger.NewGetLogFilesFromAllInstancesParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationLoggerGetLogFilesFromAllInstancesResult(appCli.Logger.GetLogFilesFromAllInstances(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationLoggerGetLogFilesFromAllInstancesParamFlags registers all flags needed to fill params
func registerOperationLoggerGetLogFilesFromAllInstancesParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationLoggerGetLogFilesFromAllInstancesResult parses request result and return the string content
func parseOperationLoggerGetLogFilesFromAllInstancesResult(resp0 *logger.GetLogFilesFromAllInstancesOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*logger.GetLogFilesFromAllInstancesOK)
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