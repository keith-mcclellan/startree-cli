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

// makeOperationBrokerGetBrokersForTenantV2Cmd returns a cmd to handle operation getBrokersForTenantV2
func makeOperationBrokerGetBrokersForTenantV2Cmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getBrokersForTenantV2",
		Short: `List brokers for a given tenant`,
		RunE:  runOperationBrokerGetBrokersForTenantV2,
	}

	if err := registerOperationBrokerGetBrokersForTenantV2ParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBrokerGetBrokersForTenantV2 uses cmd flags to call endpoint api
func runOperationBrokerGetBrokersForTenantV2(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broker.NewGetBrokersForTenantV2Params()
	if err, _ := retrieveOperationBrokerGetBrokersForTenantV2StateFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBrokerGetBrokersForTenantV2TenantNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBrokerGetBrokersForTenantV2Result(appCli.Broker.GetBrokersForTenantV2(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBrokerGetBrokersForTenantV2ParamFlags registers all flags needed to fill params
func registerOperationBrokerGetBrokersForTenantV2ParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBrokerGetBrokersForTenantV2StateParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBrokerGetBrokersForTenantV2TenantNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBrokerGetBrokersForTenantV2StateParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationBrokerGetBrokersForTenantV2TenantNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	tenantNameDescription := `Required. Name of the tenant`

	var tenantNameFlagName string
	if cmdPrefix == "" {
		tenantNameFlagName = "tenantName"
	} else {
		tenantNameFlagName = fmt.Sprintf("%v.tenantName", cmdPrefix)
	}

	var tenantNameFlagDefault string

	_ = cmd.PersistentFlags().String(tenantNameFlagName, tenantNameFlagDefault, tenantNameDescription)

	return nil
}

func retrieveOperationBrokerGetBrokersForTenantV2StateFlag(m *broker.GetBrokersForTenantV2Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationBrokerGetBrokersForTenantV2TenantNameFlag(m *broker.GetBrokersForTenantV2Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("tenantName") {

		var tenantNameFlagName string
		if cmdPrefix == "" {
			tenantNameFlagName = "tenantName"
		} else {
			tenantNameFlagName = fmt.Sprintf("%v.tenantName", cmdPrefix)
		}

		tenantNameFlagValue, err := cmd.Flags().GetString(tenantNameFlagName)
		if err != nil {
			return err, false
		}
		m.TenantName = tenantNameFlagValue

	}
	return nil, retAdded
}

// parseOperationBrokerGetBrokersForTenantV2Result parses request result and return the string content
func parseOperationBrokerGetBrokersForTenantV2Result(resp0 *broker.GetBrokersForTenantV2OK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*broker.GetBrokersForTenantV2OK)
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
