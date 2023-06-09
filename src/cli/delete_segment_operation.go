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

// makeOperationSegmentDeleteSegmentCmd returns a cmd to handle operation deleteSegment
func makeOperationSegmentDeleteSegmentCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteSegment",
		Short: `Delete a segment`,
		RunE:  runOperationSegmentDeleteSegment,
	}

	if err := registerOperationSegmentDeleteSegmentParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSegmentDeleteSegment uses cmd flags to call endpoint api
func runOperationSegmentDeleteSegment(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := segment.NewDeleteSegmentParams()
	if err, _ := retrieveOperationSegmentDeleteSegmentRetentionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSegmentDeleteSegmentSegmentNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSegmentDeleteSegmentTableNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSegmentDeleteSegmentResult(appCli.Segment.DeleteSegment(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSegmentDeleteSegmentParamFlags registers all flags needed to fill params
func registerOperationSegmentDeleteSegmentParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSegmentDeleteSegmentRetentionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSegmentDeleteSegmentSegmentNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSegmentDeleteSegmentTableNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSegmentDeleteSegmentRetentionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSegmentDeleteSegmentSegmentNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	segmentNameDescription := `Required. Name of the segment`

	var segmentNameFlagName string
	if cmdPrefix == "" {
		segmentNameFlagName = "segmentName"
	} else {
		segmentNameFlagName = fmt.Sprintf("%v.segmentName", cmdPrefix)
	}

	var segmentNameFlagDefault string

	_ = cmd.PersistentFlags().String(segmentNameFlagName, segmentNameFlagDefault, segmentNameDescription)

	return nil
}
func registerOperationSegmentDeleteSegmentTableNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationSegmentDeleteSegmentRetentionFlag(m *segment.DeleteSegmentParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSegmentDeleteSegmentSegmentNameFlag(m *segment.DeleteSegmentParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("segmentName") {

		var segmentNameFlagName string
		if cmdPrefix == "" {
			segmentNameFlagName = "segmentName"
		} else {
			segmentNameFlagName = fmt.Sprintf("%v.segmentName", cmdPrefix)
		}

		segmentNameFlagValue, err := cmd.Flags().GetString(segmentNameFlagName)
		if err != nil {
			return err, false
		}
		m.SegmentName = segmentNameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationSegmentDeleteSegmentTableNameFlag(m *segment.DeleteSegmentParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationSegmentDeleteSegmentResult parses request result and return the string content
func parseOperationSegmentDeleteSegmentResult(resp0 *segment.DeleteSegmentOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*segment.DeleteSegmentOK)
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
