// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/spf13/cobra"
	"startree.ai/cli/models"
)

// Schema cli for TableWriteConfig

// register flags to command
func registerModelTableWriteConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTableWriteConfigEncoderClass(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTableWriteConfigEncoderConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTableWriteConfigPartitionColumns(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTableWriteConfigProducerConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTableWriteConfigProducerType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTableWriteConfigTopic(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTableWriteConfigEncoderClass(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	encoderClassDescription := ``

	var encoderClassFlagName string
	if cmdPrefix == "" {
		encoderClassFlagName = "encoderClass"
	} else {
		encoderClassFlagName = fmt.Sprintf("%v.encoderClass", cmdPrefix)
	}

	var encoderClassFlagDefault string

	_ = cmd.PersistentFlags().String(encoderClassFlagName, encoderClassFlagDefault, encoderClassDescription)

	return nil
}

func registerTableWriteConfigEncoderConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: encoderConfig map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerTableWriteConfigPartitionColumns(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: partitionColumns []string array type is not supported by go-swagger cli yet

	return nil
}

func registerTableWriteConfigProducerConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: producerConfig map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerTableWriteConfigProducerType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	producerTypeDescription := ``

	var producerTypeFlagName string
	if cmdPrefix == "" {
		producerTypeFlagName = "producerType"
	} else {
		producerTypeFlagName = fmt.Sprintf("%v.producerType", cmdPrefix)
	}

	var producerTypeFlagDefault string

	_ = cmd.PersistentFlags().String(producerTypeFlagName, producerTypeFlagDefault, producerTypeDescription)

	return nil
}

func registerTableWriteConfigTopic(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	topicDescription := ``

	var topicFlagName string
	if cmdPrefix == "" {
		topicFlagName = "topic"
	} else {
		topicFlagName = fmt.Sprintf("%v.topic", cmdPrefix)
	}

	var topicFlagDefault string

	_ = cmd.PersistentFlags().String(topicFlagName, topicFlagDefault, topicDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTableWriteConfigFlags(depth int, m *models.TableWriteConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, encoderClassAdded := retrieveTableWriteConfigEncoderClassFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || encoderClassAdded

	err, encoderConfigAdded := retrieveTableWriteConfigEncoderConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || encoderConfigAdded

	err, partitionColumnsAdded := retrieveTableWriteConfigPartitionColumnsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || partitionColumnsAdded

	err, producerConfigAdded := retrieveTableWriteConfigProducerConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || producerConfigAdded

	err, producerTypeAdded := retrieveTableWriteConfigProducerTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || producerTypeAdded

	err, topicAdded := retrieveTableWriteConfigTopicFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || topicAdded

	return nil, retAdded
}

func retrieveTableWriteConfigEncoderClassFlags(depth int, m *models.TableWriteConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	encoderClassFlagName := fmt.Sprintf("%v.encoderClass", cmdPrefix)
	if cmd.Flags().Changed(encoderClassFlagName) {

		var encoderClassFlagName string
		if cmdPrefix == "" {
			encoderClassFlagName = "encoderClass"
		} else {
			encoderClassFlagName = fmt.Sprintf("%v.encoderClass", cmdPrefix)
		}

		encoderClassFlagValue, err := cmd.Flags().GetString(encoderClassFlagName)
		if err != nil {
			return err, false
		}
		m.EncoderClass = encoderClassFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTableWriteConfigEncoderConfigFlags(depth int, m *models.TableWriteConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	encoderConfigFlagName := fmt.Sprintf("%v.encoderConfig", cmdPrefix)
	if cmd.Flags().Changed(encoderConfigFlagName) {
		// warning: encoderConfig map type map[string]string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveTableWriteConfigPartitionColumnsFlags(depth int, m *models.TableWriteConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	partitionColumnsFlagName := fmt.Sprintf("%v.partitionColumns", cmdPrefix)
	if cmd.Flags().Changed(partitionColumnsFlagName) {
		// warning: partitionColumns array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveTableWriteConfigProducerConfigFlags(depth int, m *models.TableWriteConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	producerConfigFlagName := fmt.Sprintf("%v.producerConfig", cmdPrefix)
	if cmd.Flags().Changed(producerConfigFlagName) {
		// warning: producerConfig map type map[string]string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveTableWriteConfigProducerTypeFlags(depth int, m *models.TableWriteConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	producerTypeFlagName := fmt.Sprintf("%v.producerType", cmdPrefix)
	if cmd.Flags().Changed(producerTypeFlagName) {

		var producerTypeFlagName string
		if cmdPrefix == "" {
			producerTypeFlagName = "producerType"
		} else {
			producerTypeFlagName = fmt.Sprintf("%v.producerType", cmdPrefix)
		}

		producerTypeFlagValue, err := cmd.Flags().GetString(producerTypeFlagName)
		if err != nil {
			return err, false
		}
		m.ProducerType = producerTypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTableWriteConfigTopicFlags(depth int, m *models.TableWriteConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	topicFlagName := fmt.Sprintf("%v.topic", cmdPrefix)
	if cmd.Flags().Changed(topicFlagName) {

		var topicFlagName string
		if cmdPrefix == "" {
			topicFlagName = "topic"
		} else {
			topicFlagName = fmt.Sprintf("%v.topic", cmdPrefix)
		}

		topicFlagValue, err := cmd.Flags().GetString(topicFlagName)
		if err != nil {
			return err, false
		}
		m.Topic = topicFlagValue

		retAdded = true
	}

	return nil, retAdded
}
