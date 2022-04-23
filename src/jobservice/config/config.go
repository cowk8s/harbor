package config

const (
	jobServiceProtocol = "JOB_SERVICE_PROTOCOL"
	jobServicePort     = "JOB_SERVICE_PORT"
	jobServiceHTTPCert = "JOB_SERVICE_CERT"
)

var DefaultConfig = &Con

type Configuration struct {
	// Protocol server listening on: https/http
	Protocol string `yaml:"protocol"`

	Port uint `yaml:"port"`
}

// HTTPSConfig keeps additional configurations when using https protocol
type HTTPSConfig struct {
	Cert string `yaml:"cert"`
	Key  string `yaml:"key"`
}
