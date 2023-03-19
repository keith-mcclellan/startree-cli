// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FormDataMultiPart form data multi part
//
// swagger:model FormDataMultiPart
type FormDataMultiPart struct {

	// body parts
	BodyParts []*BodyPart `json:"bodyParts"`

	// content disposition
	ContentDisposition *ContentDisposition `json:"contentDisposition,omitempty"`

	// entity
	Entity interface{} `json:"entity,omitempty"`

	// fields
	Fields map[string][]FormDataBodyPart `json:"fields,omitempty"`

	// headers
	Headers map[string][]string `json:"headers,omitempty"`

	// media type
	MediaType *MediaType `json:"mediaType,omitempty"`

	// message body workers
	MessageBodyWorkers MessageBodyWorkers `json:"messageBodyWorkers,omitempty"`

	// parameterized headers
	ParameterizedHeaders map[string][]ParameterizedHeader `json:"parameterizedHeaders,omitempty"`

	// parent
	Parent *MultiPart `json:"parent,omitempty"`

	// providers
	Providers Providers `json:"providers,omitempty"`
}

// Validate validates this form data multi part
func (m *FormDataMultiPart) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBodyParts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentDisposition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameterizedHeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FormDataMultiPart) validateBodyParts(formats strfmt.Registry) error {
	if swag.IsZero(m.BodyParts) { // not required
		return nil
	}

	for i := 0; i < len(m.BodyParts); i++ {
		if swag.IsZero(m.BodyParts[i]) { // not required
			continue
		}

		if m.BodyParts[i] != nil {
			if err := m.BodyParts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bodyParts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bodyParts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FormDataMultiPart) validateContentDisposition(formats strfmt.Registry) error {
	if swag.IsZero(m.ContentDisposition) { // not required
		return nil
	}

	if m.ContentDisposition != nil {
		if err := m.ContentDisposition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contentDisposition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contentDisposition")
			}
			return err
		}
	}

	return nil
}

func (m *FormDataMultiPart) validateFields(formats strfmt.Registry) error {
	if swag.IsZero(m.Fields) { // not required
		return nil
	}

	for k := range m.Fields {

		if err := validate.Required("fields"+"."+k, "body", m.Fields[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.Fields[k]); i++ {

			if err := m.Fields[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fields" + "." + k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fields" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *FormDataMultiPart) validateMediaType(formats strfmt.Registry) error {
	if swag.IsZero(m.MediaType) { // not required
		return nil
	}

	if m.MediaType != nil {
		if err := m.MediaType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mediaType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mediaType")
			}
			return err
		}
	}

	return nil
}

func (m *FormDataMultiPart) validateParameterizedHeaders(formats strfmt.Registry) error {
	if swag.IsZero(m.ParameterizedHeaders) { // not required
		return nil
	}

	for k := range m.ParameterizedHeaders {

		if err := validate.Required("parameterizedHeaders"+"."+k, "body", m.ParameterizedHeaders[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.ParameterizedHeaders[k]); i++ {

			if err := m.ParameterizedHeaders[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameterizedHeaders" + "." + k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameterizedHeaders" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *FormDataMultiPart) validateParent(formats strfmt.Registry) error {
	if swag.IsZero(m.Parent) { // not required
		return nil
	}

	if m.Parent != nil {
		if err := m.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this form data multi part based on the context it is used
func (m *FormDataMultiPart) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBodyParts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContentDisposition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMediaType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParameterizedHeaders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FormDataMultiPart) contextValidateBodyParts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BodyParts); i++ {

		if m.BodyParts[i] != nil {
			if err := m.BodyParts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bodyParts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bodyParts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FormDataMultiPart) contextValidateContentDisposition(ctx context.Context, formats strfmt.Registry) error {

	if m.ContentDisposition != nil {
		if err := m.ContentDisposition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contentDisposition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contentDisposition")
			}
			return err
		}
	}

	return nil
}

func (m *FormDataMultiPart) contextValidateFields(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Fields {

		for i := 0; i < len(m.Fields[k]); i++ {

			if err := m.Fields[k][i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fields" + "." + k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fields" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *FormDataMultiPart) contextValidateMediaType(ctx context.Context, formats strfmt.Registry) error {

	if m.MediaType != nil {
		if err := m.MediaType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mediaType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mediaType")
			}
			return err
		}
	}

	return nil
}

func (m *FormDataMultiPart) contextValidateParameterizedHeaders(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.ParameterizedHeaders {

		for i := 0; i < len(m.ParameterizedHeaders[k]); i++ {

			if err := m.ParameterizedHeaders[k][i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameterizedHeaders" + "." + k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameterizedHeaders" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *FormDataMultiPart) contextValidateParent(ctx context.Context, formats strfmt.Registry) error {

	if m.Parent != nil {
		if err := m.Parent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FormDataMultiPart) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FormDataMultiPart) UnmarshalBinary(b []byte) error {
	var res FormDataMultiPart
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}