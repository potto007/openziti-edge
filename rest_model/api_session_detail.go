// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry Inc.
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

// APISessionDetail An API Session object
//
// swagger:model apiSessionDetail
type APISessionDetail struct {
	BaseEntity

	// auth queries
	// Required: true
	AuthQueries AuthQueryList `json:"authQueries"`

	// authenticator Id
	// Required: true
	AuthenticatorID *string `json:"authenticatorId"`

	// cached last activity at
	// Format: date-time
	CachedLastActivityAt strfmt.DateTime `json:"cachedLastActivityAt,omitempty"`

	// config types
	// Required: true
	ConfigTypes []string `json:"configTypes"`

	// identity
	// Required: true
	Identity *EntityRef `json:"identity"`

	// identity Id
	// Required: true
	IdentityID *string `json:"identityId"`

	// ip address
	// Required: true
	IPAddress *string `json:"ipAddress"`

	// is mfa complete
	// Required: true
	IsMfaComplete *bool `json:"isMfaComplete"`

	// is mfa required
	// Required: true
	IsMfaRequired *bool `json:"isMfaRequired"`

	// last activity at
	// Format: date-time
	LastActivityAt strfmt.DateTime `json:"lastActivityAt,omitempty"`

	// token
	// Required: true
	Token *string `json:"token"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *APISessionDetail) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseEntity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseEntity = aO0

	// AO1
	var dataAO1 struct {
		AuthQueries AuthQueryList `json:"authQueries"`

		AuthenticatorID *string `json:"authenticatorId"`

		CachedLastActivityAt strfmt.DateTime `json:"cachedLastActivityAt,omitempty"`

		ConfigTypes []string `json:"configTypes"`

		Identity *EntityRef `json:"identity"`

		IdentityID *string `json:"identityId"`

		IPAddress *string `json:"ipAddress"`

		IsMfaComplete *bool `json:"isMfaComplete"`

		IsMfaRequired *bool `json:"isMfaRequired"`

		LastActivityAt strfmt.DateTime `json:"lastActivityAt,omitempty"`

		Token *string `json:"token"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AuthQueries = dataAO1.AuthQueries

	m.AuthenticatorID = dataAO1.AuthenticatorID

	m.CachedLastActivityAt = dataAO1.CachedLastActivityAt

	m.ConfigTypes = dataAO1.ConfigTypes

	m.Identity = dataAO1.Identity

	m.IdentityID = dataAO1.IdentityID

	m.IPAddress = dataAO1.IPAddress

	m.IsMfaComplete = dataAO1.IsMfaComplete

	m.IsMfaRequired = dataAO1.IsMfaRequired

	m.LastActivityAt = dataAO1.LastActivityAt

	m.Token = dataAO1.Token

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m APISessionDetail) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseEntity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		AuthQueries AuthQueryList `json:"authQueries"`

		AuthenticatorID *string `json:"authenticatorId"`

		CachedLastActivityAt strfmt.DateTime `json:"cachedLastActivityAt,omitempty"`

		ConfigTypes []string `json:"configTypes"`

		Identity *EntityRef `json:"identity"`

		IdentityID *string `json:"identityId"`

		IPAddress *string `json:"ipAddress"`

		IsMfaComplete *bool `json:"isMfaComplete"`

		IsMfaRequired *bool `json:"isMfaRequired"`

		LastActivityAt strfmt.DateTime `json:"lastActivityAt,omitempty"`

		Token *string `json:"token"`
	}

	dataAO1.AuthQueries = m.AuthQueries

	dataAO1.AuthenticatorID = m.AuthenticatorID

	dataAO1.CachedLastActivityAt = m.CachedLastActivityAt

	dataAO1.ConfigTypes = m.ConfigTypes

	dataAO1.Identity = m.Identity

	dataAO1.IdentityID = m.IdentityID

	dataAO1.IPAddress = m.IPAddress

	dataAO1.IsMfaComplete = m.IsMfaComplete

	dataAO1.IsMfaRequired = m.IsMfaRequired

	dataAO1.LastActivityAt = m.LastActivityAt

	dataAO1.Token = m.Token

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this api session detail
func (m *APISessionDetail) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseEntity
	if err := m.BaseEntity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthQueries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticatorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCachedLastActivityAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsMfaComplete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsMfaRequired(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastActivityAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APISessionDetail) validateAuthQueries(formats strfmt.Registry) error {

	if err := validate.Required("authQueries", "body", m.AuthQueries); err != nil {
		return err
	}

	if err := m.AuthQueries.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("authQueries")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("authQueries")
		}
		return err
	}

	return nil
}

func (m *APISessionDetail) validateAuthenticatorID(formats strfmt.Registry) error {

	if err := validate.Required("authenticatorId", "body", m.AuthenticatorID); err != nil {
		return err
	}

	return nil
}

func (m *APISessionDetail) validateCachedLastActivityAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CachedLastActivityAt) { // not required
		return nil
	}

	if err := validate.FormatOf("cachedLastActivityAt", "body", "date-time", m.CachedLastActivityAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APISessionDetail) validateConfigTypes(formats strfmt.Registry) error {

	if err := validate.Required("configTypes", "body", m.ConfigTypes); err != nil {
		return err
	}

	return nil
}

func (m *APISessionDetail) validateIdentity(formats strfmt.Registry) error {

	if err := validate.Required("identity", "body", m.Identity); err != nil {
		return err
	}

	if m.Identity != nil {
		if err := m.Identity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identity")
			}
			return err
		}
	}

	return nil
}

func (m *APISessionDetail) validateIdentityID(formats strfmt.Registry) error {

	if err := validate.Required("identityId", "body", m.IdentityID); err != nil {
		return err
	}

	return nil
}

func (m *APISessionDetail) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("ipAddress", "body", m.IPAddress); err != nil {
		return err
	}

	return nil
}

func (m *APISessionDetail) validateIsMfaComplete(formats strfmt.Registry) error {

	if err := validate.Required("isMfaComplete", "body", m.IsMfaComplete); err != nil {
		return err
	}

	return nil
}

func (m *APISessionDetail) validateIsMfaRequired(formats strfmt.Registry) error {

	if err := validate.Required("isMfaRequired", "body", m.IsMfaRequired); err != nil {
		return err
	}

	return nil
}

func (m *APISessionDetail) validateLastActivityAt(formats strfmt.Registry) error {

	if swag.IsZero(m.LastActivityAt) { // not required
		return nil
	}

	if err := validate.FormatOf("lastActivityAt", "body", "date-time", m.LastActivityAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APISessionDetail) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this api session detail based on the context it is used
func (m *APISessionDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseEntity
	if err := m.BaseEntity.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuthQueries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APISessionDetail) contextValidateAuthQueries(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AuthQueries.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("authQueries")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("authQueries")
		}
		return err
	}

	return nil
}

func (m *APISessionDetail) contextValidateIdentity(ctx context.Context, formats strfmt.Registry) error {

	if m.Identity != nil {
		if err := m.Identity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APISessionDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APISessionDetail) UnmarshalBinary(b []byte) error {
	var res APISessionDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
