// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"startree.ai/cli/client/instance"

	"github.com/spf13/cobra"
)

// makeOperationInstanceDropInstanceCmd returns a cmd to handle operation dropInstance
func makeOperationInstanceDropInstanceCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "dropInstance",
		Short: `Drop an instance`,
		RunE:  runOperationInstanceDropInstance,
	}

	if err := registerOperationInstanceDropInstanceParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationInstanceDropInstance uses cmd flags to call endpoint api
func runOperationInstanceDropInstance(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := instance.NewDropInstanceParams()
	if err, _ := retrieveOperationInstanceDropInstanceInstanceNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationInstanceDropInstanceResult(appCli.Instance.DropInstance(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationInstanceDropInstanceParamFlags registers all flags needed to fill params
func registerOperationInstanceDropInstanceParamFlags(cmd *cobra.Command) error {
	if err := registerOperationInstanceDropInstanceInstanceNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationInstanceDropInstanceInstanceNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationInstanceDropInstanceInstanceNameFlag(m *instance.DropInstanceParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationInstanceDropInstanceResult parses request result and return the string content
func parseOperationInstanceDropInstanceResult(resp0 *instance.DropInstanceOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning dropInstanceOK is not supported

		// Non schema case: warning dropInstanceNotFound is not supported

		// Non schema case: warning dropInstanceConflict is not supported

		// Non schema case: warning dropInstanceInternalServerError is not supported

		return "", respErr
	}

	// warning: non schema response dropInstanceOK is not supported by go-swagger cli yet.

	return "", nil
}
