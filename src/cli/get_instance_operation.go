// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"startree.ai/cli/client/instance"

	"github.com/spf13/cobra"
)

// makeOperationInstanceGetInstanceCmd returns a cmd to handle operation getInstance
func makeOperationInstanceGetInstanceCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getInstance",
		Short: ``,
		RunE:  runOperationInstanceGetInstance,
	}

	if err := registerOperationInstanceGetInstanceParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationInstanceGetInstance uses cmd flags to call endpoint api
func runOperationInstanceGetInstance(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := instance.NewGetInstanceParams()
	if err, _ := retrieveOperationInstanceGetInstanceInstanceNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationInstanceGetInstanceResult(appCli.Instance.GetInstance(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationInstanceGetInstanceParamFlags registers all flags needed to fill params
func registerOperationInstanceGetInstanceParamFlags(cmd *cobra.Command) error {
	if err := registerOperationInstanceGetInstanceInstanceNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationInstanceGetInstanceInstanceNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	instanceNameDescription := `Required. Instance name`

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

func retrieveOperationInstanceGetInstanceInstanceNameFlag(m *instance.GetInstanceParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationInstanceGetInstanceResult parses request result and return the string content
func parseOperationInstanceGetInstanceResult(resp0 *instance.GetInstanceOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning getInstanceOK is not supported

		// Non schema case: warning getInstanceNotFound is not supported

		// Non schema case: warning getInstanceInternalServerError is not supported

		return "", respErr
	}

	// warning: non schema response getInstanceOK is not supported by go-swagger cli yet.

	return "", nil
}
