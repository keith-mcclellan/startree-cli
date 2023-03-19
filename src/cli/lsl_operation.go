// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"startree.ai/cli/client/zookeeper"

	"github.com/spf13/cobra"
)

// makeOperationZookeeperLslCmd returns a cmd to handle operation lsl
func makeOperationZookeeperLslCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "lsl",
		Short: ``,
		RunE:  runOperationZookeeperLsl,
	}

	if err := registerOperationZookeeperLslParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationZookeeperLsl uses cmd flags to call endpoint api
func runOperationZookeeperLsl(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := zookeeper.NewLslParams()
	if err, _ := retrieveOperationZookeeperLslPathFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationZookeeperLslResult(appCli.Zookeeper.Lsl(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationZookeeperLslParamFlags registers all flags needed to fill params
func registerOperationZookeeperLslParamFlags(cmd *cobra.Command) error {
	if err := registerOperationZookeeperLslPathParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationZookeeperLslPathParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationZookeeperLslPathFlag(m *zookeeper.LslParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationZookeeperLslResult parses request result and return the string content
func parseOperationZookeeperLslResult(resp0 *zookeeper.LslOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning lslOK is not supported

		// Non schema case: warning lslNotFound is not supported

		// Non schema case: warning lslInternalServerError is not supported

		return "", respErr
	}

	// warning: non schema response lslOK is not supported by go-swagger cli yet.

	return "", nil
}
