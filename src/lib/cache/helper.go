package cache

import (
	"context"
	"sync"
	"time"
)

var (
	fetchOrSaveMu = keyMutex{m: &sync.Map{}}
)

func FetchOrSave(ctx context.Context, c Cache, key string, value interface{}, builder func() (interface{}, error), expiration ...time.Duration) error {
	err :=
}