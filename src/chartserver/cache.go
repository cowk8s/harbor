package chartserver

import (
	"math"
	"time"

	beego_cache "github.com/astaxie/beego/cache"

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

// Initialize the cache driver based on the config
func initCacheDriver(cacheConfig *ChartCacheConfig) beego_cache.Cache {
	switch cacheConfig.DriverType {
	case cacheDriverMem:

	}
}

// backoff: fast->slow->fast
func backoff(count int) int {
	f := 5 - math.Abs((float64)(count)-5)
	return (int)(math.Pow(2, f))
}
