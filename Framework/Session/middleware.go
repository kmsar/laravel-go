package Session

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IHttp"
	"github.com/kmsar/laravel-go/Framework/Contracts/IPipeline"
	"github.com/kmsar/laravel-go/Framework/Contracts/ISession"
)

func StartSession(session ISession.Session, request IHttp.IHttpRequest, next IPipeline.Pipe) interface{} {
	session.Start()
	return next(request)
}
