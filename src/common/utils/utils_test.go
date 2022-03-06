package utils

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTestTCPConn(t *testing.T) {
	server := httptest.NewServer(nil)
	defer server.Close()
	addr := strings.TrimPrefix(server.URL, "http://")
	if err := TestTCPConn(addr, 60, 2); err != nil {
		t.Fatalf("failed to test tcp connection of %s: %v", addr, err)
	}
}
