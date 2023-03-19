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

// Schema cli for ComplexTypeConfig

// register flags to command
func registerModelComplexTypeConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerComplexTypeConfigCollectionNotUnnestedToJSON(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerComplexTypeConfigDelimiter(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerComplexTypeConfigFieldsToUnnest(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerComplexTypeConfigPrefixesToRename(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerComplexTypeConfigCollectionNotUnnestedToJSON(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	collectionNotUnnestedToJsonDescription := `Enum: ["NONE","NON_PRIMITIVE","ALL"]. `

	var collectionNotUnnestedToJsonFlagName string
	if cmdPrefix == "" {
		collectionNotUnnestedToJsonFlagName = "collectionNotUnnestedToJson"
	} else {
		collectionNotUnnestedToJsonFlagName = fmt.Sprintf("%v.collectionNotUnnestedToJson", cmdPrefix)
	}

	var collectionNotUnnestedToJsonFlagDefault string

	_ = cmd.PersistentFlags().String(collectionNotUnnestedToJsonFlagName, collectionNotUnnestedToJsonFlagDefault, collectionNotUnnestedToJsonDescription)

	if err := cmd.RegisterFlagCompletionFunc(collectionNotUnnestedToJsonFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["NONE","NON_PRIMITIVE","ALL"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerComplexTypeConfigDelimiter(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	delimiterDescription := ``

	var delimiterFlagName string
	if cmdPrefix == "" {
		delimiterFlagName = "delimiter"
	} else {
		delimiterFlagName = fmt.Sprintf("%v.delimiter", cmdPrefix)
	}

	var delimiterFlagDefault string

	_ = cmd.PersistentFlags().String(delimiterFlagName, delimiterFlagDefault, delimiterDescription)

	return nil
}

func registerComplexTypeConfigFieldsToUnnest(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: fieldsToUnnest []string array type is not supported by go-swagger cli yet

	return nil
}

func registerComplexTypeConfigPrefixesToRename(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: prefixesToRename map[string]string map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelComplexTypeConfigFlags(depth int, m *models.ComplexTypeConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, collectionNotUnnestedToJsonAdded := retrieveComplexTypeConfigCollectionNotUnnestedToJSONFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || collectionNotUnnestedToJsonAdded

	err, delimiterAdded := retrieveComplexTypeConfigDelimiterFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || delimiterAdded

	err, fieldsToUnnestAdded := retrieveComplexTypeConfigFieldsToUnnestFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || fieldsToUnnestAdded

	err, prefixesToRenameAdded := retrieveComplexTypeConfigPrefixesToRenameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || prefixesToRenameAdded

	return nil, retAdded
}

func retrieveComplexTypeConfigCollectionNotUnnestedToJSONFlags(depth int, m *models.ComplexTypeConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	collectionNotUnnestedToJsonFlagName := fmt.Sprintf("%v.collectionNotUnnestedToJson", cmdPrefix)
	if cmd.Flags().Changed(collectionNotUnnestedToJsonFlagName) {

		var collectionNotUnnestedToJsonFlagName string
		if cmdPrefix == "" {
			collectionNotUnnestedToJsonFlagName = "collectionNotUnnestedToJson"
		} else {
			collectionNotUnnestedToJsonFlagName = fmt.Sprintf("%v.collectionNotUnnestedToJson", cmdPrefix)
		}

		collectionNotUnnestedToJsonFlagValue, err := cmd.Flags().GetString(collectionNotUnnestedToJsonFlagName)
		if err != nil {
			return err, false
		}
		m.CollectionNotUnnestedToJSON = collectionNotUnnestedToJsonFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveComplexTypeConfigDelimiterFlags(depth int, m *models.ComplexTypeConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	delimiterFlagName := fmt.Sprintf("%v.delimiter", cmdPrefix)
	if cmd.Flags().Changed(delimiterFlagName) {

		var delimiterFlagName string
		if cmdPrefix == "" {
			delimiterFlagName = "delimiter"
		} else {
			delimiterFlagName = fmt.Sprintf("%v.delimiter", cmdPrefix)
		}

		delimiterFlagValue, err := cmd.Flags().GetString(delimiterFlagName)
		if err != nil {
			return err, false
		}
		m.Delimiter = delimiterFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveComplexTypeConfigFieldsToUnnestFlags(depth int, m *models.ComplexTypeConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	fieldsToUnnestFlagName := fmt.Sprintf("%v.fieldsToUnnest", cmdPrefix)
	if cmd.Flags().Changed(fieldsToUnnestFlagName) {
		// warning: fieldsToUnnest array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveComplexTypeConfigPrefixesToRenameFlags(depth int, m *models.ComplexTypeConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	prefixesToRenameFlagName := fmt.Sprintf("%v.prefixesToRename", cmdPrefix)
	if cmd.Flags().Changed(prefixesToRenameFlagName) {
		// warning: prefixesToRename map type map[string]string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}