// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"startree.ai/cli/client/health"

	"github.com/spf13/cobra"
)

// makeOperationHealthCheckHealthCmd returns a cmd to handle operation checkHealth
func makeOperationHealthCheckHealthCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "checkHealth",
		Short: ``,
		RunE:  runOperationHealthCheckHealth,
	}

	if err := registerOperationHealthCheckHealthParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationHealthCheckHealth uses cmd flags to call endpoint api
func runOperationHealthCheckHealth(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := health.NewCheckHealthParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationHealthCheckHealthResult(appCli.Health.CheckHealth(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationHealthCheckHealthParamFlags registers all flags needed to fill params
func registerOperationHealthCheckHealthParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationHealthCheckHealthResult parses request result and return the string content
func parseOperationHealthCheckHealthResult(resp0 *health.CheckHealthOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning checkHealthOK is not supported

		return "", respErr
	}

	// warning: non schema response checkHealthOK is not supported by go-swagger cli yet.

	return "", nil
}
