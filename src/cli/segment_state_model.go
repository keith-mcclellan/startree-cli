// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"
	"startree.ai/cli/models"

	"github.com/spf13/cobra"
)

// Schema cli for SegmentState

// register flags to command
func registerModelSegmentStateFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSegmentStateConsumerInfo(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSegmentStateErrorInfo(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSegmentStateExternalView(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSegmentStateIdealState(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSegmentStateSegmentSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSegmentStateConsumerInfo(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var consumerInfoFlagName string
	if cmdPrefix == "" {
		consumerInfoFlagName = "consumerInfo"
	} else {
		consumerInfoFlagName = fmt.Sprintf("%v.consumerInfo", cmdPrefix)
	}

	if err := registerModelSegmentConsumerInfoFlags(depth+1, consumerInfoFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerSegmentStateErrorInfo(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var errorInfoFlagName string
	if cmdPrefix == "" {
		errorInfoFlagName = "errorInfo"
	} else {
		errorInfoFlagName = fmt.Sprintf("%v.errorInfo", cmdPrefix)
	}

	if err := registerModelSegmentErrorInfoFlags(depth+1, errorInfoFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerSegmentStateExternalView(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	externalViewDescription := ``

	var externalViewFlagName string
	if cmdPrefix == "" {
		externalViewFlagName = "externalView"
	} else {
		externalViewFlagName = fmt.Sprintf("%v.externalView", cmdPrefix)
	}

	var externalViewFlagDefault string

	_ = cmd.PersistentFlags().String(externalViewFlagName, externalViewFlagDefault, externalViewDescription)

	return nil
}

func registerSegmentStateIdealState(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idealStateDescription := ``

	var idealStateFlagName string
	if cmdPrefix == "" {
		idealStateFlagName = "idealState"
	} else {
		idealStateFlagName = fmt.Sprintf("%v.idealState", cmdPrefix)
	}

	var idealStateFlagDefault string

	_ = cmd.PersistentFlags().String(idealStateFlagName, idealStateFlagDefault, idealStateDescription)

	return nil
}

func registerSegmentStateSegmentSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	segmentSizeDescription := ``

	var segmentSizeFlagName string
	if cmdPrefix == "" {
		segmentSizeFlagName = "segmentSize"
	} else {
		segmentSizeFlagName = fmt.Sprintf("%v.segmentSize", cmdPrefix)
	}

	var segmentSizeFlagDefault string

	_ = cmd.PersistentFlags().String(segmentSizeFlagName, segmentSizeFlagDefault, segmentSizeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSegmentStateFlags(depth int, m *models.SegmentState, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, consumerInfoAdded := retrieveSegmentStateConsumerInfoFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || consumerInfoAdded

	err, errorInfoAdded := retrieveSegmentStateErrorInfoFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorInfoAdded

	err, externalViewAdded := retrieveSegmentStateExternalViewFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || externalViewAdded

	err, idealStateAdded := retrieveSegmentStateIdealStateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idealStateAdded

	err, segmentSizeAdded := retrieveSegmentStateSegmentSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || segmentSizeAdded

	return nil, retAdded
}

func retrieveSegmentStateConsumerInfoFlags(depth int, m *models.SegmentState, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	consumerInfoFlagName := fmt.Sprintf("%v.consumerInfo", cmdPrefix)
	if cmd.Flags().Changed(consumerInfoFlagName) {
		// info: complex object consumerInfo SegmentConsumerInfo is retrieved outside this Changed() block
	}
	consumerInfoFlagValue := m.ConsumerInfo
	if swag.IsZero(consumerInfoFlagValue) {
		consumerInfoFlagValue = &models.SegmentConsumerInfo{}
	}

	err, consumerInfoAdded := retrieveModelSegmentConsumerInfoFlags(depth+1, consumerInfoFlagValue, consumerInfoFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || consumerInfoAdded
	if consumerInfoAdded {
		m.ConsumerInfo = consumerInfoFlagValue
	}

	return nil, retAdded
}

func retrieveSegmentStateErrorInfoFlags(depth int, m *models.SegmentState, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	errorInfoFlagName := fmt.Sprintf("%v.errorInfo", cmdPrefix)
	if cmd.Flags().Changed(errorInfoFlagName) {
		// info: complex object errorInfo SegmentErrorInfo is retrieved outside this Changed() block
	}
	errorInfoFlagValue := m.ErrorInfo
	if swag.IsZero(errorInfoFlagValue) {
		errorInfoFlagValue = &models.SegmentErrorInfo{}
	}

	err, errorInfoAdded := retrieveModelSegmentErrorInfoFlags(depth+1, errorInfoFlagValue, errorInfoFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorInfoAdded
	if errorInfoAdded {
		m.ErrorInfo = errorInfoFlagValue
	}

	return nil, retAdded
}

func retrieveSegmentStateExternalViewFlags(depth int, m *models.SegmentState, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	externalViewFlagName := fmt.Sprintf("%v.externalView", cmdPrefix)
	if cmd.Flags().Changed(externalViewFlagName) {

		var externalViewFlagName string
		if cmdPrefix == "" {
			externalViewFlagName = "externalView"
		} else {
			externalViewFlagName = fmt.Sprintf("%v.externalView", cmdPrefix)
		}

		externalViewFlagValue, err := cmd.Flags().GetString(externalViewFlagName)
		if err != nil {
			return err, false
		}
		m.ExternalView = externalViewFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSegmentStateIdealStateFlags(depth int, m *models.SegmentState, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idealStateFlagName := fmt.Sprintf("%v.idealState", cmdPrefix)
	if cmd.Flags().Changed(idealStateFlagName) {

		var idealStateFlagName string
		if cmdPrefix == "" {
			idealStateFlagName = "idealState"
		} else {
			idealStateFlagName = fmt.Sprintf("%v.idealState", cmdPrefix)
		}

		idealStateFlagValue, err := cmd.Flags().GetString(idealStateFlagName)
		if err != nil {
			return err, false
		}
		m.IdealState = idealStateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSegmentStateSegmentSizeFlags(depth int, m *models.SegmentState, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	segmentSizeFlagName := fmt.Sprintf("%v.segmentSize", cmdPrefix)
	if cmd.Flags().Changed(segmentSizeFlagName) {

		var segmentSizeFlagName string
		if cmdPrefix == "" {
			segmentSizeFlagName = "segmentSize"
		} else {
			segmentSizeFlagName = fmt.Sprintf("%v.segmentSize", cmdPrefix)
		}

		segmentSizeFlagValue, err := cmd.Flags().GetString(segmentSizeFlagName)
		if err != nil {
			return err, false
		}
		m.SegmentSize = segmentSizeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
