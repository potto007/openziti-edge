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

// ListServiceUpdatesReader is a Reader for the ListServiceUpdates structure.
type ListServiceUpdatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServiceUpdatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServiceUpdatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListServiceUpdatesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListServiceUpdatesOK creates a ListServiceUpdatesOK with default headers values
func NewListServiceUpdatesOK() *ListServiceUpdatesOK {
	return &ListServiceUpdatesOK{}
}

/* ListServiceUpdatesOK describes a response with status code 200, with default header values.

Data indicating necessary service updates
*/
type ListServiceUpdatesOK struct {
	Payload *rest_model.ListCurrentAPISessionServiceUpdatesEnvelope
}

func (o *ListServiceUpdatesOK) Error() string {
	return fmt.Sprintf("[GET /current-api-session/service-updates][%d] listServiceUpdatesOK  %+v", 200, o.Payload)
}
func (o *ListServiceUpdatesOK) GetPayload() *rest_model.ListCurrentAPISessionServiceUpdatesEnvelope {
	return o.Payload
}

func (o *ListServiceUpdatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListCurrentAPISessionServiceUpdatesEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServiceUpdatesUnauthorized creates a ListServiceUpdatesUnauthorized with default headers values
func NewListServiceUpdatesUnauthorized() *ListServiceUpdatesUnauthorized {
	return &ListServiceUpdatesUnauthorized{}
}

/* ListServiceUpdatesUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type ListServiceUpdatesUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListServiceUpdatesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /current-api-session/service-updates][%d] listServiceUpdatesUnauthorized  %+v", 401, o.Payload)
}
func (o *ListServiceUpdatesUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListServiceUpdatesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
