// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"startree.ai/cli/client/broker"

	"github.com/spf13/cobra"
)

// makeOperationBrokerToggleQueryRateLimitingCmd returns a cmd to handle operation toggleQueryRateLimiting
func makeOperationBrokerToggleQueryRateLimitingCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "toggleQueryRateLimiting",
		Short: `Enable/disable the query rate limiting for a broker instance`,
		RunE:  runOperationBrokerToggleQueryRateLimiting,
	}

	if err := registerOperationBrokerToggleQueryRateLimitingParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBrokerToggleQueryRateLimiting uses cmd flags to call endpoint api
func runOperationBrokerToggleQueryRateLimiting(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broker.NewToggleQueryRateLimitingParams()
	if err, _ := retrieveOperationBrokerToggleQueryRateLimitingInstanceNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBrokerToggleQueryRateLimitingStateFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBrokerToggleQueryRateLimitingResult(appCli.Broker.ToggleQueryRateLimiting(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBrokerToggleQueryRateLimitingParamFlags registers all flags needed to fill params
func registerOperationBrokerToggleQueryRateLimitingParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBrokerToggleQueryRateLimitingInstanceNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBrokerToggleQueryRateLimitingStateParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBrokerToggleQueryRateLimitingInstanceNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	instanceNameDescription := `Required. Broker instance name`

	var instanceNameFlagName string
	if cmdPrefix == "" {
		instanceNameFlagName = "instanceName"
	} else {
		instanceNameFlagName = fmt.Sprintf("%v.instanceName", cmdPrefix)
	}

	var instanceNameFlagDefault string

	_ = cmd.PersistentFlags().String(instanceNameFlagName, instanceNameFlagDefault, instanceNameDescription)

	return nil
}
func registerOperationBrokerToggleQueryRateLimitingStateParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	stateDescription := `Enum: ["ENABLE","DISABLE"]. Required. ENABLE|DISABLE`

	var stateFlagName string
	if cmdPrefix == "" {
		stateFlagName = "state"
	} else {
		stateFlagName = fmt.Sprintf("%v.state", cmdPrefix)
	}

	var stateFlagDefault string

	_ = cmd.PersistentFlags().String(stateFlagName, stateFlagDefault, stateDescription)

	if err := cmd.RegisterFlagCompletionFunc(stateFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["ENABLE","DISABLE"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func retrieveOperationBrokerToggleQueryRateLimitingInstanceNameFlag(m *broker.ToggleQueryRateLimitingParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("instanceName") {

		var instanceNameFlagName string
		if cmdPrefix == "" {
			instanceNameFlagName = "instanceName"
		} else {
			instanceNameFlagName = fmt.Sprintf("%v.instanceName", cmdPrefix)
		}

		instanceNameFlagValue, err := cmd.Flags().GetString(instanceNameFlagName)
		if err != nil {
			return err, false
		}
		m.InstanceName = instanceNameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationBrokerToggleQueryRateLimitingStateFlag(m *broker.ToggleQueryRateLimitingParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.State = stateFlagValue

	}
	return nil, retAdded
}

// parseOperationBrokerToggleQueryRateLimitingResult parses request result and return the string content
func parseOperationBrokerToggleQueryRateLimitingResult(resp0 *broker.ToggleQueryRateLimitingOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning toggleQueryRateLimitingOK is not supported

		// Non schema case: warning toggleQueryRateLimitingBadRequest is not supported

		// Non schema case: warning toggleQueryRateLimitingNotFound is not supported

		// Non schema case: warning toggleQueryRateLimitingInternalServerError is not supported

		return "", respErr
	}

	// warning: non schema response toggleQueryRateLimitingOK is not supported by go-swagger cli yet.

	return "", nil
}
