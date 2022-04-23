package exporter

import "github.com/prometheus/client_golang/prometheus"

const (
	healthURL           string = "/api/v2.0/health"
	healthCollectorName string = "HealthCollector"
)

var (
	harborHealth = typedDesc{
		desc:      newDesc("", "health", "Running status of Harbor"),
		valueType: prometheus.GaugeValue,
	}
	harborComponentsHealth = typedDesc{
		desc:      newDescWithLables("", "up", "Running status of Harbor component", "component"),
		valueType: prometheus.GaugeValue,
	}
)

// NewHealthCollect ...
func NewHealthCollect(cli *HarborClient) *H

// HealthCollector is the Heartbeat
type HealthCollector struct {
	*HarborClient
}

func (hc *HealthCollector) Describe(c chan<- *prometheus.Desc) {
	c <- harborHealth.Desc()
	c <- harborComponentsHealth.Desc()
}

func (hc *HealthCollector) Collect(c chan<- prometheus.Metric) {
	for _, m := range hc.getHealthStatus() {
		c <- m
	}
}

type responseHealth struct {
	Status     string              `json:"status"`
	Components []responseComponent `json:"components"`
}

type responseComponent struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func healthy(h string) float64 {
	if h == "healthy" {
		return 1
	}
	return 0
}
