package Session

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IHttp"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IPipeline"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/ISession"
)

func StartSession(session ISession.Session, request IHttp.IHttpRequest, next IPipeline.Pipe) interface{} {
	session.Start()
	return next(request)
}
