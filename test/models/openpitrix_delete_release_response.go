// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDeleteReleaseResponse openpitrix delete release response
// swagger:model openpitrixDeleteReleaseResponse
type OpenpitrixDeleteReleaseResponse struct {

	// release name
	ReleaseName string `json:"release_name,omitempty"`
}

// Validate validates this openpitrix delete release response
func (m *OpenpitrixDeleteReleaseResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDeleteReleaseResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDeleteReleaseResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDeleteReleaseResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
