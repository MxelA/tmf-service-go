// Code generated by go-swagger; DO NOT EDIT.

package models

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

// ServiceOrderUpdate
// Skipped properties: id,href,orderDate,jeopardyAlert,errorMessage,milestone,@baseType,@schemaLocation,@type,cancellationDate,cancellationReason,category,completionDate,startDate
//
// swagger:model ServiceOrder_Update
type ServiceOrderUpdate struct {

	// A free-text description of the service order
	Description string `json:"description,omitempty"`

	// Expected delivery date amended by the provider
	// Format: date-time
	ExpectedCompletionDate strfmt.DateTime `json:"expectedCompletionDate,omitempty"`

	// ID given by the consumer to facilitate searches
	ExternalID string `json:"externalId,omitempty"`

	// external reference
	ExternalReference []*ExternalReference `json:"externalReference"`

	// Extra-information about the order; e.g. useful to add extra delivery information that could be useful for a human process
	Note []*Note `json:"note"`

	// Contact attached to the order to send back information regarding this order
	NotificationContact string `json:"notificationContact,omitempty"`

	// A list of service orders related to this order (e.g. prerequisite, dependent on)
	OrderRelationship []*ServiceOrderRelationship `json:"orderRelationship"`

	// Can be used by consumers to prioritize orders in a Service Order Management system
	Priority string `json:"priority,omitempty"`

	// A list of parties which are involved in this order and the role they are playing
	RelatedParty []*RelatedParty `json:"relatedParty"`

	// Requested delivery date from the requestors perspective
	// Format: date-time
	RequestedCompletionDate strfmt.DateTime `json:"requestedCompletionDate,omitempty"`

	// Order start date wished by the requestor
	// Format: date-time
	RequestedStartDate strfmt.DateTime `json:"requestedStartDate,omitempty"`

	// A list of service order items to be processed by this order
	ServiceOrderItem []*ServiceOrderItem `json:"serviceOrderItem"`

	// State of the order: described in the state-machine diagram
	State ServiceOrderStateType `json:"state,omitempty"`
}

// Validate validates this service order update
func (m *ServiceOrderUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpectedCompletionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderRelationship(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelatedParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedCompletionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceOrderItem(formats); err != nil {
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

func (m *ServiceOrderUpdate) validateExpectedCompletionDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpectedCompletionDate) { // not required
		return nil
	}

	if err := validate.FormatOf("expectedCompletionDate", "body", "date-time", m.ExpectedCompletionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServiceOrderUpdate) validateExternalReference(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalReference) { // not required
		return nil
	}

	for i := 0; i < len(m.ExternalReference); i++ {
		if swag.IsZero(m.ExternalReference[i]) { // not required
			continue
		}

		if m.ExternalReference[i] != nil {
			if err := m.ExternalReference[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("externalReference" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceOrderUpdate) validateNote(formats strfmt.Registry) error {
	if swag.IsZero(m.Note) { // not required
		return nil
	}

	for i := 0; i < len(m.Note); i++ {
		if swag.IsZero(m.Note[i]) { // not required
			continue
		}

		if m.Note[i] != nil {
			if err := m.Note[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("note" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceOrderUpdate) validateOrderRelationship(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderRelationship) { // not required
		return nil
	}

	for i := 0; i < len(m.OrderRelationship); i++ {
		if swag.IsZero(m.OrderRelationship[i]) { // not required
			continue
		}

		if m.OrderRelationship[i] != nil {
			if err := m.OrderRelationship[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orderRelationship" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceOrderUpdate) validateRelatedParty(formats strfmt.Registry) error {
	if swag.IsZero(m.RelatedParty) { // not required
		return nil
	}

	for i := 0; i < len(m.RelatedParty); i++ {
		if swag.IsZero(m.RelatedParty[i]) { // not required
			continue
		}

		if m.RelatedParty[i] != nil {
			if err := m.RelatedParty[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relatedParty" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceOrderUpdate) validateRequestedCompletionDate(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestedCompletionDate) { // not required
		return nil
	}

	if err := validate.FormatOf("requestedCompletionDate", "body", "date-time", m.RequestedCompletionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServiceOrderUpdate) validateRequestedStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestedStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("requestedStartDate", "body", "date-time", m.RequestedStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServiceOrderUpdate) validateServiceOrderItem(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceOrderItem) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceOrderItem); i++ {
		if swag.IsZero(m.ServiceOrderItem[i]) { // not required
			continue
		}

		if m.ServiceOrderItem[i] != nil {
			if err := m.ServiceOrderItem[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceOrderItem" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceOrderUpdate) validateState(formats strfmt.Registry) error {
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

// ContextValidate validate this service order update based on the context it is used
func (m *ServiceOrderUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExternalReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNote(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrderRelationship(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelatedParty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceOrderItem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOrderUpdate) contextValidateExternalReference(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExternalReference); i++ {

		if m.ExternalReference[i] != nil {

			if swag.IsZero(m.ExternalReference[i]) { // not required
				return nil
			}

			if err := m.ExternalReference[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("externalReference" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceOrderUpdate) contextValidateNote(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Note); i++ {

		if m.Note[i] != nil {

			if swag.IsZero(m.Note[i]) { // not required
				return nil
			}

			if err := m.Note[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("note" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceOrderUpdate) contextValidateOrderRelationship(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OrderRelationship); i++ {

		if m.OrderRelationship[i] != nil {

			if swag.IsZero(m.OrderRelationship[i]) { // not required
				return nil
			}

			if err := m.OrderRelationship[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orderRelationship" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceOrderUpdate) contextValidateRelatedParty(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RelatedParty); i++ {

		if m.RelatedParty[i] != nil {

			if swag.IsZero(m.RelatedParty[i]) { // not required
				return nil
			}

			if err := m.RelatedParty[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relatedParty" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceOrderUpdate) contextValidateServiceOrderItem(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServiceOrderItem); i++ {

		if m.ServiceOrderItem[i] != nil {

			if swag.IsZero(m.ServiceOrderItem[i]) { // not required
				return nil
			}

			if err := m.ServiceOrderItem[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceOrderItem" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceOrderUpdate) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ServiceOrderUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOrderUpdate) UnmarshalBinary(b []byte) error {
	var res ServiceOrderUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
