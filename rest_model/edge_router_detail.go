// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package rest_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EdgeRouterDetail A detail edge router resource
// Example: {"_links":{"edge-router-policies":{"href":"./edge-routers/b0766b8d-bd1a-4d28-8415-639b29d3c83d/edge-routers"},"self":{"href":"./edge-routers/b0766b8d-bd1a-4d28-8415-639b29d3c83d"}},"createdAt":"2020-03-16T17:13:31.5807454Z","enrollmentCreatedAt":"2020-03-16T17:13:31.5777637Z","enrollmentExpiresAt":"2020-03-16T17:18:31.5777637Z","enrollmentJwt":"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbSI6ImVyb3R0IiwiZXhwIjoxNTg0Mzc5MTExLCJpc3MiOiJodHRwczovL 2xvY2FsaG9zdDoxMjgwIiwianRpIjoiMzBhMWYwZWEtZDM5Yi00YWFlLWI4NTItMzA0Y2YxYzMwZDFmIiwic3ViIjoiYjA3NjZiOGQtYmQxYS00ZDI 4LTg0MTUtNjM5YjI5ZDNjODNkIn0.UsyQhCPORQ5tQnYWY7S88LNvV9iFS5Hy-P4aJaClZzEICobKgnQoyQblJcdMvk3cGKwyFqAnQtt0tDZkb8tHz Vqyv6bilHcAFuMRrdwXRqdXquabSN5geu2qBUnyzL7Mf2X85if8sbMida6snB4oLZsVRF3CRn4ODBJdeiVJ_Z4rgD-zW2IwtXPApT7ALyiiw2cN4EH 8pqQ7tpZKqztE0PGEbBQFPGKUFnm7oXyvSUo17EsFJUv5gUlBzfKKGolh5io4ptp22HZrqsqSnqDSOnYEZHonr5Yljuwiktrlh-JKiK6GGns5OAJMP dO9lgM4yHSpF2ILbqhWMV93Y3zMOg","enrollmentToken":"30a1f0ea-d39b-4aae-b852-304cf1c30d1f","fingerprint":null,"hostname":"","id":"b0766b8d-bd1a-4d28-8415-639b29d3c83d","isOnline":false,"isTunnelerEnabled":false,"isVerified":false,"name":"TestRouter-e33c837f-3222-4b40-bcd6-b3458fd5156e","roleAttributes":["eastCoast","sales","test"],"supportedProtocols":{},"tags":{},"updatedAt":"2020-03-16T17:13:31.5807454Z"}
//
// swagger:model edgeRouterDetail
type EdgeRouterDetail struct {
	BaseEntity

	CommonEdgeRouterProperties

	// enrollment created at
	// Format: date-time
	EnrollmentCreatedAt *strfmt.DateTime `json:"enrollmentCreatedAt,omitempty"`

	// enrollment expires at
	// Format: date-time
	EnrollmentExpiresAt *strfmt.DateTime `json:"enrollmentExpiresAt,omitempty"`

	// enrollment jwt
	EnrollmentJwt *string `json:"enrollmentJwt,omitempty"`

	// enrollment token
	EnrollmentToken *string `json:"enrollmentToken,omitempty"`

	// fingerprint
	Fingerprint string `json:"fingerprint,omitempty"`

	// is tunneler enabled
	// Required: true
	IsTunnelerEnabled *bool `json:"isTunnelerEnabled"`

	// is verified
	// Required: true
	IsVerified *bool `json:"isVerified"`

	// role attributes
	// Required: true
	RoleAttributes *Attributes `json:"roleAttributes"`

	// version info
	VersionInfo *VersionInfo `json:"versionInfo,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *EdgeRouterDetail) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseEntity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseEntity = aO0

	// AO1
	var aO1 CommonEdgeRouterProperties
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.CommonEdgeRouterProperties = aO1

	// AO2
	var dataAO2 struct {
		EnrollmentCreatedAt *strfmt.DateTime `json:"enrollmentCreatedAt,omitempty"`

		EnrollmentExpiresAt *strfmt.DateTime `json:"enrollmentExpiresAt,omitempty"`

		EnrollmentJwt *string `json:"enrollmentJwt,omitempty"`

		EnrollmentToken *string `json:"enrollmentToken,omitempty"`

		Fingerprint string `json:"fingerprint,omitempty"`

		IsTunnelerEnabled *bool `json:"isTunnelerEnabled"`

		IsVerified *bool `json:"isVerified"`

		RoleAttributes *Attributes `json:"roleAttributes"`

		VersionInfo *VersionInfo `json:"versionInfo,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO2); err != nil {
		return err
	}

	m.EnrollmentCreatedAt = dataAO2.EnrollmentCreatedAt

	m.EnrollmentExpiresAt = dataAO2.EnrollmentExpiresAt

	m.EnrollmentJwt = dataAO2.EnrollmentJwt

	m.EnrollmentToken = dataAO2.EnrollmentToken

	m.Fingerprint = dataAO2.Fingerprint

	m.IsTunnelerEnabled = dataAO2.IsTunnelerEnabled

	m.IsVerified = dataAO2.IsVerified

	m.RoleAttributes = dataAO2.RoleAttributes

	m.VersionInfo = dataAO2.VersionInfo

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m EdgeRouterDetail) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.BaseEntity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.CommonEdgeRouterProperties)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	var dataAO2 struct {
		EnrollmentCreatedAt *strfmt.DateTime `json:"enrollmentCreatedAt,omitempty"`

		EnrollmentExpiresAt *strfmt.DateTime `json:"enrollmentExpiresAt,omitempty"`

		EnrollmentJwt *string `json:"enrollmentJwt,omitempty"`

		EnrollmentToken *string `json:"enrollmentToken,omitempty"`

		Fingerprint string `json:"fingerprint,omitempty"`

		IsTunnelerEnabled *bool `json:"isTunnelerEnabled"`

		IsVerified *bool `json:"isVerified"`

		RoleAttributes *Attributes `json:"roleAttributes"`

		VersionInfo *VersionInfo `json:"versionInfo,omitempty"`
	}

	dataAO2.EnrollmentCreatedAt = m.EnrollmentCreatedAt

	dataAO2.EnrollmentExpiresAt = m.EnrollmentExpiresAt

	dataAO2.EnrollmentJwt = m.EnrollmentJwt

	dataAO2.EnrollmentToken = m.EnrollmentToken

	dataAO2.Fingerprint = m.Fingerprint

	dataAO2.IsTunnelerEnabled = m.IsTunnelerEnabled

	dataAO2.IsVerified = m.IsVerified

	dataAO2.RoleAttributes = m.RoleAttributes

	dataAO2.VersionInfo = m.VersionInfo

	jsonDataAO2, errAO2 := swag.WriteJSON(dataAO2)
	if errAO2 != nil {
		return nil, errAO2
	}
	_parts = append(_parts, jsonDataAO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this edge router detail
func (m *EdgeRouterDetail) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseEntity
	if err := m.BaseEntity.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with CommonEdgeRouterProperties
	if err := m.CommonEdgeRouterProperties.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnrollmentCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnrollmentExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsTunnelerEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsVerified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeRouterDetail) validateEnrollmentCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.EnrollmentCreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("enrollmentCreatedAt", "body", "date-time", m.EnrollmentCreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EdgeRouterDetail) validateEnrollmentExpiresAt(formats strfmt.Registry) error {

	if swag.IsZero(m.EnrollmentExpiresAt) { // not required
		return nil
	}

	if err := validate.FormatOf("enrollmentExpiresAt", "body", "date-time", m.EnrollmentExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EdgeRouterDetail) validateIsTunnelerEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isTunnelerEnabled", "body", m.IsTunnelerEnabled); err != nil {
		return err
	}

	return nil
}

func (m *EdgeRouterDetail) validateIsVerified(formats strfmt.Registry) error {

	if err := validate.Required("isVerified", "body", m.IsVerified); err != nil {
		return err
	}

	return nil
}

func (m *EdgeRouterDetail) validateRoleAttributes(formats strfmt.Registry) error {

	if err := validate.Required("roleAttributes", "body", m.RoleAttributes); err != nil {
		return err
	}

	if m.RoleAttributes != nil {
		if err := m.RoleAttributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roleAttributes")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeRouterDetail) validateVersionInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.VersionInfo) { // not required
		return nil
	}

	if m.VersionInfo != nil {
		if err := m.VersionInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versionInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this edge router detail based on the context it is used
func (m *EdgeRouterDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseEntity
	if err := m.BaseEntity.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with CommonEdgeRouterProperties
	if err := m.CommonEdgeRouterProperties.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoleAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersionInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeRouterDetail) contextValidateRoleAttributes(ctx context.Context, formats strfmt.Registry) error {

	if m.RoleAttributes != nil {
		if err := m.RoleAttributes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roleAttributes")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeRouterDetail) contextValidateVersionInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.VersionInfo != nil {
		if err := m.VersionInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versionInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EdgeRouterDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeRouterDetail) UnmarshalBinary(b []byte) error {
	var res EdgeRouterDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
