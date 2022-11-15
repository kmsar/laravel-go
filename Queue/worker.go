package Queue

import (
	"fmt"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Parallel"

	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IDatabase"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IQueue"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/ISerialize"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Logs"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Exceptions"

	"runtime/debug"
	"time"
)

type Worker struct {
	name             string
	queue            IQueue.Queue
	closeChan        chan bool
	exceptionHandler IExeption.ExceptionHandler
	workers          *Parallel.Workers
	config           WorkerConfig

	db              IDatabase.DBConnection
	serializer      ISerialize.ClassSerializer
	failedJobsTable string
	dbIsReady       bool
}

type WorkerParam struct {
	handler         IExeption.ExceptionHandler
	db              IDatabase.DBConnection
	failedJobsTable string
	config          WorkerConfig
	serializer      ISerialize.ClassSerializer
}

func NewWorker(name string, queue IQueue.Queue, param WorkerParam) IQueue.QueueWorker {
	return &Worker{
		db:               param.db,
		dbIsReady:        true,
		failedJobsTable:  param.failedJobsTable,
		serializer:       param.serializer,
		name:             name,
		queue:            queue,
		closeChan:        make(chan bool),
		exceptionHandler: param.handler,
		config:           param.config,
	}
}

func (worker *Worker) workQueue(queue IQueue.Queue) {
	defer func() {
		if err := recover(); err != nil {
			Logs.WithException(Exceptions.WithRecover(err, nil)).Error("worker.workQueue failed")
		}
	}()
	var msgPipe = queue.Listen(worker.config.Queue...)
	Logs.Default().Info(fmt.Sprintf("queue.Worker.workQueue: %s worker is working...", worker.name))
	for {
		select {
		case msg := <-msgPipe:
			var err = worker.workers.Handle(func() {
				var job = msg.Job
				Logs.Default().WithField("job", job).Debug(fmt.Sprintf("queue.Worker.workQueue: processing job"))
				if err := worker.handleJob(job); err != nil {
					Logs.Default().WithField("job", job).Debug(fmt.Sprintf("queue.Worker.workQueue: Failed to process job"))
					job.Fail(err)
					if (job.GetMaxTries() > 0 && job.GetAttemptsNum() >= job.GetMaxTries()) || job.GetAttemptsNum() >= worker.config.Tries { // 达到最大尝试次数
						// 保存到死信队列
						if saveErr := worker.saveOnFailedJobs(msg.Job); saveErr != nil {
							panic(err)
						}
					} else {
						// 放回队列中重试
						if err = queue.Later(time.Now().Add(time.Second*time.Duration(job.GetRetryInterval())), job); err != nil {
							Logs.WithError(err).Warn("queue.Worker.workQueue: job release failed")
							panic(err)
						}
					}
					msg.Ack()
					worker.exceptionHandler.Handle(JobException{Exception: Exceptions.WithError(err, Support.Fields{
						"msg": msg,
					})})
				} else {
					Logs.Default().WithField("job", job).Debug(fmt.Sprintf("queue.Worker.workQueue: Processing job succeeded"))
					msg.Ack()
				}
			})
			if err != nil {
				return
			}
		case <-worker.closeChan:
			queue.Stop()
			return
		}
	}
}

func (worker *Worker) Work() {
	worker.workers = Parallel.NewWorkers(worker.config.Processes)
	worker.workQueue(worker.queue)
}

func (worker *Worker) Stop() {
	worker.closeChan <- true
	worker.workers.Stop()
}

// saveOnFailedJobs 保存死信
func (worker *Worker) saveOnFailedJobs(job IQueue.Job) (err error) {
	if worker.dbIsReady && worker.db != nil {
		_, err = worker.db.Exec(
			fmt.Sprintf("insert into %s (connection, queue, payload, exception) values ('%s','%s','%s','%s')",
				worker.failedJobsTable,
				job.GetConnectionName(),
				job.GetQueue(),
				worker.serializer.Serialize(job),
				debug.Stack(),
			),
		)
		if err != nil {
			Logs.WithError(err).Warn("queue.Worker.saveOnFailedJobs: Failed to save to database")
			worker.dbIsReady = false
		}
	}

	if err != nil || !worker.dbIsReady { // 如果没有配置数据库死信，或者保存到数据库失败了
		if err = worker.queue.Push(job, fmt.Sprintf("deaded_%s", job.GetQueue())); err != nil {
			Logs.WithError(err).Error("queue.Worker.saveOnFailedJobs: failed to save")
		}
	}
	return
}

func (worker *Worker) handleJob(job IQueue.Job) (err error) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			err = Exceptions.ResolveException(panicValue)
		}
	}()

	job.IncrementAttemptsNum()
	job.Handle()

	return nil
}
