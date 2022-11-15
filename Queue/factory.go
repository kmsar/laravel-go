package Queue

import (
	"fmt"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IQueue"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Exceptions"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Field"
)

type Factory struct {
	queues       map[string]IQueue.Queue
	queueDrivers map[string]IQueue.QueueDriver
	config       Config
	serializer   IQueue.JobSerializer
}

func (factory *Factory) Connection(name ...string) IQueue.Queue {
	if len(name) > 0 {
		return factory.Queue(name[0])
	}

	return factory.Queue(factory.config.Defaults.Connection)
}

func (factory *Factory) Extend(name string, driver IQueue.QueueDriver) {
	factory.queueDrivers[name] = driver
}

func (factory *Factory) Queue(name string) IQueue.Queue {
	if queue, exists := factory.queues[name]; exists {
		return queue
	}

	config := factory.config.Connections[name]
	driver := Field.GetStringField(config, "driver")
	if config["default"] == nil {
		config["default"] = factory.config.Defaults.Queue
	}

	if queueDriver, exists := factory.queueDrivers[driver]; exists {
		factory.queues[name] = queueDriver(name, config, factory.serializer)
		return factory.queues[name]
	}

	panic(DriverException{
		Exception: Exceptions.New(fmt.Sprintf("unsupported queue driver：%s", driver), config),
	})
}
