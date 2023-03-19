// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"startree.ai/cli/client/zookeeper"

	"github.com/spf13/cobra"
)

// makeOperationZookeeperGetDataCmd returns a cmd to handle operation getData
func makeOperationZookeeperGetDataCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getData",
		Short: ``,
		RunE:  runOperationZookeeperGetData,
	}

	if err := registerOperationZookeeperGetDataParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationZookeeperGetData uses cmd flags to call endpoint api
func runOperationZookeeperGetData(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := zookeeper.NewGetDataParams()
	if err, _ := retrieveOperationZookeeperGetDataPathFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationZookeeperGetDataResult(appCli.Zookeeper.GetData(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationZookeeperGetDataParamFlags registers all flags needed to fill params
func registerOperationZookeeperGetDataParamFlags(cmd *cobra.Command) error {
	if err := registerOperationZookeeperGetDataPathParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationZookeeperGetDataPathParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationZookeeperGetDataPathFlag(m *zookeeper.GetDataParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationZookeeperGetDataResult parses request result and return the string content
func parseOperationZookeeperGetDataResult(resp0 *zookeeper.GetDataOK, resp1 *zookeeper.GetDataNoContent, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning getDataOK is not supported

		// Non schema case: warning getDataNoContent is not supported

		// Non schema case: warning getDataNotFound is not supported

		// Non schema case: warning getDataInternalServerError is not supported

		return "", respErr
	}

	// warning: non schema response getDataOK is not supported by go-swagger cli yet.

	// warning: non schema response getDataNoContent is not supported by go-swagger cli yet.

	return "", nil
}
