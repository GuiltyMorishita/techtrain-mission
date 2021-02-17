// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserUpdateRequest user update request
//
// swagger:model UserUpdateRequest
type UserUpdateRequest struct {

	// ユーザ名
	Name string `json:"name,omitempty"`
}

// Validate validates this user update request
func (m *UserUpdateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user update request based on context it is used
func (m *UserUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserUpdateRequest) UnmarshalBinary(b []byte) error {
	var res UserUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
