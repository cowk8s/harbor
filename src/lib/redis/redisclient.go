package redis

import (
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

var knownPool sync.Map
var m sync.Mutex

// PoolParam ...
type PoolParam struct {
	PoolMaxIdle     int
	PoolMaxActive   int
	PoolIdleTimeout time.Duration

	DialConnectionTimeout time.Duration
	DialReadTimeout       time.Duration
	DialWriteTimeout      time.Duration
}

// GetRedisPool get a named redis pool
// supported rawurl
// redis://user:pass@redis_host:port/db
// redis+sentinel://user:pass@redis_sentinel1:port1,redis_sentinel2:port2/monitor_name/db?idle_timeout_seconds=100
func GetRedisPool(name string, rawurl string, param *PoolParam) (*redis.Pool, error) {
	if p, ok := knownPool.Load(name); ok {
		return p.(*redis.Pool), nil
	}
	m.Lock()
	defer m.Unlock()
	// load again in case multi threads
	if p, ok := knownPool.Load(name); ok {
		return p.(*redis.Pool), nil
	}

	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, fmt.Errorf("bad redis url: %s, %s, %s", name, rawurl, err)
	}

	if param == nil {
		param = &PoolParam{
			PoolMaxIdle:           0,
			PoolMaxActive:         1,
			PoolIdleTimeout:       time.Minute,
			DialConnectionTimeout: time.Second,
			DialReadTimeout:       time.Second,
			DialWriteTimeout:      time.Second,
		}
	}
}
