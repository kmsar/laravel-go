package Queue

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IConfig"
	"github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"
	"github.com/kmsar/laravel-go/Framework/Contracts/IExeption"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
	"github.com/kmsar/laravel-go/Framework/Contracts/IQueue"
	"github.com/kmsar/laravel-go/Framework/Contracts/ISerialize"
	"github.com/kmsar/laravel-go/Framework/Queue/drivers"
)

type ServiceProvider struct {
	app     IFoundation.IApplication
	workers []IQueue.QueueWorker
}

func (this *ServiceProvider) Register(application IFoundation.IApplication) {

	application.NamedSingleton("queue.factory", func(config IConfig.Config, serializer IQueue.JobSerializer) IQueue.QueueFactory {
		return &Factory{
			serializer: serializer,
			queues:     map[string]IQueue.Queue{},
			queueDrivers: map[string]IQueue.QueueDriver{
				//"kafka": drivers.KafkaDriver,
				"nsq": drivers.NsqDriver,
			},
			config: config.Get("queue").(Config),
		}
	})
	application.NamedSingleton("queue", func(factory IQueue.QueueFactory) IQueue.Queue {
		return factory.Connection()
	})
	application.NamedSingleton("job.serializer", func(serializer ISerialize.ClassSerializer) IQueue.JobSerializer {
		return NewJobSerializer(serializer)
	})
	this.app = application
}

func (this *ServiceProvider) Start() error {
	this.runWorkers()
	return nil
}
func (this *ServiceProvider) Boot(application IFoundation.IApplication) {
	//TODO implement me
	panic("implement me")
}

// runWorkers 运行所有 worker
func (this *ServiceProvider) runWorkers() {
	this.app.Call(func(factory IQueue.QueueFactory, config IConfig.Config, handler IExeption.ExceptionHandler, db IDatabase.DBFactory, serializer ISerialize.ClassSerializer) {
		var (
			queueConfig = config.Get("queue").(Config)
			env         = this.app.Environment()
		)

		if queueConfig.Workers[env] != nil {
			for name, workerConfig := range queueConfig.Workers[env] {
				worker := NewWorker(name, factory.Connection(workerConfig.Connection), WorkerParam{
					handler:         handler,
					db:              db.Connection(queueConfig.Failed.Database),
					failedJobsTable: queueConfig.Failed.Table,
					config:          workerConfig,
					serializer:      serializer,
				})
				this.workers = append(this.workers, worker)
				go worker.Work()
			}
		}
	})
}

func (this *ServiceProvider) Stop() {
	for _, worker := range this.workers {
		worker.Stop()
	}
}
