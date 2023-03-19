// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/spf13/cobra"
	"startree.ai/cli/models"
)

// Schema cli for JSONIndexConfig

// register flags to command
func registerModelJSONIndexConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerJSONIndexConfigDisableCrossArrayUnnest(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerJSONIndexConfigExcludeArray(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerJSONIndexConfigExcludeFields(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerJSONIndexConfigExcludePaths(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerJSONIndexConfigIncludePaths(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerJSONIndexConfigMaxLevels(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerJSONIndexConfigDisableCrossArrayUnnest(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	disableCrossArrayUnnestDescription := ``

	var disableCrossArrayUnnestFlagName string
	if cmdPrefix == "" {
		disableCrossArrayUnnestFlagName = "disableCrossArrayUnnest"
	} else {
		disableCrossArrayUnnestFlagName = fmt.Sprintf("%v.disableCrossArrayUnnest", cmdPrefix)
	}

	var disableCrossArrayUnnestFlagDefault bool

	_ = cmd.PersistentFlags().Bool(disableCrossArrayUnnestFlagName, disableCrossArrayUnnestFlagDefault, disableCrossArrayUnnestDescription)

	return nil
}

func registerJSONIndexConfigExcludeArray(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	excludeArrayDescription := ``

	var excludeArrayFlagName string
	if cmdPrefix == "" {
		excludeArrayFlagName = "excludeArray"
	} else {
		excludeArrayFlagName = fmt.Sprintf("%v.excludeArray", cmdPrefix)
	}

	var excludeArrayFlagDefault bool

	_ = cmd.PersistentFlags().Bool(excludeArrayFlagName, excludeArrayFlagDefault, excludeArrayDescription)

	return nil
}

func registerJSONIndexConfigExcludeFields(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: excludeFields []string array type is not supported by go-swagger cli yet

	return nil
}

func registerJSONIndexConfigExcludePaths(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: excludePaths []string array type is not supported by go-swagger cli yet

	return nil
}

func registerJSONIndexConfigIncludePaths(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: includePaths []string array type is not supported by go-swagger cli yet

	return nil
}

func registerJSONIndexConfigMaxLevels(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	maxLevelsDescription := ``

	var maxLevelsFlagName string
	if cmdPrefix == "" {
		maxLevelsFlagName = "maxLevels"
	} else {
		maxLevelsFlagName = fmt.Sprintf("%v.maxLevels", cmdPrefix)
	}

	var maxLevelsFlagDefault int32

	_ = cmd.PersistentFlags().Int32(maxLevelsFlagName, maxLevelsFlagDefault, maxLevelsDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelJSONIndexConfigFlags(depth int, m *models.JSONIndexConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, disableCrossArrayUnnestAdded := retrieveJSONIndexConfigDisableCrossArrayUnnestFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || disableCrossArrayUnnestAdded

	err, excludeArrayAdded := retrieveJSONIndexConfigExcludeArrayFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || excludeArrayAdded

	err, excludeFieldsAdded := retrieveJSONIndexConfigExcludeFieldsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || excludeFieldsAdded

	err, excludePathsAdded := retrieveJSONIndexConfigExcludePathsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || excludePathsAdded

	err, includePathsAdded := retrieveJSONIndexConfigIncludePathsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || includePathsAdded

	err, maxLevelsAdded := retrieveJSONIndexConfigMaxLevelsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || maxLevelsAdded

	return nil, retAdded
}

func retrieveJSONIndexConfigDisableCrossArrayUnnestFlags(depth int, m *models.JSONIndexConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	disableCrossArrayUnnestFlagName := fmt.Sprintf("%v.disableCrossArrayUnnest", cmdPrefix)
	if cmd.Flags().Changed(disableCrossArrayUnnestFlagName) {

		var disableCrossArrayUnnestFlagName string
		if cmdPrefix == "" {
			disableCrossArrayUnnestFlagName = "disableCrossArrayUnnest"
		} else {
			disableCrossArrayUnnestFlagName = fmt.Sprintf("%v.disableCrossArrayUnnest", cmdPrefix)
		}

		disableCrossArrayUnnestFlagValue, err := cmd.Flags().GetBool(disableCrossArrayUnnestFlagName)
		if err != nil {
			return err, false
		}
		m.DisableCrossArrayUnnest = disableCrossArrayUnnestFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveJSONIndexConfigExcludeArrayFlags(depth int, m *models.JSONIndexConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	excludeArrayFlagName := fmt.Sprintf("%v.excludeArray", cmdPrefix)
	if cmd.Flags().Changed(excludeArrayFlagName) {

		var excludeArrayFlagName string
		if cmdPrefix == "" {
			excludeArrayFlagName = "excludeArray"
		} else {
			excludeArrayFlagName = fmt.Sprintf("%v.excludeArray", cmdPrefix)
		}

		excludeArrayFlagValue, err := cmd.Flags().GetBool(excludeArrayFlagName)
		if err != nil {
			return err, false
		}
		m.ExcludeArray = excludeArrayFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveJSONIndexConfigExcludeFieldsFlags(depth int, m *models.JSONIndexConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	excludeFieldsFlagName := fmt.Sprintf("%v.excludeFields", cmdPrefix)
	if cmd.Flags().Changed(excludeFieldsFlagName) {
		// warning: excludeFields array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveJSONIndexConfigExcludePathsFlags(depth int, m *models.JSONIndexConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	excludePathsFlagName := fmt.Sprintf("%v.excludePaths", cmdPrefix)
	if cmd.Flags().Changed(excludePathsFlagName) {
		// warning: excludePaths array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveJSONIndexConfigIncludePathsFlags(depth int, m *models.JSONIndexConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	includePathsFlagName := fmt.Sprintf("%v.includePaths", cmdPrefix)
	if cmd.Flags().Changed(includePathsFlagName) {
		// warning: includePaths array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveJSONIndexConfigMaxLevelsFlags(depth int, m *models.JSONIndexConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	maxLevelsFlagName := fmt.Sprintf("%v.maxLevels", cmdPrefix)
	if cmd.Flags().Changed(maxLevelsFlagName) {

		var maxLevelsFlagName string
		if cmdPrefix == "" {
			maxLevelsFlagName = "maxLevels"
		} else {
			maxLevelsFlagName = fmt.Sprintf("%v.maxLevels", cmdPrefix)
		}

		maxLevelsFlagValue, err := cmd.Flags().GetInt32(maxLevelsFlagName)
		if err != nil {
			return err, false
		}
		m.MaxLevels = maxLevelsFlagValue

		retAdded = true
	}

	return nil, retAdded
}