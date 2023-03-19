// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/spf13/cobra"
	"startree.ai/cli/models"
)

// Schema cli for InstanceReplicaGroupPartitionConfig

// register flags to command
func registerModelInstanceReplicaGroupPartitionConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerInstanceReplicaGroupPartitionConfigMinimizeDataMovement(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInstanceReplicaGroupPartitionConfigNumInstances(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInstanceReplicaGroupPartitionConfigNumInstancesPerPartition(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInstanceReplicaGroupPartitionConfigNumInstancesPerReplicaGroup(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInstanceReplicaGroupPartitionConfigNumPartitions(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInstanceReplicaGroupPartitionConfigNumReplicaGroups(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInstanceReplicaGroupPartitionConfigReplicaGroupBased(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerInstanceReplicaGroupPartitionConfigMinimizeDataMovement(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	minimizeDataMovementDescription := ``

	var minimizeDataMovementFlagName string
	if cmdPrefix == "" {
		minimizeDataMovementFlagName = "minimizeDataMovement"
	} else {
		minimizeDataMovementFlagName = fmt.Sprintf("%v.minimizeDataMovement", cmdPrefix)
	}

	var minimizeDataMovementFlagDefault bool

	_ = cmd.PersistentFlags().Bool(minimizeDataMovementFlagName, minimizeDataMovementFlagDefault, minimizeDataMovementDescription)

	return nil
}

func registerInstanceReplicaGroupPartitionConfigNumInstances(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	numInstancesDescription := ``

	var numInstancesFlagName string
	if cmdPrefix == "" {
		numInstancesFlagName = "numInstances"
	} else {
		numInstancesFlagName = fmt.Sprintf("%v.numInstances", cmdPrefix)
	}

	var numInstancesFlagDefault int32

	_ = cmd.PersistentFlags().Int32(numInstancesFlagName, numInstancesFlagDefault, numInstancesDescription)

	return nil
}

func registerInstanceReplicaGroupPartitionConfigNumInstancesPerPartition(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	numInstancesPerPartitionDescription := ``

	var numInstancesPerPartitionFlagName string
	if cmdPrefix == "" {
		numInstancesPerPartitionFlagName = "numInstancesPerPartition"
	} else {
		numInstancesPerPartitionFlagName = fmt.Sprintf("%v.numInstancesPerPartition", cmdPrefix)
	}

	var numInstancesPerPartitionFlagDefault int32

	_ = cmd.PersistentFlags().Int32(numInstancesPerPartitionFlagName, numInstancesPerPartitionFlagDefault, numInstancesPerPartitionDescription)

	return nil
}

func registerInstanceReplicaGroupPartitionConfigNumInstancesPerReplicaGroup(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	numInstancesPerReplicaGroupDescription := ``

	var numInstancesPerReplicaGroupFlagName string
	if cmdPrefix == "" {
		numInstancesPerReplicaGroupFlagName = "numInstancesPerReplicaGroup"
	} else {
		numInstancesPerReplicaGroupFlagName = fmt.Sprintf("%v.numInstancesPerReplicaGroup", cmdPrefix)
	}

	var numInstancesPerReplicaGroupFlagDefault int32

	_ = cmd.PersistentFlags().Int32(numInstancesPerReplicaGroupFlagName, numInstancesPerReplicaGroupFlagDefault, numInstancesPerReplicaGroupDescription)

	return nil
}

func registerInstanceReplicaGroupPartitionConfigNumPartitions(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	numPartitionsDescription := ``

	var numPartitionsFlagName string
	if cmdPrefix == "" {
		numPartitionsFlagName = "numPartitions"
	} else {
		numPartitionsFlagName = fmt.Sprintf("%v.numPartitions", cmdPrefix)
	}

	var numPartitionsFlagDefault int32

	_ = cmd.PersistentFlags().Int32(numPartitionsFlagName, numPartitionsFlagDefault, numPartitionsDescription)

	return nil
}

func registerInstanceReplicaGroupPartitionConfigNumReplicaGroups(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	numReplicaGroupsDescription := ``

	var numReplicaGroupsFlagName string
	if cmdPrefix == "" {
		numReplicaGroupsFlagName = "numReplicaGroups"
	} else {
		numReplicaGroupsFlagName = fmt.Sprintf("%v.numReplicaGroups", cmdPrefix)
	}

	var numReplicaGroupsFlagDefault int32

	_ = cmd.PersistentFlags().Int32(numReplicaGroupsFlagName, numReplicaGroupsFlagDefault, numReplicaGroupsDescription)

	return nil
}

func registerInstanceReplicaGroupPartitionConfigReplicaGroupBased(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	replicaGroupBasedDescription := ``

	var replicaGroupBasedFlagName string
	if cmdPrefix == "" {
		replicaGroupBasedFlagName = "replicaGroupBased"
	} else {
		replicaGroupBasedFlagName = fmt.Sprintf("%v.replicaGroupBased", cmdPrefix)
	}

	var replicaGroupBasedFlagDefault bool

	_ = cmd.PersistentFlags().Bool(replicaGroupBasedFlagName, replicaGroupBasedFlagDefault, replicaGroupBasedDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelInstanceReplicaGroupPartitionConfigFlags(depth int, m *models.InstanceReplicaGroupPartitionConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, minimizeDataMovementAdded := retrieveInstanceReplicaGroupPartitionConfigMinimizeDataMovementFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || minimizeDataMovementAdded

	err, numInstancesAdded := retrieveInstanceReplicaGroupPartitionConfigNumInstancesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || numInstancesAdded

	err, numInstancesPerPartitionAdded := retrieveInstanceReplicaGroupPartitionConfigNumInstancesPerPartitionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || numInstancesPerPartitionAdded

	err, numInstancesPerReplicaGroupAdded := retrieveInstanceReplicaGroupPartitionConfigNumInstancesPerReplicaGroupFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || numInstancesPerReplicaGroupAdded

	err, numPartitionsAdded := retrieveInstanceReplicaGroupPartitionConfigNumPartitionsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || numPartitionsAdded

	err, numReplicaGroupsAdded := retrieveInstanceReplicaGroupPartitionConfigNumReplicaGroupsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || numReplicaGroupsAdded

	err, replicaGroupBasedAdded := retrieveInstanceReplicaGroupPartitionConfigReplicaGroupBasedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || replicaGroupBasedAdded

	return nil, retAdded
}

func retrieveInstanceReplicaGroupPartitionConfigMinimizeDataMovementFlags(depth int, m *models.InstanceReplicaGroupPartitionConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	minimizeDataMovementFlagName := fmt.Sprintf("%v.minimizeDataMovement", cmdPrefix)
	if cmd.Flags().Changed(minimizeDataMovementFlagName) {

		var minimizeDataMovementFlagName string
		if cmdPrefix == "" {
			minimizeDataMovementFlagName = "minimizeDataMovement"
		} else {
			minimizeDataMovementFlagName = fmt.Sprintf("%v.minimizeDataMovement", cmdPrefix)
		}

		minimizeDataMovementFlagValue, err := cmd.Flags().GetBool(minimizeDataMovementFlagName)
		if err != nil {
			return err, false
		}
		m.MinimizeDataMovement = &minimizeDataMovementFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveInstanceReplicaGroupPartitionConfigNumInstancesFlags(depth int, m *models.InstanceReplicaGroupPartitionConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	numInstancesFlagName := fmt.Sprintf("%v.numInstances", cmdPrefix)
	if cmd.Flags().Changed(numInstancesFlagName) {

		var numInstancesFlagName string
		if cmdPrefix == "" {
			numInstancesFlagName = "numInstances"
		} else {
			numInstancesFlagName = fmt.Sprintf("%v.numInstances", cmdPrefix)
		}

		numInstancesFlagValue, err := cmd.Flags().GetInt32(numInstancesFlagName)
		if err != nil {
			return err, false
		}
		m.NumInstances = numInstancesFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveInstanceReplicaGroupPartitionConfigNumInstancesPerPartitionFlags(depth int, m *models.InstanceReplicaGroupPartitionConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	numInstancesPerPartitionFlagName := fmt.Sprintf("%v.numInstancesPerPartition", cmdPrefix)
	if cmd.Flags().Changed(numInstancesPerPartitionFlagName) {

		var numInstancesPerPartitionFlagName string
		if cmdPrefix == "" {
			numInstancesPerPartitionFlagName = "numInstancesPerPartition"
		} else {
			numInstancesPerPartitionFlagName = fmt.Sprintf("%v.numInstancesPerPartition", cmdPrefix)
		}

		numInstancesPerPartitionFlagValue, err := cmd.Flags().GetInt32(numInstancesPerPartitionFlagName)
		if err != nil {
			return err, false
		}
		m.NumInstancesPerPartition = numInstancesPerPartitionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveInstanceReplicaGroupPartitionConfigNumInstancesPerReplicaGroupFlags(depth int, m *models.InstanceReplicaGroupPartitionConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	numInstancesPerReplicaGroupFlagName := fmt.Sprintf("%v.numInstancesPerReplicaGroup", cmdPrefix)
	if cmd.Flags().Changed(numInstancesPerReplicaGroupFlagName) {

		var numInstancesPerReplicaGroupFlagName string
		if cmdPrefix == "" {
			numInstancesPerReplicaGroupFlagName = "numInstancesPerReplicaGroup"
		} else {
			numInstancesPerReplicaGroupFlagName = fmt.Sprintf("%v.numInstancesPerReplicaGroup", cmdPrefix)
		}

		numInstancesPerReplicaGroupFlagValue, err := cmd.Flags().GetInt32(numInstancesPerReplicaGroupFlagName)
		if err != nil {
			return err, false
		}
		m.NumInstancesPerReplicaGroup = numInstancesPerReplicaGroupFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveInstanceReplicaGroupPartitionConfigNumPartitionsFlags(depth int, m *models.InstanceReplicaGroupPartitionConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	numPartitionsFlagName := fmt.Sprintf("%v.numPartitions", cmdPrefix)
	if cmd.Flags().Changed(numPartitionsFlagName) {

		var numPartitionsFlagName string
		if cmdPrefix == "" {
			numPartitionsFlagName = "numPartitions"
		} else {
			numPartitionsFlagName = fmt.Sprintf("%v.numPartitions", cmdPrefix)
		}

		numPartitionsFlagValue, err := cmd.Flags().GetInt32(numPartitionsFlagName)
		if err != nil {
			return err, false
		}
		m.NumPartitions = numPartitionsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveInstanceReplicaGroupPartitionConfigNumReplicaGroupsFlags(depth int, m *models.InstanceReplicaGroupPartitionConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	numReplicaGroupsFlagName := fmt.Sprintf("%v.numReplicaGroups", cmdPrefix)
	if cmd.Flags().Changed(numReplicaGroupsFlagName) {

		var numReplicaGroupsFlagName string
		if cmdPrefix == "" {
			numReplicaGroupsFlagName = "numReplicaGroups"
		} else {
			numReplicaGroupsFlagName = fmt.Sprintf("%v.numReplicaGroups", cmdPrefix)
		}

		numReplicaGroupsFlagValue, err := cmd.Flags().GetInt32(numReplicaGroupsFlagName)
		if err != nil {
			return err, false
		}
		m.NumReplicaGroups = numReplicaGroupsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveInstanceReplicaGroupPartitionConfigReplicaGroupBasedFlags(depth int, m *models.InstanceReplicaGroupPartitionConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	replicaGroupBasedFlagName := fmt.Sprintf("%v.replicaGroupBased", cmdPrefix)
	if cmd.Flags().Changed(replicaGroupBasedFlagName) {

		var replicaGroupBasedFlagName string
		if cmdPrefix == "" {
			replicaGroupBasedFlagName = "replicaGroupBased"
		} else {
			replicaGroupBasedFlagName = fmt.Sprintf("%v.replicaGroupBased", cmdPrefix)
		}

		replicaGroupBasedFlagValue, err := cmd.Flags().GetBool(replicaGroupBasedFlagName)
		if err != nil {
			return err, false
		}
		m.ReplicaGroupBased = &replicaGroupBasedFlagValue

		retAdded = true
	}

	return nil, retAdded
}
