package Events

import (
	"fmt"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Parallel"

	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IEvent"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"
	"sync"
)

func NewDispatcher(handler IExeption.ExceptionHandler) IEvent.EventDispatcher {
	return &EventDispatcher{
		eventListenersMap: sync.Map{},
		exceptionHandler:  handler,
	}
}

type EventDispatcher struct {
	eventListenersMap sync.Map

	// 依赖异常处理器
	exceptionHandler IExeption.ExceptionHandler
}

func (dispatcher *EventDispatcher) Register(name string, listener IEvent.EventListener) {
	dispatcher.eventListenersMap.Store(name, append(dispatcher.getListeners(name), listener))
}
func (dispatcher *EventDispatcher) getListeners(name string) []IEvent.EventListener {
	if value, exists := dispatcher.eventListenersMap.Load(name); exists {
		return value.([]IEvent.EventListener)
	}
	return nil
}

func (dispatcher *EventDispatcher) Dispatch(event IEvent.Event) {
	if e, isSync := event.(IEvent.SyncEvent); isSync && e.Sync() {
		// 同步执行事件
		dispatcher.handleEvent(event)
	} else {
		// 协程执行
		go func() {
			dispatcher.handleEvent(event)
		}()
	}
}

func (dispatcher *EventDispatcher) exceptionHandle(err interface{}, event IEvent.Event) {
	if err != nil {
		dispatcher.exceptionHandler.Handle(EventException{
			error:  fmt.Errorf("%v", err),
			fields: nil,
			event:  event,
		})
	}
}

func (dispatcher *EventDispatcher) handleEvent(event IEvent.Event) {
	defer func() {
		dispatcher.exceptionHandle(recover(), event)
	}()

	listeners := dispatcher.getListeners(event.Event())
	parallelInstance := Parallel.NewParallel(len(listeners))

	for _, listener := range listeners {
		_ = parallelInstance.Add(func() interface{} {
			listener.Handle(event)
			return nil
		})
	}

	for _, result := range parallelInstance.Wait() {
		dispatcher.exceptionHandle(result, event)
	}
}
