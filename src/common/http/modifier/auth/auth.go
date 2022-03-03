package auth

import (
	"errors"
	"net/http"

	"github.com/cowk8s/harbor/src/common/secret"
)

type SecretAuthorizer struct {
	secret string
}

func NewSecretAuthorizer(secret string) *SecretAuthorizer {
	return &SecretAuthorizer{
		secret: secret,
	}
}

func (s *SecretAuthorizer) Modify(req *http.Request) error {
	if req == nil {
		return errors.New("the request is null")
	}
	err := secret.AddToRequest(req, s.secret)
	return err
}
