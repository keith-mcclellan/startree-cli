// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"startree.ai/cli/client/logger"

	"github.com/spf13/cobra"
)

// makeOperationLoggerDownloadLogFileFromInstanceCmd returns a cmd to handle operation downloadLogFileFromInstance
func makeOperationLoggerDownloadLogFileFromInstanceCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "downloadLogFileFromInstance",
		Short: ``,
		RunE:  runOperationLoggerDownloadLogFileFromInstance,
	}

	if err := registerOperationLoggerDownloadLogFileFromInstanceParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationLoggerDownloadLogFileFromInstance uses cmd flags to call endpoint api
func runOperationLoggerDownloadLogFileFromInstance(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := logger.NewDownloadLogFileFromInstanceParams()
	if err, _ := retrieveOperationLoggerDownloadLogFileFromInstanceFilePathFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationLoggerDownloadLogFileFromInstanceInstanceNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationLoggerDownloadLogFileFromInstanceResult(appCli.Logger.DownloadLogFileFromInstance(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationLoggerDownloadLogFileFromInstanceParamFlags registers all flags needed to fill params
func registerOperationLoggerDownloadLogFileFromInstanceParamFlags(cmd *cobra.Command) error {
	if err := registerOperationLoggerDownloadLogFileFromInstanceFilePathParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationLoggerDownloadLogFileFromInstanceInstanceNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationLoggerDownloadLogFileFromInstanceFilePathParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filePathDescription := `Required. Log file path`

	var filePathFlagName string
	if cmdPrefix == "" {
		filePathFlagName = "filePath"
	} else {
		filePathFlagName = fmt.Sprintf("%v.filePath", cmdPrefix)
	}

	var filePathFlagDefault string

	_ = cmd.PersistentFlags().String(filePathFlagName, filePathFlagDefault, filePathDescription)

	return nil
}
func registerOperationLoggerDownloadLogFileFromInstanceInstanceNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	instanceNameDescription := `Required. Instance Name`

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

func retrieveOperationLoggerDownloadLogFileFromInstanceFilePathFlag(m *logger.DownloadLogFileFromInstanceParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filePath") {

		var filePathFlagName string
		if cmdPrefix == "" {
			filePathFlagName = "filePath"
		} else {
			filePathFlagName = fmt.Sprintf("%v.filePath", cmdPrefix)
		}

		filePathFlagValue, err := cmd.Flags().GetString(filePathFlagName)
		if err != nil {
			return err, false
		}
		m.FilePath = filePathFlagValue

	}
	return nil, retAdded
}
func retrieveOperationLoggerDownloadLogFileFromInstanceInstanceNameFlag(m *logger.DownloadLogFileFromInstanceParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationLoggerDownloadLogFileFromInstanceResult parses request result and return the string content
func parseOperationLoggerDownloadLogFileFromInstanceResult(respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning downloadLogFileFromInstance default is not supported

		return "", respErr
	}
	return "", nil
}
