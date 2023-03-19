// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"startree.ai/cli/client/zookeeper"

	"github.com/spf13/cobra"
)

// makeOperationZookeeperLsCmd returns a cmd to handle operation ls
func makeOperationZookeeperLsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: ``,
		RunE:  runOperationZookeeperLs,
	}

	if err := registerOperationZookeeperLsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationZookeeperLs uses cmd flags to call endpoint api
func runOperationZookeeperLs(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := zookeeper.NewLsParams()
	if err, _ := retrieveOperationZookeeperLsPathFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationZookeeperLsResult(appCli.Zookeeper.Ls(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationZookeeperLsParamFlags registers all flags needed to fill params
func registerOperationZookeeperLsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationZookeeperLsPathParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationZookeeperLsPathParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	pathDescription := `Required. Zookeeper Path, must start with /`

	var pathFlagName string
	if cmdPrefix == "" {
		pathFlagName = "path"
	} else {
		pathFlagName = fmt.Sprintf("%v.path", cmdPrefix)
	}

	var pathFlagDefault string

	_ = cmd.PersistentFlags().String(pathFlagName, pathFlagDefault, pathDescription)

	return nil
}

func retrieveOperationZookeeperLsPathFlag(m *zookeeper.LsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("path") {

		var pathFlagName string
		if cmdPrefix == "" {
			pathFlagName = "path"
		} else {
			pathFlagName = fmt.Sprintf("%v.path", cmdPrefix)
		}

		pathFlagValue, err := cmd.Flags().GetString(pathFlagName)
		if err != nil {
			return err, false
		}
		m.Path = pathFlagValue

	}
	return nil, retAdded
}

// parseOperationZookeeperLsResult parses request result and return the string content
func parseOperationZookeeperLsResult(resp0 *zookeeper.LsOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning lsOK is not supported

		// Non schema case: warning lsNotFound is not supported

		// Non schema case: warning lsInternalServerError is not supported

		return "", respErr
	}

	// warning: non schema response lsOK is not supported by go-swagger cli yet.

	return "", nil
}
