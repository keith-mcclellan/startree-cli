// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"startree.ai/cli/client/segment"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSegmentGetReloadJobStatusCmd returns a cmd to handle operation getReloadJobStatus
func makeOperationSegmentGetReloadJobStatusCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getReloadJobStatus",
		Short: `Get status for a submitted reload operation`,
		RunE:  runOperationSegmentGetReloadJobStatus,
	}

	if err := registerOperationSegmentGetReloadJobStatusParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSegmentGetReloadJobStatus uses cmd flags to call endpoint api
func runOperationSegmentGetReloadJobStatus(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := segment.NewGetReloadJobStatusParams()
	if err, _ := retrieveOperationSegmentGetReloadJobStatusJobIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSegmentGetReloadJobStatusResult(appCli.Segment.GetReloadJobStatus(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSegmentGetReloadJobStatusParamFlags registers all flags needed to fill params
func registerOperationSegmentGetReloadJobStatusParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSegmentGetReloadJobStatusJobIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSegmentGetReloadJobStatusJobIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	jobIdDescription := `Required. Reload job id`

	var jobIdFlagName string
	if cmdPrefix == "" {
		jobIdFlagName = "jobId"
	} else {
		jobIdFlagName = fmt.Sprintf("%v.jobId", cmdPrefix)
	}

	var jobIdFlagDefault string

	_ = cmd.PersistentFlags().String(jobIdFlagName, jobIdFlagDefault, jobIdDescription)

	return nil
}

func retrieveOperationSegmentGetReloadJobStatusJobIDFlag(m *segment.GetReloadJobStatusParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("jobId") {

		var jobIdFlagName string
		if cmdPrefix == "" {
			jobIdFlagName = "jobId"
		} else {
			jobIdFlagName = fmt.Sprintf("%v.jobId", cmdPrefix)
		}

		jobIdFlagValue, err := cmd.Flags().GetString(jobIdFlagName)
		if err != nil {
			return err, false
		}
		m.JobID = jobIdFlagValue

	}
	return nil, retAdded
}

// parseOperationSegmentGetReloadJobStatusResult parses request result and return the string content
func parseOperationSegmentGetReloadJobStatusResult(resp0 *segment.GetReloadJobStatusOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*segment.GetReloadJobStatusOK)
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
