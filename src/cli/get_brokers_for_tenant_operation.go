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

// makeOperationBrokerGetBrokersForTenantCmd returns a cmd to handle operation getBrokersForTenant
func makeOperationBrokerGetBrokersForTenantCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getBrokersForTenant",
		Short: `List brokers for a given tenant`,
		RunE:  runOperationBrokerGetBrokersForTenant,
	}

	if err := registerOperationBrokerGetBrokersForTenantParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBrokerGetBrokersForTenant uses cmd flags to call endpoint api
func runOperationBrokerGetBrokersForTenant(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broker.NewGetBrokersForTenantParams()
	if err, _ := retrieveOperationBrokerGetBrokersForTenantStateFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBrokerGetBrokersForTenantTenantNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBrokerGetBrokersForTenantResult(appCli.Broker.GetBrokersForTenant(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBrokerGetBrokersForTenantParamFlags registers all flags needed to fill params
func registerOperationBrokerGetBrokersForTenantParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBrokerGetBrokersForTenantStateParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBrokerGetBrokersForTenantTenantNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBrokerGetBrokersForTenantStateParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationBrokerGetBrokersForTenantTenantNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationBrokerGetBrokersForTenantStateFlag(m *broker.GetBrokersForTenantParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationBrokerGetBrokersForTenantTenantNameFlag(m *broker.GetBrokersForTenantParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationBrokerGetBrokersForTenantResult parses request result and return the string content
func parseOperationBrokerGetBrokersForTenantResult(resp0 *broker.GetBrokersForTenantOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*broker.GetBrokersForTenantOK)
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
