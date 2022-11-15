package Session

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IEvent"
	"github.com/kmsar/laravel-go/Framework/Contracts/ISession"
	"github.com/kmsar/laravel-go/Framework/Http"
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
