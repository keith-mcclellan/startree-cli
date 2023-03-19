// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/spf13/cobra"
	"startree.ai/cli/models"
)

// Schema cli for TransformConfig

// register flags to command
func registerModelTransformConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTransformConfigColumnName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTransformConfigTransformFunction(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTransformConfigColumnName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	columnNameDescription := ``

	var columnNameFlagName string
	if cmdPrefix == "" {
		columnNameFlagName = "columnName"
	} else {
		columnNameFlagName = fmt.Sprintf("%v.columnName", cmdPrefix)
	}

	var columnNameFlagDefault string

	_ = cmd.PersistentFlags().String(columnNameFlagName, columnNameFlagDefault, columnNameDescription)

	return nil
}

func registerTransformConfigTransformFunction(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	transformFunctionDescription := ``

	var transformFunctionFlagName string
	if cmdPrefix == "" {
		transformFunctionFlagName = "transformFunction"
	} else {
		transformFunctionFlagName = fmt.Sprintf("%v.transformFunction", cmdPrefix)
	}

	var transformFunctionFlagDefault string

	_ = cmd.PersistentFlags().String(transformFunctionFlagName, transformFunctionFlagDefault, transformFunctionDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTransformConfigFlags(depth int, m *models.TransformConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, columnNameAdded := retrieveTransformConfigColumnNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || columnNameAdded

	err, transformFunctionAdded := retrieveTransformConfigTransformFunctionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || transformFunctionAdded

	return nil, retAdded
}

func retrieveTransformConfigColumnNameFlags(depth int, m *models.TransformConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	columnNameFlagName := fmt.Sprintf("%v.columnName", cmdPrefix)
	if cmd.Flags().Changed(columnNameFlagName) {

		var columnNameFlagName string
		if cmdPrefix == "" {
			columnNameFlagName = "columnName"
		} else {
			columnNameFlagName = fmt.Sprintf("%v.columnName", cmdPrefix)
		}

		columnNameFlagValue, err := cmd.Flags().GetString(columnNameFlagName)
		if err != nil {
			return err, false
		}
		m.ColumnName = columnNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTransformConfigTransformFunctionFlags(depth int, m *models.TransformConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	transformFunctionFlagName := fmt.Sprintf("%v.transformFunction", cmdPrefix)
	if cmd.Flags().Changed(transformFunctionFlagName) {

		var transformFunctionFlagName string
		if cmdPrefix == "" {
			transformFunctionFlagName = "transformFunction"
		} else {
			transformFunctionFlagName = fmt.Sprintf("%v.transformFunction", cmdPrefix)
		}

		transformFunctionFlagValue, err := cmd.Flags().GetString(transformFunctionFlagName)
		if err != nil {
			return err, false
		}
		m.TransformFunction = transformFunctionFlagValue

		retAdded = true
	}

	return nil, retAdded
}