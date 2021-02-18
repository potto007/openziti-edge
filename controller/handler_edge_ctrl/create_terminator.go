/*
	Copyright NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package handler_edge_ctrl

import (
	"github.com/golang/protobuf/proto"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/edge/controller/env"
	"github.com/openziti/edge/controller/persistence"
	"github.com/openziti/edge/edge_common"
	"github.com/openziti/edge/pb/edge_ctrl_pb"
	"github.com/openziti/fabric/controller/network"
	"github.com/openziti/foundation/channel2"
	"github.com/pkg/errors"
	"math"
)

type createTerminatorHandler struct {
	baseRequestHandler
}

func NewCreateTerminatorHandler(appEnv *env.AppEnv, ch channel2.Channel) channel2.ReceiveHandler {
	return &createTerminatorHandler{
		baseRequestHandler{
			ch:     ch,
			appEnv: appEnv,
		},
	}
}

func (self *createTerminatorHandler) ContentType() int32 {
	return int32(edge_ctrl_pb.ContentType_CreateTerminatorRequestType)
}

func (self *createTerminatorHandler) Label() string {
	return "create.terminator"
}

func (self *createTerminatorHandler) sendResponse(ctx *CreateTerminatorRequestContext, terminatorId string) {
	log := pfxlog.ContextLogger(self.ch.Label())

	responseMsg := channel2.NewMessage(int32(edge_ctrl_pb.ContentType_CreateTerminatorResponseType), []byte(terminatorId))
	responseMsg.ReplyTo(ctx.msg)
	if err := self.ch.Send(responseMsg); err != nil {
		log.WithError(err).WithField("token", ctx.req.SessionToken).Error("failed to send create circuit response")
	}
}

func (self *createTerminatorHandler) HandleReceive(msg *channel2.Message, ch channel2.Channel) {
	req := &edge_ctrl_pb.CreateTerminatorRequest{}
	if err := proto.Unmarshal(msg.Body, req); err != nil {
		pfxlog.ContextLogger(ch.Label()).WithError(err).Error("could not unmarshal CreateTerminatorRequest")
		return
	}

	ctx := &CreateTerminatorRequestContext{
		baseRequestContext: baseRequestContext{handler: self, msg: msg},
		req:                req,
	}

	go self.CreateTerminator(ctx)
}

func (self *createTerminatorHandler) CreateTerminator(ctx *CreateTerminatorRequestContext) {
	log := pfxlog.ContextLogger(self.ch.Label()).WithField("token", ctx.req.SessionToken)

	if !ctx.loadRouter() {
		return
	}
	ctx.loadSession(ctx.req.SessionToken)
	ctx.checkSessionType(persistence.SessionTypeBind)
	ctx.checkSessionFingerprints(ctx.req.Fingerprints)
	ctx.verifyEdgeRouterAccess()
	ctx.loadService()

	if ctx.err != nil {
		self.returnError(ctx, ctx.err)
		return
	}

	log = log.WithField("service", ctx.service.Name)

	if ctx.req.Cost > math.MaxUint16 {
		self.returnError(ctx, errors.Errorf("invalid cost %v. cost must be between 0 and %v inclusive", ctx.req.Cost, math.MaxUint16))
		return
	}

	terminator := &network.Terminator{
		Service:        ctx.session.ServiceId,
		Router:         ctx.sourceRouter.Id,
		Binding:        edge_common.Binding,
		Address:        "hosted:" + ctx.session.Token,
		Identity:       ctx.req.Identity,
		IdentitySecret: ctx.req.IdentitySecret,
		PeerData:       ctx.req.PeerData,
		Precedence:     ctx.req.GetXtPrecedence(),
		Cost:           uint16(ctx.req.Cost),
	}

	n := self.appEnv.GetHostController().GetNetwork()
	id, err := n.Terminators.Create(terminator)
	if err != nil {
		self.returnError(ctx, err)
		return
	}

	log.WithField("terminator", id).Info("created terminator")

	self.sendResponse(ctx, id)
}

type CreateTerminatorRequestContext struct {
	baseRequestContext
	req *edge_ctrl_pb.CreateTerminatorRequest
}

func (self *CreateTerminatorRequestContext) GetSessionToken() string {
	return self.req.SessionToken
}
