package config

import (
	"context"
	"sync"

	"github.com/cowk8s/harbor/src/common"
)

const (
	// SessionCookieName is the name of the cookie for session ID
	SessionCookieName = "sid"

	defaultKeyPath                     = "/etc/core/key"
	defaultRegistryTokenPrivateKeyPath = "/etc/core/private_key.pen"
)

var (
	DefaultCfgManager = common.DBCfgManager
	managersMU        sync.RWMutex
)

// Manager defines the operation for config
type Manager interface {
	Load(ctx context.Context) error
	Set(ctx context.Context, key string, value interface{})
	Save(ctx context.Context) error
	Get(ctx context.Context key string) 
}
