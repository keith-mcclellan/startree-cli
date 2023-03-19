// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"
	"startree.ai/cli/models"

	"github.com/spf13/cobra"
)

// Schema cli for FormDataBodyPart

// register flags to command
func registerModelFormDataBodyPartFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerFormDataBodyPartContentDisposition(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFormDataBodyPartEntity(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFormDataBodyPartFormDataContentDisposition(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFormDataBodyPartHeaders(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFormDataBodyPartMediaType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFormDataBodyPartMessageBodyWorkers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFormDataBodyPartName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFormDataBodyPartParameterizedHeaders(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFormDataBodyPartParent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFormDataBodyPartProviders(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFormDataBodyPartSimple(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFormDataBodyPartValue(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerFormDataBodyPartContentDisposition(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var contentDispositionFlagName string
	if cmdPrefix == "" {
		contentDispositionFlagName = "contentDisposition"
	} else {
		contentDispositionFlagName = fmt.Sprintf("%v.contentDisposition", cmdPrefix)
	}

	if err := registerModelContentDispositionFlags(depth+1, contentDispositionFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerFormDataBodyPartEntity(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: entity interface{} map type is not supported by go-swagger cli yet

	return nil
}

func registerFormDataBodyPartFormDataContentDisposition(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var formDataContentDispositionFlagName string
	if cmdPrefix == "" {
		formDataContentDispositionFlagName = "formDataContentDisposition"
	} else {
		formDataContentDispositionFlagName = fmt.Sprintf("%v.formDataContentDisposition", cmdPrefix)
	}

	if err := registerModelFormDataContentDispositionFlags(depth+1, formDataContentDispositionFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerFormDataBodyPartHeaders(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: headers map[string][]string map type is not supported by go-swagger cli yet

	return nil
}

func registerFormDataBodyPartMediaType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var mediaTypeFlagName string
	if cmdPrefix == "" {
		mediaTypeFlagName = "mediaType"
	} else {
		mediaTypeFlagName = fmt.Sprintf("%v.mediaType", cmdPrefix)
	}

	if err := registerModelMediaTypeFlags(depth+1, mediaTypeFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerFormDataBodyPartMessageBodyWorkers(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: messageBodyWorkers MessageBodyWorkers map type is not supported by go-swagger cli yet

	return nil
}

func registerFormDataBodyPartName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerFormDataBodyPartParameterizedHeaders(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: parameterizedHeaders map[string][]ParameterizedHeader map type is not supported by go-swagger cli yet

	return nil
}

func registerFormDataBodyPartParent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var parentFlagName string
	if cmdPrefix == "" {
		parentFlagName = "parent"
	} else {
		parentFlagName = fmt.Sprintf("%v.parent", cmdPrefix)
	}

	if err := registerModelMultiPartFlags(depth+1, parentFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerFormDataBodyPartProviders(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: providers Providers map type is not supported by go-swagger cli yet

	return nil
}

func registerFormDataBodyPartSimple(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	simpleDescription := ``

	var simpleFlagName string
	if cmdPrefix == "" {
		simpleFlagName = "simple"
	} else {
		simpleFlagName = fmt.Sprintf("%v.simple", cmdPrefix)
	}

	var simpleFlagDefault bool

	_ = cmd.PersistentFlags().Bool(simpleFlagName, simpleFlagDefault, simpleDescription)

	return nil
}

func registerFormDataBodyPartValue(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	valueDescription := ``

	var valueFlagName string
	if cmdPrefix == "" {
		valueFlagName = "value"
	} else {
		valueFlagName = fmt.Sprintf("%v.value", cmdPrefix)
	}

	var valueFlagDefault string

	_ = cmd.PersistentFlags().String(valueFlagName, valueFlagDefault, valueDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelFormDataBodyPartFlags(depth int, m *models.FormDataBodyPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, contentDispositionAdded := retrieveFormDataBodyPartContentDispositionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || contentDispositionAdded

	err, entityAdded := retrieveFormDataBodyPartEntityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || entityAdded

	err, formDataContentDispositionAdded := retrieveFormDataBodyPartFormDataContentDispositionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || formDataContentDispositionAdded

	err, headersAdded := retrieveFormDataBodyPartHeadersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || headersAdded

	err, mediaTypeAdded := retrieveFormDataBodyPartMediaTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || mediaTypeAdded

	err, messageBodyWorkersAdded := retrieveFormDataBodyPartMessageBodyWorkersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || messageBodyWorkersAdded

	err, nameAdded := retrieveFormDataBodyPartNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, parameterizedHeadersAdded := retrieveFormDataBodyPartParameterizedHeadersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || parameterizedHeadersAdded

	err, parentAdded := retrieveFormDataBodyPartParentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || parentAdded

	err, providersAdded := retrieveFormDataBodyPartProvidersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || providersAdded

	err, simpleAdded := retrieveFormDataBodyPartSimpleFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || simpleAdded

	err, valueAdded := retrieveFormDataBodyPartValueFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || valueAdded

	return nil, retAdded
}

func retrieveFormDataBodyPartContentDispositionFlags(depth int, m *models.FormDataBodyPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	contentDispositionFlagName := fmt.Sprintf("%v.contentDisposition", cmdPrefix)
	if cmd.Flags().Changed(contentDispositionFlagName) {
		// info: complex object contentDisposition ContentDisposition is retrieved outside this Changed() block
	}
	contentDispositionFlagValue := m.ContentDisposition
	if swag.IsZero(contentDispositionFlagValue) {
		contentDispositionFlagValue = &models.ContentDisposition{}
	}

	err, contentDispositionAdded := retrieveModelContentDispositionFlags(depth+1, contentDispositionFlagValue, contentDispositionFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || contentDispositionAdded
	if contentDispositionAdded {
		m.ContentDisposition = contentDispositionFlagValue
	}

	return nil, retAdded
}

func retrieveFormDataBodyPartEntityFlags(depth int, m *models.FormDataBodyPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	entityFlagName := fmt.Sprintf("%v.entity", cmdPrefix)
	if cmd.Flags().Changed(entityFlagName) {
		// warning: entity map type interface{} is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveFormDataBodyPartFormDataContentDispositionFlags(depth int, m *models.FormDataBodyPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	formDataContentDispositionFlagName := fmt.Sprintf("%v.formDataContentDisposition", cmdPrefix)
	if cmd.Flags().Changed(formDataContentDispositionFlagName) {
		// info: complex object formDataContentDisposition FormDataContentDisposition is retrieved outside this Changed() block
	}
	formDataContentDispositionFlagValue := m.FormDataContentDisposition
	if swag.IsZero(formDataContentDispositionFlagValue) {
		formDataContentDispositionFlagValue = &models.FormDataContentDisposition{}
	}

	err, formDataContentDispositionAdded := retrieveModelFormDataContentDispositionFlags(depth+1, formDataContentDispositionFlagValue, formDataContentDispositionFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || formDataContentDispositionAdded
	if formDataContentDispositionAdded {
		m.FormDataContentDisposition = formDataContentDispositionFlagValue
	}

	return nil, retAdded
}

func retrieveFormDataBodyPartHeadersFlags(depth int, m *models.FormDataBodyPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	headersFlagName := fmt.Sprintf("%v.headers", cmdPrefix)
	if cmd.Flags().Changed(headersFlagName) {
		// warning: headers map type map[string][]string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveFormDataBodyPartMediaTypeFlags(depth int, m *models.FormDataBodyPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	mediaTypeFlagName := fmt.Sprintf("%v.mediaType", cmdPrefix)
	if cmd.Flags().Changed(mediaTypeFlagName) {
		// info: complex object mediaType MediaType is retrieved outside this Changed() block
	}
	mediaTypeFlagValue := m.MediaType
	if swag.IsZero(mediaTypeFlagValue) {
		mediaTypeFlagValue = &models.MediaType{}
	}

	err, mediaTypeAdded := retrieveModelMediaTypeFlags(depth+1, mediaTypeFlagValue, mediaTypeFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || mediaTypeAdded
	if mediaTypeAdded {
		m.MediaType = mediaTypeFlagValue
	}

	return nil, retAdded
}

func retrieveFormDataBodyPartMessageBodyWorkersFlags(depth int, m *models.FormDataBodyPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	messageBodyWorkersFlagName := fmt.Sprintf("%v.messageBodyWorkers", cmdPrefix)
	if cmd.Flags().Changed(messageBodyWorkersFlagName) {
		// warning: messageBodyWorkers map type MessageBodyWorkers is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveFormDataBodyPartNameFlags(depth int, m *models.FormDataBodyPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveFormDataBodyPartParameterizedHeadersFlags(depth int, m *models.FormDataBodyPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	parameterizedHeadersFlagName := fmt.Sprintf("%v.parameterizedHeaders", cmdPrefix)
	if cmd.Flags().Changed(parameterizedHeadersFlagName) {
		// warning: parameterizedHeaders map type map[string][]ParameterizedHeader is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveFormDataBodyPartParentFlags(depth int, m *models.FormDataBodyPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	parentFlagName := fmt.Sprintf("%v.parent", cmdPrefix)
	if cmd.Flags().Changed(parentFlagName) {
		// info: complex object parent MultiPart is retrieved outside this Changed() block
	}
	parentFlagValue := m.Parent
	if swag.IsZero(parentFlagValue) {
		parentFlagValue = &models.MultiPart{}
	}

	err, parentAdded := retrieveModelMultiPartFlags(depth+1, parentFlagValue, parentFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || parentAdded
	if parentAdded {
		m.Parent = parentFlagValue
	}

	return nil, retAdded
}

func retrieveFormDataBodyPartProvidersFlags(depth int, m *models.FormDataBodyPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	providersFlagName := fmt.Sprintf("%v.providers", cmdPrefix)
	if cmd.Flags().Changed(providersFlagName) {
		// warning: providers map type Providers is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveFormDataBodyPartSimpleFlags(depth int, m *models.FormDataBodyPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	simpleFlagName := fmt.Sprintf("%v.simple", cmdPrefix)
	if cmd.Flags().Changed(simpleFlagName) {

		var simpleFlagName string
		if cmdPrefix == "" {
			simpleFlagName = "simple"
		} else {
			simpleFlagName = fmt.Sprintf("%v.simple", cmdPrefix)
		}

		simpleFlagValue, err := cmd.Flags().GetBool(simpleFlagName)
		if err != nil {
			return err, false
		}
		m.Simple = simpleFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveFormDataBodyPartValueFlags(depth int, m *models.FormDataBodyPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	valueFlagName := fmt.Sprintf("%v.value", cmdPrefix)
	if cmd.Flags().Changed(valueFlagName) {

		var valueFlagName string
		if cmdPrefix == "" {
			valueFlagName = "value"
		} else {
			valueFlagName = fmt.Sprintf("%v.value", cmdPrefix)
		}

		valueFlagValue, err := cmd.Flags().GetString(valueFlagName)
		if err != nil {
			return err, false
		}
		m.Value = valueFlagValue

		retAdded = true
	}

	return nil, retAdded
}