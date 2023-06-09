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

// Schema cli for Tenant

// register flags to command
func registerModelTenantFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTenantNumberOfInstances(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTenantOfflineInstances(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTenantRealtimeInstances(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTenantTenantName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTenantTenantRole(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTenantNumberOfInstances(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	numberOfInstancesDescription := ``

	var numberOfInstancesFlagName string
	if cmdPrefix == "" {
		numberOfInstancesFlagName = "numberOfInstances"
	} else {
		numberOfInstancesFlagName = fmt.Sprintf("%v.numberOfInstances", cmdPrefix)
	}

	var numberOfInstancesFlagDefault int32

	_ = cmd.PersistentFlags().Int32(numberOfInstancesFlagName, numberOfInstancesFlagDefault, numberOfInstancesDescription)

	return nil
}

func registerTenantOfflineInstances(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	offlineInstancesDescription := ``

	var offlineInstancesFlagName string
	if cmdPrefix == "" {
		offlineInstancesFlagName = "offlineInstances"
	} else {
		offlineInstancesFlagName = fmt.Sprintf("%v.offlineInstances", cmdPrefix)
	}

	var offlineInstancesFlagDefault int32

	_ = cmd.PersistentFlags().Int32(offlineInstancesFlagName, offlineInstancesFlagDefault, offlineInstancesDescription)

	return nil
}

func registerTenantRealtimeInstances(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	realtimeInstancesDescription := ``

	var realtimeInstancesFlagName string
	if cmdPrefix == "" {
		realtimeInstancesFlagName = "realtimeInstances"
	} else {
		realtimeInstancesFlagName = fmt.Sprintf("%v.realtimeInstances", cmdPrefix)
	}

	var realtimeInstancesFlagDefault int32

	_ = cmd.PersistentFlags().Int32(realtimeInstancesFlagName, realtimeInstancesFlagDefault, realtimeInstancesDescription)

	return nil
}

func registerTenantTenantName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	tenantNameDescription := `Required. `

	var tenantNameFlagName string
	if cmdPrefix == "" {
		tenantNameFlagName = "tenantName"
	} else {
		tenantNameFlagName = fmt.Sprintf("%v.tenantName", cmdPrefix)
	}

	var tenantNameFlagDefault string

	_ = cmd.PersistentFlags().String(tenantNameFlagName, tenantNameFlagDefault, tenantNameDescription)

	return nil
}

func registerTenantTenantRole(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	tenantRoleDescription := `Enum: ["SERVER","BROKER","MINION"]. Required. `

	var tenantRoleFlagName string
	if cmdPrefix == "" {
		tenantRoleFlagName = "tenantRole"
	} else {
		tenantRoleFlagName = fmt.Sprintf("%v.tenantRole", cmdPrefix)
	}

	var tenantRoleFlagDefault string

	_ = cmd.PersistentFlags().String(tenantRoleFlagName, tenantRoleFlagDefault, tenantRoleDescription)

	if err := cmd.RegisterFlagCompletionFunc(tenantRoleFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["SERVER","BROKER","MINION"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTenantFlags(depth int, m *models.Tenant, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, numberOfInstancesAdded := retrieveTenantNumberOfInstancesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || numberOfInstancesAdded

	err, offlineInstancesAdded := retrieveTenantOfflineInstancesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || offlineInstancesAdded

	err, realtimeInstancesAdded := retrieveTenantRealtimeInstancesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || realtimeInstancesAdded

	err, tenantNameAdded := retrieveTenantTenantNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tenantNameAdded

	err, tenantRoleAdded := retrieveTenantTenantRoleFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tenantRoleAdded

	return nil, retAdded
}

func retrieveTenantNumberOfInstancesFlags(depth int, m *models.Tenant, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	numberOfInstancesFlagName := fmt.Sprintf("%v.numberOfInstances", cmdPrefix)
	if cmd.Flags().Changed(numberOfInstancesFlagName) {

		var numberOfInstancesFlagName string
		if cmdPrefix == "" {
			numberOfInstancesFlagName = "numberOfInstances"
		} else {
			numberOfInstancesFlagName = fmt.Sprintf("%v.numberOfInstances", cmdPrefix)
		}

		numberOfInstancesFlagValue, err := cmd.Flags().GetInt32(numberOfInstancesFlagName)
		if err != nil {
			return err, false
		}
		m.NumberOfInstances = numberOfInstancesFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTenantOfflineInstancesFlags(depth int, m *models.Tenant, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	offlineInstancesFlagName := fmt.Sprintf("%v.offlineInstances", cmdPrefix)
	if cmd.Flags().Changed(offlineInstancesFlagName) {

		var offlineInstancesFlagName string
		if cmdPrefix == "" {
			offlineInstancesFlagName = "offlineInstances"
		} else {
			offlineInstancesFlagName = fmt.Sprintf("%v.offlineInstances", cmdPrefix)
		}

		offlineInstancesFlagValue, err := cmd.Flags().GetInt32(offlineInstancesFlagName)
		if err != nil {
			return err, false
		}
		m.OfflineInstances = offlineInstancesFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTenantRealtimeInstancesFlags(depth int, m *models.Tenant, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	realtimeInstancesFlagName := fmt.Sprintf("%v.realtimeInstances", cmdPrefix)
	if cmd.Flags().Changed(realtimeInstancesFlagName) {

		var realtimeInstancesFlagName string
		if cmdPrefix == "" {
			realtimeInstancesFlagName = "realtimeInstances"
		} else {
			realtimeInstancesFlagName = fmt.Sprintf("%v.realtimeInstances", cmdPrefix)
		}

		realtimeInstancesFlagValue, err := cmd.Flags().GetInt32(realtimeInstancesFlagName)
		if err != nil {
			return err, false
		}
		m.RealtimeInstances = realtimeInstancesFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTenantTenantNameFlags(depth int, m *models.Tenant, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	tenantNameFlagName := fmt.Sprintf("%v.tenantName", cmdPrefix)
	if cmd.Flags().Changed(tenantNameFlagName) {

		var tenantNameFlagName string
		if cmdPrefix == "" {
			tenantNameFlagName = "tenantName"
		} else {
			tenantNameFlagName = fmt.Sprintf("%v.tenantName", cmdPrefix)
		}

		tenantNameFlagValue, err := cmd.Flags().GetString(tenantNameFlagName)
		if err != nil {
			return err, false
		}
		m.TenantName = tenantNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTenantTenantRoleFlags(depth int, m *models.Tenant, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	tenantRoleFlagName := fmt.Sprintf("%v.tenantRole", cmdPrefix)
	if cmd.Flags().Changed(tenantRoleFlagName) {

		var tenantRoleFlagName string
		if cmdPrefix == "" {
			tenantRoleFlagName = "tenantRole"
		} else {
			tenantRoleFlagName = fmt.Sprintf("%v.tenantRole", cmdPrefix)
		}

		tenantRoleFlagValue, err := cmd.Flags().GetString(tenantRoleFlagName)
		if err != nil {
			return err, false
		}
		m.TenantRole = tenantRoleFlagValue

		retAdded = true
	}

	return nil, retAdded
}
