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

package current_api_session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge/rest_model"
)

// ListCurrentAPISessionCertificatesReader is a Reader for the ListCurrentAPISessionCertificates structure.
type ListCurrentAPISessionCertificatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCurrentAPISessionCertificatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCurrentAPISessionCertificatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListCurrentAPISessionCertificatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListCurrentAPISessionCertificatesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListCurrentAPISessionCertificatesOK creates a ListCurrentAPISessionCertificatesOK with default headers values
func NewListCurrentAPISessionCertificatesOK() *ListCurrentAPISessionCertificatesOK {
	return &ListCurrentAPISessionCertificatesOK{}
}

/* ListCurrentAPISessionCertificatesOK describes a response with status code 200, with default header values.

A list of the current API Session's certificate
*/
type ListCurrentAPISessionCertificatesOK struct {
	Payload *rest_model.ListCurrentAPISessionCertificatesEnvelope
}

func (o *ListCurrentAPISessionCertificatesOK) Error() string {
	return fmt.Sprintf("[GET /current-api-session/certificates][%d] listCurrentApiSessionCertificatesOK  %+v", 200, o.Payload)
}
func (o *ListCurrentAPISessionCertificatesOK) GetPayload() *rest_model.ListCurrentAPISessionCertificatesEnvelope {
	return o.Payload
}

func (o *ListCurrentAPISessionCertificatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListCurrentAPISessionCertificatesEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCurrentAPISessionCertificatesBadRequest creates a ListCurrentAPISessionCertificatesBadRequest with default headers values
func NewListCurrentAPISessionCertificatesBadRequest() *ListCurrentAPISessionCertificatesBadRequest {
	return &ListCurrentAPISessionCertificatesBadRequest{}
}

/* ListCurrentAPISessionCertificatesBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type ListCurrentAPISessionCertificatesBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListCurrentAPISessionCertificatesBadRequest) Error() string {
	return fmt.Sprintf("[GET /current-api-session/certificates][%d] listCurrentApiSessionCertificatesBadRequest  %+v", 400, o.Payload)
}
func (o *ListCurrentAPISessionCertificatesBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListCurrentAPISessionCertificatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCurrentAPISessionCertificatesUnauthorized creates a ListCurrentAPISessionCertificatesUnauthorized with default headers values
func NewListCurrentAPISessionCertificatesUnauthorized() *ListCurrentAPISessionCertificatesUnauthorized {
	return &ListCurrentAPISessionCertificatesUnauthorized{}
}

/* ListCurrentAPISessionCertificatesUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type ListCurrentAPISessionCertificatesUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListCurrentAPISessionCertificatesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /current-api-session/certificates][%d] listCurrentApiSessionCertificatesUnauthorized  %+v", 401, o.Payload)
}
func (o *ListCurrentAPISessionCertificatesUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListCurrentAPISessionCertificatesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
