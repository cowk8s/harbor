package registryctl

import (
	"os"

	"github.com/cowk8s/harbor/src/common"
	"github.com/cowk8s/harbor/src/lib/log"
)

var (
	// RegistryCtlClient is a client for registry controller
	RegistryCtlClient client.Client
)

// Init ...
func Init() {
	initRegistryCtlClient()
}

func initRegistryCtlClient() {
	registryCtlURL := os.Getenv("REGISTRY_CONTROLLER_URL")
	if len(registryCtlURL) == 0 {
		registryCtlURL = common.DefaultRegistryControllerEndpoint
	}

	log.Infof("initializing client for registry %s ...", registryCtlURL)
	cfg := &client.Config{
		Secret: os.Getenv("JOBSERVICE_SECRET"),
	}
	RegistryCtlClient = client.NewClient(registryCtlURL, cfg)
}
