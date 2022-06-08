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

// DetailConfigTypeReader is a Reader for the DetailConfigType structure.
type DetailConfigTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetailConfigTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDetailConfigTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDetailConfigTypeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDetailConfigTypeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDetailConfigTypeOK creates a DetailConfigTypeOK with default headers values
func NewDetailConfigTypeOK() *DetailConfigTypeOK {
	return &DetailConfigTypeOK{}
}

/* DetailConfigTypeOK describes a response with status code 200, with default header values.

A singular config-type resource
*/
type DetailConfigTypeOK struct {
	Payload *rest_model.DetailConfigTypeEnvelope
}

func (o *DetailConfigTypeOK) Error() string {
	return fmt.Sprintf("[GET /config-types/{id}][%d] detailConfigTypeOK  %+v", 200, o.Payload)
}
func (o *DetailConfigTypeOK) GetPayload() *rest_model.DetailConfigTypeEnvelope {
	return o.Payload
}

func (o *DetailConfigTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.DetailConfigTypeEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailConfigTypeUnauthorized creates a DetailConfigTypeUnauthorized with default headers values
func NewDetailConfigTypeUnauthorized() *DetailConfigTypeUnauthorized {
	return &DetailConfigTypeUnauthorized{}
}

/* DetailConfigTypeUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type DetailConfigTypeUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailConfigTypeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /config-types/{id}][%d] detailConfigTypeUnauthorized  %+v", 401, o.Payload)
}
func (o *DetailConfigTypeUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailConfigTypeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailConfigTypeNotFound creates a DetailConfigTypeNotFound with default headers values
func NewDetailConfigTypeNotFound() *DetailConfigTypeNotFound {
	return &DetailConfigTypeNotFound{}
}

/* DetailConfigTypeNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type DetailConfigTypeNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailConfigTypeNotFound) Error() string {
	return fmt.Sprintf("[GET /config-types/{id}][%d] detailConfigTypeNotFound  %+v", 404, o.Payload)
}
func (o *DetailConfigTypeNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailConfigTypeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
