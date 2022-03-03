package http

import (
	"crypto/tls"
	"fmt"
	"os"
	"strings"
)

const (
	// Internal TLS ENV
	internalTLSEnable        = "INTERNAL_TLS_ENABLED"
	internalVerifyClientCert = "INTERNAL_VERIFY_CLIENT_CERT"
	internalTLSKeyPath       = "INTERNAL_TLS_KEY_PATH"
	internalTLSCertPath      = "INTERNAL_TLS_CERT_PATH"
	internalTrustCAPath      = "INTERNAL_TLS_TRUST_CA_PATH"
)

func InternalTLSEnabled() bool {
	return strings.ToLower(os.Getenv(internalTLSEnable)) == "true"
}

// GetInternalCertPair used to get internal cert and key pair from environment
func GetInternalCertPair() (tls.Certificate, error) {
	crtPath := os.Getenv(internalTLSCertPath)
	keyPath := os.Getenv(internalTLSKeyPath)
	cert, err := tls.LoadX509KeyPair(crtPath, keyPath)
	return cert, err
}

func GetInternalTLSConfig() (*tls.Config, error) {
	// genrate key pair
	cert, err := GetInternalCertPair()
	if err != nil {
		return nil, fmt.Errorf("internal TLS enabled but can't get cert file %w", err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
	}, nil
}
