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

// makeOperationTableSetInstancePartitionsCmd returns a cmd to handle operation setInstancePartitions
func makeOperationTableSetInstancePartitionsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "setInstancePartitions",
		Short: ``,
		RunE:  runOperationTableSetInstancePartitions,
	}

	if err := registerOperationTableSetInstancePartitionsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTableSetInstancePartitions uses cmd flags to call endpoint api
func runOperationTableSetInstancePartitions(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := table.NewSetInstancePartitionsParams()
	if err, _ := retrieveOperationTableSetInstancePartitionsBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTableSetInstancePartitionsTableNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationTableSetInstancePartitionsResult(appCli.Table.SetInstancePartitions(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationTableSetInstancePartitionsParamFlags registers all flags needed to fill params
func registerOperationTableSetInstancePartitionsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTableSetInstancePartitionsBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTableSetInstancePartitionsTableNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTableSetInstancePartitionsBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationTableSetInstancePartitionsTableNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	tableNameDescription := `Required. Name of the table`

	var tableNameFlagName string
	if cmdPrefix == "" {
		tableNameFlagName = "tableName"
	} else {
		tableNameFlagName = fmt.Sprintf("%v.tableName", cmdPrefix)
	}

	var tableNameFlagDefault string

	_ = cmd.PersistentFlags().String(tableNameFlagName, tableNameFlagDefault, tableNameDescription)

	return nil
}

func retrieveOperationTableSetInstancePartitionsBodyFlag(m *table.SetInstancePartitionsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationTableSetInstancePartitionsTableNameFlag(m *table.SetInstancePartitionsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("tableName") {

		var tableNameFlagName string
		if cmdPrefix == "" {
			tableNameFlagName = "tableName"
		} else {
			tableNameFlagName = fmt.Sprintf("%v.tableName", cmdPrefix)
		}

		tableNameFlagValue, err := cmd.Flags().GetString(tableNameFlagName)
		if err != nil {
			return err, false
		}
		m.TableName = tableNameFlagValue

	}
	return nil, retAdded
}

// parseOperationTableSetInstancePartitionsResult parses request result and return the string content
func parseOperationTableSetInstancePartitionsResult(resp0 *table.SetInstancePartitionsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*table.SetInstancePartitionsOK)
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
