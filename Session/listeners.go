package Session

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IEvent"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/ISession"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Http"
)

type RequestBeforeListener struct {
}

func (this *RequestBeforeListener) Handle(event IEvent.Event) {
}

type RequestAfterListener struct {
}

// Handle 如果开启了 session 那么请求结束时保存 session
func (this *RequestAfterListener) Handle(event IEvent.Event) {
	if responseBeforeEvent, ok := event.(*Http.ResponseBefore); ok {
		if session, isSession := responseBeforeEvent.Request().Get("session").(ISession.Session); isSession {
			if session.IsStarted() {
				session.Save()
			}
		}
	}
}
