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

package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge/rest_model"
)

// DetailIdentityOKCode is the HTTP code returned for type DetailIdentityOK
const DetailIdentityOKCode int = 200

/*DetailIdentityOK A single identity

swagger:response detailIdentityOK
*/
type DetailIdentityOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.DetailIdentityEnvelope `json:"body,omitempty"`
}

// NewDetailIdentityOK creates DetailIdentityOK with default headers values
func NewDetailIdentityOK() *DetailIdentityOK {

	return &DetailIdentityOK{}
}

// WithPayload adds the payload to the detail identity o k response
func (o *DetailIdentityOK) WithPayload(payload *rest_model.DetailIdentityEnvelope) *DetailIdentityOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail identity o k response
func (o *DetailIdentityOK) SetPayload(payload *rest_model.DetailIdentityEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailIdentityOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailIdentityUnauthorizedCode is the HTTP code returned for type DetailIdentityUnauthorized
const DetailIdentityUnauthorizedCode int = 401

/*DetailIdentityUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response detailIdentityUnauthorized
*/
type DetailIdentityUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailIdentityUnauthorized creates DetailIdentityUnauthorized with default headers values
func NewDetailIdentityUnauthorized() *DetailIdentityUnauthorized {

	return &DetailIdentityUnauthorized{}
}

// WithPayload adds the payload to the detail identity unauthorized response
func (o *DetailIdentityUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailIdentityUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail identity unauthorized response
func (o *DetailIdentityUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailIdentityUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailIdentityNotFoundCode is the HTTP code returned for type DetailIdentityNotFound
const DetailIdentityNotFoundCode int = 404

/*DetailIdentityNotFound The requested resource does not exist

swagger:response detailIdentityNotFound
*/
type DetailIdentityNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailIdentityNotFound creates DetailIdentityNotFound with default headers values
func NewDetailIdentityNotFound() *DetailIdentityNotFound {

	return &DetailIdentityNotFound{}
}

// WithPayload adds the payload to the detail identity not found response
func (o *DetailIdentityNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailIdentityNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail identity not found response
func (o *DetailIdentityNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailIdentityNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
