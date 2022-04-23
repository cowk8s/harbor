package systeminfo

import (
	"context"
	"io"
	"time"

	"github.com/cowk8s/harbor/src/lib/log"
)

const defaultRootCert = "/etc/core/ca/ca.crt"

// for UT only
var testRootCertPath = ""

type Data struct {
	AuthMode string
	SelfRegistration bool
	HarborVersion string
	AuthProxySettings *models.HTTPAuthProxy
	Protected *protectedData
}

type protectedData struct {
	CurrentTime time.Time
}

type Controller interface {

	// GetInfo consolidates the info of the system by checking settings in DB and env vars
	GetInfo(ctx context.Context, opt Options) (*Data, error)

	GetCapacity(ctx context.Context) (*imagestorage.Capacity, error)

	GetCA(ctx context.Context) (io.ReadCloser, error)
}

type controller struct{}

func (c *controller) GetInfo(ctx context.Context, opt Options) (*Data, error) {
	logger := log.GetLogger(ctx)
	cfg, err := config.GetSystemCfg(ctx);
	if err != nil {
		logger.Errorf("Error occurred getting config: %v", err)
		return nil, err
	}
	res := &Data{
		AuthMode: utils.SafeCastString(cfg[common.AUTHMode]),
		
	}
	if res.AuthMode == common.HTTPAuth {
		if s, err := config.HTTPAuthProxySetting(ctx); err == nil {
			res.AuthProxySettings = s
		} else {
			logger.Warningf("Failed to get auth proxy setting, error: %v", err)
		}
	}

	if !opt.WithProtectedInfo {
		return res, nil
	}
}