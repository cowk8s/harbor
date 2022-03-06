package health

import (
	"context"
	"sort"
	"time"

	"github.com/docker/distribution/health"
)

var (
	timeout  = 60 * time.Second
	registry = map[string]health.Checker{}
	Ctk      = NewController()
)

func NewController() Controller {
	return &controller{}
}

// Controller defines the health related operations
type Controller interface {
	GetHealth(ctx context.Context) *OverallHealthStatus
}

type controller struct{}

func (c *controller) GetHealth(ctx context.Context) *OverallHealthStatus {
	var isHealthy healthy = true
	components := []*ComponentHealthStatus{}
	ch := make(chan *ComponentHealthStatus, len(registry))
	for name, checker := range registry {
		go check(name, checker, timeout, ch)
	}
	for i := 0; i < len(registry); i++ {
		componentStatus := <-ch
		if len(componentStatus.Error) != 0 {
			isHealthy = false
		}
		components = append(components, componentStatus)
	}

	sort.Slice(components, func(i, j int) bool { return components[i].Name < components[j].Name })

	return &OverallHealthStatus{
		Status:     isHealthy.String(),
		Components: components,
	}
}

func check(name string, checker health.Checker,
	timeout time.Duration, c chan *ComponentHealthStatus) {
	statusChan := make(chan *ComponentHealthStatus)
	go func() {
		err := checker.Check()
		var healthy healthy = err == nil
		status := &ComponentHealthStatus{
			Name:   name,
			Status: healthy.String(),
		}
		if !healthy {
			status.Error = err.Error()
		}
		statusChan <- status
	}()

	select {
	case status := <-statusChan:
		c <- status
	case <-time.After(timeout):
		var healthy healthy = false
		c <- &ComponentHealthStatus{
			Name:   name,
			Status: healthy.String(),
			Error:  "failed to check the health status: timeout",
		}
	}
}
