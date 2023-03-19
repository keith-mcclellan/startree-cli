// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"startree.ai/cli/models"
)

// Schema cli for MetricFieldSpec

// register flags to command
func registerModelMetricFieldSpecFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerMetricFieldSpecDataType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMetricFieldSpecDefaultNullValue(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMetricFieldSpecDefaultNullValueString(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMetricFieldSpecMaxLength(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMetricFieldSpecName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMetricFieldSpecSingleValueField(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMetricFieldSpecTransformFunction(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMetricFieldSpecVirtualColumnProvider(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerMetricFieldSpecDataType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	dataTypeDescription := `Enum: ["INT","LONG","FLOAT","DOUBLE","BIG_DECIMAL","BOOLEAN","TIMESTAMP","STRING","JSON","BYTES","STRUCT","MAP","LIST"]. `

	var dataTypeFlagName string
	if cmdPrefix == "" {
		dataTypeFlagName = "dataType"
	} else {
		dataTypeFlagName = fmt.Sprintf("%v.dataType", cmdPrefix)
	}

	var dataTypeFlagDefault string

	_ = cmd.PersistentFlags().String(dataTypeFlagName, dataTypeFlagDefault, dataTypeDescription)

	if err := cmd.RegisterFlagCompletionFunc(dataTypeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["INT","LONG","FLOAT","DOUBLE","BIG_DECIMAL","BOOLEAN","TIMESTAMP","STRING","JSON","BYTES","STRUCT","MAP","LIST"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerMetricFieldSpecDefaultNullValue(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: defaultNullValue interface{} map type is not supported by go-swagger cli yet

	return nil
}

func registerMetricFieldSpecDefaultNullValueString(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	defaultNullValueStringDescription := ``

	var defaultNullValueStringFlagName string
	if cmdPrefix == "" {
		defaultNullValueStringFlagName = "defaultNullValueString"
	} else {
		defaultNullValueStringFlagName = fmt.Sprintf("%v.defaultNullValueString", cmdPrefix)
	}

	var defaultNullValueStringFlagDefault string

	_ = cmd.PersistentFlags().String(defaultNullValueStringFlagName, defaultNullValueStringFlagDefault, defaultNullValueStringDescription)

	return nil
}

func registerMetricFieldSpecMaxLength(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	maxLengthDescription := ``

	var maxLengthFlagName string
	if cmdPrefix == "" {
		maxLengthFlagName = "maxLength"
	} else {
		maxLengthFlagName = fmt.Sprintf("%v.maxLength", cmdPrefix)
	}

	var maxLengthFlagDefault int32

	_ = cmd.PersistentFlags().Int32(maxLengthFlagName, maxLengthFlagDefault, maxLengthDescription)

	return nil
}

func registerMetricFieldSpecName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerMetricFieldSpecSingleValueField(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	singleValueFieldDescription := ``

	var singleValueFieldFlagName string
	if cmdPrefix == "" {
		singleValueFieldFlagName = "singleValueField"
	} else {
		singleValueFieldFlagName = fmt.Sprintf("%v.singleValueField", cmdPrefix)
	}

	var singleValueFieldFlagDefault bool

	_ = cmd.PersistentFlags().Bool(singleValueFieldFlagName, singleValueFieldFlagDefault, singleValueFieldDescription)

	return nil
}

func registerMetricFieldSpecTransformFunction(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerMetricFieldSpecVirtualColumnProvider(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	virtualColumnProviderDescription := ``

	var virtualColumnProviderFlagName string
	if cmdPrefix == "" {
		virtualColumnProviderFlagName = "virtualColumnProvider"
	} else {
		virtualColumnProviderFlagName = fmt.Sprintf("%v.virtualColumnProvider", cmdPrefix)
	}

	var virtualColumnProviderFlagDefault string

	_ = cmd.PersistentFlags().String(virtualColumnProviderFlagName, virtualColumnProviderFlagDefault, virtualColumnProviderDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelMetricFieldSpecFlags(depth int, m *models.MetricFieldSpec, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataTypeAdded := retrieveMetricFieldSpecDataTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataTypeAdded

	err, defaultNullValueAdded := retrieveMetricFieldSpecDefaultNullValueFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || defaultNullValueAdded

	err, defaultNullValueStringAdded := retrieveMetricFieldSpecDefaultNullValueStringFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || defaultNullValueStringAdded

	err, maxLengthAdded := retrieveMetricFieldSpecMaxLengthFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || maxLengthAdded

	err, nameAdded := retrieveMetricFieldSpecNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, singleValueFieldAdded := retrieveMetricFieldSpecSingleValueFieldFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || singleValueFieldAdded

	err, transformFunctionAdded := retrieveMetricFieldSpecTransformFunctionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || transformFunctionAdded

	err, virtualColumnProviderAdded := retrieveMetricFieldSpecVirtualColumnProviderFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || virtualColumnProviderAdded

	return nil, retAdded
}

func retrieveMetricFieldSpecDataTypeFlags(depth int, m *models.MetricFieldSpec, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataTypeFlagName := fmt.Sprintf("%v.dataType", cmdPrefix)
	if cmd.Flags().Changed(dataTypeFlagName) {

		var dataTypeFlagName string
		if cmdPrefix == "" {
			dataTypeFlagName = "dataType"
		} else {
			dataTypeFlagName = fmt.Sprintf("%v.dataType", cmdPrefix)
		}

		dataTypeFlagValue, err := cmd.Flags().GetString(dataTypeFlagName)
		if err != nil {
			return err, false
		}
		m.DataType = dataTypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMetricFieldSpecDefaultNullValueFlags(depth int, m *models.MetricFieldSpec, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	defaultNullValueFlagName := fmt.Sprintf("%v.defaultNullValue", cmdPrefix)
	if cmd.Flags().Changed(defaultNullValueFlagName) {
		// warning: defaultNullValue map type interface{} is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveMetricFieldSpecDefaultNullValueStringFlags(depth int, m *models.MetricFieldSpec, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	defaultNullValueStringFlagName := fmt.Sprintf("%v.defaultNullValueString", cmdPrefix)
	if cmd.Flags().Changed(defaultNullValueStringFlagName) {

		var defaultNullValueStringFlagName string
		if cmdPrefix == "" {
			defaultNullValueStringFlagName = "defaultNullValueString"
		} else {
			defaultNullValueStringFlagName = fmt.Sprintf("%v.defaultNullValueString", cmdPrefix)
		}

		defaultNullValueStringFlagValue, err := cmd.Flags().GetString(defaultNullValueStringFlagName)
		if err != nil {
			return err, false
		}
		m.DefaultNullValueString = defaultNullValueStringFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMetricFieldSpecMaxLengthFlags(depth int, m *models.MetricFieldSpec, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	maxLengthFlagName := fmt.Sprintf("%v.maxLength", cmdPrefix)
	if cmd.Flags().Changed(maxLengthFlagName) {

		var maxLengthFlagName string
		if cmdPrefix == "" {
			maxLengthFlagName = "maxLength"
		} else {
			maxLengthFlagName = fmt.Sprintf("%v.maxLength", cmdPrefix)
		}

		maxLengthFlagValue, err := cmd.Flags().GetInt32(maxLengthFlagName)
		if err != nil {
			return err, false
		}
		m.MaxLength = maxLengthFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMetricFieldSpecNameFlags(depth int, m *models.MetricFieldSpec, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMetricFieldSpecSingleValueFieldFlags(depth int, m *models.MetricFieldSpec, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	singleValueFieldFlagName := fmt.Sprintf("%v.singleValueField", cmdPrefix)
	if cmd.Flags().Changed(singleValueFieldFlagName) {

		var singleValueFieldFlagName string
		if cmdPrefix == "" {
			singleValueFieldFlagName = "singleValueField"
		} else {
			singleValueFieldFlagName = fmt.Sprintf("%v.singleValueField", cmdPrefix)
		}

		singleValueFieldFlagValue, err := cmd.Flags().GetBool(singleValueFieldFlagName)
		if err != nil {
			return err, false
		}
		m.SingleValueField = singleValueFieldFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMetricFieldSpecTransformFunctionFlags(depth int, m *models.MetricFieldSpec, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveMetricFieldSpecVirtualColumnProviderFlags(depth int, m *models.MetricFieldSpec, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	virtualColumnProviderFlagName := fmt.Sprintf("%v.virtualColumnProvider", cmdPrefix)
	if cmd.Flags().Changed(virtualColumnProviderFlagName) {

		var virtualColumnProviderFlagName string
		if cmdPrefix == "" {
			virtualColumnProviderFlagName = "virtualColumnProvider"
		} else {
			virtualColumnProviderFlagName = fmt.Sprintf("%v.virtualColumnProvider", cmdPrefix)
		}

		virtualColumnProviderFlagValue, err := cmd.Flags().GetString(virtualColumnProviderFlagName)
		if err != nil {
			return err, false
		}
		m.VirtualColumnProvider = virtualColumnProviderFlagValue

		retAdded = true
	}

	return nil, retAdded
}
