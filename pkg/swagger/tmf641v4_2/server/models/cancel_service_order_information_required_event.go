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

// CancelServiceOrderInformationRequiredEvent The notification data structure
//
// swagger:model CancelServiceOrderInformationRequiredEvent
type CancelServiceOrderInformationRequiredEvent struct {

	// The correlation id for this event.
	CorrelationID string `json:"correlationId,omitempty"`

	// An explnatory of the event.
	Description string `json:"description,omitempty"`

	// The domain of the event.
	Domain string `json:"domain,omitempty"`

	// The event payload linked to the involved resource object
	Event *CancelServiceOrderInformationRequiredEventPayload `json:"event,omitempty"`

	// The identifier of the notification.
	EventID string `json:"eventId,omitempty"`

	// Time of the event occurrence.
	// Format: date-time
	EventTime strfmt.DateTime `json:"eventTime,omitempty"`

	// The type of the notification.
	EventType string `json:"eventType,omitempty"`

	// The path identifying the object field concerned by this notification.
	FieldPath string `json:"fieldPath,omitempty"`

	// A priority.
	Priority string `json:"priority,omitempty"`

	// The time the event occured.
	// Format: date-time
	TimeOcurred strfmt.DateTime `json:"timeOcurred,omitempty"`

	// The title of the event.
	Title string `json:"title,omitempty"`
}

// Validate validates this cancel service order information required event
func (m *CancelServiceOrderInformationRequiredEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeOcurred(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CancelServiceOrderInformationRequiredEvent) validateEvent(formats strfmt.Registry) error {
	if swag.IsZero(m.Event) { // not required
		return nil
	}

	if m.Event != nil {
		if err := m.Event.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event")
			}
			return err
		}
	}

	return nil
}

func (m *CancelServiceOrderInformationRequiredEvent) validateEventTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EventTime) { // not required
		return nil
	}

	if err := validate.FormatOf("eventTime", "body", "date-time", m.EventTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CancelServiceOrderInformationRequiredEvent) validateTimeOcurred(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeOcurred) { // not required
		return nil
	}

	if err := validate.FormatOf("timeOcurred", "body", "date-time", m.TimeOcurred.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cancel service order information required event based on the context it is used
func (m *CancelServiceOrderInformationRequiredEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEvent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CancelServiceOrderInformationRequiredEvent) contextValidateEvent(ctx context.Context, formats strfmt.Registry) error {

	if m.Event != nil {

		if swag.IsZero(m.Event) { // not required
			return nil
		}

		if err := m.Event.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CancelServiceOrderInformationRequiredEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CancelServiceOrderInformationRequiredEvent) UnmarshalBinary(b []byte) error {
	var res CancelServiceOrderInformationRequiredEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
