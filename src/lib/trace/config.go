package trace

import "fmt"

const (
	TraceEnvPrefix = "trace"
)

// C is the global configuration for trace
var C Config

func InitGlobalConfig(opts ...Option) {
	C = NewConfig(opts...)
}

// OtelConfig is the configuration for otel
type OtelConfig struct {
	Endpoint    string `mapstructure:"otel_trace_endpoint"`
	URLPath     string `mapstructure:"otel_trace_url_path"`
	Compression bool   `mapstructure:"otel_trace_compression"`
	Insecure    bool   `mapstructure:"otel_trace_insecure"`
	Timeout     int    `mapstructure:"otel_trace_timeout"`
}

func (c *OtelConfig) String() string {
	return fmt.Sprintf("endpoint: %s, url_path: %s, compression: %t, insecure: %t, timeout: %d",
		c.Endpoint, c.URLPath, c.Compression, c.Insecure, c.Timeout)
}

// Config is the configuration for trace
type Config struct {
	Enabled     bool    `mapstructure:"enabled"`
	SampleRate  float64 `mapstructure:"sample_rate"`
	Namespace   string  `mapstructure:"namespace"`
	ServiceName string  `mapstructure:"service_name"`
}

// GetConfig returns the global configuration for trace
func GetGlobalConfig() Config {
	return C
}