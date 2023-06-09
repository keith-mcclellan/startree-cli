// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"startree.ai/cli/client/auth"

	"github.com/spf13/cobra"
)

// makeOperationAuthVerifyCmd returns a cmd to handle operation verify
func makeOperationAuthVerifyCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "verify",
		Short: ``,
		RunE:  runOperationAuthVerify,
	}

	if err := registerOperationAuthVerifyParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationAuthVerify uses cmd flags to call endpoint api
func runOperationAuthVerify(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := auth.NewVerifyParams()
	if err, _ := retrieveOperationAuthVerifyAccessTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationAuthVerifyEndpointURLFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationAuthVerifyTableNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationAuthVerifyResult(appCli.Auth.Verify(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationAuthVerifyParamFlags registers all flags needed to fill params
func registerOperationAuthVerifyParamFlags(cmd *cobra.Command) error {
	if err := registerOperationAuthVerifyAccessTypeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationAuthVerifyEndpointURLParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationAuthVerifyTableNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationAuthVerifyAccessTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	accessTypeDescription := `Enum: ["CREATE","READ","UPDATE","DELETE"]. API access type`

	var accessTypeFlagName string
	if cmdPrefix == "" {
		accessTypeFlagName = "accessType"
	} else {
		accessTypeFlagName = fmt.Sprintf("%v.accessType", cmdPrefix)
	}

	var accessTypeFlagDefault string

	_ = cmd.PersistentFlags().String(accessTypeFlagName, accessTypeFlagDefault, accessTypeDescription)

	if err := cmd.RegisterFlagCompletionFunc(accessTypeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["CREATE","READ","UPDATE","DELETE"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}
func registerOperationAuthVerifyEndpointURLParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	endpointUrlDescription := `Endpoint URL`

	var endpointUrlFlagName string
	if cmdPrefix == "" {
		endpointUrlFlagName = "endpointUrl"
	} else {
		endpointUrlFlagName = fmt.Sprintf("%v.endpointUrl", cmdPrefix)
	}

	var endpointUrlFlagDefault string

	_ = cmd.PersistentFlags().String(endpointUrlFlagName, endpointUrlFlagDefault, endpointUrlDescription)

	return nil
}
func registerOperationAuthVerifyTableNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	tableNameDescription := `Table name without type`

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

func retrieveOperationAuthVerifyAccessTypeFlag(m *auth.VerifyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("accessType") {

		var accessTypeFlagName string
		if cmdPrefix == "" {
			accessTypeFlagName = "accessType"
		} else {
			accessTypeFlagName = fmt.Sprintf("%v.accessType", cmdPrefix)
		}

		accessTypeFlagValue, err := cmd.Flags().GetString(accessTypeFlagName)
		if err != nil {
			return err, false
		}
		m.AccessType = &accessTypeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationAuthVerifyEndpointURLFlag(m *auth.VerifyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("endpointUrl") {

		var endpointUrlFlagName string
		if cmdPrefix == "" {
			endpointUrlFlagName = "endpointUrl"
		} else {
			endpointUrlFlagName = fmt.Sprintf("%v.endpointUrl", cmdPrefix)
		}

		endpointUrlFlagValue, err := cmd.Flags().GetString(endpointUrlFlagName)
		if err != nil {
			return err, false
		}
		m.EndpointURL = &endpointUrlFlagValue

	}
	return nil, retAdded
}
func retrieveOperationAuthVerifyTableNameFlag(m *auth.VerifyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.TableName = &tableNameFlagValue

	}
	return nil, retAdded
}

// parseOperationAuthVerifyResult parses request result and return the string content
func parseOperationAuthVerifyResult(resp0 *auth.VerifyOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning verifyOK is not supported

		// Non schema case: warning verifyInternalServerError is not supported

		return "", respErr
	}

	// warning: non schema response verifyOK is not supported by go-swagger cli yet.

	return "", nil
}
