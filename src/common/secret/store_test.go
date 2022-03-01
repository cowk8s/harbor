package secret

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	store := NewStore(map[string]string{
		"secret1": "username1",
	})

	assert.False(t, store.IsValid("invalid_secret"))
	assert.True(t, store.IsValid("secret1"))
}

func TestGetUsername(t *testing.T) {
	store := NewStore(map[string]string{
		"secret1": "username1",
	})

	assert.Equal(t, "", store.GetUsername("invalid_secret"))
	assert.Equal(t, "username1", store.GetUsername("secret1"))
}
