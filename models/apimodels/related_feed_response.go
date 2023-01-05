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

// RelatedFeedResponse related_feed structure
//
// swagger:model related_feed_response
type RelatedFeedResponse struct {

	// related feed
	RelatedFeed []*RelatedFeed `json:"related_feed"`

	// page state
	PageState *PaginationData `json:"page_state,omitempty"`
}

// Validate validates this related feed response
func (m *RelatedFeedResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRelatedFeed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelatedFeedResponse) validateRelatedFeed(formats strfmt.Registry) error {

	if swag.IsZero(m.RelatedFeed) { // not required
		return nil
	}

	for i := 0; i < len(m.RelatedFeed); i++ {
		if swag.IsZero(m.RelatedFeed[i]) { // not required
			continue
		}

		if m.RelatedFeed[i] != nil {
			if err := m.RelatedFeed[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("related_feed" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RelatedFeedResponse) validatePageState(formats strfmt.Registry) error {

	if swag.IsZero(m.PageState) { // not required
		return nil
	}

	if m.PageState != nil {
		if err := m.PageState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("page_state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RelatedFeedResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelatedFeedResponse) UnmarshalBinary(b []byte) error {
	var res RelatedFeedResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
