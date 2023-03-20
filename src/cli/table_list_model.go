// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/spf13/cobra"
	"startree.ai/cli/models"
)

// Schema cli for TableList

// register flags to command
func registerModelTableListFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTableListTables(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTableListTables(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: tables []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTableListFlags(depth int, m *models.TableList, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, tablesAdded := retrieveTableListTablesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tablesAdded

	return nil, retAdded
}

func retrieveTableListTablesFlags(depth int, m *models.TableList, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	tablesFlagName := fmt.Sprintf("%v.tables", cmdPrefix)
	if cmd.Flags().Changed(tablesFlagName) {
		// warning: tables array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
