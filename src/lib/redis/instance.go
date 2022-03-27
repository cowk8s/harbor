package redis

import (
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	// instance is a global redis client
	_instance *redis.Client
	_once sync.Once
)

func Instance() *redis.Client {
	_once.Do(func() {
		url := os.Getenv("_REDIS_URL_REG")
		if url == "" {
			url = "redis://localhost:6379/1"
		}

		c, err := 
	})
}
