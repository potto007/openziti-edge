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

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge/rest_model"
)

// UpdateConfigReader is a Reader for the UpdateConfig structure.
type UpdateConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateConfigOK creates a UpdateConfigOK with default headers values
func NewUpdateConfigOK() *UpdateConfigOK {
	return &UpdateConfigOK{}
}

/* UpdateConfigOK describes a response with status code 200, with default header values.

The update request was successful and the resource has been altered
*/
type UpdateConfigOK struct {
	Payload *rest_model.Empty
}

func (o *UpdateConfigOK) Error() string {
	return fmt.Sprintf("[PUT /configs/{id}][%d] updateConfigOK  %+v", 200, o.Payload)
}
func (o *UpdateConfigOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *UpdateConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConfigBadRequest creates a UpdateConfigBadRequest with default headers values
func NewUpdateConfigBadRequest() *UpdateConfigBadRequest {
	return &UpdateConfigBadRequest{}
}

/* UpdateConfigBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type UpdateConfigBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *UpdateConfigBadRequest) Error() string {
	return fmt.Sprintf("[PUT /configs/{id}][%d] updateConfigBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateConfigBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdateConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConfigUnauthorized creates a UpdateConfigUnauthorized with default headers values
func NewUpdateConfigUnauthorized() *UpdateConfigUnauthorized {
	return &UpdateConfigUnauthorized{}
}

/* UpdateConfigUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type UpdateConfigUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *UpdateConfigUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /configs/{id}][%d] updateConfigUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateConfigUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdateConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConfigNotFound creates a UpdateConfigNotFound with default headers values
func NewUpdateConfigNotFound() *UpdateConfigNotFound {
	return &UpdateConfigNotFound{}
}

/* UpdateConfigNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type UpdateConfigNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *UpdateConfigNotFound) Error() string {
	return fmt.Sprintf("[PUT /configs/{id}][%d] updateConfigNotFound  %+v", 404, o.Payload)
}
func (o *UpdateConfigNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *UpdateConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
