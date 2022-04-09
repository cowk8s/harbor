package lib

// NewWorkerPool creates a new worker pool with specified size
func NewWorkerPool(size int32) *WorkerPool {
	wp := &WorkerPool{}
	wp.queue = make(chan struct{}, size)
	return wp
}

// WorkerPool controls the concurrency limit of task/process
type WorkerPool struct {
	queue chan struct{}
}

// GetWorker hangs until a free worker available
func (w *WorkerPool) GetWorker() {
	w.queue <- struct{}{}
}

// ReleaseWorker hangs until the worker return back into the pool
func (w *WorkerPool) ReleaseWorker() {
	<-w.queue
}