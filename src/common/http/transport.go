package http

import "net/http"

const (
	// InsecureTransport used to get the insecure http Transport
	InsecureTransport = iota
	// SecureTransport used to get the external secure http Transport
	SecureTransport
)

var (
	secureHTTPTransport   http.RoundTripper
	insecureHTTPTransport http.RoundTripper
)

// TransportConfig is the configuration for http transport
type TransportConfig struct {
	Insecure bool
}

// TransportOption is the option for http transport
type TransportOption func(*TransportConfig)

func GetHttpTransport() http.RoundTripper {
	return secureHTTPTransport
}
