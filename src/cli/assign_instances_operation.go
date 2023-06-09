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

// makeOperationTableAssignInstancesCmd returns a cmd to handle operation assignInstances
func makeOperationTableAssignInstancesCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "assignInstances",
		Short: ``,
		RunE:  runOperationTableAssignInstances,
	}

	if err := registerOperationTableAssignInstancesParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTableAssignInstances uses cmd flags to call endpoint api
func runOperationTableAssignInstances(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := table.NewAssignInstancesParams()
	if err, _ := retrieveOperationTableAssignInstancesDryRunFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTableAssignInstancesTableNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTableAssignInstancesTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationTableAssignInstancesResult(appCli.Table.AssignInstances(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationTableAssignInstancesParamFlags registers all flags needed to fill params
func registerOperationTableAssignInstancesParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTableAssignInstancesDryRunParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTableAssignInstancesTableNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTableAssignInstancesTypeParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTableAssignInstancesDryRunParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	dryRunDescription := `Whether to do dry-run`

	var dryRunFlagName string
	if cmdPrefix == "" {
		dryRunFlagName = "dryRun"
	} else {
		dryRunFlagName = fmt.Sprintf("%v.dryRun", cmdPrefix)
	}

	var dryRunFlagDefault bool

	_ = cmd.PersistentFlags().Bool(dryRunFlagName, dryRunFlagDefault, dryRunDescription)

	return nil
}
func registerOperationTableAssignInstancesTableNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationTableAssignInstancesTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	typeDescription := `Enum: ["OFFLINE","CONSUMING","COMPLETED"]. OFFLINE|CONSUMING|COMPLETED`

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	if err := cmd.RegisterFlagCompletionFunc(typeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["OFFLINE","CONSUMING","COMPLETED"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func retrieveOperationTableAssignInstancesDryRunFlag(m *table.AssignInstancesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("dryRun") {

		var dryRunFlagName string
		if cmdPrefix == "" {
			dryRunFlagName = "dryRun"
		} else {
			dryRunFlagName = fmt.Sprintf("%v.dryRun", cmdPrefix)
		}

		dryRunFlagValue, err := cmd.Flags().GetBool(dryRunFlagName)
		if err != nil {
			return err, false
		}
		m.DryRun = &dryRunFlagValue

	}
	return nil, retAdded
}
func retrieveOperationTableAssignInstancesTableNameFlag(m *table.AssignInstancesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationTableAssignInstancesTypeFlag(m *table.AssignInstancesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationTableAssignInstancesResult parses request result and return the string content
func parseOperationTableAssignInstancesResult(resp0 *table.AssignInstancesOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*table.AssignInstancesOK)
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
