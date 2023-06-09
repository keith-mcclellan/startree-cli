// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"startree.ai/cli/client/table"

	"github.com/spf13/cobra"
)

// makeOperationTableGetConsumingSegmentsInfoCmd returns a cmd to handle operation getConsumingSegmentsInfo
func makeOperationTableGetConsumingSegmentsInfoCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getConsumingSegmentsInfo",
		Short: `Gets the status of consumers from all servers.Note that the partitionToOffsetMap has been deprecated and will be removed in the next release. The info is now embedded within each partition's state as currentOffsetsMap.`,
		RunE:  runOperationTableGetConsumingSegmentsInfo,
	}

	if err := registerOperationTableGetConsumingSegmentsInfoParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTableGetConsumingSegmentsInfo uses cmd flags to call endpoint api
func runOperationTableGetConsumingSegmentsInfo(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := table.NewGetConsumingSegmentsInfoParams()
	if err, _ := retrieveOperationTableGetConsumingSegmentsInfoTableNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationTableGetConsumingSegmentsInfoResult(appCli.Table.GetConsumingSegmentsInfo(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationTableGetConsumingSegmentsInfoParamFlags registers all flags needed to fill params
func registerOperationTableGetConsumingSegmentsInfoParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTableGetConsumingSegmentsInfoTableNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTableGetConsumingSegmentsInfoTableNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	tableNameDescription := `Required. Realtime table name with or without type`

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

func retrieveOperationTableGetConsumingSegmentsInfoTableNameFlag(m *table.GetConsumingSegmentsInfoParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationTableGetConsumingSegmentsInfoResult parses request result and return the string content
func parseOperationTableGetConsumingSegmentsInfoResult(resp0 *table.GetConsumingSegmentsInfoOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning getConsumingSegmentsInfoOK is not supported

		// Non schema case: warning getConsumingSegmentsInfoNotFound is not supported

		// Non schema case: warning getConsumingSegmentsInfoInternalServerError is not supported

		return "", respErr
	}

	// warning: non schema response getConsumingSegmentsInfoOK is not supported by go-swagger cli yet.

	return "", nil
}
