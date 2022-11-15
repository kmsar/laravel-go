package Queue

import (
	"errors"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IQueue"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/ISerialize"
)

var (
	JobUnserializeError = errors.New("unserialize job failed")
)

type Serializer struct {
	serializer ISerialize.ClassSerializer
}

func NewJobSerializer(serializer ISerialize.ClassSerializer) IQueue.JobSerializer {
	return &Serializer{serializer: serializer}
}

func (this *Serializer) Serializer(job IQueue.Job) string {
	return this.serializer.Serialize(job)
}

func (this *Serializer) Unserialize(serialized string) (IQueue.Job, error) {
	var result, err = this.serializer.Parse(serialized)
	if err != nil {
		return nil, err
	}

	if job, isJob := result.(IQueue.Job); isJob {
		return job, nil
	}

	return nil, JobUnserializeError
}
