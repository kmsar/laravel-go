
#### Example
```go
func calculator(a int, b int) {
	fmt.Printf("%d\n", a + b)
}

func main() {
	bus := EventBus.New();
	bus.Subscribe("main:calculator", calculator);
	bus.Publish("main:calculator", 20, 40);
	bus.Unsubscribe("main:calculator", calculator);
}
```

#### Implemented methods
* **New()**
* **Subscribe()**
* **SubscribeOnce()**
* **HasCallback()**
* **Unsubscribe()**
* **Publish()**
* **SubscribeAsync()**
* **SubscribeOnceAsync()**
* **WaitAsync()**

#### New()
New returns new EventBus with empty handlers.
```go
bus := EventBus.New();
```

#### Subscribe(topic string, fn interface{}) error
Subscribe to a topic. Returns error if `fn` is not a function.
```go
func Handler() { ... }
...
bus.Subscribe("topic:handler", Handler)
```

#### SubscribeOnce(topic string, fn interface{}) error
Subscribe to a topic once. Handler will be removed after executing. Returns error if `fn` is not a function.
```go
func HelloWorld() { ... }
...
bus.SubscribeOnce("topic:handler", HelloWorld)
```

#### Unsubscribe(topic string, fn interface{}) error
Remove callback defined for a topic. Returns error if there are no callbacks subscribed to the topic.
```go
bus.Unsubscribe("topic:handler", HelloWord);
```

#### HasCallback(topic string) bool
Returns true if exists any callback subscribed to the topic.

#### Publish(topic string, args ...interface{})
Publish executes callback defined for a topic. Any additional argument will be transferred to the callback.
```go
func Handler(str string) { ... }
...
bus.Subscribe("topic:handler", Handler)
...
bus.Publish("topic:handler", "Hello, World!");
```

#### SubscribeAsync(topic string, fn interface{}, transactional bool)
Subscribe to a topic with an asynchronous callback. Returns error if `fn` is not a function.
```go
func slowCalculator(a, b int) {
	time.Sleep(3 * time.Second)
	fmt.Printf("%d\n", a + b)
}

bus := EventBus.New()
bus.SubscribeAsync("main:slow_calculator", slowCalculator, false)

bus.Publish("main:slow_calculator", 20, 60)

fmt.Println("start: do some stuff while waiting for a result")
fmt.Println("end: do some stuff while waiting for a result")

bus.WaitAsync() // wait for all async callbacks to complete

fmt.Println("do some stuff after waiting for result")
```
Transactional determines whether subsequent callbacks for a topic are run serially (true) or concurrently(false)

#### SubscribeOnceAsync(topic string, args ...interface{})
SubscribeOnceAsync works like SubscribeOnce except the callback to executed asynchronously

####  WaitAsync()
WaitAsync waits for all async callbacks to complete.

#### Cross Process Events
Works with two rpc services:
- a client service to listen to remotely published events from a server
- a server service to listen to client subscriptions

server.go
```go
func main() {
    server := NewServer(":2010", "/_server_bus_", New())
    server.Start()
    // ...
    server.EventBus().Publish("main:calculator", 4, 6)
    // ...
    server.Stop()
}
```

client.go
```go
func main() {
    client := NewClient(":2015", "/_client_bus_", New())
    client.Start()
    client.Subscribe("main:calculator", calculator, ":2010", "/_server_bus_")
    // ...
    client.Stop()
}
```

