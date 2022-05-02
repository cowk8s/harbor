package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/cowk8s/harbor/src/registryctl/config"
)

// RegistryCtl for registry controller
type RegistryCtl struct {
	ServerConf config.Configuration
	Handler    http.Handler
}

// Start the registry controller
func (s *RegistryCtl) Start() {
	regCtl := &http.Server{
		Addr:      ":" + s.ServerConf.Port,
		Handler:   s.Handler,
		TLSConfig: common_http.NewServerTLSConfig(),
	}

}

func main() {
	configPath := flag.String("c", "", "Specify registryCtl configuration file path")
	flag.Parse()

	if configPath == nil || len(*configPath) == 0 {
		flag.Usage()
		log.Fatal("Config file should be specified")
	}
	if err := config.DefaultConfig.Load(*configPath, true); err != nil {
		log.Fatalf("Failed to load configurations with errors: %s\n", err)
	}

	regCtl := &RegistryCtl{
		ServerConf: *config.DefaultConfig,
		Handler:    handlers.NewHandlerChain(*config.DefaultConfig),
	}
	regCtl.Start()
}
