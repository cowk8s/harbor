package models

const (
	// JobPending ...
	JobPending string = "pending"
	// JobRunning ...
	JobRunning string = "running"
	// JobError ...
	JobError string = "error"
	// JobStopped ...
	JobStopped string = "stopped"
	// JobFinished ...
	JobFinished string = "finished"
	// JobCanceled ...
	JobCanceled string = "canceled"
	// JobRetrying indicate the job needs to be retried, it will be scheduled to the end of job queue by statemachine after an interval.
	JobRetrying string = "retrying"
	// JobContinue is the status returned by statehandler to tell statemachine to move to next possible state based on trasition table.
	JobContinue string = "_continue"
	// JobScheduled ...
	JobScheduled string = "scheduled"
)
