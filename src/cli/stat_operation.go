// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"startree.ai/cli/client/zookeeper"

	"github.com/spf13/cobra"
)

// makeOperationZookeeperStatCmd returns a cmd to handle operation stat
func makeOperationZookeeperStatCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "stat",
		Short: ` Use this api to fetch additional details of a znode such as creation time, modified time, numChildren etc `,
		RunE:  runOperationZookeeperStat,
	}

	if err := registerOperationZookeeperStatParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationZookeeperStat uses cmd flags to call endpoint api
func runOperationZookeeperStat(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := zookeeper.NewStatParams()
	if err, _ := retrieveOperationZookeeperStatPathFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationZookeeperStatResult(appCli.Zookeeper.Stat(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationZookeeperStatParamFlags registers all flags needed to fill params
func registerOperationZookeeperStatParamFlags(cmd *cobra.Command) error {
	if err := registerOperationZookeeperStatPathParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationZookeeperStatPathParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationZookeeperStatPathFlag(m *zookeeper.StatParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationZookeeperStatResult parses request result and return the string content
func parseOperationZookeeperStatResult(resp0 *zookeeper.StatOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning statOK is not supported

		// Non schema case: warning statNotFound is not supported

		// Non schema case: warning statInternalServerError is not supported

		return "", respErr
	}

	// warning: non schema response statOK is not supported by go-swagger cli yet.

	return "", nil
}
