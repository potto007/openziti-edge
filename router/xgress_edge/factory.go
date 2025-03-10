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

package xgress_edge

import (
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/channel/v2"
	"github.com/openziti/edge/router/fabric"
	"github.com/openziti/edge/router/handler_edge_ctrl"
	"github.com/openziti/edge/router/internal/apiproxy"
	"github.com/openziti/edge/router/internal/edgerouter"
	"github.com/openziti/fabric/router"
	"github.com/openziti/fabric/router/env"
	"github.com/openziti/fabric/router/xgress"
	"github.com/openziti/foundation/v2/versions"
	"github.com/openziti/identity"
	"github.com/openziti/metrics"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type Factory struct {
	id               *identity.TokenId
	ctrls            env.NetworkControllers
	enabled          bool
	routerConfig     *router.Config
	edgeRouterConfig *edgerouter.Config
	hostedServices   *hostedServiceRegistry
	stateManager     fabric.StateManager
	versionProvider  versions.VersionProvider
	certChecker      *CertExpirationChecker
	metricsRegistry  metrics.Registry
}

func (factory *Factory) GetNetworkControllers() env.NetworkControllers {
	return factory.ctrls
}

func (factory *Factory) Enabled() bool {
	return factory.enabled
}

const (
	WsType = "ws"
)

func (factory *Factory) BindChannel(binding channel.Binding) error {
	binding.AddTypedReceiveHandler(handler_edge_ctrl.NewHelloHandler(factory.edgeRouterConfig.EdgeListeners))

	binding.AddTypedReceiveHandler(handler_edge_ctrl.NewSessionRemovedHandler(factory.stateManager))

	binding.AddTypedReceiveHandler(handler_edge_ctrl.NewApiSessionAddedHandler(factory.stateManager, binding))
	binding.AddTypedReceiveHandler(handler_edge_ctrl.NewApiSessionRemovedHandler(factory.stateManager))
	binding.AddTypedReceiveHandler(handler_edge_ctrl.NewApiSessionUpdatedHandler(factory.stateManager))

	binding.AddTypedReceiveHandler(handler_edge_ctrl.NewExtendEnrollmentCertsHandler(factory.routerConfig.Id, func() {
		factory.certChecker.CertsUpdated()
	}))

	return nil
}

func (factory *Factory) NotifyOfReconnect(ch channel.Channel) {
	go factory.stateManager.ValidateSessions(ch, factory.edgeRouterConfig.SessionValidateChunkSize, factory.edgeRouterConfig.SessionValidateMinInterval, factory.edgeRouterConfig.SessionValidateMaxInterval)
}

func (factory *Factory) GetTraceDecoders() []channel.TraceMessageDecoder {
	return nil
}

func (factory *Factory) Run(env env.RouterEnv) error {
	factory.ctrls = env.GetNetworkControllers()

	factory.stateManager.StartHeartbeat(env, factory.edgeRouterConfig.HeartbeatIntervalSeconds, env.GetCloseNotify())

	factory.certChecker = NewCertExpirationChecker(factory.routerConfig.Id, factory.edgeRouterConfig, env.GetNetworkControllers(), env.GetCloseNotify())

	go func() {
		if err := factory.certChecker.Run(); err != nil {
			pfxlog.Logger().WithError(err).Error("error while running certchecker")
		}
	}()

	return nil
}

func (factory *Factory) LoadConfig(configMap map[interface{}]interface{}) error {
	_, factory.enabled = configMap["edge"]

	if !factory.enabled {
		return nil
	}

	var err error
	config := edgerouter.NewConfig(factory.routerConfig)
	if err = config.LoadConfigFromMap(configMap); err != nil {
		return err
	}

	factory.id = config.RouterConfig.Id

	factory.edgeRouterConfig = config
	go apiproxy.Start(config)

	return nil
}

// NewFactory constructs a new Edge Xgress Factory instance
func NewFactory(routerConfig *router.Config, versionProvider versions.VersionProvider, stateManager fabric.StateManager, metricsRegistry metrics.Registry) *Factory {
	factory := &Factory{
		hostedServices:  &hostedServiceRegistry{},
		stateManager:    stateManager,
		versionProvider: versionProvider,
		routerConfig:    routerConfig,
		metricsRegistry: metricsRegistry,
	}
	return factory
}

// CreateListener creates a new Edge Xgress listener
func (factory *Factory) CreateListener(optionsData xgress.OptionsData) (xgress.Listener, error) {
	if !factory.enabled {
		return nil, errors.New("edge listener enabled but required configuration section [edge] is missing")
	}

	options := &Options{}
	if err := options.load(optionsData); err != nil {
		return nil, err
	}

	pfxlog.Logger().Debugf("xgress edge listener options: %v", options.ToLoggableString())

	versionInfo := factory.versionProvider.AsVersionInfo()
	versionHeader, err := factory.versionProvider.EncoderDecoder().Encode(versionInfo)

	if err != nil {
		return nil, fmt.Errorf("could not generate version header: %v", err)
	}

	headers := map[int32][]byte{
		channel.HelloVersionHeader: versionHeader,
	}

	return newListener(factory.id, factory, options, headers), nil
}

// CreateDialer creates a new Edge Xgress dialer
func (factory *Factory) CreateDialer(optionsData xgress.OptionsData) (xgress.Dialer, error) {
	if !factory.enabled {
		return nil, errors.New("edge listener enabled but required configuration section [edge] is missing")
	}

	options := &Options{}
	if err := options.load(optionsData); err != nil {
		return nil, err
	}

	return newDialer(factory, options), nil
}

type Options struct {
	xgress.Options
	channelOptions          *channel.Options
	lookupApiSessionTimeout time.Duration
	lookupSessionTimeout    time.Duration
}

func (options *Options) ToLoggableString() string {
	buf := strings.Builder{}
	buf.WriteString(fmt.Sprintf("mtu=%v\n", options.Mtu))
	buf.WriteString(fmt.Sprintf("randomDrops=%v\n", options.RandomDrops))
	buf.WriteString(fmt.Sprintf("drop1InN=%v\n", options.Drop1InN))
	buf.WriteString(fmt.Sprintf("txQueueSize=%v\n", options.TxQueueSize))
	buf.WriteString(fmt.Sprintf("txPortalStartSize=%v\n", options.TxPortalStartSize))
	buf.WriteString(fmt.Sprintf("txPortalMaxSize=%v\n", options.TxPortalMaxSize))
	buf.WriteString(fmt.Sprintf("txPortalMinSize=%v\n", options.TxPortalMinSize))
	buf.WriteString(fmt.Sprintf("txPortalIncreaseThresh=%v\n", options.TxPortalIncreaseThresh))
	buf.WriteString(fmt.Sprintf("txPortalIncreaseScale=%v\n", options.TxPortalIncreaseScale))
	buf.WriteString(fmt.Sprintf("txPortalRetxThresh=%v\n", options.TxPortalRetxThresh))
	buf.WriteString(fmt.Sprintf("txPortalRetxScale=%v\n", options.TxPortalRetxScale))
	buf.WriteString(fmt.Sprintf("txPortalDupAckThresh=%v\n", options.TxPortalDupAckThresh))
	buf.WriteString(fmt.Sprintf("txPortalDupAckScale=%v\n", options.TxPortalDupAckScale))
	buf.WriteString(fmt.Sprintf("rxBufferSize=%v\n", options.RxBufferSize))
	buf.WriteString(fmt.Sprintf("retxStartMs=%v\n", options.RetxStartMs))
	buf.WriteString(fmt.Sprintf("retxScale=%v\n", options.RetxScale))
	buf.WriteString(fmt.Sprintf("retxAddMs=%v\n", options.RetxAddMs))
	buf.WriteString(fmt.Sprintf("maxCloseWait=%v\n", options.MaxCloseWait))
	buf.WriteString(fmt.Sprintf("getCircuitTimeout=%v\n", options.GetCircuitTimeout))

	buf.WriteString(fmt.Sprintf("lookupApiSessionTimeout=%v\n", options.lookupApiSessionTimeout))
	buf.WriteString(fmt.Sprintf("lookupSessionTimeout=%v\n", options.lookupSessionTimeout))

	buf.WriteString(fmt.Sprintf("channel.outQueueSize=%v\n", options.channelOptions.OutQueueSize))
	buf.WriteString(fmt.Sprintf("channel.connectTimeout=%v\n", options.channelOptions.ConnectTimeout))
	buf.WriteString(fmt.Sprintf("channel.maxOutstandingConnects=%v\n", options.channelOptions.MaxOutstandingConnects))
	buf.WriteString(fmt.Sprintf("channel.maxQueuedConnects=%v\n", options.channelOptions.MaxQueuedConnects))

	return buf.String()
}

func (options *Options) load(data xgress.OptionsData) error {
	o, err := xgress.LoadOptions(data)
	if err != nil {
		return errors.Wrap(err, "error loading options")
	}
	options.Options = *o
	options.lookupSessionTimeout = 5 * time.Second
	options.lookupApiSessionTimeout = 5 * time.Second

	if value, found := data["options"]; found {
		data = value.(map[interface{}]interface{})

		var err error
		options.channelOptions, err = channel.LoadOptions(data)
		if err != nil {
			return err
		}
		if err := options.channelOptions.Validate(); err != nil {
			return fmt.Errorf("error loading options for [edge/options]: %v", err)
		}

		if value, found := data["lookupSessionTimeout"]; found {
			timeout, err := time.ParseDuration(value.(string))
			if err != nil {
				return errors.Wrap(err, "invalid 'lookupSessionTimeout' value")
			}
			options.lookupSessionTimeout = timeout
		}

		if value, found := data["lookupApiSessionTimeout"]; found {
			timeout, err := time.ParseDuration(value.(string))
			if err != nil {
				return errors.Wrap(err, "invalid 'lookupApiSessionTimeout' value")
			}
			options.lookupApiSessionTimeout = timeout
		}
	} else {
		options.channelOptions = channel.DefaultOptions()
	}
	return nil
}
