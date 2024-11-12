// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CancelOrder A Order cancel is a type of task which  can  be used to place a request to cancel an order
//
// swagger:model CancelOrder
type CancelOrder struct {

	// When sub-classing, this defines the super-class
	AtBaseType string `json:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	// Format: uri
	AtSchemaLocation strfmt.URI `json:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class Extensible name
	AtType string `json:"@type,omitempty"`

	// Reason why the order is cancelled.
	CancellationReason string `json:"cancellationReason,omitempty"`

	// Date when the order is cancelled.
	// Format: date-time
	EffectiveCancellationDate strfmt.DateTime `json:"effectiveCancellationDate,omitempty"`

	// Hyperlink to access the cancellation request
	Href string `json:"href,omitempty"`

	// id of the cancellation request (this is not an order id)
	ID string `json:"id,omitempty"`

	// Date when the submitter wants the order to be cancelled
	// Format: date-time
	RequestedCancellationDate strfmt.DateTime `json:"requestedCancellationDate,omitempty"`

	// Tracks the lifecycle status of the cancellation request, such as Acknowledged, Rejected, InProgress, Pending and so on.
	State TaskStateType `json:"state,omitempty"`
}

// Validate validates this cancel order
func (m *CancelOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAtSchemaLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectiveCancellationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedCancellationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CancelOrder) validateAtSchemaLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.AtSchemaLocation) { // not required
		return nil
	}

	if err := validate.FormatOf("@schemaLocation", "body", "uri", m.AtSchemaLocation.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CancelOrder) validateEffectiveCancellationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EffectiveCancellationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("effectiveCancellationDate", "body", "date-time", m.EffectiveCancellationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CancelOrder) validateRequestedCancellationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestedCancellationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("requestedCancellationDate", "body", "date-time", m.RequestedCancellationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CancelOrder) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// ContextValidate validate this cancel order based on the context it is used
func (m *CancelOrder) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CancelOrder) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CancelOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CancelOrder) UnmarshalBinary(b []byte) error {
	var res CancelOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}