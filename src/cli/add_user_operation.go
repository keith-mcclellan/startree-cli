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

// makeOperationUserAddUserCmd returns a cmd to handle operation addUser
func makeOperationUserAddUserCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "addUser",
		Short: `Add a user`,
		RunE:  runOperationUserAddUser,
	}

	if err := registerOperationUserAddUserParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationUserAddUser uses cmd flags to call endpoint api
func runOperationUserAddUser(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := user.NewAddUserParams()
	if err, _ := retrieveOperationUserAddUserBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationUserAddUserResult(appCli.User.AddUser(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationUserAddUserParamFlags registers all flags needed to fill params
func registerOperationUserAddUserParamFlags(cmd *cobra.Command) error {
	if err := registerOperationUserAddUserBodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationUserAddUserBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	bodyDescription := ``

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	var bodyFlagDefault string

	_ = cmd.PersistentFlags().String(bodyFlagName, bodyFlagDefault, bodyDescription)

	return nil
}

func retrieveOperationUserAddUserBodyFlag(m *user.AddUserParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {

		var bodyFlagName string
		if cmdPrefix == "" {
			bodyFlagName = "body"
		} else {
			bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
		}

		bodyFlagValue, err := cmd.Flags().GetString(bodyFlagName)
		if err != nil {
			return err, false
		}
		m.Body = bodyFlagValue

	}
	return nil, retAdded
}

// parseOperationUserAddUserResult parses request result and return the string content
func parseOperationUserAddUserResult(resp0 *user.AddUserOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*user.AddUserOK)
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
