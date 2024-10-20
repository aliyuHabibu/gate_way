// Code generated by go-swagger; DO NOT EDIT.

package model

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

// PaymentsPagination payments pagination
//
// swagger:model paymentsPagination
type PaymentsPagination struct {

	// cursor
	// Example: 8c857501-e67d-4a0b-98d9-46e461b42c09
	// Required: true
	Cursor string `json:"cursor"`

	// limit
	// Example: 50
	// Required: true
	Limit int64 `json:"limit"`

	// results
	// Required: true
	Results []*Payment `json:"results"`
}

// Validate validates this payments pagination
func (m *PaymentsPagination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCursor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentsPagination) validateCursor(formats strfmt.Registry) error {

	if err := validate.RequiredString("cursor", "body", m.Cursor); err != nil {
		return err
	}

	return nil
}

func (m *PaymentsPagination) validateLimit(formats strfmt.Registry) error {

	if err := validate.Required("limit", "body", int64(m.Limit)); err != nil {
		return err
	}

	return nil
}

func (m *PaymentsPagination) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("results", "body", m.Results); err != nil {
		return err
	}

	for i := 0; i < len(m.Results); i++ {
		if swag.IsZero(m.Results[i]) { // not required
			continue
		}

		if m.Results[i] != nil {
			if err := m.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this payments pagination based on the context it is used
func (m *PaymentsPagination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentsPagination) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Results); i++ {

		if m.Results[i] != nil {
			if err := m.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentsPagination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentsPagination) UnmarshalBinary(b []byte) error {
	var res PaymentsPagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
