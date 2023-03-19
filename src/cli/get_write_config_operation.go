// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"startree.ai/cli/client/write_api"

	"github.com/spf13/cobra"
)

// makeOperationWriteAPIGetWriteConfigCmd returns a cmd to handle operation getWriteConfig
func makeOperationWriteAPIGetWriteConfigCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getWriteConfig",
		Short: `Gets a config for specific table. May contain Kafka producer configs`,
		RunE:  runOperationWriteAPIGetWriteConfig,
	}

	if err := registerOperationWriteAPIGetWriteConfigParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationWriteAPIGetWriteConfig uses cmd flags to call endpoint api
func runOperationWriteAPIGetWriteConfig(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := write_api.NewGetWriteConfigParams()
	if err, _ := retrieveOperationWriteAPIGetWriteConfigTableFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationWriteAPIGetWriteConfigResult(appCli.WriteAPI.GetWriteConfig(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationWriteAPIGetWriteConfigParamFlags registers all flags needed to fill params
func registerOperationWriteAPIGetWriteConfigParamFlags(cmd *cobra.Command) error {
	if err := registerOperationWriteAPIGetWriteConfigTableParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationWriteAPIGetWriteConfigTableParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	tableDescription := `Required. `

	var tableFlagName string
	if cmdPrefix == "" {
		tableFlagName = "table"
	} else {
		tableFlagName = fmt.Sprintf("%v.table", cmdPrefix)
	}

	var tableFlagDefault string

	_ = cmd.PersistentFlags().String(tableFlagName, tableFlagDefault, tableDescription)

	return nil
}

func retrieveOperationWriteAPIGetWriteConfigTableFlag(m *write_api.GetWriteConfigParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("table") {

		var tableFlagName string
		if cmdPrefix == "" {
			tableFlagName = "table"
		} else {
			tableFlagName = fmt.Sprintf("%v.table", cmdPrefix)
		}

		tableFlagValue, err := cmd.Flags().GetString(tableFlagName)
		if err != nil {
			return err, false
		}
		m.Table = tableFlagValue

	}
	return nil, retAdded
}

// parseOperationWriteAPIGetWriteConfigResult parses request result and return the string content
func parseOperationWriteAPIGetWriteConfigResult(respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning getWriteConfig default is not supported

		return "", respErr
	}
	return "", nil
}
