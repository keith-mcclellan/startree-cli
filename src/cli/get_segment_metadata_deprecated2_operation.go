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

// makeOperationSegmentGetSegmentMetadataDeprecated2Cmd returns a cmd to handle operation getSegmentMetadataDeprecated2
func makeOperationSegmentGetSegmentMetadataDeprecated2Cmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getSegmentMetadataDeprecated2",
		Short: `Get the metadata for a segment (deprecated, use 'GET /segments/{tableName}/{segmentName}/metadata' instead)`,
		RunE:  runOperationSegmentGetSegmentMetadataDeprecated2,
	}

	if err := registerOperationSegmentGetSegmentMetadataDeprecated2ParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSegmentGetSegmentMetadataDeprecated2 uses cmd flags to call endpoint api
func runOperationSegmentGetSegmentMetadataDeprecated2(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := segment.NewGetSegmentMetadataDeprecated2Params()
	if err, _ := retrieveOperationSegmentGetSegmentMetadataDeprecated2SegmentNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSegmentGetSegmentMetadataDeprecated2StateFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSegmentGetSegmentMetadataDeprecated2TableNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSegmentGetSegmentMetadataDeprecated2TypeFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSegmentGetSegmentMetadataDeprecated2Result(appCli.Segment.GetSegmentMetadataDeprecated2(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSegmentGetSegmentMetadataDeprecated2ParamFlags registers all flags needed to fill params
func registerOperationSegmentGetSegmentMetadataDeprecated2ParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSegmentGetSegmentMetadataDeprecated2SegmentNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSegmentGetSegmentMetadataDeprecated2StateParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSegmentGetSegmentMetadataDeprecated2TableNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSegmentGetSegmentMetadataDeprecated2TypeParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSegmentGetSegmentMetadataDeprecated2SegmentNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSegmentGetSegmentMetadataDeprecated2StateParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	stateDescription := `MUST be null`

	var stateFlagName string
	if cmdPrefix == "" {
		stateFlagName = "state"
	} else {
		stateFlagName = fmt.Sprintf("%v.state", cmdPrefix)
	}

	var stateFlagDefault string

	_ = cmd.PersistentFlags().String(stateFlagName, stateFlagDefault, stateDescription)

	return nil
}
func registerOperationSegmentGetSegmentMetadataDeprecated2TableNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSegmentGetSegmentMetadataDeprecated2TypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	typeDescription := `OFFLINE|REALTIME`

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	return nil
}

func retrieveOperationSegmentGetSegmentMetadataDeprecated2SegmentNameFlag(m *segment.GetSegmentMetadataDeprecated2Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSegmentGetSegmentMetadataDeprecated2StateFlag(m *segment.GetSegmentMetadataDeprecated2Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("state") {

		var stateFlagName string
		if cmdPrefix == "" {
			stateFlagName = "state"
		} else {
			stateFlagName = fmt.Sprintf("%v.state", cmdPrefix)
		}

		stateFlagValue, err := cmd.Flags().GetString(stateFlagName)
		if err != nil {
			return err, false
		}
		m.State = &stateFlagValue

	}
	return nil, retAdded
}
func retrieveOperationSegmentGetSegmentMetadataDeprecated2TableNameFlag(m *segment.GetSegmentMetadataDeprecated2Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSegmentGetSegmentMetadataDeprecated2TypeFlag(m *segment.GetSegmentMetadataDeprecated2Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("type") {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = &typeFlagValue

	}
	return nil, retAdded
}

// parseOperationSegmentGetSegmentMetadataDeprecated2Result parses request result and return the string content
func parseOperationSegmentGetSegmentMetadataDeprecated2Result(resp0 *segment.GetSegmentMetadataDeprecated2OK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*segment.GetSegmentMetadataDeprecated2OK)
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
