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

// Schema cli for Schema

// register flags to command
func registerModelSchemaFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSchemaDateTimeFieldSpecs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSchemaDimensionFieldSpecs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSchemaMetricFieldSpecs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSchemaPrimaryKeyColumns(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSchemaSchemaName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSchemaTimeFieldSpec(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSchemaDateTimeFieldSpecs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: dateTimeFieldSpecs []*DateTimeFieldSpec array type is not supported by go-swagger cli yet

	return nil
}

func registerSchemaDimensionFieldSpecs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: dimensionFieldSpecs []*DimensionFieldSpec array type is not supported by go-swagger cli yet

	return nil
}

func registerSchemaMetricFieldSpecs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: metricFieldSpecs []*MetricFieldSpec array type is not supported by go-swagger cli yet

	return nil
}

func registerSchemaPrimaryKeyColumns(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: primaryKeyColumns []string array type is not supported by go-swagger cli yet

	return nil
}

func registerSchemaSchemaName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	schemaNameDescription := ``

	var schemaNameFlagName string
	if cmdPrefix == "" {
		schemaNameFlagName = "schemaName"
	} else {
		schemaNameFlagName = fmt.Sprintf("%v.schemaName", cmdPrefix)
	}

	var schemaNameFlagDefault string

	_ = cmd.PersistentFlags().String(schemaNameFlagName, schemaNameFlagDefault, schemaNameDescription)

	return nil
}

func registerSchemaTimeFieldSpec(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var timeFieldSpecFlagName string
	if cmdPrefix == "" {
		timeFieldSpecFlagName = "timeFieldSpec"
	} else {
		timeFieldSpecFlagName = fmt.Sprintf("%v.timeFieldSpec", cmdPrefix)
	}

	if err := registerModelTimeFieldSpecFlags(depth+1, timeFieldSpecFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSchemaFlags(depth int, m *models.Schema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dateTimeFieldSpecsAdded := retrieveSchemaDateTimeFieldSpecsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dateTimeFieldSpecsAdded

	err, dimensionFieldSpecsAdded := retrieveSchemaDimensionFieldSpecsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dimensionFieldSpecsAdded

	err, metricFieldSpecsAdded := retrieveSchemaMetricFieldSpecsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || metricFieldSpecsAdded

	err, primaryKeyColumnsAdded := retrieveSchemaPrimaryKeyColumnsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || primaryKeyColumnsAdded

	err, schemaNameAdded := retrieveSchemaSchemaNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || schemaNameAdded

	err, timeFieldSpecAdded := retrieveSchemaTimeFieldSpecFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || timeFieldSpecAdded

	return nil, retAdded
}

func retrieveSchemaDateTimeFieldSpecsFlags(depth int, m *models.Schema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dateTimeFieldSpecsFlagName := fmt.Sprintf("%v.dateTimeFieldSpecs", cmdPrefix)
	if cmd.Flags().Changed(dateTimeFieldSpecsFlagName) {
		// warning: dateTimeFieldSpecs array type []*DateTimeFieldSpec is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveSchemaDimensionFieldSpecsFlags(depth int, m *models.Schema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dimensionFieldSpecsFlagName := fmt.Sprintf("%v.dimensionFieldSpecs", cmdPrefix)
	if cmd.Flags().Changed(dimensionFieldSpecsFlagName) {
		// warning: dimensionFieldSpecs array type []*DimensionFieldSpec is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveSchemaMetricFieldSpecsFlags(depth int, m *models.Schema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	metricFieldSpecsFlagName := fmt.Sprintf("%v.metricFieldSpecs", cmdPrefix)
	if cmd.Flags().Changed(metricFieldSpecsFlagName) {
		// warning: metricFieldSpecs array type []*MetricFieldSpec is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveSchemaPrimaryKeyColumnsFlags(depth int, m *models.Schema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	primaryKeyColumnsFlagName := fmt.Sprintf("%v.primaryKeyColumns", cmdPrefix)
	if cmd.Flags().Changed(primaryKeyColumnsFlagName) {
		// warning: primaryKeyColumns array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveSchemaSchemaNameFlags(depth int, m *models.Schema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	schemaNameFlagName := fmt.Sprintf("%v.schemaName", cmdPrefix)
	if cmd.Flags().Changed(schemaNameFlagName) {

		var schemaNameFlagName string
		if cmdPrefix == "" {
			schemaNameFlagName = "schemaName"
		} else {
			schemaNameFlagName = fmt.Sprintf("%v.schemaName", cmdPrefix)
		}

		schemaNameFlagValue, err := cmd.Flags().GetString(schemaNameFlagName)
		if err != nil {
			return err, false
		}
		m.SchemaName = schemaNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSchemaTimeFieldSpecFlags(depth int, m *models.Schema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	timeFieldSpecFlagName := fmt.Sprintf("%v.timeFieldSpec", cmdPrefix)
	if cmd.Flags().Changed(timeFieldSpecFlagName) {
		// info: complex object timeFieldSpec TimeFieldSpec is retrieved outside this Changed() block
	}
	timeFieldSpecFlagValue := m.TimeFieldSpec
	if swag.IsZero(timeFieldSpecFlagValue) {
		timeFieldSpecFlagValue = &models.TimeFieldSpec{}
	}

	err, timeFieldSpecAdded := retrieveModelTimeFieldSpecFlags(depth+1, timeFieldSpecFlagValue, timeFieldSpecFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || timeFieldSpecAdded
	if timeFieldSpecAdded {
		m.TimeFieldSpec = timeFieldSpecFlagValue
	}

	return nil, retAdded
}