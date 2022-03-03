package health

import (
	"net/http"

	"time"

	"github.com/docker/distribution/health"
)

func HTTPStatusCodeHealthChecker(method string, url string, header http.Header,
	timeout time.Duration, statusCode int) health.Checker {
	return health.CheckFunc(func() error {
		return nil
	})
}
