package chartserver

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"math"
	"time"

	beego_cache "github.com/astaxie/beego/cache"
	hlog "github.com/cowk8s/harbor/src/lib/log"

	// Enable redis cache adaptor
	_ "github.com/astaxie/beego/cache/redis"
)

const (
	standardExpireTime       = 3600 * time.Second
	redisENVKey              = "_REDIS_URL_CORE"
	cacheDriverENVKey        = "CHART_CACHE_DRIVER" // "memory" or "redis"
	cacheDriverMem           = "memory"
	cacheDriverRedis         = "redis"
	cacheDriverRedisSentinel = "redis_sentinel"
	cacheCollectionName      = "helm_chart_cache"
	maxTry                   = 10
)

// ChartCache is designed to cache some processed data for repeated accessing
// to improve the performace
type ChartCache struct {
	// Cache driver
	cache beego_cache.Cache

	// keep the driver type
	driverType string

	// To indicate if the chart cache is enabled
	isEnabled bool
}


