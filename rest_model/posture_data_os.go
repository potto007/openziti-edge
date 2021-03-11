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
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostureDataOs posture data os
//
// swagger:model postureDataOs
type PostureDataOs struct {
	PostureDataBase

	// build
	// Required: true
	Build *string `json:"build"`

	// type
	// Required: true
	Type *string `json:"type"`

	// version
	// Required: true
	Version *string `json:"version"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PostureDataOs) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PostureDataBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PostureDataBase = aO0

	// AO1
	var dataAO1 struct {
		Build *string `json:"build"`

		Type *string `json:"type"`

		Version *string `json:"version"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Build = dataAO1.Build

	m.Type = dataAO1.Type

	m.Version = dataAO1.Version

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PostureDataOs) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PostureDataBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Build *string `json:"build"`

		Type *string `json:"type"`

		Version *string `json:"version"`
	}

	dataAO1.Build = m.Build

	dataAO1.Type = m.Type

	dataAO1.Version = m.Version

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this posture data os
func (m *PostureDataOs) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PostureDataBase
	if err := m.PostureDataBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuild(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostureDataOs) validateBuild(formats strfmt.Registry) error {

	if err := validate.Required("build", "body", m.Build); err != nil {
		return err
	}

	return nil
}

func (m *PostureDataOs) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *PostureDataOs) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostureDataOs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostureDataOs) UnmarshalBinary(b []byte) error {
	var res PostureDataOs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
