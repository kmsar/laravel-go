package drivers

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IQueue"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Logs"
	"time"
)

func NsqDriver(name string, config Support.Fields, serializer IQueue.JobSerializer) IQueue.Queue {
	var (
		nsqConfig *nsq.Config
		ok        bool
	)
	if nsqConfig, ok = config["config"].(*nsq.Config); !ok {
		nsqConfig = nsq.NewConfig()
	}

	return &Nsq{
		config:          nsqConfig,
		name:            name,
		address:         config["address"].(string),
		lookupAddresses: config["lookup_addresses"].([]string),
		defaultQueue:    config["default"].(string),
		serializer:      serializer,
		consumers:       make(map[string]*nsq.Consumer),
	}
}

type Nsq struct {
	name            string
	address         string
	lookupAddresses []string
	defaultQueue    string
	serializer      IQueue.JobSerializer
	stopped         bool
	config          *nsq.Config
	consumers       map[string]*nsq.Consumer
	producer        *nsq.Producer
}

func (this *Nsq) getQueue(queues []string, queue string) string {
	if len(queues) > 0 {
		return queues[0]
	}
	if queue != "" {
		return queue
	}
	return this.defaultQueue
}

func (this *Nsq) getConsumer(queue string) *nsq.Consumer {
	if this.consumers[queue] != nil {
		return this.consumers[queue]
	}
	consumer, err := nsq.NewConsumer(queue, "channel", this.config)
	if err != nil {
		panic(err)
	}

	this.consumers[queue] = consumer
	return consumer
}

func (this *Nsq) getProducer() *nsq.Producer {
	if this.producer != nil {
		return this.producer
	}

	producer, err := nsq.NewProducer(this.address, this.config)

	if err != nil {
		panic(err)
	}

	this.producer = producer

	return this.producer
}

func (this *Nsq) Push(job IQueue.Job, queue ...string) error {
	return this.PushOn(this.getQueue(queue, job.GetQueue()), job)
}

func (this *Nsq) PushOn(queue string, job IQueue.Job) error {
	job.SetQueue(queue)
	return this.getProducer().Publish(queue, []byte(this.serializer.Serializer(job)))
}

func (this *Nsq) PushRaw(payload, queue string, options ...Support.Fields) error {
	return this.getProducer().Publish(queue, []byte(payload))
}

func (this *Nsq) Later(delay time.Time, job IQueue.Job, queue ...string) error {
	return this.LaterOn(this.getQueue(queue, job.GetQueue()), delay, job)
}

func (this *Nsq) LaterOn(queue string, delay time.Time, job IQueue.Job) error {
	job.SetQueue(queue)

	return this.getProducer().DeferredPublish(queue, delay.Sub(time.Now()), []byte(this.serializer.Serializer(job)))
}

func (this *Nsq) GetConnectionName() string {
	return this.name
}

func (this *Nsq) Release(job IQueue.Job, delay ...int) error {
	delayAt := time.Now()
	if len(delay) > 0 {
		delayAt = delayAt.Add(time.Second * time.Duration(delay[0]))
	}

	return this.Later(delayAt, job)
}

func (this *Nsq) Stop() {
	this.stopped = true
	for _, consumer := range this.consumers {
		consumer.Stop()
	}
}

func (this *Nsq) Listen(queue ...string) chan IQueue.Msg {
	this.stopped = false
	ch := make(chan IQueue.Msg)

	for _, name := range queue {
		this.consume(this.getConsumer(name), ch)
	}

	return ch
}

func (this *Nsq) consume(consumer *nsq.Consumer, ch chan IQueue.Msg) {
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		var job, err = this.serializer.Unserialize(string(message.Body))
		if err != nil {
			Logs.WithError(err).WithField("msg", string(message.Body)).Error("nsq.consume: Unserialize job failed")
			return nil
		}
		ackChan := make(chan error)
		ch <- IQueue.Msg{
			Ack: func() {
				ackChan <- nil
			},
			Job: job,
		}
		return <-ackChan
	}))

	if len(this.lookupAddresses) > 0 && this.lookupAddresses[0] != "" {
		if err := consumer.ConnectToNSQLookupds(this.lookupAddresses); err != nil {
			panic(err)
		}
	} else {
		if err := consumer.ConnectToNSQD(this.address); err != nil {
			panic(err)
		}
	}

}
