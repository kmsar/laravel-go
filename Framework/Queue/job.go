package Queue

import (
	"errors"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Support/Random"

	"time"
)

var JobFailedErr = errors.New("job failed")

type Error string

func (e Error) Error() string {
	return string(e)
}

type Job struct {
	UUID          string         `json:"uuid,omitempty"`
	CreatedAt     int64          `json:"created_at,omitempty"`
	Queue         string         `json:"queue,omitempty"`
	Connection    string         `json:"connection,omitempty"`
	Tries         int            `json:"tries,omitempty"`
	MaxTries      int            `json:"max_tries,omitempty"`
	IsDelete      bool           `json:"is_delete,omitempty"`
	Options       Support.Fields `json:"options,omitempty"`
	IsRelease     bool           `json:"is_released,omitempty"`
	Error         *Error         `json:"error,omitempty"`
	Timeout       int
	RetryInterval int
}

func BaseJob(queue string) *Job {
	return &Job{
		UUID:          Random.RandStr(30),
		CreatedAt:     time.Now().Unix(),
		Queue:         queue,
		Tries:         0,
		MaxTries:      0,
		RetryInterval: 3,
	}
}

func (job *Job) Handle() {
}

func (job *Job) Uuid() string {
	return job.UUID
}

func (job *Job) GetOptions() Support.Fields {
	return job.Options
}

func (job *Job) IsReleased() bool {
	return job.IsRelease
}

func (job *Job) IsDeleted() bool {
	return job.IsDelete
}

func (job *Job) IsDeletedOrReleased() bool {
	return job.IsDelete || job.IsRelease
}

func (job *Job) Attempts() int {
	return job.Tries
}

func (job *Job) HasFailed() bool {
	return job.Error != nil
}

func (job *Job) MarkAsFailed() {
	job.Fail(JobFailedErr)
}

func (job *Job) Fail(err error) {
	var e Error
	e = Error(err.Error())
	job.Error = &e
}

func (job *Job) GetMaxTries() int {
	return job.MaxTries
}

func (job *Job) GetAttemptsNum() int {
	return job.Tries
}

func (job *Job) GetRetryInterval() int {
	return job.RetryInterval
}

func (job *Job) IncrementAttemptsNum() {
	job.Tries++
}

func (job *Job) GetTimeout() int {
	return job.Timeout
}

func (job *Job) GetConnectionName() string {
	return job.Connection
}

func (job *Job) GetQueue() string {
	return job.Queue
}

func (job *Job) SetQueue(queue string) {
	job.Queue = queue
}
