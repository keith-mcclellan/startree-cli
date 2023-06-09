// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/spf13/cobra"
	"startree.ai/cli/models"
)

// Schema cli for QueryConfig

// register flags to command
func registerModelQueryConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerQueryConfigDisableGroovy(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerQueryConfigExpressionOverrideMap(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerQueryConfigTimeoutMs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerQueryConfigUseApproximateFunction(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerQueryConfigDisableGroovy(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	disableGroovyDescription := ``

	var disableGroovyFlagName string
	if cmdPrefix == "" {
		disableGroovyFlagName = "disableGroovy"
	} else {
		disableGroovyFlagName = fmt.Sprintf("%v.disableGroovy", cmdPrefix)
	}

	var disableGroovyFlagDefault bool

	_ = cmd.PersistentFlags().Bool(disableGroovyFlagName, disableGroovyFlagDefault, disableGroovyDescription)

	return nil
}

func registerQueryConfigExpressionOverrideMap(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: expressionOverrideMap map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerQueryConfigTimeoutMs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	timeoutMsDescription := ``

	var timeoutMsFlagName string
	if cmdPrefix == "" {
		timeoutMsFlagName = "timeoutMs"
	} else {
		timeoutMsFlagName = fmt.Sprintf("%v.timeoutMs", cmdPrefix)
	}

	var timeoutMsFlagDefault int64

	_ = cmd.PersistentFlags().Int64(timeoutMsFlagName, timeoutMsFlagDefault, timeoutMsDescription)

	return nil
}

func registerQueryConfigUseApproximateFunction(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	useApproximateFunctionDescription := ``

	var useApproximateFunctionFlagName string
	if cmdPrefix == "" {
		useApproximateFunctionFlagName = "useApproximateFunction"
	} else {
		useApproximateFunctionFlagName = fmt.Sprintf("%v.useApproximateFunction", cmdPrefix)
	}

	var useApproximateFunctionFlagDefault bool

	_ = cmd.PersistentFlags().Bool(useApproximateFunctionFlagName, useApproximateFunctionFlagDefault, useApproximateFunctionDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelQueryConfigFlags(depth int, m *models.QueryConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, disableGroovyAdded := retrieveQueryConfigDisableGroovyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || disableGroovyAdded

	err, expressionOverrideMapAdded := retrieveQueryConfigExpressionOverrideMapFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || expressionOverrideMapAdded

	err, timeoutMsAdded := retrieveQueryConfigTimeoutMsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || timeoutMsAdded

	err, useApproximateFunctionAdded := retrieveQueryConfigUseApproximateFunctionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || useApproximateFunctionAdded

	return nil, retAdded
}

func retrieveQueryConfigDisableGroovyFlags(depth int, m *models.QueryConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	disableGroovyFlagName := fmt.Sprintf("%v.disableGroovy", cmdPrefix)
	if cmd.Flags().Changed(disableGroovyFlagName) {

		var disableGroovyFlagName string
		if cmdPrefix == "" {
			disableGroovyFlagName = "disableGroovy"
		} else {
			disableGroovyFlagName = fmt.Sprintf("%v.disableGroovy", cmdPrefix)
		}

		disableGroovyFlagValue, err := cmd.Flags().GetBool(disableGroovyFlagName)
		if err != nil {
			return err, false
		}
		m.DisableGroovy = &disableGroovyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveQueryConfigExpressionOverrideMapFlags(depth int, m *models.QueryConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	expressionOverrideMapFlagName := fmt.Sprintf("%v.expressionOverrideMap", cmdPrefix)
	if cmd.Flags().Changed(expressionOverrideMapFlagName) {
		// warning: expressionOverrideMap map type map[string]string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveQueryConfigTimeoutMsFlags(depth int, m *models.QueryConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	timeoutMsFlagName := fmt.Sprintf("%v.timeoutMs", cmdPrefix)
	if cmd.Flags().Changed(timeoutMsFlagName) {

		var timeoutMsFlagName string
		if cmdPrefix == "" {
			timeoutMsFlagName = "timeoutMs"
		} else {
			timeoutMsFlagName = fmt.Sprintf("%v.timeoutMs", cmdPrefix)
		}

		timeoutMsFlagValue, err := cmd.Flags().GetInt64(timeoutMsFlagName)
		if err != nil {
			return err, false
		}
		m.TimeoutMs = timeoutMsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveQueryConfigUseApproximateFunctionFlags(depth int, m *models.QueryConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	useApproximateFunctionFlagName := fmt.Sprintf("%v.useApproximateFunction", cmdPrefix)
	if cmd.Flags().Changed(useApproximateFunctionFlagName) {

		var useApproximateFunctionFlagName string
		if cmdPrefix == "" {
			useApproximateFunctionFlagName = "useApproximateFunction"
		} else {
			useApproximateFunctionFlagName = fmt.Sprintf("%v.useApproximateFunction", cmdPrefix)
		}

		useApproximateFunctionFlagValue, err := cmd.Flags().GetBool(useApproximateFunctionFlagName)
		if err != nil {
			return err, false
		}
		m.UseApproximateFunction = &useApproximateFunctionFlagValue

		retAdded = true
	}

	return nil, retAdded
}
