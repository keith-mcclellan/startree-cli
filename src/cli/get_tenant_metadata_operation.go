// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"startree.ai/cli/client/tenant"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationTenantGetTenantMetadataCmd returns a cmd to handle operation getTenantMetadata
func makeOperationTenantGetTenantMetadataCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getTenantMetadata",
		Short: ``,
		RunE:  runOperationTenantGetTenantMetadata,
	}

	if err := registerOperationTenantGetTenantMetadataParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTenantGetTenantMetadata uses cmd flags to call endpoint api
func runOperationTenantGetTenantMetadata(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := tenant.NewGetTenantMetadataParams()
	if err, _ := retrieveOperationTenantGetTenantMetadataTenantNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTenantGetTenantMetadataTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationTenantGetTenantMetadataResult(appCli.Tenant.GetTenantMetadata(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationTenantGetTenantMetadataParamFlags registers all flags needed to fill params
func registerOperationTenantGetTenantMetadataParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTenantGetTenantMetadataTenantNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTenantGetTenantMetadataTypeParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTenantGetTenantMetadataTenantNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	tenantNameDescription := `Required. Tenant name`

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
func registerOperationTenantGetTenantMetadataTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	typeDescription := `Enum: ["SERVER","BROKER"]. tenant type`

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	if err := cmd.RegisterFlagCompletionFunc(typeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["SERVER","BROKER"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func retrieveOperationTenantGetTenantMetadataTenantNameFlag(m *tenant.GetTenantMetadataParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("tenantName") {

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

	}
	return nil, retAdded
}
func retrieveOperationTenantGetTenantMetadataTypeFlag(m *tenant.GetTenantMetadataParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("type") {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = &typeFlagValue

	}
	return nil, retAdded
}

// parseOperationTenantGetTenantMetadataResult parses request result and return the string content
func parseOperationTenantGetTenantMetadataResult(resp0 *tenant.GetTenantMetadataOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*tenant.GetTenantMetadataOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning getTenantMetadataNotFound is not supported

		// Non schema case: warning getTenantMetadataInternalServerError is not supported

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
