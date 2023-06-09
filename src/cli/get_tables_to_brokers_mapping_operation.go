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

// makeOperationBrokerGetTablesToBrokersMappingCmd returns a cmd to handle operation getTablesToBrokersMapping
func makeOperationBrokerGetTablesToBrokersMappingCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getTablesToBrokersMapping",
		Short: `List tables to brokers mappings`,
		RunE:  runOperationBrokerGetTablesToBrokersMapping,
	}

	if err := registerOperationBrokerGetTablesToBrokersMappingParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBrokerGetTablesToBrokersMapping uses cmd flags to call endpoint api
func runOperationBrokerGetTablesToBrokersMapping(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broker.NewGetTablesToBrokersMappingParams()
	if err, _ := retrieveOperationBrokerGetTablesToBrokersMappingStateFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBrokerGetTablesToBrokersMappingResult(appCli.Broker.GetTablesToBrokersMapping(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBrokerGetTablesToBrokersMappingParamFlags registers all flags needed to fill params
func registerOperationBrokerGetTablesToBrokersMappingParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBrokerGetTablesToBrokersMappingStateParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBrokerGetTablesToBrokersMappingStateParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationBrokerGetTablesToBrokersMappingStateFlag(m *broker.GetTablesToBrokersMappingParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationBrokerGetTablesToBrokersMappingResult parses request result and return the string content
func parseOperationBrokerGetTablesToBrokersMappingResult(resp0 *broker.GetTablesToBrokersMappingOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*broker.GetTablesToBrokersMappingOK)
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
