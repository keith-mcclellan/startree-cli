// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"startree.ai/cli/models"

	"github.com/spf13/cobra"
)

// Schema cli for FieldConfig

// register flags to command
func registerModelFieldConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerFieldConfigCompressionCodec(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFieldConfigEncodingType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFieldConfigIndexType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFieldConfigIndexTypes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFieldConfigName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFieldConfigProperties(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFieldConfigTimestampConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerFieldConfigCompressionCodec(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	compressionCodecDescription := `Enum: ["PASS_THROUGH","SNAPPY","ZSTANDARD","LZ4"]. `

	var compressionCodecFlagName string
	if cmdPrefix == "" {
		compressionCodecFlagName = "compressionCodec"
	} else {
		compressionCodecFlagName = fmt.Sprintf("%v.compressionCodec", cmdPrefix)
	}

	var compressionCodecFlagDefault string

	_ = cmd.PersistentFlags().String(compressionCodecFlagName, compressionCodecFlagDefault, compressionCodecDescription)

	if err := cmd.RegisterFlagCompletionFunc(compressionCodecFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["PASS_THROUGH","SNAPPY","ZSTANDARD","LZ4"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerFieldConfigEncodingType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	encodingTypeDescription := `Enum: ["RAW","DICTIONARY"]. `

	var encodingTypeFlagName string
	if cmdPrefix == "" {
		encodingTypeFlagName = "encodingType"
	} else {
		encodingTypeFlagName = fmt.Sprintf("%v.encodingType", cmdPrefix)
	}

	var encodingTypeFlagDefault string

	_ = cmd.PersistentFlags().String(encodingTypeFlagName, encodingTypeFlagDefault, encodingTypeDescription)

	if err := cmd.RegisterFlagCompletionFunc(encodingTypeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["RAW","DICTIONARY"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerFieldConfigIndexType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	indexTypeDescription := `Enum: ["INVERTED","SORTED","TEXT","FST","H3","JSON","TIMESTAMP","RANGE"]. `

	var indexTypeFlagName string
	if cmdPrefix == "" {
		indexTypeFlagName = "indexType"
	} else {
		indexTypeFlagName = fmt.Sprintf("%v.indexType", cmdPrefix)
	}

	var indexTypeFlagDefault string

	_ = cmd.PersistentFlags().String(indexTypeFlagName, indexTypeFlagDefault, indexTypeDescription)

	if err := cmd.RegisterFlagCompletionFunc(indexTypeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["INVERTED","SORTED","TEXT","FST","H3","JSON","TIMESTAMP","RANGE"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerFieldConfigIndexTypes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: indexTypes []string array type is not supported by go-swagger cli yet

	return nil
}

func registerFieldConfigName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Required. `

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

func registerFieldConfigProperties(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: properties map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerFieldConfigTimestampConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var timestampConfigFlagName string
	if cmdPrefix == "" {
		timestampConfigFlagName = "timestampConfig"
	} else {
		timestampConfigFlagName = fmt.Sprintf("%v.timestampConfig", cmdPrefix)
	}

	if err := registerModelTimestampConfigFlags(depth+1, timestampConfigFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelFieldConfigFlags(depth int, m *models.FieldConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, compressionCodecAdded := retrieveFieldConfigCompressionCodecFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || compressionCodecAdded

	err, encodingTypeAdded := retrieveFieldConfigEncodingTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || encodingTypeAdded

	err, indexTypeAdded := retrieveFieldConfigIndexTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || indexTypeAdded

	err, indexTypesAdded := retrieveFieldConfigIndexTypesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || indexTypesAdded

	err, nameAdded := retrieveFieldConfigNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, propertiesAdded := retrieveFieldConfigPropertiesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || propertiesAdded

	err, timestampConfigAdded := retrieveFieldConfigTimestampConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || timestampConfigAdded

	return nil, retAdded
}

func retrieveFieldConfigCompressionCodecFlags(depth int, m *models.FieldConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	compressionCodecFlagName := fmt.Sprintf("%v.compressionCodec", cmdPrefix)
	if cmd.Flags().Changed(compressionCodecFlagName) {

		var compressionCodecFlagName string
		if cmdPrefix == "" {
			compressionCodecFlagName = "compressionCodec"
		} else {
			compressionCodecFlagName = fmt.Sprintf("%v.compressionCodec", cmdPrefix)
		}

		compressionCodecFlagValue, err := cmd.Flags().GetString(compressionCodecFlagName)
		if err != nil {
			return err, false
		}
		m.CompressionCodec = compressionCodecFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveFieldConfigEncodingTypeFlags(depth int, m *models.FieldConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	encodingTypeFlagName := fmt.Sprintf("%v.encodingType", cmdPrefix)
	if cmd.Flags().Changed(encodingTypeFlagName) {

		var encodingTypeFlagName string
		if cmdPrefix == "" {
			encodingTypeFlagName = "encodingType"
		} else {
			encodingTypeFlagName = fmt.Sprintf("%v.encodingType", cmdPrefix)
		}

		encodingTypeFlagValue, err := cmd.Flags().GetString(encodingTypeFlagName)
		if err != nil {
			return err, false
		}
		m.EncodingType = encodingTypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveFieldConfigIndexTypeFlags(depth int, m *models.FieldConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	indexTypeFlagName := fmt.Sprintf("%v.indexType", cmdPrefix)
	if cmd.Flags().Changed(indexTypeFlagName) {

		var indexTypeFlagName string
		if cmdPrefix == "" {
			indexTypeFlagName = "indexType"
		} else {
			indexTypeFlagName = fmt.Sprintf("%v.indexType", cmdPrefix)
		}

		indexTypeFlagValue, err := cmd.Flags().GetString(indexTypeFlagName)
		if err != nil {
			return err, false
		}
		m.IndexType = indexTypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveFieldConfigIndexTypesFlags(depth int, m *models.FieldConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	indexTypesFlagName := fmt.Sprintf("%v.indexTypes", cmdPrefix)
	if cmd.Flags().Changed(indexTypesFlagName) {
		// warning: indexTypes array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveFieldConfigNameFlags(depth int, m *models.FieldConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveFieldConfigPropertiesFlags(depth int, m *models.FieldConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	propertiesFlagName := fmt.Sprintf("%v.properties", cmdPrefix)
	if cmd.Flags().Changed(propertiesFlagName) {
		// warning: properties map type map[string]string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveFieldConfigTimestampConfigFlags(depth int, m *models.FieldConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	timestampConfigFlagName := fmt.Sprintf("%v.timestampConfig", cmdPrefix)
	if cmd.Flags().Changed(timestampConfigFlagName) {
		// info: complex object timestampConfig TimestampConfig is retrieved outside this Changed() block
	}
	timestampConfigFlagValue := m.TimestampConfig
	if swag.IsZero(timestampConfigFlagValue) {
		timestampConfigFlagValue = &models.TimestampConfig{}
	}

	err, timestampConfigAdded := retrieveModelTimestampConfigFlags(depth+1, timestampConfigFlagValue, timestampConfigFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || timestampConfigAdded
	if timestampConfigAdded {
		m.TimestampConfig = timestampConfigFlagValue
	}

	return nil, retAdded
}