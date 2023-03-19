// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/spf13/cobra"
	"startree.ai/cli/models"
)

// Schema cli for PinotTaskConfig

// register flags to command
func registerModelPinotTaskConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPinotTaskConfigConfigs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPinotTaskConfigTableName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPinotTaskConfigTaskID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPinotTaskConfigTaskType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPinotTaskConfigConfigs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: configs map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerPinotTaskConfigTableName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	tableNameDescription := ``

	var tableNameFlagName string
	if cmdPrefix == "" {
		tableNameFlagName = "tableName"
	} else {
		tableNameFlagName = fmt.Sprintf("%v.tableName", cmdPrefix)
	}

	var tableNameFlagDefault string

	_ = cmd.PersistentFlags().String(tableNameFlagName, tableNameFlagDefault, tableNameDescription)

	return nil
}

func registerPinotTaskConfigTaskID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	taskIdDescription := ``

	var taskIdFlagName string
	if cmdPrefix == "" {
		taskIdFlagName = "taskId"
	} else {
		taskIdFlagName = fmt.Sprintf("%v.taskId", cmdPrefix)
	}

	var taskIdFlagDefault string

	_ = cmd.PersistentFlags().String(taskIdFlagName, taskIdFlagDefault, taskIdDescription)

	return nil
}

func registerPinotTaskConfigTaskType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	taskTypeDescription := ``

	var taskTypeFlagName string
	if cmdPrefix == "" {
		taskTypeFlagName = "taskType"
	} else {
		taskTypeFlagName = fmt.Sprintf("%v.taskType", cmdPrefix)
	}

	var taskTypeFlagDefault string

	_ = cmd.PersistentFlags().String(taskTypeFlagName, taskTypeFlagDefault, taskTypeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPinotTaskConfigFlags(depth int, m *models.PinotTaskConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, configsAdded := retrievePinotTaskConfigConfigsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || configsAdded

	err, tableNameAdded := retrievePinotTaskConfigTableNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tableNameAdded

	err, taskIdAdded := retrievePinotTaskConfigTaskIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || taskIdAdded

	err, taskTypeAdded := retrievePinotTaskConfigTaskTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || taskTypeAdded

	return nil, retAdded
}

func retrievePinotTaskConfigConfigsFlags(depth int, m *models.PinotTaskConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	configsFlagName := fmt.Sprintf("%v.configs", cmdPrefix)
	if cmd.Flags().Changed(configsFlagName) {
		// warning: configs map type map[string]string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrievePinotTaskConfigTableNameFlags(depth int, m *models.PinotTaskConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	tableNameFlagName := fmt.Sprintf("%v.tableName", cmdPrefix)
	if cmd.Flags().Changed(tableNameFlagName) {

		var tableNameFlagName string
		if cmdPrefix == "" {
			tableNameFlagName = "tableName"
		} else {
			tableNameFlagName = fmt.Sprintf("%v.tableName", cmdPrefix)
		}

		tableNameFlagValue, err := cmd.Flags().GetString(tableNameFlagName)
		if err != nil {
			return err, false
		}
		m.TableName = tableNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePinotTaskConfigTaskIDFlags(depth int, m *models.PinotTaskConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	taskIdFlagName := fmt.Sprintf("%v.taskId", cmdPrefix)
	if cmd.Flags().Changed(taskIdFlagName) {

		var taskIdFlagName string
		if cmdPrefix == "" {
			taskIdFlagName = "taskId"
		} else {
			taskIdFlagName = fmt.Sprintf("%v.taskId", cmdPrefix)
		}

		taskIdFlagValue, err := cmd.Flags().GetString(taskIdFlagName)
		if err != nil {
			return err, false
		}
		m.TaskID = taskIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePinotTaskConfigTaskTypeFlags(depth int, m *models.PinotTaskConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	taskTypeFlagName := fmt.Sprintf("%v.taskType", cmdPrefix)
	if cmd.Flags().Changed(taskTypeFlagName) {

		var taskTypeFlagName string
		if cmdPrefix == "" {
			taskTypeFlagName = "taskType"
		} else {
			taskTypeFlagName = fmt.Sprintf("%v.taskType", cmdPrefix)
		}

		taskTypeFlagValue, err := cmd.Flags().GetString(taskTypeFlagName)
		if err != nil {
			return err, false
		}
		m.TaskType = taskTypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}