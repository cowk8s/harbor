package config

import (
	"context"
	"errors"

	"github.com/cowk8s/harbor/src/lib/log"
)

// It contains all user related configurations, each of user related settings requires a context provided

// GetSystemCfg returns the all configurations
func GetSystemCfg(ctx context.Context) (map[string]interface{}, error) {
	sysCfg := DefaultMgr().GetAll(ctx)
	if len(sysCfg) == 0 {
		return nil, errors.New("can not load system config, the database might be down")
	}
	return sysCfg, nil
}

// AuthMode ...
func AuthMode(ctx context.Context) (string, error) {
	mgr := DefaultMgr()
	err := mgr.Load(ctx)
	if err != nil {
		log.Errorf("failed to load config, error %v", err)
		return "db_auth", err
	}
	return mgr.Get(ctx, common.AUTHMode).GetString(), nil
}
