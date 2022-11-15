package Queue

import "github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"

type WorkerConfig struct {
	Connection string
	Queue      []string
	Processes  int
	Tries      int
	Timeout    int
}

type Workers map[string]WorkerConfig

type FailedJobs struct {
	Database string
	Table    string
}

type Defaults struct {
	Connection string
	Queue      string
}

type Config struct {
	Defaults    Defaults
	Connections map[string]Support.Fields
	Failed      FailedJobs
	Workers     map[string]Workers
}
