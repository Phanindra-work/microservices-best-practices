// Code generated by go-swagger; DO NOT EDIT.

package apimodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonStudents Register Student
//
// swagger:model common_students
type CommonStudents struct {

	// students
	Students []*Student `json:"students"`
}

// Validate validates this common students
func (m *CommonStudents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStudents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonStudents) validateStudents(formats strfmt.Registry) error {

	if swag.IsZero(m.Students) { // not required
		return nil
	}

	for i := 0; i < len(m.Students); i++ {
		if swag.IsZero(m.Students[i]) { // not required
			continue
		}

		if m.Students[i] != nil {
			if err := m.Students[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("students" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommonStudents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonStudents) UnmarshalBinary(b []byte) error {
	var res CommonStudents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}