// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/spf13/cobra"
	"startree.ai/cli/models"
)

// Schema cli for TenantMetadata

// register flags to command
func registerModelTenantMetadataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTenantMetadataBrokerInstances(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTenantMetadataOfflineServerInstances(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTenantMetadataRealtimeServerInstances(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTenantMetadataServerInstances(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTenantMetadataTenantName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTenantMetadataBrokerInstances(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: BrokerInstances []string array type is not supported by go-swagger cli yet

	return nil
}

func registerTenantMetadataOfflineServerInstances(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: OfflineServerInstances []string array type is not supported by go-swagger cli yet

	return nil
}

func registerTenantMetadataRealtimeServerInstances(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: RealtimeServerInstances []string array type is not supported by go-swagger cli yet

	return nil
}

func registerTenantMetadataServerInstances(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: ServerInstances []string array type is not supported by go-swagger cli yet

	return nil
}

func registerTenantMetadataTenantName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	tenantNameDescription := ``

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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTenantMetadataFlags(depth int, m *models.TenantMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, brokerInstancesAdded := retrieveTenantMetadataBrokerInstancesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || brokerInstancesAdded

	err, offlineServerInstancesAdded := retrieveTenantMetadataOfflineServerInstancesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || offlineServerInstancesAdded

	err, realtimeServerInstancesAdded := retrieveTenantMetadataRealtimeServerInstancesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || realtimeServerInstancesAdded

	err, serverInstancesAdded := retrieveTenantMetadataServerInstancesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || serverInstancesAdded

	err, tenantNameAdded := retrieveTenantMetadataTenantNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tenantNameAdded

	return nil, retAdded
}

func retrieveTenantMetadataBrokerInstancesFlags(depth int, m *models.TenantMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	brokerInstancesFlagName := fmt.Sprintf("%v.BrokerInstances", cmdPrefix)
	if cmd.Flags().Changed(brokerInstancesFlagName) {
		// warning: BrokerInstances array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveTenantMetadataOfflineServerInstancesFlags(depth int, m *models.TenantMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	offlineServerInstancesFlagName := fmt.Sprintf("%v.OfflineServerInstances", cmdPrefix)
	if cmd.Flags().Changed(offlineServerInstancesFlagName) {
		// warning: OfflineServerInstances array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveTenantMetadataRealtimeServerInstancesFlags(depth int, m *models.TenantMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	realtimeServerInstancesFlagName := fmt.Sprintf("%v.RealtimeServerInstances", cmdPrefix)
	if cmd.Flags().Changed(realtimeServerInstancesFlagName) {
		// warning: RealtimeServerInstances array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveTenantMetadataServerInstancesFlags(depth int, m *models.TenantMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	serverInstancesFlagName := fmt.Sprintf("%v.ServerInstances", cmdPrefix)
	if cmd.Flags().Changed(serverInstancesFlagName) {
		// warning: ServerInstances array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveTenantMetadataTenantNameFlags(depth int, m *models.TenantMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
