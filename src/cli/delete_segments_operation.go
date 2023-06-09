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

// makeOperationSegmentDeleteSegmentsCmd returns a cmd to handle operation deleteSegments
func makeOperationSegmentDeleteSegmentsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteSegments",
		Short: `Delete the segments in the JSON array payload`,
		RunE:  runOperationSegmentDeleteSegments,
	}

	if err := registerOperationSegmentDeleteSegmentsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSegmentDeleteSegments uses cmd flags to call endpoint api
func runOperationSegmentDeleteSegments(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := segment.NewDeleteSegmentsParams()
	if err, _ := retrieveOperationSegmentDeleteSegmentsBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSegmentDeleteSegmentsRetentionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSegmentDeleteSegmentsTableNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSegmentDeleteSegmentsResult(appCli.Segment.DeleteSegments(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSegmentDeleteSegmentsParamFlags registers all flags needed to fill params
func registerOperationSegmentDeleteSegmentsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSegmentDeleteSegmentsBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSegmentDeleteSegmentsRetentionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSegmentDeleteSegmentsTableNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSegmentDeleteSegmentsBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	bodyDescription := ``

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	var bodyFlagDefault []string

	_ = cmd.PersistentFlags().StringSlice(bodyFlagName, bodyFlagDefault, bodyDescription)

	return nil
}
func registerOperationSegmentDeleteSegmentsRetentionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	retentionDescription := `Retention period for the table segments (e.g. 12h, 3d); If not set, the retention period will default to the first config that's not null: the table config, then to cluster setting, then '7d'. Using 0d or -1d will instantly delete segments without retention`

	var retentionFlagName string
	if cmdPrefix == "" {
		retentionFlagName = "retention"
	} else {
		retentionFlagName = fmt.Sprintf("%v.retention", cmdPrefix)
	}

	var retentionFlagDefault string

	_ = cmd.PersistentFlags().String(retentionFlagName, retentionFlagDefault, retentionDescription)

	return nil
}
func registerOperationSegmentDeleteSegmentsTableNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	tableNameDescription := `Required. Name of the table`

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

func retrieveOperationSegmentDeleteSegmentsBodyFlag(m *segment.DeleteSegmentsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {

		var bodyFlagName string
		if cmdPrefix == "" {
			bodyFlagName = "body"
		} else {
			bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
		}

		bodyFlagValues, err := cmd.Flags().GetStringSlice(bodyFlagName)
		if err != nil {
			return err, false
		}
		m.Body = bodyFlagValues

	}
	return nil, retAdded
}
func retrieveOperationSegmentDeleteSegmentsRetentionFlag(m *segment.DeleteSegmentsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("retention") {

		var retentionFlagName string
		if cmdPrefix == "" {
			retentionFlagName = "retention"
		} else {
			retentionFlagName = fmt.Sprintf("%v.retention", cmdPrefix)
		}

		retentionFlagValue, err := cmd.Flags().GetString(retentionFlagName)
		if err != nil {
			return err, false
		}
		m.Retention = &retentionFlagValue

	}
	return nil, retAdded
}
func retrieveOperationSegmentDeleteSegmentsTableNameFlag(m *segment.DeleteSegmentsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.TableName = tableNameFlagValue

	}
	return nil, retAdded
}

// parseOperationSegmentDeleteSegmentsResult parses request result and return the string content
func parseOperationSegmentDeleteSegmentsResult(resp0 *segment.DeleteSegmentsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*segment.DeleteSegmentsOK)
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
