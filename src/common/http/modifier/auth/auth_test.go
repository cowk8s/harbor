package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAuthorizeOfSecretAuthorizer(t *testing.T) {
	secret := "secret"
	authorizer := NewSecretAuthorizer(secret)

	// nil request
	require.NotNil(t, authorizer.Modify(nil))

	req, err := http.NewRequest("", "", nil)
	require.Nil(t, err)
	require.Nil(t, authorizer.Modify(req))
}
