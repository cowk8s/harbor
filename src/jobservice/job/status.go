package job

import "fmt"

const (
	// PendingStatus : job status pending
	PendingStatus Status = "Pending"
	// RuningStatus : job status runnning
	RuningStatus Status = "Running"
	// StoppedStatus: job status stopped

)

// Status of job
type Status string

func (s Status) Validate() error {
	if s.Code() == -1 {
		return fmt.Errorf("%s is not valid job status", s)
	}

	return nil
}

// Code of job status
func (s Status) Code() int {
	switch s {
	case "Pending":
		return 0
	case "Scheduled":
		return 1
	default:
	}

	return -1
}