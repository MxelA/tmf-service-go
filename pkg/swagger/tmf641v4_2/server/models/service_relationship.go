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

// ServiceRelationship service relationship
//
// swagger:model ServiceRelationship
type ServiceRelationship struct {

	// When sub-classing, this defines the super-class
	AtBaseType string `json:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	// Format: uri
	AtSchemaLocation strfmt.URI `json:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class Extensible name
	AtType string `json:"@type,omitempty"`

	// relationship type
	// Required: true
	RelationshipType *string `json:"relationshipType"`

	// service
	Service *ServiceRefOrValue `json:"service,omitempty"`

	// service relationship characteristic
	ServiceRelationshipCharacteristic []*Characteristic `json:"serviceRelationshipCharacteristic"`
}

// Validate validates this service relationship
func (m *ServiceRelationship) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAtSchemaLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationshipType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceRelationshipCharacteristic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceRelationship) validateAtSchemaLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.AtSchemaLocation) { // not required
		return nil
	}

	if err := validate.FormatOf("@schemaLocation", "body", "uri", m.AtSchemaLocation.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServiceRelationship) validateRelationshipType(formats strfmt.Registry) error {

	if err := validate.Required("relationshipType", "body", m.RelationshipType); err != nil {
		return err
	}

	return nil
}

func (m *ServiceRelationship) validateService(formats strfmt.Registry) error {
	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceRelationship) validateServiceRelationshipCharacteristic(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceRelationshipCharacteristic) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceRelationshipCharacteristic); i++ {
		if swag.IsZero(m.ServiceRelationshipCharacteristic[i]) { // not required
			continue
		}

		if m.ServiceRelationshipCharacteristic[i] != nil {
			if err := m.ServiceRelationshipCharacteristic[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceRelationshipCharacteristic" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this service relationship based on the context it is used
func (m *ServiceRelationship) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceRelationshipCharacteristic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceRelationship) contextValidateService(ctx context.Context, formats strfmt.Registry) error {

	if m.Service != nil {

		if swag.IsZero(m.Service) { // not required
			return nil
		}

		if err := m.Service.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceRelationship) contextValidateServiceRelationshipCharacteristic(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServiceRelationshipCharacteristic); i++ {

		if m.ServiceRelationshipCharacteristic[i] != nil {

			if swag.IsZero(m.ServiceRelationshipCharacteristic[i]) { // not required
				return nil
			}

			if err := m.ServiceRelationshipCharacteristic[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceRelationshipCharacteristic" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceRelationship) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceRelationship) UnmarshalBinary(b []byte) error {
	var res ServiceRelationship
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}