// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WritePayload write payload
//
// swagger:model WritePayload
type WritePayload struct {

	// column names
	ColumnNames []string `json:"columnNames"`

	// rows
	Rows [][]interface{} `json:"rows"`

	// values
	Values []map[string]interface{} `json:"values"`
}

// Validate validates this write payload
func (m *WritePayload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this write payload based on context it is used
func (m *WritePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WritePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritePayload) UnmarshalBinary(b []byte) error {
	var res WritePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}