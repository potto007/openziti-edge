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

package edge_router

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReEnrollEdgeRouterHandlerFunc turns a function with the right signature into a re enroll edge router handler
type ReEnrollEdgeRouterHandlerFunc func(ReEnrollEdgeRouterParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReEnrollEdgeRouterHandlerFunc) Handle(params ReEnrollEdgeRouterParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReEnrollEdgeRouterHandler interface for that can handle valid re enroll edge router params
type ReEnrollEdgeRouterHandler interface {
	Handle(ReEnrollEdgeRouterParams, interface{}) middleware.Responder
}

// NewReEnrollEdgeRouter creates a new http.Handler for the re enroll edge router operation
func NewReEnrollEdgeRouter(ctx *middleware.Context, handler ReEnrollEdgeRouterHandler) *ReEnrollEdgeRouter {
	return &ReEnrollEdgeRouter{Context: ctx, Handler: handler}
}

/* ReEnrollEdgeRouter swagger:route POST /edge-routers/{id}/re-enroll Edge Router reEnrollEdgeRouter

Re-enroll an edge router

Removes current certificate based authentication mechanisms and reverts the edge router into a state where enrollment must be performed.
The router retains all other properties and associations. If the router is currently connected, it will be disconnected and any
attemps to reconnect will fail until the enrollment process is completed with the newly generated JWT.

If the edge router has an existing outstanding enrollment JWT it will be replaced. The previous JWT will no longer be usable to
complete the enrollment process.


*/
type ReEnrollEdgeRouter struct {
	Context *middleware.Context
	Handler ReEnrollEdgeRouterHandler
}

func (o *ReEnrollEdgeRouter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewReEnrollEdgeRouterParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
