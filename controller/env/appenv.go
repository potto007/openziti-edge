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

package env

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/tls"
	"fmt"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	openApiMiddleware "github.com/go-openapi/runtime/middleware"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/lucsky/cuid"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/channel/v2"
	clientServer "github.com/openziti/edge-api/rest_client_api_server"
	clientOperations "github.com/openziti/edge-api/rest_client_api_server/operations"
	managementServer "github.com/openziti/edge-api/rest_management_api_server"
	managementOperations "github.com/openziti/edge-api/rest_management_api_server/operations"
	"github.com/openziti/edge-api/rest_model"
	edgeConfig "github.com/openziti/edge/controller/config"
	"github.com/openziti/edge/controller/internal/permissions"
	"github.com/openziti/edge/controller/model"
	"github.com/openziti/edge/controller/persistence"
	"github.com/openziti/edge/controller/response"
	"github.com/openziti/edge/eid"
	"github.com/openziti/edge/events"
	"github.com/openziti/edge/internal/cert"
	"github.com/openziti/edge/internal/jwtsigner"
	"github.com/openziti/fabric/controller/api"
	"github.com/openziti/fabric/controller/network"
	"github.com/openziti/fabric/controller/xctrl"
	"github.com/openziti/fabric/controller/xmgmt"
	"github.com/openziti/foundation/v2/errorz"
	"github.com/openziti/foundation/v2/stringz"
	"github.com/openziti/identity"
	"github.com/openziti/metrics"
	"github.com/openziti/sdk-golang/ziti"
	"github.com/openziti/storage/boltz"
	"github.com/openziti/xweb/v2"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/xeipuuv/gojsonschema"
	"io"
	"net/http"
	"strings"
	"time"
)

var _ model.Env = &AppEnv{}

const ZitiSession = "zt-session"

type AppEnv struct {
	BoltStores *persistence.Stores
	Managers   *model.Managers
	Config     *edgeConfig.Config

	Versions *ziti.Versions

	ApiServerCsrSigner     cert.Signer
	ApiClientCsrSigner     cert.Signer
	ControlClientCsrSigner cert.Signer

	FingerprintGenerator cert.FingerprintGenerator
	AuthRegistry         model.AuthRegistry
	EnrollRegistry       model.EnrollmentRegistry
	Broker               *Broker
	HostController       HostController
	ManagementApi        *managementOperations.ZitiEdgeManagementAPI
	ClientApi            *clientOperations.ZitiEdgeClientAPI
	IdentityRefreshMap   cmap.ConcurrentMap[string, time.Time]
	identityRefreshMeter metrics.Meter
	StartupTime          time.Time
	InstanceId           string
	enrollmentSigner     jwtsigner.Signer
	TraceManager         *TraceManager
	EventDispatcher      *events.Dispatcher
}

func (ae *AppEnv) GetApiServerCsrSigner() cert.Signer {
	return ae.ApiServerCsrSigner
}

func (ae *AppEnv) GetControlClientCsrSigner() cert.Signer {
	return ae.ControlClientCsrSigner
}

func (ae *AppEnv) GetApiClientCsrSigner() cert.Signer {
	return ae.ApiClientCsrSigner
}

func (ae *AppEnv) GetHostController() model.HostController {
	return ae.HostController
}

func (ae *AppEnv) GetManagers() *model.Managers {
	return ae.Managers
}

func (ae *AppEnv) GetConfig() *edgeConfig.Config {
	return ae.Config
}

func (ae *AppEnv) GetJwtSigner() jwtsigner.Signer {
	return ae.enrollmentSigner
}

func (ae *AppEnv) GetDbProvider() persistence.DbProvider {
	return ae.HostController.GetNetwork()
}

func (ae *AppEnv) GetStores() *persistence.Stores {
	return ae.BoltStores
}

func (ae *AppEnv) GetAuthRegistry() model.AuthRegistry {
	return ae.AuthRegistry
}

func (ae *AppEnv) GetEnrollRegistry() model.EnrollmentRegistry {
	return ae.EnrollRegistry
}

func (ae *AppEnv) IsEdgeRouterOnline(id string) bool {
	return ae.Broker.IsEdgeRouterOnline(id)
}

func (ae *AppEnv) GetMetricsRegistry() metrics.Registry {
	return ae.HostController.GetNetwork().GetMetricsRegistry()
}

func (ae *AppEnv) GetFingerprintGenerator() cert.FingerprintGenerator {
	return ae.FingerprintGenerator
}

type HostController interface {
	RegisterAgentBindHandler(bindHandler channel.BindHandler)
	RegisterXctrl(x xctrl.Xctrl) error
	RegisterXmgmt(x xmgmt.Xmgmt) error
	GetXWebInstance() xweb.Instance
	GetNetwork() *network.Network
	GetCloseNotifyChannel() <-chan struct{}
	Shutdown()
	Identity() identity.Identity
	IsRaftEnabled() bool
}

type Schemes struct {
	Association             *BasicEntitySchema
	Authenticator           *BasicEntitySchema
	AuthenticatorSelf       *BasicEntitySchema
	Ca                      *BasicEntitySchema
	Config                  *BasicEntitySchema
	ConfigType              *BasicEntitySchema
	Enroller                *BasicEntitySchema
	EnrollEr                *BasicEntitySchema
	EnrollUpdb              *BasicEntitySchema
	EdgeRouter              *BasicEntitySchema
	EdgeRouterPolicy        *BasicEntitySchema
	TransitRouter           *BasicEntitySchema
	Identity                *IdentityEntitySchema
	Service                 *BasicEntitySchema
	ServiceEdgeRouterPolicy *BasicEntitySchema
	ServicePolicy           *BasicEntitySchema
	Session                 *BasicEntitySchema
	Terminator              *BasicEntitySchema
}

func (s Schemes) GetEnrollErPost() *gojsonschema.Schema {
	return s.EnrollEr.Post
}

func (s Schemes) GetEnrollUpdbPost() *gojsonschema.Schema {
	return s.EnrollUpdb.Post
}

type IdentityEntitySchema struct {
	Post           *gojsonschema.Schema
	Patch          *gojsonschema.Schema
	Put            *gojsonschema.Schema
	ServiceConfigs *gojsonschema.Schema
}

type BasicEntitySchema struct {
	Post  *gojsonschema.Schema
	Patch *gojsonschema.Schema
	Put   *gojsonschema.Schema
}

type AppHandler func(ae *AppEnv, rc *response.RequestContext)

type AppMiddleware func(*AppEnv, http.Handler) http.Handler

type authorizer struct {
}

const (
	EventualEventsGauge = "eventual.events"
)

func (a authorizer) Authorize(request *http.Request, principal interface{}) error {
	//principal is an API Session
	_, ok := principal.(*model.ApiSession)

	if !ok {
		pfxlog.Logger().Error("principal expected to be an ApiSession and was not")
		return errorz.NewUnauthorized()
	}

	rc, err := GetRequestContextFromHttpContext(request)

	if rc == nil || err != nil {
		pfxlog.Logger().WithError(err).Error("attempting to retrieve request context failed")
		return errorz.NewUnauthorized()
	}

	if rc.Identity == nil {
		return errorz.NewUnauthorized()
	}

	return nil
}

func (ae *AppEnv) FillRequestContext(rc *response.RequestContext) error {
	rc.SessionToken = ae.GetSessionTokenFromRequest(rc.Request)
	logger := pfxlog.Logger()

	if rc.SessionToken != "" {
		_, err := uuid.Parse(rc.SessionToken)
		if err != nil {
			logger.WithError(err).Debug("failed to parse session id")
			rc.SessionToken = ""
		} else {
			logger.Tracef("authorizing request using session id '%v'", rc.SessionToken)
		}

	}

	if rc.SessionToken != "" {
		var err error
		rc.ApiSession, err = ae.GetManagers().ApiSession.ReadByToken(rc.SessionToken)
		if err != nil {
			logger.WithError(err).Debugf("looking up ApiConfig session for %s resulted in an error, request will continue unauthenticated", rc.SessionToken)
			rc.ApiSession = nil
			rc.SessionToken = ""
		}
	}

	if rc.ApiSession != nil {
		//updates for api session timeouts
		ae.GetManagers().ApiSession.MarkLastActivityById(rc.ApiSession.Id)

		var err error
		rc.Identity, err = ae.GetManagers().Identity.Read(rc.ApiSession.IdentityId)
		if err != nil {
			if boltz.IsErrNotFoundErr(err) {
				apiErr := errorz.NewUnauthorized()
				apiErr.Cause = fmt.Errorf("associated identity %s not found", rc.ApiSession.IdentityId)
				apiErr.AppendCause = true
				return apiErr
			} else {
				return err
			}
		}
	}

	if rc.Identity != nil {
		var err error
		rc.AuthPolicy, err = ae.GetManagers().AuthPolicy.Read(rc.Identity.AuthPolicyId)

		if err != nil {
			if boltz.IsErrNotFoundErr(err) {
				apiErr := errorz.NewUnauthorized()
				apiErr.Cause = fmt.Errorf("associated auth policy %s not found", rc.Identity.AuthPolicyId)
				apiErr.AppendCause = true
				return apiErr
			} else {
				return err
			}
		}

		if rc.AuthPolicy == nil {
			err := fmt.Errorf("unahndled scenario, nil auth policy [%s] found on identity [%s]", rc.Identity.AuthPolicyId, rc.Identity.Id)
			logger.Error(err)
			return err
		}

		ProcessAuthQueries(ae, rc)

		isPartialAuth := len(rc.AuthQueries) > 0

		if isPartialAuth {
			rc.ActivePermissions = append(rc.ActivePermissions, permissions.PartiallyAuthenticatePermission)
		} else {
			rc.ActivePermissions = append(rc.ActivePermissions, permissions.AuthenticatedPermission)
		}

		if rc.Identity.IsAdmin || rc.Identity.IsDefaultAdmin {
			rc.ActivePermissions = append(rc.ActivePermissions, permissions.AdminPermission)
		}
	}

	return nil
}

func NewAuthQueryZitiMfa() *rest_model.AuthQueryDetail {
	provider := rest_model.MfaProvidersZiti
	return &rest_model.AuthQueryDetail{
		TypeID:     "MFA",
		Format:     rest_model.MfaFormatsAlphaNumeric,
		HTTPMethod: http.MethodPost,
		HTTPURL:    "./authenticate/mfa",
		MaxLength:  model.TotpMaxLength,
		MinLength:  model.TotpMinLength,
		Provider:   &provider,
	}
}

func NewAuthQueryExtJwt(url string) *rest_model.AuthQueryDetail {
	return &rest_model.AuthQueryDetail{
		HTTPURL: url,
		TypeID:  "EXT-JWT",
	}
}

// ProcessAuthQueries will inspect a response.RequestContext and set the AuthQueries
// with the current outstanding authentication queries.
func ProcessAuthQueries(ae *AppEnv, rc *response.RequestContext) {
	if rc.ApiSession == nil || rc.AuthPolicy == nil {
		return
	}

	totpRequired := rc.ApiSession.MfaRequired || rc.AuthPolicy.Secondary.RequireTotp

	if totpRequired && !rc.ApiSession.MfaComplete {
		rc.AuthQueries = append(rc.AuthQueries, NewAuthQueryZitiMfa())
	}

	if rc.AuthPolicy.Secondary.RequiredExtJwtSigner != nil {
		extJwtAuthVal := ae.GetAuthRegistry().GetByMethod(model.AuthMethodExtJwt)
		extJwtAuth := extJwtAuthVal.(*model.AuthModuleExtJwt)
		if extJwtAuth != nil {
			authResult, err := extJwtAuth.ProcessSecondary(model.NewAuthContextHttp(rc.Request, model.AuthMethodExtJwt, nil, rc.NewChangeContext()))

			if err != nil || !authResult.IsSuccessful() {
				signer, err := ae.Managers.ExternalJwtSigner.Read(*rc.AuthPolicy.Secondary.RequiredExtJwtSigner)
				authUrl := ""
				if err == nil {
					authUrl = stringz.OrEmpty(signer.ExternalAuthUrl)
				}

				rc.AuthQueries = append(rc.AuthQueries, NewAuthQueryExtJwt(authUrl))

			}
		}
	}
}

func NewAppEnv(c *edgeConfig.Config, host HostController) *AppEnv {
	clientSpec, err := loads.Embedded(clientServer.SwaggerJSON, clientServer.FlatSwaggerJSON)
	if err != nil {
		pfxlog.Logger().Fatalln(err)
	}

	managementSpec, err := loads.Embedded(managementServer.SwaggerJSON, managementServer.FlatSwaggerJSON)
	if err != nil {
		pfxlog.Logger().Fatalln(err)
	}

	clientApi := clientOperations.NewZitiEdgeClientAPI(clientSpec)
	clientApi.ServeError = ServeError

	managementApi := managementOperations.NewZitiEdgeManagementAPI(managementSpec)
	managementApi.ServeError = ServeError

	ae := &AppEnv{
		Config: c,
		Versions: &ziti.Versions{
			Api:           "1.0.0",
			EnrollmentApi: "1.0.0",
		},
		HostController:     host,
		InstanceId:         cuid.New(),
		AuthRegistry:       &model.AuthProcessorRegistryImpl{},
		EnrollRegistry:     &model.EnrollmentRegistryImpl{},
		ManagementApi:      managementApi,
		ClientApi:          clientApi,
		IdentityRefreshMap: cmap.New[time.Time](),
		StartupTime:        time.Now().UTC(),
	}

	ae.identityRefreshMeter = ae.GetHostController().GetNetwork().GetMetricsRegistry().Meter("identity.refresh")

	clientApi.APIAuthorizer = authorizer{}
	managementApi.APIAuthorizer = authorizer{}

	noOpConsumer := runtime.ConsumerFunc(func(reader io.Reader, data interface{}) error {
		return nil //do nothing
	})

	//enrollment consumer, leave content unread, allow modules to read
	clientApi.ApplicationXPemFileConsumer = noOpConsumer
	clientApi.ApplicationPkcs10Consumer = noOpConsumer
	clientApi.ApplicationXPemFileProducer = &PemProducer{}
	clientApi.TextYamlProducer = &YamlProducer{}
	clientApi.ZtSessionAuth = func(token string) (principal interface{}, err error) {
		principal, err = ae.GetManagers().ApiSession.ReadByToken(token)

		if err != nil {
			if !boltz.IsErrNotFoundErr(err) {
				pfxlog.Logger().WithError(err).Errorf("encountered error checking for session that was not expected; returning masking unauthorized response")
			}

			return nil, errorz.NewUnauthorized()
		}

		return principal, nil
	}

	managementApi.TextYamlProducer = &YamlProducer{}
	managementApi.ZtSessionAuth = clientApi.ZtSessionAuth

	ae.ApiClientCsrSigner = cert.NewClientSigner(ae.Config.Enrollment.SigningCert.Cert().Leaf, ae.Config.Enrollment.SigningCert.Cert().PrivateKey)
	ae.ApiServerCsrSigner = cert.NewServerSigner(ae.Config.Enrollment.SigningCert.Cert().Leaf, ae.Config.Enrollment.SigningCert.Cert().PrivateKey)
	ae.ControlClientCsrSigner = cert.NewClientSigner(ae.Config.Enrollment.SigningCert.Cert().Leaf, ae.Config.Enrollment.SigningCert.Cert().PrivateKey)

	ae.FingerprintGenerator = cert.NewFingerprintGenerator()

	if err != nil {
		log := pfxlog.Logger()
		log.WithField("cause", err).Fatal("could not load schemas")
	}

	return ae
}

func (ae *AppEnv) InitPersistence() error {
	var err error

	ae.BoltStores, err = persistence.NewBoltStores(ae.HostController.GetNetwork())
	if err != nil {
		return err
	}

	if err = persistence.RunMigrations(ae.GetDbProvider().GetDb(), ae.BoltStores); err != nil {
		return err
	}

	ae.BoltStores.EventualEventer.AddListener(persistence.EventualEventAddedName, func(i ...interface{}) {
		if len(i) == 0 {
			pfxlog.Logger().Errorf("could not update metrics for %s gauge on add, event argument length was 0", EventualEventsGauge)
			return
		}

		if event, ok := i[0].(*persistence.EventualEventAdded); ok {
			gauge := ae.GetHostController().GetNetwork().GetMetricsRegistry().Gauge(EventualEventsGauge)
			gauge.Update(event.Total)
		} else {
			pfxlog.Logger().Errorf("could not update metrics for %s gauge on add, event argument was %T expected *EventualEventAdded", EventualEventsGauge, i[0])
		}
	})
	ae.BoltStores.EventualEventer.AddListener(persistence.EventualEventRemovedName, func(i ...interface{}) {
		if len(i) == 0 {
			pfxlog.Logger().Errorf("could not update metrics for %s gauge on remove, event argument length was 0", EventualEventsGauge)
			return
		}

		if event, ok := i[0].(*persistence.EventualEventRemoved); ok {
			gauge := ae.GetHostController().GetNetwork().GetMetricsRegistry().Gauge(EventualEventsGauge)
			gauge.Update(event.Total)
		} else {
			pfxlog.Logger().Errorf("could not update metrics for %s gauge on remove, event argument was %T expected *EventualEventRemoved", EventualEventsGauge, i[0])
		}
	})

	ae.Managers = model.InitEntityManagers(ae)
	ae.EventDispatcher = events.NewDispatcher(ae.GetHostController().GetNetwork(), ae.GetDbProvider(), ae.BoltStores, ae.GetHostController().GetCloseNotifyChannel())

	persistence.ServiceEvents.AddServiceEventHandler(ae.HandleServiceEvent)
	ae.BoltStores.Identity.AddEntityIdListener(ae.IdentityRefreshMap.Remove, boltz.EntityDeletedAsync)

	return err
}

func getJwtSigningMethod(cert *tls.Certificate) jwt.SigningMethod {

	var sm jwt.SigningMethod = jwt.SigningMethodNone

	switch cert.Leaf.PublicKey.(type) {
	case *ecdsa.PublicKey:
		key := cert.Leaf.PublicKey.(*ecdsa.PublicKey)
		switch key.Params().BitSize {
		case jwt.SigningMethodES256.CurveBits:
			sm = jwt.SigningMethodES256
		case jwt.SigningMethodES384.CurveBits:
			sm = jwt.SigningMethodES384
		case jwt.SigningMethodES512.CurveBits:
			sm = jwt.SigningMethodES512
		default:
			pfxlog.Logger().Panic("unsupported EC key size: ", key.Params().BitSize)
		}
	case *rsa.PublicKey:
		sm = jwt.SigningMethodRS256
	default:
		pfxlog.Logger().Panic("unknown certificate type, unable to determine signing method")
	}

	return sm
}

func (ae *AppEnv) GetSessionTokenFromRequest(r *http.Request) string {
	return r.Header.Get(ZitiSession)
}

func (ae *AppEnv) CreateRequestContext(rw http.ResponseWriter, r *http.Request) *response.RequestContext {
	rid := eid.New()

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewReader(body))

	requestContext := &response.RequestContext{
		Id:                rid,
		ResponseWriter:    rw,
		Request:           r,
		Body:              body,
		Identity:          nil,
		ApiSession:        nil,
		ActivePermissions: []string{},
		StartTime:         time.Now(),
	}

	requestContext.Responder = response.NewResponder(requestContext)

	return requestContext
}

func GetRequestContextFromHttpContext(r *http.Request) (*response.RequestContext, error) {
	val := r.Context().Value(api.ZitiContextKey)
	if val == nil {
		return nil, fmt.Errorf("value for key %s no found in context", api.ZitiContextKey)
	}

	requestContext := val.(*response.RequestContext)

	if requestContext == nil {
		return nil, fmt.Errorf("value for key %s is not a request context", api.ZitiContextKey)
	}

	return requestContext, nil
}

// getMetricTimerName returns a metric timer name based on the incoming HTTP request's URL and method.
// Unique ids are removed from the URL and replaced with :id and :subid to group metrics from the same
// endpoint that happen to be working on different ids.
func getMetricTimerName(r *http.Request) string {
	cleanUrl := r.URL.Path

	rc, _ := api.GetRequestContextFromHttpContext(r)

	if rc != nil {
		if id, err := rc.GetEntityId(); err == nil && id != "" {
			cleanUrl = strings.Replace(cleanUrl, id, ":id", -1)
		}

		if subid, err := rc.GetEntitySubId(); err == nil && subid != "" {
			cleanUrl = strings.Replace(cleanUrl, subid, ":subid", -1)
		}
	}

	return fmt.Sprintf("%s.%s", cleanUrl, r.Method)
}

func (ae *AppEnv) IsAllowed(responderFunc func(ae *AppEnv, rc *response.RequestContext), request *http.Request, entityId string, entitySubId string, permissions ...permissions.Resolver) openApiMiddleware.Responder {
	return openApiMiddleware.ResponderFunc(func(writer http.ResponseWriter, producer runtime.Producer) {

		rc, err := GetRequestContextFromHttpContext(request)

		if rc == nil {
			rc = ae.CreateRequestContext(writer, request)
		}

		rc.SetProducer(producer)
		rc.SetEntityId(entityId)
		rc.SetEntitySubId(entitySubId)

		if err != nil {
			pfxlog.Logger().WithError(err).Error("could not retrieve request context")
			rc.RespondWithError(err)
			return
		}

		for _, permission := range permissions {
			if !permission.IsAllowed(rc.ActivePermissions...) {
				rc.RespondWithApiError(errorz.NewUnauthorized())
				return
			}
		}

		responderFunc(ae, rc)

		if !rc.StartTime.IsZero() {
			timer := ae.GetHostController().GetNetwork().GetMetricsRegistry().Timer(getMetricTimerName(rc.Request))
			timer.UpdateSince(rc.StartTime)
		} else {
			pfxlog.Logger().WithFields(map[string]interface{}{
				"url": request.URL,
			}).Warn("could not mark metrics for REST ApiConfig endpoint, request context start time is zero")
		}
	})
}

func (ae *AppEnv) HandleServiceEvent(event *persistence.ServiceEvent) {
	ae.HandleServiceUpdatedEventForIdentityId(event.IdentityId)
}

func (ae *AppEnv) HandleServiceUpdatedEventForIdentityId(identityId string) {
	ae.IdentityRefreshMap.Set(identityId, time.Now().UTC())
	ae.identityRefreshMeter.Mark(1)
}

func (ae *AppEnv) SetEnrollmentSigningCert(serverCert *tls.Certificate) {
	signMethod := getJwtSigningMethod(serverCert)
	ae.enrollmentSigner = jwtsigner.New(ae.Config.Api.Address, signMethod, serverCert.PrivateKey)
}
