// Code generated by go-swagger; DO NOT EDIT.

package apimodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FeedPost feed post
//
// swagger:model feed_post
type FeedPost struct {

	// actor
	Actor string `json:"actor,omitempty"`

	// verb
	Verb string `json:"verb,omitempty"`

	// object
	Object string `json:"object,omitempty"`

	// target
	Target string `json:"target,omitempty"`

	// createdat
	// Format: date-time
	Createdat strfmt.DateTime `json:"createdat,omitempty"`
}

// Validate validates this feed post
func (m *FeedPost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeedPost) validateCreatedat(formats strfmt.Registry) error {

	if swag.IsZero(m.Createdat) { // not required
		return nil
	}

	if err := validate.FormatOf("createdat", "body", "date-time", m.Createdat.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FeedPost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeedPost) UnmarshalBinary(b []byte) error {
	var res FeedPost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
