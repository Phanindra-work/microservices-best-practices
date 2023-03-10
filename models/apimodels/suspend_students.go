// Code generated by go-swagger; DO NOT EDIT.

package apimodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SuspendStudents Suspend Students
//
// swagger:model suspend_students
type SuspendStudents struct {

	// student id
	StudentID string `json:"student_id,omitempty"`
}

// Validate validates this suspend students
func (m *SuspendStudents) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SuspendStudents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SuspendStudents) UnmarshalBinary(b []byte) error {
	var res SuspendStudents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
