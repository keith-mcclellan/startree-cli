// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"startree.ai/cli/client/schema"

	"github.com/spf13/cobra"
)

// makeOperationSchemaDeleteSchemaCmd returns a cmd to handle operation deleteSchema
func makeOperationSchemaDeleteSchemaCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteSchema",
		Short: `Deletes a schema by name`,
		RunE:  runOperationSchemaDeleteSchema,
	}

	if err := registerOperationSchemaDeleteSchemaParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSchemaDeleteSchema uses cmd flags to call endpoint api
func runOperationSchemaDeleteSchema(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := schema.NewDeleteSchemaParams()
	if err, _ := retrieveOperationSchemaDeleteSchemaSchemaNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSchemaDeleteSchemaResult(appCli.Schema.DeleteSchema(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSchemaDeleteSchemaParamFlags registers all flags needed to fill params
func registerOperationSchemaDeleteSchemaParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSchemaDeleteSchemaSchemaNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSchemaDeleteSchemaSchemaNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	schemaNameDescription := `Required. Schema name`

	var schemaNameFlagName string
	if cmdPrefix == "" {
		schemaNameFlagName = "schemaName"
	} else {
		schemaNameFlagName = fmt.Sprintf("%v.schemaName", cmdPrefix)
	}

	var schemaNameFlagDefault string

	_ = cmd.PersistentFlags().String(schemaNameFlagName, schemaNameFlagDefault, schemaNameDescription)

	return nil
}

func retrieveOperationSchemaDeleteSchemaSchemaNameFlag(m *schema.DeleteSchemaParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("schemaName") {

		var schemaNameFlagName string
		if cmdPrefix == "" {
			schemaNameFlagName = "schemaName"
		} else {
			schemaNameFlagName = fmt.Sprintf("%v.schemaName", cmdPrefix)
		}

		schemaNameFlagValue, err := cmd.Flags().GetString(schemaNameFlagName)
		if err != nil {
			return err, false
		}
		m.SchemaName = schemaNameFlagValue

	}
	return nil, retAdded
}

// parseOperationSchemaDeleteSchemaResult parses request result and return the string content
func parseOperationSchemaDeleteSchemaResult(resp0 *schema.DeleteSchemaOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning deleteSchemaOK is not supported

		// Non schema case: warning deleteSchemaNotFound is not supported

		// Non schema case: warning deleteSchemaConflict is not supported

		// Non schema case: warning deleteSchemaInternalServerError is not supported

		return "", respErr
	}

	// warning: non schema response deleteSchemaOK is not supported by go-swagger cli yet.

	return "", nil
}