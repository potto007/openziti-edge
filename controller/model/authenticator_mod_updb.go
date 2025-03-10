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

package model

import (
	"encoding/base64"
	"errors"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/edge/controller/apierror"
	"github.com/openziti/edge/controller/persistence"
	"github.com/openziti/foundation/v2/errorz"
	cmap "github.com/orcaman/concurrent-map/v2"
	"time"
)

var _ AuthProcessor = &AuthModuleUpdb{}

type AuthModuleUpdb struct {
	env                       Env
	method                    string
	attemptsByAuthenticatorId cmap.ConcurrentMap[string, int64]
}

func NewAuthModuleUpdb(env Env) *AuthModuleUpdb {
	return &AuthModuleUpdb{
		env:                       env,
		method:                    "password",
		attemptsByAuthenticatorId: cmap.New[int64](),
	}
}

func (module *AuthModuleUpdb) CanHandle(method string) bool {
	return method == module.method
}

func (module *AuthModuleUpdb) Process(context AuthContext) (AuthResult, error) {
	logger := pfxlog.Logger().WithField("authMethod", module.method)

	data := context.GetData()

	username := ""
	password := ""

	if usernameVal := data["username"]; usernameVal != nil {
		username = usernameVal.(string)
	}
	if passwordVal := data["password"]; passwordVal != nil {
		password = passwordVal.(string)
	}

	if username == "" || password == "" {
		return nil, errorz.NewCouldNotValidate(errors.New("username and password fields are required"))
	}

	logger = logger.WithField("username", username)
	authenticator, err := module.env.GetManagers().Authenticator.ReadByUsername(username)

	if err != nil {
		logger.WithError(err).Error("could not authenticate, authenticator lookup by username errored")
		return nil, err
	}

	if authenticator == nil {
		logger.WithError(err).Error("could not authenticate, authenticator lookup returned nil")
		return nil, apierror.NewInvalidAuth()
	}

	logger = logger.
		WithField("authenticatorId", authenticator.Id).
		WithField("identityId", authenticator.IdentityId)

	authPolicy, identity, err := getAuthPolicyByIdentityId(module.env, module.method, authenticator.Id, authenticator.IdentityId)

	if err != nil {
		logger.WithError(err).Errorf("could not look up auth policy by identity id")
		return nil, apierror.NewInvalidAuth()
	}

	if authPolicy == nil {
		logger.Error("auth policy look up returned nil")
		return nil, apierror.NewInvalidAuth()
	}

	if identity.Disabled {
		logger.
			WithField("disabledAt", identity.DisabledAt).
			WithField("disabledUntil", identity.DisabledUntil).
			Error("authentication failed, identity is disabled")
		return nil, apierror.NewInvalidAuth()
	}

	logger = logger.WithField("authPolicyId", authPolicy.Id)

	if !authPolicy.Primary.Updb.Allowed {
		logger.Error("auth policy does not allow updb authentication")
		return nil, apierror.NewInvalidAuth()
	}

	attempts := int64(0)
	module.attemptsByAuthenticatorId.Upsert(authenticator.Id, 0, func(exist bool, prevAttempts int64, newValue int64) int64 {
		if exist {
			attempts = prevAttempts + 1
			return attempts
		}

		return 0
	})

	if authPolicy.Primary.Updb.MaxAttempts != persistence.UpdbUnlimitedAttemptsLimit && attempts > authPolicy.Primary.Updb.MaxAttempts {
		logger.WithField("attempts", attempts).WithField("maxAttempts", authPolicy.Primary.Updb.MaxAttempts).Error("updb auth failed, max attempts exceeded")

		duration := time.Duration(authPolicy.Primary.Updb.LockoutDurationMinutes) * time.Minute
		if err = module.env.GetManagers().Identity.Disable(authenticator.IdentityId, duration, context.GetChangeContext()); err != nil {
			logger.WithError(err).Error("could not lock identity, unhandled error")
		}

		return nil, apierror.NewInvalidAuth()
	}

	updb := authenticator.ToUpdb()

	salt, err := decodeSalt(updb.Salt)

	if err != nil {
		return nil, apierror.NewInvalidAuth()
	}

	hr := module.env.GetManagers().Authenticator.ReHashPassword(password, salt)

	if updb.Password != hr.Password {

		return nil, apierror.NewInvalidAuth()
	}

	return &AuthResultBase{
		identity:        identity,
		identityId:      updb.IdentityId,
		authenticator:   authenticator,
		authenticatorId: authenticator.Id,
		env:             module.env,
	}, nil
}

func decodeSalt(s string) ([]byte, error) {
	salt := make([]byte, 1024)
	n, err := base64.StdEncoding.Decode(salt, []byte(s))

	if err != nil {
		return nil, err
	}
	return salt[:n], nil
}
