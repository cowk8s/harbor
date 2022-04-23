package exporter

import "github.com/prometheus/client_golang/prometheus"

// JobServiceCollectorName ...
const JobServiceCollectorName = "JobServiceCollector"

var (
	jobServiceTaskQueueSize = typedDesc{
		desc:      newDescWithLables("", "task_queue_size", "Total number of tasks", "type"),
		valueType: prometheus.GaugeValue,
	}
	jobServiceTaskQueueLatency = typedDesc{
		desc:      newDescWithLables("", "task_queue_latency", "how long ago the next job to be processed was enqueued", "type"),
		valueType: prometheus.GaugeValue,
	}
	jobServiceConcurrency = typedDesc{
		desc:      newDescWithLables("", "task_concurrency", "Total number of concurrency on a pool", "type", "pool"),
		valueType: prometheus.GaugeValue,
	}
	jobServiceScheduledJobTotal = typedDesc{
		desc:      newDesc("", "task_scheduled_total", "total number of scheduled job"),
		valueType: prometheus.GaugeValue,
	}
)
