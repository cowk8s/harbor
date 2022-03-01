package secret

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rc := m.Run()
	if rc != 0 {
		os.Exit(rc)
	}
}

func TestFromRequest(t *testing.T) {
	assert := assert.New(t)
	secret := "mysecret"
	req, _ := http.NewRequest("GET", "http://test.com", nil)
	req.Header.Add("Authorization", "Harbor-Secret "+secret)
	assert.Equal(secret, FromRequest(req))
	assert.Equal("", FromRequest(nil))
}

func TestAddToRequest(t *testing.T) {
	assert := assert.New(t)
	secret := "mysecret"
	req, _ := http.NewRequest("GET", "http://test.com", nil)
	err := AddToRequest(req, secret)
	assert.Nil(err)
	assert.Equal(secret, FromRequest(req))
	assert.NotNil(AddToRequest(nil, secret))
}
