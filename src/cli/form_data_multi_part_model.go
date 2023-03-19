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

// Schema cli for FormDataMultiPart

// register flags to command
func registerModelFormDataMultiPartFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerFormDataMultiPartBodyParts(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFormDataMultiPartContentDisposition(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFormDataMultiPartEntity(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFormDataMultiPartFields(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFormDataMultiPartHeaders(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFormDataMultiPartMediaType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFormDataMultiPartMessageBodyWorkers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFormDataMultiPartParameterizedHeaders(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFormDataMultiPartParent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFormDataMultiPartProviders(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerFormDataMultiPartBodyParts(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: bodyParts []*BodyPart array type is not supported by go-swagger cli yet

	return nil
}

func registerFormDataMultiPartContentDisposition(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerFormDataMultiPartEntity(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: entity interface{} map type is not supported by go-swagger cli yet

	return nil
}

func registerFormDataMultiPartFields(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: fields map[string][]FormDataBodyPart map type is not supported by go-swagger cli yet

	return nil
}

func registerFormDataMultiPartHeaders(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: headers map[string][]string map type is not supported by go-swagger cli yet

	return nil
}

func registerFormDataMultiPartMediaType(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerFormDataMultiPartMessageBodyWorkers(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: messageBodyWorkers MessageBodyWorkers map type is not supported by go-swagger cli yet

	return nil
}

func registerFormDataMultiPartParameterizedHeaders(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: parameterizedHeaders map[string][]ParameterizedHeader map type is not supported by go-swagger cli yet

	return nil
}

func registerFormDataMultiPartParent(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerFormDataMultiPartProviders(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: providers Providers map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelFormDataMultiPartFlags(depth int, m *models.FormDataMultiPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, bodyPartsAdded := retrieveFormDataMultiPartBodyPartsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || bodyPartsAdded

	err, contentDispositionAdded := retrieveFormDataMultiPartContentDispositionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || contentDispositionAdded

	err, entityAdded := retrieveFormDataMultiPartEntityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || entityAdded

	err, fieldsAdded := retrieveFormDataMultiPartFieldsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || fieldsAdded

	err, headersAdded := retrieveFormDataMultiPartHeadersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || headersAdded

	err, mediaTypeAdded := retrieveFormDataMultiPartMediaTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || mediaTypeAdded

	err, messageBodyWorkersAdded := retrieveFormDataMultiPartMessageBodyWorkersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || messageBodyWorkersAdded

	err, parameterizedHeadersAdded := retrieveFormDataMultiPartParameterizedHeadersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || parameterizedHeadersAdded

	err, parentAdded := retrieveFormDataMultiPartParentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || parentAdded

	err, providersAdded := retrieveFormDataMultiPartProvidersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || providersAdded

	return nil, retAdded
}

func retrieveFormDataMultiPartBodyPartsFlags(depth int, m *models.FormDataMultiPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	bodyPartsFlagName := fmt.Sprintf("%v.bodyParts", cmdPrefix)
	if cmd.Flags().Changed(bodyPartsFlagName) {
		// warning: bodyParts array type []*BodyPart is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveFormDataMultiPartContentDispositionFlags(depth int, m *models.FormDataMultiPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveFormDataMultiPartEntityFlags(depth int, m *models.FormDataMultiPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveFormDataMultiPartFieldsFlags(depth int, m *models.FormDataMultiPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	fieldsFlagName := fmt.Sprintf("%v.fields", cmdPrefix)
	if cmd.Flags().Changed(fieldsFlagName) {
		// warning: fields map type map[string][]FormDataBodyPart is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveFormDataMultiPartHeadersFlags(depth int, m *models.FormDataMultiPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveFormDataMultiPartMediaTypeFlags(depth int, m *models.FormDataMultiPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveFormDataMultiPartMessageBodyWorkersFlags(depth int, m *models.FormDataMultiPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveFormDataMultiPartParameterizedHeadersFlags(depth int, m *models.FormDataMultiPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveFormDataMultiPartParentFlags(depth int, m *models.FormDataMultiPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveFormDataMultiPartProvidersFlags(depth int, m *models.FormDataMultiPart, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
