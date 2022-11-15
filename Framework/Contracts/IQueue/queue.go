package IQueue

import "github.com/kmsar/laravel-go/Framework/Contracts/Support"
import "time"

type QueueFactory interface {

	// Connection Resolve a queue connection instance.
	Connection(name ...string) Queue

	// Extend Add a queue connection resolver.
	Extend(name string, driver QueueDriver)
}

type Msg struct {
	Ack Ack
	Job Job
}

type Ack func()

// QueueDriver Get queue connection with given information.
type QueueDriver func(name string, config Support.Fields, serializer JobSerializer) Queue

type Queue interface {

	// Push a new job onto the queue.
	Push(job Job, queue ...string) error

	// PushOn push a new job onto the queue.
	PushOn(queue string, job Job) error

	// PushRaw push a raw payload onto the queue.
	PushRaw(payload, queue string, options ...Support.Fields) error

	// Later push a new job onto the queue after a delay.
	Later(delay time.Time, job Job, queue ...string) error

	// LaterOn push a new job onto the queue after a delay.
	LaterOn(queue string, delay time.Time, job Job) error

	// GetConnectionName Get the connection name for the queue.
	GetConnectionName() string

	// Release release the job back into the queue.
	// Accepts a delay specified in seconds.
	Release(job Job, delay ...int) error

	// Listen listen to the given queue.
	Listen(queue ...string) chan Msg

	// Stop close queue.
	Stop()
}

type Job interface {

	// Uuid Get the UUID of the job.
	Uuid() string

	// GetOptions Get the decoded body of the job.
	GetOptions() Support.Fields

	// Handle 执行工作
	// the job.
	Handle()

	// IsReleased Determine if the job was released back into the queue.
	IsReleased() bool

	// IsDeleted Determine if the job has been deleted.
	IsDeleted() bool

	// IsDeletedOrReleased Determine if the job has been deleted or released.
	IsDeletedOrReleased() bool

	// Attempts Get the number of times the job has been attempted.
	Attempts() int

	// HasFailed Determine if the job has been marked as a failure.
	HasFailed() bool

	// MarkAsFailed Mark the job as "failed".
	MarkAsFailed()

	// Fail Delete the job, call the "failed" method, and raise the failed job event.
	Fail(err error)

	// GetMaxTries Get the max number of times to attempt a job.
	GetMaxTries() int

	// GetRetryInterval Get the interval between retry interval tasks, in seconds
	GetRetryInterval() int

	// GetAttemptsNum Get the number of times to attempt a job.
	GetAttemptsNum() int

	// IncrementAttemptsNum increase the number of attempts.
	IncrementAttemptsNum()

	// GetTimeout Get the number of seconds the job can run.
	GetTimeout() int

	// GetConnectionName Get the name of the connection the job belongs to.
	GetConnectionName() string

	// GetQueue Get the name of the queue the job belongs to.
	GetQueue() string

	// SetQueue Sets the name of the queue to which the job belongs.
	SetQueue(queue string)
}

type QueueWorker interface {

	// Work perform work.
	Work()

	// Stop  working.
	Stop()
}

type JobSerializer interface {

	// Serializer Serialize "Job instance" to string
	Serializer(job Job) string

	// Unserialize Convert the serialized string to a "Job instance"
	Unserialize(serialized string) (Job, error)
}

type ShouldQueue interface {

	// ShouldQueue Determine whether to queue.
	ShouldQueue() bool
}

type ShouldBeUnique interface {

	// ShouldBeUnique determine whether it is unique.
	ShouldBeUnique() bool
}

type ShouldBeEncrypted interface {

	// ShouldBeEncrypted Determine whether to encrypt.
	ShouldBeEncrypted() bool
}
