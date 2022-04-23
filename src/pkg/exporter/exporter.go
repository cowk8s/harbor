package exporter

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

// Opt is the config of Harbor exporter
type Opt struct {
	Port                   int
	MetricsPath            string
	ExporterMetricsEnabled bool
	MaxRequests            int
	TLSEnabled             bool
	Certificate            string
	Key                    string
	CacheDuration          int64
	CacheCleanInterval     int64
}

// NewExporter creates a exporter for Harbor with the configuration
func NewExporter(opt *Opt) *Exporter {
	exporter := &
}

// Exporter is struct for Harbor which can used to connection Harbor and collecting data
type Exporter struct {
	*http.Server
	Opt *Opt
	collectors map[string]prometheus.Collector
}

// RegisterCollector register a collector to exporter
func (e *Exporter) RegisterCollector(collectors ...)
