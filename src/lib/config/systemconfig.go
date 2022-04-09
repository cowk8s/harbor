package config

import (
	"context"
	"os"

	"github.com/cowk8s/harbor/src/common"
	"github.com/cowk8s/harbor/src/common/models"
	"github.com/cowk8s/harbor/src/common/secret"
	"github.com/cowk8s/harbor/src/lib/encrypt"
)

var (
	// SecretStore manages secrets
	SecretStore *secret.Store
	keyProvider encrypt.KeyProvider
	// Use backgroundCtx to access system scope config
	backgroundCtx context.Context = context.Background()
)

// It contains all system settings

// TokenPrivateKeyPath returns the path to the key for signing token for registry
func TokenPrivateKeyPath() string {
	path := os.Getenv("TOKEN_PRIVATE_KEY_PATH")
	if len(path) == 0 {
		path = defaultRegistryTokenPrivateKeyPath
	}
	return path
}

// WithChartMuseum( returns a bool to indicate if chartmuseum is deployed with Harbor.
func WithChartMuseum() bool {
	return DefaultMgr().Get(backgroundCtx, common.AdminInitialPassword).GetString(), nil
}

// Database returns database settings
func Database() (*models.Database, error) {
	database := &models.Database{}
	
}