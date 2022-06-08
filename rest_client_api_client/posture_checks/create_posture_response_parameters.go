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

package posture_checks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge/rest_model"
)

// NewCreatePostureResponseParams creates a new CreatePostureResponseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePostureResponseParams() *CreatePostureResponseParams {
	return &CreatePostureResponseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePostureResponseParamsWithTimeout creates a new CreatePostureResponseParams object
// with the ability to set a timeout on a request.
func NewCreatePostureResponseParamsWithTimeout(timeout time.Duration) *CreatePostureResponseParams {
	return &CreatePostureResponseParams{
		timeout: timeout,
	}
}

// NewCreatePostureResponseParamsWithContext creates a new CreatePostureResponseParams object
// with the ability to set a context for a request.
func NewCreatePostureResponseParamsWithContext(ctx context.Context) *CreatePostureResponseParams {
	return &CreatePostureResponseParams{
		Context: ctx,
	}
}

// NewCreatePostureResponseParamsWithHTTPClient creates a new CreatePostureResponseParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePostureResponseParamsWithHTTPClient(client *http.Client) *CreatePostureResponseParams {
	return &CreatePostureResponseParams{
		HTTPClient: client,
	}
}

/* CreatePostureResponseParams contains all the parameters to send to the API endpoint
   for the create posture response operation.

   Typically these are written to a http.Request.
*/
type CreatePostureResponseParams struct {

	/* PostureResponse.

	   A Posture Response
	*/
	PostureResponse rest_model.PostureResponseCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create posture response params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePostureResponseParams) WithDefaults() *CreatePostureResponseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create posture response params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePostureResponseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create posture response params
func (o *CreatePostureResponseParams) WithTimeout(timeout time.Duration) *CreatePostureResponseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create posture response params
func (o *CreatePostureResponseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create posture response params
func (o *CreatePostureResponseParams) WithContext(ctx context.Context) *CreatePostureResponseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create posture response params
func (o *CreatePostureResponseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create posture response params
func (o *CreatePostureResponseParams) WithHTTPClient(client *http.Client) *CreatePostureResponseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create posture response params
func (o *CreatePostureResponseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPostureResponse adds the postureResponse to the create posture response params
func (o *CreatePostureResponseParams) WithPostureResponse(postureResponse rest_model.PostureResponseCreate) *CreatePostureResponseParams {
	o.SetPostureResponse(postureResponse)
	return o
}

// SetPostureResponse adds the postureResponse to the create posture response params
func (o *CreatePostureResponseParams) SetPostureResponse(postureResponse rest_model.PostureResponseCreate) {
	o.PostureResponse = postureResponse
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePostureResponseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.PostureResponse); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
