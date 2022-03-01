package health

import (
	"time"
	"github.com/docker/distribution/health"
)

var (
	timeout = 60 * time.Second
	registry = map[string]health.Checker{}
	Ctl = 
)