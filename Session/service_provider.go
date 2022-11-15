package Session

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConfig"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IEncryption"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IEvent"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFoundation"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IHttp"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IRedis"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/ISession"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Session/stores"
)

type ServiceProvider struct {
	app    IFoundation.IApplication
	config Config
}

func (this *ServiceProvider) Register(application IFoundation.IApplication) {
	this.app = application

	application.NamedBind("session", func(
		config IConfig.Config,
		request IHttp.IHttpRequest,
		encryptor IEncryption.Encryptor,
		redis IRedis.RedisFactory,
	) ISession.Session {
		if session, isSession := request.Get("session").(ISession.Session); isSession {
			return session
		}

		sessionConfig := config.Get("session").(Config)
		var store ISession.SessionStore

		switch sessionConfig.Driver {
		case "cookie":
			if sessionConfig.Encrypt {
				store = stores.CookieStore(sessionConfig.Name, sessionConfig.Lifetime, request, encryptor)
			} else {
				store = stores.CookieStore(sessionConfig.Name, sessionConfig.Lifetime, request, nil)
			}

		case "redis":
			store = stores.RedisStore(sessionConfig.Key, sessionConfig.Lifetime, redis.Connection(sessionConfig.Connection))
		}

		session := New(sessionConfig, request, store)

		request.Set("session", session)
		return session
	})
}

func (this *ServiceProvider) Start() error {
	this.app.Call(func(dispatcher IEvent.EventDispatcher) {
		dispatcher.Register("RESPONSE_BEFORE", &RequestAfterListener{})
	})
	return nil
}

func (this *ServiceProvider) Stop() {
}
