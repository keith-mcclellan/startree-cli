// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"startree.ai/cli/client/table"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationTableAddTableCmd returns a cmd to handle operation addTable
func makeOperationTableAddTableCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "addTable",
		Short: `Adds a table`,
		RunE:  runOperationTableAddTable,
	}

	if err := registerOperationTableAddTableParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTableAddTable uses cmd flags to call endpoint api
func runOperationTableAddTable(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := table.NewAddTableParams()
	if err, _ := retrieveOperationTableAddTableBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTableAddTableValidationTypesToSkipFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationTableAddTableResult(appCli.Table.AddTable(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationTableAddTableParamFlags registers all flags needed to fill params
func registerOperationTableAddTableParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTableAddTableBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTableAddTableValidationTypesToSkipParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTableAddTableBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	bodyDescription := ``

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	var bodyFlagDefault string

	_ = cmd.PersistentFlags().String(bodyFlagName, bodyFlagDefault, bodyDescription)

	return nil
}
func registerOperationTableAddTableValidationTypesToSkipParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	validationTypesToSkipDescription := `comma separated list of validation type(s) to skip. supported types: (ALL|TASK|UPSERT)`

	var validationTypesToSkipFlagName string
	if cmdPrefix == "" {
		validationTypesToSkipFlagName = "validationTypesToSkip"
	} else {
		validationTypesToSkipFlagName = fmt.Sprintf("%v.validationTypesToSkip", cmdPrefix)
	}

	var validationTypesToSkipFlagDefault string

	_ = cmd.PersistentFlags().String(validationTypesToSkipFlagName, validationTypesToSkipFlagDefault, validationTypesToSkipDescription)

	return nil
}

func retrieveOperationTableAddTableBodyFlag(m *table.AddTableParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {

		var bodyFlagName string
		if cmdPrefix == "" {
			bodyFlagName = "body"
		} else {
			bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
		}

		bodyFlagValue, err := cmd.Flags().GetString(bodyFlagName)
		if err != nil {
			return err, false
		}
		m.Body = bodyFlagValue

	}
	return nil, retAdded
}
func retrieveOperationTableAddTableValidationTypesToSkipFlag(m *table.AddTableParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("validationTypesToSkip") {

		var validationTypesToSkipFlagName string
		if cmdPrefix == "" {
			validationTypesToSkipFlagName = "validationTypesToSkip"
		} else {
			validationTypesToSkipFlagName = fmt.Sprintf("%v.validationTypesToSkip", cmdPrefix)
		}

		validationTypesToSkipFlagValue, err := cmd.Flags().GetString(validationTypesToSkipFlagName)
		if err != nil {
			return err, false
		}
		m.ValidationTypesToSkip = &validationTypesToSkipFlagValue

	}
	return nil, retAdded
}

// parseOperationTableAddTableResult parses request result and return the string content
func parseOperationTableAddTableResult(resp0 *table.AddTableOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*table.AddTableOK)
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