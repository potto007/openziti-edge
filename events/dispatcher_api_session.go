/*
	Copyright NetFoundry Inc.

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

package events

import (
	"fmt"
	"github.com/openziti/edge/controller/persistence"
	"github.com/openziti/foundation/v2/stringz"
	"github.com/openziti/storage/boltz"
	"github.com/pkg/errors"
	"reflect"
	"time"
)

func (self *Dispatcher) AddApiSessionEventHandler(handler ApiSessionEventHandler) {
	self.apiSessionEventHandlers.Append(handler)
}

func (self *Dispatcher) RemoveApiSessionEventHandler(handler ApiSessionEventHandler) {
	self.apiSessionEventHandlers.DeleteIf(func(val ApiSessionEventHandler) bool {
		if val == handler {
			return true
		}
		if w, ok := val.(ApiSessionEventHandlerWrapper); ok {
			return w.IsWrapping(handler)
		}
		return false
	})
}

func (self *Dispatcher) initApiSessionEvents(stores *persistence.Stores) {
	stores.ApiSession.AddEntityEventListenerF(self.apiSessionCreated, boltz.EntityCreated)
	stores.ApiSession.AddEntityEventListenerF(self.apiSessionDeleted, boltz.EntityDeleted)
}

func (self *Dispatcher) apiSessionCreated(apiSession *persistence.ApiSession) {
	event := &ApiSessionEvent{
		Namespace:  ApiSessionEventNS,
		EventType:  ApiSessionEventTypeCreated,
		Id:         apiSession.Id,
		Timestamp:  time.Now(),
		Token:      apiSession.Token,
		IdentityId: apiSession.IdentityId,
		IpAddress:  apiSession.IPAddress,
	}

	for _, handler := range self.apiSessionEventHandlers.Value() {
		go handler.AcceptApiSessionEvent(event)
	}
}

func (self *Dispatcher) apiSessionDeleted(apiSession *persistence.ApiSession) {
	event := &ApiSessionEvent{
		Namespace:  ApiSessionEventNS,
		EventType:  ApiSessionEventTypeDeleted,
		Id:         apiSession.Id,
		Timestamp:  time.Now(),
		Token:      apiSession.Token,
		IdentityId: apiSession.IdentityId,
		IpAddress:  apiSession.IPAddress,
	}

	for _, handler := range self.apiSessionEventHandlers.Value() {
		go handler.AcceptApiSessionEvent(event)
	}
}

func (self *Dispatcher) registerApiSessionEventHandler(val interface{}, config map[string]interface{}) error {
	handler, ok := val.(ApiSessionEventHandler)

	if !ok {
		return errors.Errorf("type %v doesn't implement github.com/openziti/edge/events/ApiSessionEventHandler interface.", reflect.TypeOf(val))
	}

	var includeList []string
	if includeVar, ok := config["include"]; ok {
		if includeStr, ok := includeVar.(string); ok {
			includeList = append(includeList, includeStr)
		} else if includeIntfList, ok := includeVar.([]interface{}); ok {
			for _, val := range includeIntfList {
				includeList = append(includeList, fmt.Sprintf("%v", val))
			}
		} else {
			return errors.Errorf("invalid type %v for %v include configuration", reflect.TypeOf(includeVar), ApiSessionEventNS)
		}
	}

	if len(includeList) == 0 || (len(includeList) == 2 && stringz.ContainsAll(includeList, ApiSessionEventTypeCreated, ApiSessionEventTypeDeleted)) {
		self.AddApiSessionEventHandler(handler)
	} else {
		for _, include := range includeList {
			if include != ApiSessionEventTypeCreated && include != ApiSessionEventTypeDeleted {
				return errors.Errorf("invalid include %v for %v. valid values are ['created', 'deleted']", include, ApiSessionEventNS)
			}
		}

		self.AddApiSessionEventHandler(&apiSessionEventAdapter{
			wrapped:     handler,
			includeList: includeList,
		})
	}

	return nil
}

func (self *Dispatcher) unregisterApiSessionEventHandler(val interface{}) {
	if handler, ok := val.(ApiSessionEventHandler); ok {
		self.RemoveApiSessionEventHandler(handler)
	}
}

type apiSessionEventAdapter struct {
	wrapped     ApiSessionEventHandler
	includeList []string
}

func (adapter *apiSessionEventAdapter) AcceptApiSessionEvent(event *ApiSessionEvent) {
	if stringz.Contains(adapter.includeList, event.EventType) {
		adapter.wrapped.AcceptApiSessionEvent(event)
	}
}

func (self *apiSessionEventAdapter) IsWrapping(value ApiSessionEventHandler) bool {
	if self.wrapped == value {
		return true
	}
	if w, ok := self.wrapped.(ApiSessionEventHandlerWrapper); ok {
		return w.IsWrapping(value)
	}
	return false
}
