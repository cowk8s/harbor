package job

import "context"

// Context is combination of BaseContext and other job specified resources.
// Context will be the real execution context for one job.
type Context interface {
	// Build the context based on the parent context
	//
	// A new job context will be generated based on the current context
	// for the provided job.
	//
	// Returns:
	// new Context based on the parent one
	// error if meet any problems
	Build(tracker Tracker) (Context, error)

	// Get property from the context
	//
	// prop string : key of the context property
	//
	// Returns:
	// 	The data of the specified context property if have
	// bool to indicate if the property existing
	Get(prop string) (interface{}, bool)


}

// ContextInitializer is a func to initialize the concrete job context
type ContextInitializer func(ctx context.Context) (Context, error)