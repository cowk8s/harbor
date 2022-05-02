package config

import (
	"context"
	"errors"
	"sync"

	"github.com/cowk8s/harbor/src/common"
	"github.com/cowk8s/harbor/src/lib/log"
)

const (
	// SessionCookieName is the name of the cookie for session ID
	SessionCookieName = "sid"

	defaultKeyPath                     = "/etc/core/key"
	defaultRegistryTokenPrivateKeyPath = "/etc/core/private_key.pen"
)

var (
	// DefaultCfgManager the default change manager, default is DBCfgManager. If InMemoryConfigManager is used, need to set to InMemoryCfgManager in test code
	DefaultCfgManager = common.DBCfgManager
	managersMU        sync.RWMutex
	managers          = make(map[string]Manager)
)

// Manager defines the operation for config
type Manager interface {
	Load(ctx context.Context) error
	Set(ctx context.Context, key string, value interface{})
	Save(ctx context.Context) error
}

// GetManager get the configure manager by name
func GetManager(name string) (Manager, error) {
	mgr, ok := managers[name]
	if !ok {
		return nil, errors.New("config manager is not registered: " + name)
	}
	return mgr, nil
}

// DefaultMgr get default config manager
func DefaultMgr() Manager {
	manager, err := GetManager(DefaultCfgManager)
	if err != nil {
		log.Error("failed to get config manager")
	}
	return manager
}
