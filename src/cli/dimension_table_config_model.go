// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/spf13/cobra"
	"startree.ai/cli/models"
)

// Schema cli for DimensionTableConfig

// register flags to command
func registerModelDimensionTableConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDimensionTableConfigDisablePreload(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDimensionTableConfigDisablePreload(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	disablePreloadDescription := `Required. `

	var disablePreloadFlagName string
	if cmdPrefix == "" {
		disablePreloadFlagName = "disablePreload"
	} else {
		disablePreloadFlagName = fmt.Sprintf("%v.disablePreload", cmdPrefix)
	}

	var disablePreloadFlagDefault bool

	_ = cmd.PersistentFlags().Bool(disablePreloadFlagName, disablePreloadFlagDefault, disablePreloadDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDimensionTableConfigFlags(depth int, m *models.DimensionTableConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, disablePreloadAdded := retrieveDimensionTableConfigDisablePreloadFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || disablePreloadAdded

	return nil, retAdded
}

func retrieveDimensionTableConfigDisablePreloadFlags(depth int, m *models.DimensionTableConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	disablePreloadFlagName := fmt.Sprintf("%v.disablePreload", cmdPrefix)
	if cmd.Flags().Changed(disablePreloadFlagName) {

		var disablePreloadFlagName string
		if cmdPrefix == "" {
			disablePreloadFlagName = "disablePreload"
		} else {
			disablePreloadFlagName = fmt.Sprintf("%v.disablePreload", cmdPrefix)
		}

		disablePreloadFlagValue, err := cmd.Flags().GetBool(disablePreloadFlagName)
		if err != nil {
			return err, false
		}
		m.DisablePreload = disablePreloadFlagValue

		retAdded = true
	}

	return nil, retAdded
}
