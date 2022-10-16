// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Cable cable
//
// swagger:model Cable
type Cable struct {

	// a terminations
	ATerminations []*GenericObject `json:"a_terminations"`

	// b terminations
	BTerminations []*GenericObject `json:"b_terminations"`

	// Color
	// Max Length: 6
	// Pattern: ^[0-9a-f]{6}$
	Color string `json:"color,omitempty"`

	// Created
	// Read Only: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	// Max Length: 100
	Label string `json:"label,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	// Length
	Length *float64 `json:"length,omitempty"`

	// length unit
	LengthUnit *CableLengthUnit `json:"length_unit,omitempty"`

	// status
	Status *CableStatus `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags,omitempty"`

	// tenant
	Tenant *NestedTenant `json:"tenant,omitempty"`

	// Type
	// Enum: [cat3 cat5 cat5e cat6 cat6a cat7 cat7a cat8 dac-active dac-passive mrj21-trunk coaxial mmf mmf-om1 mmf-om2 mmf-om3 mmf-om4 mmf-om5 smf smf-os1 smf-os2 aoc power]
	Type string `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this cable
func (m *Cable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateATerminations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBTerminations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLengthUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cable) validateATerminations(formats strfmt.Registry) error {
	if swag.IsZero(m.ATerminations) { // not required
		return nil
	}

	for i := 0; i < len(m.ATerminations); i++ {
		if swag.IsZero(m.ATerminations[i]) { // not required
			continue
		}

		if m.ATerminations[i] != nil {
			if err := m.ATerminations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("a_terminations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("a_terminations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Cable) validateBTerminations(formats strfmt.Registry) error {
	if swag.IsZero(m.BTerminations) { // not required
		return nil
	}

	for i := 0; i < len(m.BTerminations); i++ {
		if swag.IsZero(m.BTerminations[i]) { // not required
			continue
		}

		if m.BTerminations[i] != nil {
			if err := m.BTerminations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("b_terminations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("b_terminations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Cable) validateColor(formats strfmt.Registry) error {
	if swag.IsZero(m.Color) { // not required
		return nil
	}

	if err := validate.MaxLength("color", "body", m.Color, 6); err != nil {
		return err
	}

	if err := validate.Pattern("color", "body", m.Color, `^[0-9a-f]{6}$`); err != nil {
		return err
	}

	return nil
}

func (m *Cable) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Cable) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 100); err != nil {
		return err
	}

	return nil
}

func (m *Cable) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Cable) validateLengthUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.LengthUnit) { // not required
		return nil
	}

	if m.LengthUnit != nil {
		if err := m.LengthUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("length_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("length_unit")
			}
			return err
		}
	}

	return nil
}

func (m *Cable) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *Cable) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Cable) validateTenant(formats strfmt.Registry) error {
	if swag.IsZero(m.Tenant) { // not required
		return nil
	}

	if m.Tenant != nil {
		if err := m.Tenant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

var cableTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cat3","cat5","cat5e","cat6","cat6a","cat7","cat7a","cat8","dac-active","dac-passive","mrj21-trunk","coaxial","mmf","mmf-om1","mmf-om2","mmf-om3","mmf-om4","mmf-om5","smf","smf-os1","smf-os2","aoc","power"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cableTypeTypePropEnum = append(cableTypeTypePropEnum, v)
	}
}

const (

	// CableTypeCat3 captures enum value "cat3"
	CableTypeCat3 string = "cat3"

	// CableTypeCat5 captures enum value "cat5"
	CableTypeCat5 string = "cat5"

	// CableTypeCat5e captures enum value "cat5e"
	CableTypeCat5e string = "cat5e"

	// CableTypeCat6 captures enum value "cat6"
	CableTypeCat6 string = "cat6"

	// CableTypeCat6a captures enum value "cat6a"
	CableTypeCat6a string = "cat6a"

	// CableTypeCat7 captures enum value "cat7"
	CableTypeCat7 string = "cat7"

	// CableTypeCat7a captures enum value "cat7a"
	CableTypeCat7a string = "cat7a"

	// CableTypeCat8 captures enum value "cat8"
	CableTypeCat8 string = "cat8"

	// CableTypeDacDashActive captures enum value "dac-active"
	CableTypeDacDashActive string = "dac-active"

	// CableTypeDacDashPassive captures enum value "dac-passive"
	CableTypeDacDashPassive string = "dac-passive"

	// CableTypeMrj21DashTrunk captures enum value "mrj21-trunk"
	CableTypeMrj21DashTrunk string = "mrj21-trunk"

	// CableTypeCoaxial captures enum value "coaxial"
	CableTypeCoaxial string = "coaxial"

	// CableTypeMmf captures enum value "mmf"
	CableTypeMmf string = "mmf"

	// CableTypeMmfDashOm1 captures enum value "mmf-om1"
	CableTypeMmfDashOm1 string = "mmf-om1"

	// CableTypeMmfDashOm2 captures enum value "mmf-om2"
	CableTypeMmfDashOm2 string = "mmf-om2"

	// CableTypeMmfDashOm3 captures enum value "mmf-om3"
	CableTypeMmfDashOm3 string = "mmf-om3"

	// CableTypeMmfDashOm4 captures enum value "mmf-om4"
	CableTypeMmfDashOm4 string = "mmf-om4"

	// CableTypeMmfDashOm5 captures enum value "mmf-om5"
	CableTypeMmfDashOm5 string = "mmf-om5"

	// CableTypeSmf captures enum value "smf"
	CableTypeSmf string = "smf"

	// CableTypeSmfDashOs1 captures enum value "smf-os1"
	CableTypeSmfDashOs1 string = "smf-os1"

	// CableTypeSmfDashOs2 captures enum value "smf-os2"
	CableTypeSmfDashOs2 string = "smf-os2"

	// CableTypeAoc captures enum value "aoc"
	CableTypeAoc string = "aoc"

	// CableTypePower captures enum value "power"
	CableTypePower string = "power"
)

// prop value enum
func (m *Cable) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cableTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Cable) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Cable) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cable based on the context it is used
func (m *Cable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateATerminations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBTerminations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLengthUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTenant(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cable) contextValidateATerminations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ATerminations); i++ {

		if m.ATerminations[i] != nil {
			if err := m.ATerminations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("a_terminations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("a_terminations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Cable) contextValidateBTerminations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BTerminations); i++ {

		if m.BTerminations[i] != nil {
			if err := m.BTerminations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("b_terminations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("b_terminations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Cable) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *Cable) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *Cable) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Cable) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *Cable) contextValidateLengthUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.LengthUnit != nil {
		if err := m.LengthUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("length_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("length_unit")
			}
			return err
		}
	}

	return nil
}

func (m *Cable) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *Cable) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Cable) contextValidateTenant(ctx context.Context, formats strfmt.Registry) error {

	if m.Tenant != nil {
		if err := m.Tenant.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *Cable) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cable) UnmarshalBinary(b []byte) error {
	var res Cable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CableLengthUnit Length unit
//
// swagger:model CableLengthUnit
type CableLengthUnit struct {

	// label
	// Required: true
	// Enum: [Kilometers Meters Centimeters Miles Feet Inches]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [km m cm mi ft in]
	Value *string `json:"value"`
}

// Validate validates this cable length unit
func (m *CableLengthUnit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cableLengthUnitTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Kilometers","Meters","Centimeters","Miles","Feet","Inches"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cableLengthUnitTypeLabelPropEnum = append(cableLengthUnitTypeLabelPropEnum, v)
	}
}

const (

	// CableLengthUnitLabelKilometers captures enum value "Kilometers"
	CableLengthUnitLabelKilometers string = "Kilometers"

	// CableLengthUnitLabelMeters captures enum value "Meters"
	CableLengthUnitLabelMeters string = "Meters"

	// CableLengthUnitLabelCentimeters captures enum value "Centimeters"
	CableLengthUnitLabelCentimeters string = "Centimeters"

	// CableLengthUnitLabelMiles captures enum value "Miles"
	CableLengthUnitLabelMiles string = "Miles"

	// CableLengthUnitLabelFeet captures enum value "Feet"
	CableLengthUnitLabelFeet string = "Feet"

	// CableLengthUnitLabelInches captures enum value "Inches"
	CableLengthUnitLabelInches string = "Inches"
)

// prop value enum
func (m *CableLengthUnit) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cableLengthUnitTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CableLengthUnit) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("length_unit"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("length_unit"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var cableLengthUnitTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["km","m","cm","mi","ft","in"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cableLengthUnitTypeValuePropEnum = append(cableLengthUnitTypeValuePropEnum, v)
	}
}

const (

	// CableLengthUnitValueKm captures enum value "km"
	CableLengthUnitValueKm string = "km"

	// CableLengthUnitValueM captures enum value "m"
	CableLengthUnitValueM string = "m"

	// CableLengthUnitValueCm captures enum value "cm"
	CableLengthUnitValueCm string = "cm"

	// CableLengthUnitValueMi captures enum value "mi"
	CableLengthUnitValueMi string = "mi"

	// CableLengthUnitValueFt captures enum value "ft"
	CableLengthUnitValueFt string = "ft"

	// CableLengthUnitValueIn captures enum value "in"
	CableLengthUnitValueIn string = "in"
)

// prop value enum
func (m *CableLengthUnit) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cableLengthUnitTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CableLengthUnit) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("length_unit"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("length_unit"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cable length unit based on context it is used
func (m *CableLengthUnit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CableLengthUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CableLengthUnit) UnmarshalBinary(b []byte) error {
	var res CableLengthUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CableStatus Status
//
// swagger:model CableStatus
type CableStatus struct {

	// label
	// Required: true
	// Enum: [Connected Planned Decommissioning]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [connected planned decommissioning]
	Value *string `json:"value"`
}

// Validate validates this cable status
func (m *CableStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cableStatusTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Connected","Planned","Decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cableStatusTypeLabelPropEnum = append(cableStatusTypeLabelPropEnum, v)
	}
}

const (

	// CableStatusLabelConnected captures enum value "Connected"
	CableStatusLabelConnected string = "Connected"

	// CableStatusLabelPlanned captures enum value "Planned"
	CableStatusLabelPlanned string = "Planned"

	// CableStatusLabelDecommissioning captures enum value "Decommissioning"
	CableStatusLabelDecommissioning string = "Decommissioning"
)

// prop value enum
func (m *CableStatus) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cableStatusTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CableStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("status"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var cableStatusTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["connected","planned","decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cableStatusTypeValuePropEnum = append(cableStatusTypeValuePropEnum, v)
	}
}

const (

	// CableStatusValueConnected captures enum value "connected"
	CableStatusValueConnected string = "connected"

	// CableStatusValuePlanned captures enum value "planned"
	CableStatusValuePlanned string = "planned"

	// CableStatusValueDecommissioning captures enum value "decommissioning"
	CableStatusValueDecommissioning string = "decommissioning"
)

// prop value enum
func (m *CableStatus) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cableStatusTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CableStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("status"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cable status based on context it is used
func (m *CableStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CableStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CableStatus) UnmarshalBinary(b []byte) error {
	var res CableStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
