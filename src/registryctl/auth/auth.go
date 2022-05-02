package auth

import (
	"errors"
	"net/http"
)

var (
	// ErrInvalidCredential is returned when the auth token does not authenticate correctly.
	ErrInvalidCredential = errors.New("invalid authorization credential")
)

// AuthenticationHandler is an interface for authorizing a request
type AuthenticationHandler interface {

	// AuthorizeRequest ...
	AuthorizeRequest(req *http.Request) error
}
