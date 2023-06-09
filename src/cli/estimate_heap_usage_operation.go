// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"startree.ai/cli/client/upsert"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationUpsertEstimateHeapUsageCmd returns a cmd to handle operation estimateHeapUsage
func makeOperationUpsertEstimateHeapUsageCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "estimateHeapUsage",
		Short: `This API returns the estimated heap usage based on primary key column stats. This allows us to estimate table size before onboarding.`,
		RunE:  runOperationUpsertEstimateHeapUsage,
	}

	if err := registerOperationUpsertEstimateHeapUsageParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationUpsertEstimateHeapUsage uses cmd flags to call endpoint api
func runOperationUpsertEstimateHeapUsage(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := upsert.NewEstimateHeapUsageParams()
	if err, _ := retrieveOperationUpsertEstimateHeapUsageBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUpsertEstimateHeapUsageCardinalityFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUpsertEstimateHeapUsageNumPartitionsFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUpsertEstimateHeapUsagePrimaryKeySizeFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationUpsertEstimateHeapUsageResult(appCli.Upsert.EstimateHeapUsage(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationUpsertEstimateHeapUsageParamFlags registers all flags needed to fill params
func registerOperationUpsertEstimateHeapUsageParamFlags(cmd *cobra.Command) error {
	if err := registerOperationUpsertEstimateHeapUsageBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUpsertEstimateHeapUsageCardinalityParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUpsertEstimateHeapUsageNumPartitionsParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUpsertEstimateHeapUsagePrimaryKeySizeParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationUpsertEstimateHeapUsageBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	bodyDescription := ``

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	var bodyFlagDefault string

	_ = cmd.PersistentFlags().String(bodyFlagName, bodyFlagDefault, bodyDescription)

	return nil
}
func registerOperationUpsertEstimateHeapUsageCardinalityParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	cardinalityDescription := `Required. cardinality`

	var cardinalityFlagName string
	if cmdPrefix == "" {
		cardinalityFlagName = "cardinality"
	} else {
		cardinalityFlagName = fmt.Sprintf("%v.cardinality", cmdPrefix)
	}

	var cardinalityFlagDefault int64

	_ = cmd.PersistentFlags().Int64(cardinalityFlagName, cardinalityFlagDefault, cardinalityDescription)

	return nil
}
func registerOperationUpsertEstimateHeapUsageNumPartitionsParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	numPartitionsDescription := `numPartitions`

	var numPartitionsFlagName string
	if cmdPrefix == "" {
		numPartitionsFlagName = "numPartitions"
	} else {
		numPartitionsFlagName = fmt.Sprintf("%v.numPartitions", cmdPrefix)
	}

	var numPartitionsFlagDefault int32 = -1

	_ = cmd.PersistentFlags().Int32(numPartitionsFlagName, numPartitionsFlagDefault, numPartitionsDescription)

	return nil
}
func registerOperationUpsertEstimateHeapUsagePrimaryKeySizeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	primaryKeySizeDescription := `primaryKeySize`

	var primaryKeySizeFlagName string
	if cmdPrefix == "" {
		primaryKeySizeFlagName = "primaryKeySize"
	} else {
		primaryKeySizeFlagName = fmt.Sprintf("%v.primaryKeySize", cmdPrefix)
	}

	var primaryKeySizeFlagDefault int32 = -1

	_ = cmd.PersistentFlags().Int32(primaryKeySizeFlagName, primaryKeySizeFlagDefault, primaryKeySizeDescription)

	return nil
}

func retrieveOperationUpsertEstimateHeapUsageBodyFlag(m *upsert.EstimateHeapUsageParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {

		var bodyFlagName string
		if cmdPrefix == "" {
			bodyFlagName = "body"
		} else {
			bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
		}

		bodyFlagValue, err := cmd.Flags().GetString(bodyFlagName)
		if err != nil {
			return err, false
		}
		m.Body = bodyFlagValue

	}
	return nil, retAdded
}
func retrieveOperationUpsertEstimateHeapUsageCardinalityFlag(m *upsert.EstimateHeapUsageParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("cardinality") {

		var cardinalityFlagName string
		if cmdPrefix == "" {
			cardinalityFlagName = "cardinality"
		} else {
			cardinalityFlagName = fmt.Sprintf("%v.cardinality", cmdPrefix)
		}

		cardinalityFlagValue, err := cmd.Flags().GetInt64(cardinalityFlagName)
		if err != nil {
			return err, false
		}
		m.Cardinality = cardinalityFlagValue

	}
	return nil, retAdded
}
func retrieveOperationUpsertEstimateHeapUsageNumPartitionsFlag(m *upsert.EstimateHeapUsageParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("numPartitions") {

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
		m.NumPartitions = &numPartitionsFlagValue

	}
	return nil, retAdded
}
func retrieveOperationUpsertEstimateHeapUsagePrimaryKeySizeFlag(m *upsert.EstimateHeapUsageParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("primaryKeySize") {

		var primaryKeySizeFlagName string
		if cmdPrefix == "" {
			primaryKeySizeFlagName = "primaryKeySize"
		} else {
			primaryKeySizeFlagName = fmt.Sprintf("%v.primaryKeySize", cmdPrefix)
		}

		primaryKeySizeFlagValue, err := cmd.Flags().GetInt32(primaryKeySizeFlagName)
		if err != nil {
			return err, false
		}
		m.PrimaryKeySize = &primaryKeySizeFlagValue

	}
	return nil, retAdded
}

// parseOperationUpsertEstimateHeapUsageResult parses request result and return the string content
func parseOperationUpsertEstimateHeapUsageResult(resp0 *upsert.EstimateHeapUsageOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*upsert.EstimateHeapUsageOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr := fmt.Sprintf("%v", resp0.Payload)
		return string(msgStr), nil
	}

	return "", nil
}
