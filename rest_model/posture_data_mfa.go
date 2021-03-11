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

// PostureDataMfa posture data mfa
//
// swagger:model postureDataMfa
type PostureDataMfa struct {

	// api session Id
	// Required: true
	APISessionID *string `json:"apiSessionId"`

	// passed mfa
	// Required: true
	PassedMfa *bool `json:"passedMfa"`
}

// Validate validates this posture data mfa
func (m *PostureDataMfa) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPISessionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassedMfa(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostureDataMfa) validateAPISessionID(formats strfmt.Registry) error {

	if err := validate.Required("apiSessionId", "body", m.APISessionID); err != nil {
		return err
	}

	return nil
}

func (m *PostureDataMfa) validatePassedMfa(formats strfmt.Registry) error {

	if err := validate.Required("passedMfa", "body", m.PassedMfa); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostureDataMfa) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostureDataMfa) UnmarshalBinary(b []byte) error {
	var res PostureDataMfa
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
