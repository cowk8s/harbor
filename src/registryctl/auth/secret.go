package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/cowk8s/harbor/src/common/secret"
)

// HarborSecret is the prefix of the value of Authorization header.
const HarborSecret = secret.HeaderPrefix

var (
	// ErrNoSecret ...
	ErrNoSecret = errors.New("no secret auth credentials")
)

type secretHandler struct {
	secrets map[string]string
}

// NewSecretHandler creates a new authentication handler which adds
// basic authentication credentials to a request.
func NewSecretHandler(secrets map[string]string) AuthenticationHandler {
	return &secretHandler{
		secrets: secrets,
	}
}

func (s *secretHandler) AuthorizeRequest(req *http.Request) error {
	if len(s.secrets) == 0 || req == nil {
		return ErrNoSecret
	}

	auth := req.Header.Get("Authorization")
	if !strings.HasPrefix(auth, HarborSecret) {
		return ErrInvalidCredential
	}
	secInReq := strings.TrimPrefix(auth, HarborSecret)

	for _, v := range s.secrets {
		if secInReq == v {
			return nil
		}
	}

	return ErrInvalidCredential
}
