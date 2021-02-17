// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserGetResponse user get response
//
// swagger:model UserGetResponse
type UserGetResponse struct {

	// ユーザ名
	Name string `json:"name,omitempty"`
}

// Validate validates this user get response
func (m *UserGetResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user get response based on context it is used
func (m *UserGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserGetResponse) UnmarshalBinary(b []byte) error {
	var res UserGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}