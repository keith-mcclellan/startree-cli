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

// makeOperationSegmentGetServerToSegmentsMapDeprecated1Cmd returns a cmd to handle operation getServerToSegmentsMapDeprecated1
func makeOperationSegmentGetServerToSegmentsMapDeprecated1Cmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getServerToSegmentsMapDeprecated1",
		Short: `Get a map from server to segments hosted by the server (deprecated, use 'GET /segments/{tableName}/servers' instead)`,
		RunE:  runOperationSegmentGetServerToSegmentsMapDeprecated1,
	}

	if err := registerOperationSegmentGetServerToSegmentsMapDeprecated1ParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSegmentGetServerToSegmentsMapDeprecated1 uses cmd flags to call endpoint api
func runOperationSegmentGetServerToSegmentsMapDeprecated1(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := segment.NewGetServerToSegmentsMapDeprecated1Params()
	if err, _ := retrieveOperationSegmentGetServerToSegmentsMapDeprecated1StateFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSegmentGetServerToSegmentsMapDeprecated1TableNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSegmentGetServerToSegmentsMapDeprecated1TypeFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSegmentGetServerToSegmentsMapDeprecated1Result(appCli.Segment.GetServerToSegmentsMapDeprecated1(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSegmentGetServerToSegmentsMapDeprecated1ParamFlags registers all flags needed to fill params
func registerOperationSegmentGetServerToSegmentsMapDeprecated1ParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSegmentGetServerToSegmentsMapDeprecated1StateParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSegmentGetServerToSegmentsMapDeprecated1TableNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSegmentGetServerToSegmentsMapDeprecated1TypeParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSegmentGetServerToSegmentsMapDeprecated1StateParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSegmentGetServerToSegmentsMapDeprecated1TableNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSegmentGetServerToSegmentsMapDeprecated1TypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationSegmentGetServerToSegmentsMapDeprecated1StateFlag(m *segment.GetServerToSegmentsMapDeprecated1Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSegmentGetServerToSegmentsMapDeprecated1TableNameFlag(m *segment.GetServerToSegmentsMapDeprecated1Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSegmentGetServerToSegmentsMapDeprecated1TypeFlag(m *segment.GetServerToSegmentsMapDeprecated1Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationSegmentGetServerToSegmentsMapDeprecated1Result parses request result and return the string content
func parseOperationSegmentGetServerToSegmentsMapDeprecated1Result(resp0 *segment.GetServerToSegmentsMapDeprecated1OK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*segment.GetServerToSegmentsMapDeprecated1OK)
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
