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

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge/rest_model"
)

// ListServiceTerminatorsOKCode is the HTTP code returned for type ListServiceTerminatorsOK
const ListServiceTerminatorsOKCode int = 200

/*ListServiceTerminatorsOK A list of terminators

swagger:response listServiceTerminatorsOK
*/
type ListServiceTerminatorsOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListTerminatorsEnvelope `json:"body,omitempty"`
}

// NewListServiceTerminatorsOK creates ListServiceTerminatorsOK with default headers values
func NewListServiceTerminatorsOK() *ListServiceTerminatorsOK {

	return &ListServiceTerminatorsOK{}
}

// WithPayload adds the payload to the list service terminators o k response
func (o *ListServiceTerminatorsOK) WithPayload(payload *rest_model.ListTerminatorsEnvelope) *ListServiceTerminatorsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service terminators o k response
func (o *ListServiceTerminatorsOK) SetPayload(payload *rest_model.ListTerminatorsEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceTerminatorsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServiceTerminatorsBadRequestCode is the HTTP code returned for type ListServiceTerminatorsBadRequest
const ListServiceTerminatorsBadRequestCode int = 400

/*ListServiceTerminatorsBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response listServiceTerminatorsBadRequest
*/
type ListServiceTerminatorsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServiceTerminatorsBadRequest creates ListServiceTerminatorsBadRequest with default headers values
func NewListServiceTerminatorsBadRequest() *ListServiceTerminatorsBadRequest {

	return &ListServiceTerminatorsBadRequest{}
}

// WithPayload adds the payload to the list service terminators bad request response
func (o *ListServiceTerminatorsBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServiceTerminatorsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service terminators bad request response
func (o *ListServiceTerminatorsBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceTerminatorsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServiceTerminatorsUnauthorizedCode is the HTTP code returned for type ListServiceTerminatorsUnauthorized
const ListServiceTerminatorsUnauthorizedCode int = 401

/*ListServiceTerminatorsUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response listServiceTerminatorsUnauthorized
*/
type ListServiceTerminatorsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListServiceTerminatorsUnauthorized creates ListServiceTerminatorsUnauthorized with default headers values
func NewListServiceTerminatorsUnauthorized() *ListServiceTerminatorsUnauthorized {

	return &ListServiceTerminatorsUnauthorized{}
}

// WithPayload adds the payload to the list service terminators unauthorized response
func (o *ListServiceTerminatorsUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListServiceTerminatorsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service terminators unauthorized response
func (o *ListServiceTerminatorsUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceTerminatorsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
