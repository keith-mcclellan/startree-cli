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

// makeOperationTableDeleteTableCmd returns a cmd to handle operation deleteTable
func makeOperationTableDeleteTableCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteTable",
		Short: `Deletes a table`,
		RunE:  runOperationTableDeleteTable,
	}

	if err := registerOperationTableDeleteTableParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTableDeleteTable uses cmd flags to call endpoint api
func runOperationTableDeleteTable(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := table.NewDeleteTableParams()
	if err, _ := retrieveOperationTableDeleteTableRetentionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTableDeleteTableTableNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTableDeleteTableTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationTableDeleteTableResult(appCli.Table.DeleteTable(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationTableDeleteTableParamFlags registers all flags needed to fill params
func registerOperationTableDeleteTableParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTableDeleteTableRetentionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTableDeleteTableTableNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTableDeleteTableTypeParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTableDeleteTableRetentionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	retentionDescription := `Retention period for the table segments (e.g. 12h, 3d); If not set, the retention period will default to the first config that's not null: the cluster setting, then '7d'. Using 0d or -1d will instantly delete segments without retention`

	var retentionFlagName string
	if cmdPrefix == "" {
		retentionFlagName = "retention"
	} else {
		retentionFlagName = fmt.Sprintf("%v.retention", cmdPrefix)
	}

	var retentionFlagDefault string

	_ = cmd.PersistentFlags().String(retentionFlagName, retentionFlagDefault, retentionDescription)

	return nil
}
func registerOperationTableDeleteTableTableNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	tableNameDescription := `Required. Name of the table to delete`

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
func registerOperationTableDeleteTableTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	typeDescription := `realtime|offline`

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	return nil
}

func retrieveOperationTableDeleteTableRetentionFlag(m *table.DeleteTableParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("retention") {

		var retentionFlagName string
		if cmdPrefix == "" {
			retentionFlagName = "retention"
		} else {
			retentionFlagName = fmt.Sprintf("%v.retention", cmdPrefix)
		}

		retentionFlagValue, err := cmd.Flags().GetString(retentionFlagName)
		if err != nil {
			return err, false
		}
		m.Retention = &retentionFlagValue

	}
	return nil, retAdded
}
func retrieveOperationTableDeleteTableTableNameFlag(m *table.DeleteTableParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationTableDeleteTableTypeFlag(m *table.DeleteTableParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("type") {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = &typeFlagValue

	}
	return nil, retAdded
}

// parseOperationTableDeleteTableResult parses request result and return the string content
func parseOperationTableDeleteTableResult(resp0 *table.DeleteTableOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*table.DeleteTableOK)
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
