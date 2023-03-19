// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"startree.ai/cli/client/instance"

	"github.com/spf13/cobra"
)

// makeOperationInstanceUpdateInstanceTagsCmd returns a cmd to handle operation updateInstanceTags
func makeOperationInstanceUpdateInstanceTagsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "updateInstanceTags",
		Short: `Update the tags of the specified instance`,
		RunE:  runOperationInstanceUpdateInstanceTags,
	}

	if err := registerOperationInstanceUpdateInstanceTagsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationInstanceUpdateInstanceTags uses cmd flags to call endpoint api
func runOperationInstanceUpdateInstanceTags(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := instance.NewUpdateInstanceTagsParams()
	if err, _ := retrieveOperationInstanceUpdateInstanceTagsInstanceNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationInstanceUpdateInstanceTagsTagsFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationInstanceUpdateInstanceTagsUpdateBrokerResourceFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationInstanceUpdateInstanceTagsResult(appCli.Instance.UpdateInstanceTags(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationInstanceUpdateInstanceTagsParamFlags registers all flags needed to fill params
func registerOperationInstanceUpdateInstanceTagsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationInstanceUpdateInstanceTagsInstanceNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationInstanceUpdateInstanceTagsTagsParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationInstanceUpdateInstanceTagsUpdateBrokerResourceParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationInstanceUpdateInstanceTagsInstanceNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationInstanceUpdateInstanceTagsTagsParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	tagsDescription := `Required. Comma separated tags list`

	var tagsFlagName string
	if cmdPrefix == "" {
		tagsFlagName = "tags"
	} else {
		tagsFlagName = fmt.Sprintf("%v.tags", cmdPrefix)
	}

	var tagsFlagDefault string

	_ = cmd.PersistentFlags().String(tagsFlagName, tagsFlagDefault, tagsDescription)

	return nil
}
func registerOperationInstanceUpdateInstanceTagsUpdateBrokerResourceParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	updateBrokerResourceDescription := `Whether to update broker resource for broker instance`

	var updateBrokerResourceFlagName string
	if cmdPrefix == "" {
		updateBrokerResourceFlagName = "updateBrokerResource"
	} else {
		updateBrokerResourceFlagName = fmt.Sprintf("%v.updateBrokerResource", cmdPrefix)
	}

	var updateBrokerResourceFlagDefault bool

	_ = cmd.PersistentFlags().Bool(updateBrokerResourceFlagName, updateBrokerResourceFlagDefault, updateBrokerResourceDescription)

	return nil
}

func retrieveOperationInstanceUpdateInstanceTagsInstanceNameFlag(m *instance.UpdateInstanceTagsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationInstanceUpdateInstanceTagsTagsFlag(m *instance.UpdateInstanceTagsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("tags") {

		var tagsFlagName string
		if cmdPrefix == "" {
			tagsFlagName = "tags"
		} else {
			tagsFlagName = fmt.Sprintf("%v.tags", cmdPrefix)
		}

		tagsFlagValue, err := cmd.Flags().GetString(tagsFlagName)
		if err != nil {
			return err, false
		}
		m.Tags = tagsFlagValue

	}
	return nil, retAdded
}
func retrieveOperationInstanceUpdateInstanceTagsUpdateBrokerResourceFlag(m *instance.UpdateInstanceTagsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("updateBrokerResource") {

		var updateBrokerResourceFlagName string
		if cmdPrefix == "" {
			updateBrokerResourceFlagName = "updateBrokerResource"
		} else {
			updateBrokerResourceFlagName = fmt.Sprintf("%v.updateBrokerResource", cmdPrefix)
		}

		updateBrokerResourceFlagValue, err := cmd.Flags().GetBool(updateBrokerResourceFlagName)
		if err != nil {
			return err, false
		}
		m.UpdateBrokerResource = &updateBrokerResourceFlagValue

	}
	return nil, retAdded
}

// parseOperationInstanceUpdateInstanceTagsResult parses request result and return the string content
func parseOperationInstanceUpdateInstanceTagsResult(resp0 *instance.UpdateInstanceTagsOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning updateInstanceTagsOK is not supported

		// Non schema case: warning updateInstanceTagsBadRequest is not supported

		// Non schema case: warning updateInstanceTagsNotFound is not supported

		// Non schema case: warning updateInstanceTagsInternalServerError is not supported

		return "", respErr
	}

	// warning: non schema response updateInstanceTagsOK is not supported by go-swagger cli yet.

	return "", nil
}
