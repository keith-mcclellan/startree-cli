// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/spf13/cobra"
	"startree.ai/cli/models"
)

// Schema cli for InstanceInfo

// register flags to command
func registerModelInstanceInfoFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerInstanceInfoHost(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInstanceInfoInstanceName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInstanceInfoPort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerInstanceInfoHost(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	hostDescription := ``

	var hostFlagName string
	if cmdPrefix == "" {
		hostFlagName = "host"
	} else {
		hostFlagName = fmt.Sprintf("%v.host", cmdPrefix)
	}

	var hostFlagDefault string

	_ = cmd.PersistentFlags().String(hostFlagName, hostFlagDefault, hostDescription)

	return nil
}

func registerInstanceInfoInstanceName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	instanceNameDescription := ``

	var instanceNameFlagName string
	if cmdPrefix == "" {
		instanceNameFlagName = "instanceName"
	} else {
		instanceNameFlagName = fmt.Sprintf("%v.instanceName", cmdPrefix)
	}

	var instanceNameFlagDefault string

	_ = cmd.PersistentFlags().String(instanceNameFlagName, instanceNameFlagDefault, instanceNameDescription)

	return nil
}

func registerInstanceInfoPort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	portDescription := ``

	var portFlagName string
	if cmdPrefix == "" {
		portFlagName = "port"
	} else {
		portFlagName = fmt.Sprintf("%v.port", cmdPrefix)
	}

	var portFlagDefault int32

	_ = cmd.PersistentFlags().Int32(portFlagName, portFlagDefault, portDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelInstanceInfoFlags(depth int, m *models.InstanceInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, hostAdded := retrieveInstanceInfoHostFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || hostAdded

	err, instanceNameAdded := retrieveInstanceInfoInstanceNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || instanceNameAdded

	err, portAdded := retrieveInstanceInfoPortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || portAdded

	return nil, retAdded
}

func retrieveInstanceInfoHostFlags(depth int, m *models.InstanceInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	hostFlagName := fmt.Sprintf("%v.host", cmdPrefix)
	if cmd.Flags().Changed(hostFlagName) {

		var hostFlagName string
		if cmdPrefix == "" {
			hostFlagName = "host"
		} else {
			hostFlagName = fmt.Sprintf("%v.host", cmdPrefix)
		}

		hostFlagValue, err := cmd.Flags().GetString(hostFlagName)
		if err != nil {
			return err, false
		}
		m.Host = hostFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveInstanceInfoInstanceNameFlags(depth int, m *models.InstanceInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	instanceNameFlagName := fmt.Sprintf("%v.instanceName", cmdPrefix)
	if cmd.Flags().Changed(instanceNameFlagName) {

		var instanceNameFlagName string
		if cmdPrefix == "" {
			instanceNameFlagName = "instanceName"
		} else {
			instanceNameFlagName = fmt.Sprintf("%v.instanceName", cmdPrefix)
		}

		instanceNameFlagValue, err := cmd.Flags().GetString(instanceNameFlagName)
		if err != nil {
			return err, false
		}
		m.InstanceName = instanceNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveInstanceInfoPortFlags(depth int, m *models.InstanceInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	portFlagName := fmt.Sprintf("%v.port", cmdPrefix)
	if cmd.Flags().Changed(portFlagName) {

		var portFlagName string
		if cmdPrefix == "" {
			portFlagName = "port"
		} else {
			portFlagName = fmt.Sprintf("%v.port", cmdPrefix)
		}

		portFlagValue, err := cmd.Flags().GetInt32(portFlagName)
		if err != nil {
			return err, false
		}
		m.Port = portFlagValue

		retAdded = true
	}

	return nil, retAdded
}
