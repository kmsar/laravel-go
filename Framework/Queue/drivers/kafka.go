package drivers

//
//import (
//	"context"
//	"github.com/segmentio/kafka-go"
//	"github.com/kmsar/laravel-go/Framework/Contracts/IQueue"
//	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
//	"github.com/kmsar/laravel-go/Framework/Support/supports-master/logs"
//	"github.com/kmsar/laravel-go/Framework/Support/supports-master/utils"
//	"time"
//)
//
//func KafkaDriver(name string, config Support.Fields, serializer IQueue.JobSerializer) IQueue.Queue {
//	var (
//		dialer *kafka.Dialer
//		ok     bool
//	)
//
//	if dialer, ok = config["dialer"].(*kafka.Dialer); !ok {
//		dialer = &kafka.Dialer{Timeout: 10 * time.Second, DualStack: true}
//	}
//	return &Kafka{
//		name:         name,
//		brokers:      config["brokers"].([]string),
//		defaultQueue: config["default"].(string),
//		delayQueue:   config["delay"].(string),
//		serializer:   serializer,
//		dialer:       dialer,
//		readers:      make(map[string]*kafka.Reader),
//	}
//}
//
//type Kafka struct {
//	delayQueue   string
//	name         string
//	brokers      []string
//	defaultQueue string
//	serializer   IQueue.JobSerializer
//	stopped      bool
//	dialer       *kafka.Dialer
//	readers      map[string]*kafka.Reader
//	writer       *kafka.Writer
//}
//
//func (this *Kafka) getQueue(queues []string, queue string) string {
//	if len(queues) > 0 {
//		return queues[0]
//	}
//	if queue != "" {
//		return queue
//	}
//	return this.defaultQueue
//}
//
//func (this *Kafka) getReader(queue string) *kafka.Reader {
//	if this.readers[queue] != nil {
//		return this.readers[queue]
//	}
//	this.readers[queue] = kafka.NewReader(kafka.ReaderConfig{
//		Brokers: this.brokers,
//		GroupID: this.name,
//		Topic:   queue,
//		Dialer:  this.dialer,
//	})
//	return this.readers[queue]
//}
//
//func (this *Kafka) getWriter() *kafka.Writer {
//	if this.writer != nil {
//		return this.writer
//	}
//	this.writer = &kafka.Writer{
//		Addr:     kafka.TCP(this.brokers[0]),
//		Balancer: &kafka.LeastBytes{},
//	}
//	return this.writer
//}
//
//func (this *Kafka) Push(job IQueue.Job, queue ...string) error {
//	return this.PushOn(this.getQueue(queue, job.GetQueue()), job)
//}
//
//func (this *Kafka) PushOn(queue string, job IQueue.Job) error {
//	job.SetQueue(queue)
//
//	err := this.getWriter().WriteMessages(context.Background(), kafka.Message{
//		Topic: queue,
//		Key:   []byte(job.Uuid()),
//		Value: []byte(this.serializer.Serializer(job)),
//	})
//	if err != nil {
//		logs.WithError(err).WithField("job", job).Debug("push on queue failed")
//	}
//	return err
//}
//
//func (this *Kafka) PushRaw(payload, queue string, options ...Support.Fields) error {
//	err := this.getWriter().WriteMessages(context.Background(), kafka.Message{
//		Topic: queue,
//		Key:   []byte(utils.RandStr(5)),
//		Value: []byte(payload),
//	})
//	if err != nil {
//		logs.WithError(err).
//			WithField("queue", queue).
//			WithField("payload", payload).
//			Debug("push on queue failed")
//	}
//	return err
//}
//
//func (this *Kafka) Later(delay time.Time, job IQueue.Job, queue ...string) error {
//	return this.LaterOn(this.getQueue(queue, job.GetQueue()), delay, job)
//}
//
//func (this *Kafka) LaterOn(queue string, delay time.Time, job IQueue.Job) error {
//	job.SetQueue(queue)
//
//	err := this.getWriter().WriteMessages(context.Background(), kafka.Message{
//		Topic: this.delayQueue,
//		Key:   []byte(job.Uuid()),
//		Value: []byte(this.serializer.Serializer(job)),
//		Time:  delay,
//	})
//	if err != nil {
//		logs.WithError(err).WithField("job", job).Debug("push on queue failed")
//	}
//	return err
//}
//
//func (this *Kafka) GetConnectionName() string {
//	return this.name
//}
//
//func (this *Kafka) Release(job IQueue.Job, delay ...int) error {
//	delayAt := time.Now()
//	if len(delay) > 0 {
//		delayAt = delayAt.Add(time.Second * time.Duration(delay[0]))
//	}
//
//	return this.Later(delayAt, job)
//}
//
//func (this *Kafka) Stop() {
//	this.stopped = true
//}
//
//func (this *Kafka) Listen(queue ...string) chan IQueue.Msg {
//	this.stopped = false
//	ch := make(chan IQueue.Msg)
//	for _, name := range queue {
//		go this.consume(this.getReader(name), ch)
//	}
//	go this.maintainDelayQueue()
//	return ch
//}
//
//func (this *Kafka) consume(reader *kafka.Reader, ch chan IQueue.Msg) {
//	ctx := context.Background()
//	for {
//		if this.stopped {
//			break
//		}
//		msg, err := reader.FetchMessage(ctx)
//		if err != nil {
//			logs.WithError(err).WithField("config", reader.Config()).Error("kafka.consume: FetchMessage failed")
//			break
//		}
//		if msg.Time.Sub(time.Now()) > 0 {
//			logs.Default().Info("未来的消息" + msg.Time.String() + string(msg.Value))
//			continue
//		}
//		job, err := this.serializer.Unserialize(string(msg.Value))
//		if err != nil {
//			logs.WithError(err).WithField("msg", string(msg.Value)).WithField("config", reader.Config()).Error("kafka.consume: Unserialize job failed")
//			break
//		}
//		(func(message kafka.Message) {
//			ch <- contracts.Msg{
//				Ack: func() {
//					if err = reader.CommitMessages(ctx, message); err != nil {
//						logs.WithError(err).WithField("message", message).Error("kafka.consume: CommitMessages failed")
//					}
//				},
//				Job: job,
//			}
//		})(msg)
//	}
//}
//
//func (this *Kafka) maintainDelayQueue() {
//	reader := this.getReader(this.delayQueue)
//	ctx := context.Background()
//	for {
//		if this.stopped {
//			break
//		}
//		msg, err := reader.FetchMessage(ctx)
//		if err != nil {
//			logs.WithError(err).WithField("config", reader.Config()).Error("kafka.maintainDelayQueue: FetchMessage failed")
//			break
//		}
//		job, err := this.serializer.Unserialize(string(msg.Value))
//		if err != nil {
//			logs.WithError(err).WithField("msg", string(msg.Value)).WithField("config", reader.Config()).Error("kafka.consume: Unserialize job failed")
//			break
//		}
//		if msg.Time.Sub(time.Now()) > 0 { // 还没到时间
//			go (func(message kafka.Message) {
//				err = this.LaterOn(job.GetQueue(), msg.Time, job)
//				if err != nil {
//					logs.WithError(err).WithField("queue", string(message.Value)).Error("kafka.maintainDelayQueue: LaterOn failed")
//					return
//				}
//				if err = reader.CommitMessages(ctx, message); err != nil {
//					logs.WithError(err).WithField("message", string(message.Value)).Error("kafka.maintainDelayQueue: CommitMessages failed(delay)")
//				}
//			})(msg)
//		} else {
//			go (func(message kafka.Message) {
//				if err = this.PushOn(job.GetQueue(), job); err != nil {
//					logs.WithError(err).WithField("queue", job.GetQueue()).Error("kafka.maintainDelayQueue: PushOn failed")
//					return
//				}
//				if err = reader.CommitMessages(ctx, message); err != nil {
//					logs.WithError(err).WithField("message", string(message.Value)).Error("kafka.maintainDelayQueue: CommitMessages failed")
//				}
//			})(msg)
//		}
//	}
//}
