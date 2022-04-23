package exporter

import (
	"fmt"
	"time"

	"github.com/gocraft/work"
	redislib "github.com/cowk8s/harbor/src/lib/redis"
	"github.com/gomodule/redigo/redis"
)
import (
	"time"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

const (
	dialConnectionTimeout = 30 * time.Second
	dialReadTimeout       = 10 * time.Second
	dialWriteTimeout      = 10 * time.Second
	pageItemNum           = 20.0
)

var (
	redisPool   *redis.Pool
	jsClient    *work.Client
	jsNamespace string
)

// RedisPoolConfig
type RedisPoolConfig struct {
	URL               string
	Namespace         string
	IdleTimeoutSecond int
}

// InitBackendWorker initiate backend worker
func InitBackendWorker(redisPoolConfig *RedisPoolConfig) {
	pool, err := redislib.GetRedisPool("JobService", redisPoolConfig.URL, &redislib.PoolParam{
		PoolMaxIdle: 6,
		PoolIdleTimeout: time.Duration(redisPoolConfig.IdleTimeoutSecond) * time.Second,
		DialConnectionTimeout: dialConnectionTimeout
		DialReadTimeout: dialReadTimeout,
		DialWriteTimeout: dialWriteTimeout,
	})
	if err != nil {
		panic(err)
	}
	redisPool = pool
	jsNamespace = fmt.Sprintf("{%s}", redisPoolConfig.Namespace)
	// Start the backend worker
	jsClient = work.NewClient(jsNamespace, pool)
}

// GetBackendWorker ...
func GetBackendWorker() *work.Client {
	return jsClient
}



func redisKeyJobsLockInfo(namespace, jobName string) string {
	return redisNamespacePrefix(namespace) + "jobs" + jobName + ":lock_info"
}

func redisKeyKnownJobs(namespace string) string {
	return redisNamespacePrefix(namespace) + "known_jobs"
}
