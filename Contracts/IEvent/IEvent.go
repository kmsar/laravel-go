package IEvent

type Event interface {
	// Event get event name.
	Event() string
}

type SyncEvent interface {
	Event

	// Sync Determine whether to synchronize events.
	Sync() bool
}

type EventListener interface {
	// Handle  the event when the event is triggered, all listeners will
	Handle(event Event)
}

type EventDispatcher interface {

	// Register an event listener with the dispatcher.
	Register(name string, listener EventListener)

	// Dispatch Provide all listeners with an event to process.
	Dispatch(event Event)
}
