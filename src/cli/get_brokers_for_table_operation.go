// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"startree.ai/cli/client/broker"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationBrokerGetBrokersForTableCmd returns a cmd to handle operation getBrokersForTable
func makeOperationBrokerGetBrokersForTableCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getBrokersForTable",
		Short: `List brokers for a given table`,
		RunE:  runOperationBrokerGetBrokersForTable,
	}

	if err := registerOperationBrokerGetBrokersForTableParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBrokerGetBrokersForTable uses cmd flags to call endpoint api
func runOperationBrokerGetBrokersForTable(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broker.NewGetBrokersForTableParams()
	if err, _ := retrieveOperationBrokerGetBrokersForTableStateFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBrokerGetBrokersForTableTableNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBrokerGetBrokersForTableTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBrokerGetBrokersForTableResult(appCli.Broker.GetBrokersForTable(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBrokerGetBrokersForTableParamFlags registers all flags needed to fill params
func registerOperationBrokerGetBrokersForTableParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBrokerGetBrokersForTableStateParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBrokerGetBrokersForTableTableNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBrokerGetBrokersForTableTypeParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBrokerGetBrokersForTableStateParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	stateDescription := `ONLINE|OFFLINE`

	var stateFlagName string
	if cmdPrefix == "" {
		stateFlagName = "state"
	} else {
		stateFlagName = fmt.Sprintf("%v.state", cmdPrefix)
	}

	var stateFlagDefault string

	_ = cmd.PersistentFlags().String(stateFlagName, stateFlagDefault, stateDescription)

	return nil
}
func registerOperationBrokerGetBrokersForTableTableNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationBrokerGetBrokersForTableTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	typeDescription := `OFFLINE|REALTIME`

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

func retrieveOperationBrokerGetBrokersForTableStateFlag(m *broker.GetBrokersForTableParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("state") {

		var stateFlagName string
		if cmdPrefix == "" {
			stateFlagName = "state"
		} else {
			stateFlagName = fmt.Sprintf("%v.state", cmdPrefix)
		}

		stateFlagValue, err := cmd.Flags().GetString(stateFlagName)
		if err != nil {
			return err, false
		}
		m.State = &stateFlagValue

	}
	return nil, retAdded
}
func retrieveOperationBrokerGetBrokersForTableTableNameFlag(m *broker.GetBrokersForTableParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationBrokerGetBrokersForTableTypeFlag(m *broker.GetBrokersForTableParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationBrokerGetBrokersForTableResult parses request result and return the string content
func parseOperationBrokerGetBrokersForTableResult(resp0 *broker.GetBrokersForTableOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*broker.GetBrokersForTableOK)
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
