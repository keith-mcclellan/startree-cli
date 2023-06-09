// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/spf13/cobra"
	"startree.ai/cli/models"
)

// Schema cli for SegmentDebugInfo

// register flags to command
func registerModelSegmentDebugInfoFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSegmentDebugInfoSegmentName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSegmentDebugInfoServerState(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSegmentDebugInfoSegmentName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	segmentNameDescription := ``

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

func registerSegmentDebugInfoServerState(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: serverState map[string]SegmentState map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSegmentDebugInfoFlags(depth int, m *models.SegmentDebugInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, segmentNameAdded := retrieveSegmentDebugInfoSegmentNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || segmentNameAdded

	err, serverStateAdded := retrieveSegmentDebugInfoServerStateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || serverStateAdded

	return nil, retAdded
}

func retrieveSegmentDebugInfoSegmentNameFlags(depth int, m *models.SegmentDebugInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	segmentNameFlagName := fmt.Sprintf("%v.segmentName", cmdPrefix)
	if cmd.Flags().Changed(segmentNameFlagName) {

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

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSegmentDebugInfoServerStateFlags(depth int, m *models.SegmentDebugInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	serverStateFlagName := fmt.Sprintf("%v.serverState", cmdPrefix)
	if cmd.Flags().Changed(serverStateFlagName) {
		// warning: serverState map type map[string]SegmentState is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
