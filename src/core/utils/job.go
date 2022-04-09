package utils

import "sync"

var (
	cl               sync.Mutex
	jobServiceClient job.Client
)

// GetJobServiceClient returns the job service client instance.
func GetJobServiceClient() job.Client {
	cl.Lock()
	defer cl.Unlock()
	if jobServiceClient == nil {
		jobServiceClient = job.NewDefaultClient(config.InternalJobServiceURL(), config.CoreSecret())
	}
	return jobServiceClient
}
