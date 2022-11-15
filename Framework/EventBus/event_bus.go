package EventBus

import (
	"fmt"
	"reflect"
	"sort"
	"sync"
)

// BusSubscriber defines subscription-related bus behavior
type BusSubscriber interface {
	Subscribe(topic string, fn EventCallbackFunc, priority ...int) error
	SubscribeAsync(topic string, fn EventCallbackFunc, transactional bool, priority ...int) error
	SubscribeOnce(topic string, fn EventCallbackFunc, priority ...int) error
	SubscribeOnceAsync(topic string, fn EventCallbackFunc, priority ...int) error
	Unsubscribe(topic string, handler EventCallbackFunc) error
}

// BusPublisher defines publishing-related bus behavior
type BusPublisher interface {
	Publish(e *Event)
}

// BusController defines bus control behavior (checking handler's presence, synchronization)
type BusController interface {
	HasCallback(topic string) bool
	WaitAsync()
}

// Bus englobes global (subscribe, publish, control) bus behavior
type Bus interface {
	BusController
	BusSubscriber
	BusPublisher
}

// EventBus - box for subscribers and callbacks.
type EventBus struct {
	subscribers map[string][]*eventHandler
	lock        sync.Mutex // a lock for the map
	wg          sync.WaitGroup
}

// Event type holds the details of single event.
type Event struct {
	Name string
	Data interface{}
}

// EventCallbackFunc is signature of event callback function.
type EventCallbackFunc func(e *Event)

type eventHandler struct {
	callBack      EventCallbackFunc
	flagOnce      bool
	async         bool
	transactional bool
	sync.Mutex    // lock for an event handler - useful for running async callbacks serially
	priority      int
}

// New returns new EventBus with empty subscribers.
func New() Bus {
	b := &EventBus{
		make(map[string][]*eventHandler),
		sync.Mutex{},
		sync.WaitGroup{},
	}
	return Bus(b)
}

// doSubscribe handles the subscription logic and is utilized by the public Subscribe functions
func (bus *EventBus) doSubscribe(topic string, fn interface{}, handler *eventHandler) error {
	bus.lock.Lock()
	defer bus.lock.Unlock()
	if !(reflect.TypeOf(fn).Kind() == reflect.Func) {
		return fmt.Errorf("%s is not of type reflect.Func", reflect.TypeOf(fn).Kind())
	}
	bus.subscribers[topic] = append(bus.subscribers[topic], handler)
	return nil
}

// Subscribe subscribes to a topic.
// Returns error if `fn` is not a function.
func (bus *EventBus) Subscribe(topic string, fn EventCallbackFunc, priority ...int) error {
	return bus.doSubscribe(topic, fn, &eventHandler{
		callBack:      fn,
		flagOnce:      false,
		async:         false,
		transactional: false,
		Mutex:         sync.Mutex{},
		priority:      parsePriority(priority),
	})
}

// SubscribeAsync subscribes to a topic with an asynchronous callback
// Transactional determines whether subsequent callbacks for a topic are
// run serially (true) or concurrently (false)
// Returns error if `fn` is not a function.
func (bus *EventBus) SubscribeAsync(topic string, fn EventCallbackFunc, transactional bool, priority ...int) error {
	return bus.doSubscribe(topic, fn, &eventHandler{

		callBack:      fn,
		flagOnce:      false,
		async:         true,
		transactional: transactional,
		Mutex:         sync.Mutex{},
		priority:      parsePriority(priority),
	})
}

// SubscribeOnce subscribes to a topic once. Handler will be removed after executing.
// Returns error if `fn` is not a function.
func (bus *EventBus) SubscribeOnce(topic string, fn EventCallbackFunc, priority ...int) error {
	return bus.doSubscribe(topic, fn, &eventHandler{

		callBack:      fn,
		flagOnce:      true,
		async:         false,
		transactional: false,
		Mutex:         sync.Mutex{},
		priority:      parsePriority(priority),
	})
}

// SubscribeOnceAsync subscribes to a topic once with an asynchronous callback
// Handler will be removed after executing.
// Returns error if `fn` is not a function.
func (bus *EventBus) SubscribeOnceAsync(topic string, fn EventCallbackFunc, priority ...int) error {
	return bus.doSubscribe(topic, fn, &eventHandler{
		callBack:      fn,
		flagOnce:      true,
		async:         true,
		transactional: false,
		Mutex:         sync.Mutex{},
		priority:      parsePriority(priority),
	})
}

// HasCallback returns true if exists any callback subscribed to the topic.
func (bus *EventBus) HasCallback(topic string) bool {
	bus.lock.Lock()
	defer bus.lock.Unlock()
	_, ok := bus.subscribers[topic]
	if ok {
		return len(bus.subscribers[topic]) > 0
	}
	return false
}

// Unsubscribe removes callback defined for a topic.
// Returns error if there are no callbacks subscribed to the topic.
func (bus *EventBus) Unsubscribe(topic string, handler EventCallbackFunc) error {
	bus.lock.Lock()
	defer bus.lock.Unlock()
	if _, ok := bus.subscribers[topic]; ok && len(bus.subscribers[topic]) > 0 {
		bus.removeHandler(topic, bus.findHandlerIdx(topic, handler))
		return nil
	}
	return fmt.Errorf("topic %s doesn't exist", topic)
}

// Publish executes callback defined for a topic. Any additional argument will be transferred to the callback.
func (bus *EventBus) Publish(e *Event) {
	bus.lock.Lock() // will unlock if handler is not found or always after setUpPublish
	defer bus.lock.Unlock()
	if handlers, ok := bus.subscribers[e.Name]; ok && 0 < len(handlers) {
		// Handlers slice may be changed by removeHandler and Unsubscribe during iteration,
		// so make a copy and iterate the copied slice.
		copyHandlers := make([]*eventHandler, len(handlers))
		copy(copyHandlers, handlers)
		for i, handler := range copyHandlers {
			if handler.flagOnce {
				bus.removeHandler(e.Name, i)
			}
			if !handler.async {
				bus.doPublish(handler, e)
			} else {
				bus.wg.Add(1)
				if handler.transactional {
					bus.lock.Unlock()
					handler.Lock()
					bus.lock.Lock()
				}
				go bus.doPublishAsync(handler, e)
			}
		}
	}
}

func (bus *EventBus) doPublish(handler *eventHandler, e *Event) {
	handler.callBack(e)
}

func (bus *EventBus) doPublishAsync(handler *eventHandler, e *Event) {
	defer bus.wg.Done()
	if handler.transactional {
		defer handler.Unlock()
	}
	bus.doPublish(handler, e)
}

func (bus *EventBus) removeHandler(topic string, idx int) {
	if _, ok := bus.subscribers[topic]; !ok {
		return
	}
	l := len(bus.subscribers[topic])

	if !(0 <= idx && idx < l) {
		return
	}

	copy(bus.subscribers[topic][idx:], bus.subscribers[topic][idx+1:])
	bus.subscribers[topic][l-1] = nil // or the zero value of T
	bus.subscribers[topic] = bus.subscribers[topic][:l-1]
}

// IsEventExists method returns true if given event is exists in the event store
// otherwise false.
func (bus *EventBus) IsEventExists(eventName string) bool {
	_, found := bus.subscribers[eventName]
	return found
}

func (bus *EventBus) findHandlerIdx(topic string, callback EventCallbackFunc) int {

	if _, ok := bus.subscribers[topic]; ok {
		for idx, handler := range bus.subscribers[topic] {
			handlerCallback := reflect.ValueOf(handler.callBack)
			vCallback := reflect.ValueOf(callback)
			if handlerCallback.Type() == vCallback.Type() &&
				handlerCallback.Pointer() == vCallback.Pointer() {
				return idx
			}
		}
	}
	return -1
}

// WaitAsync waits for all async callbacks to complete
func (bus *EventBus) WaitAsync() {
	bus.wg.Wait()
}

func (bus *EventBus) sortEventSubscribers(eventName string) {
	if bus.HasCallback(eventName) {
		ec := bus.subscribers[eventName]
		sort.Slice(ec, func(i, j int) bool { return ec[i].priority < ec[j].priority })
	}
}

func (bus *EventBus) sortAndPublishSync(e *Event) {
	bus.sortEventSubscribers(e.Name)
	bus.Publish(e)
}

func parsePriority(priority []int) int {
	pr := 1 // default priority is 1
	if len(priority) > 0 && priority[0] > 0 {
		pr = priority[0]
	}
	return pr
}
