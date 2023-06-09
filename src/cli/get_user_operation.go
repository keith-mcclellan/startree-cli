// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"startree.ai/cli/client/user"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationUserGetUserCmd returns a cmd to handle operation getUser
func makeOperationUserGetUserCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getUser",
		Short: `Get an user in cluster`,
		RunE:  runOperationUserGetUser,
	}

	if err := registerOperationUserGetUserParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationUserGetUser uses cmd flags to call endpoint api
func runOperationUserGetUser(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := user.NewGetUserParams()
	if err, _ := retrieveOperationUserGetUserComponentFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUserGetUserUsernameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationUserGetUserResult(appCli.User.GetUser(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationUserGetUserParamFlags registers all flags needed to fill params
func registerOperationUserGetUserParamFlags(cmd *cobra.Command) error {
	if err := registerOperationUserGetUserComponentParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUserGetUserUsernameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationUserGetUserComponentParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	componentDescription := ``

	var componentFlagName string
	if cmdPrefix == "" {
		componentFlagName = "component"
	} else {
		componentFlagName = fmt.Sprintf("%v.component", cmdPrefix)
	}

	var componentFlagDefault string

	_ = cmd.PersistentFlags().String(componentFlagName, componentFlagDefault, componentDescription)

	return nil
}
func registerOperationUserGetUserUsernameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	usernameDescription := `Required. `

	var usernameFlagName string
	if cmdPrefix == "" {
		usernameFlagName = "username"
	} else {
		usernameFlagName = fmt.Sprintf("%v.username", cmdPrefix)
	}

	var usernameFlagDefault string

	_ = cmd.PersistentFlags().String(usernameFlagName, usernameFlagDefault, usernameDescription)

	return nil
}

func retrieveOperationUserGetUserComponentFlag(m *user.GetUserParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("component") {

		var componentFlagName string
		if cmdPrefix == "" {
			componentFlagName = "component"
		} else {
			componentFlagName = fmt.Sprintf("%v.component", cmdPrefix)
		}

		componentFlagValue, err := cmd.Flags().GetString(componentFlagName)
		if err != nil {
			return err, false
		}
		m.Component = &componentFlagValue

	}
	return nil, retAdded
}
func retrieveOperationUserGetUserUsernameFlag(m *user.GetUserParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("username") {

		var usernameFlagName string
		if cmdPrefix == "" {
			usernameFlagName = "username"
		} else {
			usernameFlagName = fmt.Sprintf("%v.username", cmdPrefix)
		}

		usernameFlagValue, err := cmd.Flags().GetString(usernameFlagName)
		if err != nil {
			return err, false
		}
		m.Username = usernameFlagValue

	}
	return nil, retAdded
}

// parseOperationUserGetUserResult parses request result and return the string content
func parseOperationUserGetUserResult(resp0 *user.GetUserOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*user.GetUserOK)
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
		msgStr := fmt.Sprintf("%v", resp0.Payload)
		return string(msgStr), nil
	}

	return "", nil
}
