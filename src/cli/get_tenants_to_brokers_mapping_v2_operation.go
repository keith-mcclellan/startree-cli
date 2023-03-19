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

// makeOperationBrokerGetTenantsToBrokersMappingV2Cmd returns a cmd to handle operation getTenantsToBrokersMappingV2
func makeOperationBrokerGetTenantsToBrokersMappingV2Cmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getTenantsToBrokersMappingV2",
		Short: `List tenants to brokers mappings`,
		RunE:  runOperationBrokerGetTenantsToBrokersMappingV2,
	}

	if err := registerOperationBrokerGetTenantsToBrokersMappingV2ParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBrokerGetTenantsToBrokersMappingV2 uses cmd flags to call endpoint api
func runOperationBrokerGetTenantsToBrokersMappingV2(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broker.NewGetTenantsToBrokersMappingV2Params()
	if err, _ := retrieveOperationBrokerGetTenantsToBrokersMappingV2StateFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBrokerGetTenantsToBrokersMappingV2Result(appCli.Broker.GetTenantsToBrokersMappingV2(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBrokerGetTenantsToBrokersMappingV2ParamFlags registers all flags needed to fill params
func registerOperationBrokerGetTenantsToBrokersMappingV2ParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBrokerGetTenantsToBrokersMappingV2StateParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBrokerGetTenantsToBrokersMappingV2StateParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationBrokerGetTenantsToBrokersMappingV2StateFlag(m *broker.GetTenantsToBrokersMappingV2Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationBrokerGetTenantsToBrokersMappingV2Result parses request result and return the string content
func parseOperationBrokerGetTenantsToBrokersMappingV2Result(resp0 *broker.GetTenantsToBrokersMappingV2OK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*broker.GetTenantsToBrokersMappingV2OK)
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
