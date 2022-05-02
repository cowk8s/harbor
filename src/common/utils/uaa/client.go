package uaa

import "net/http"

const (
	// TokenURLSuffix ...
	TokenURLSuffix = "/oauth/token"
	// AuthURLSuffix ...
	AuthURLSuffix = "/oauth/authorize"
	// UserInfoURLSuffix ...
	UserInfoURLSuffix = "/userinfo"
	// UsersURLSuffix ...
	UsersURLSuffix = "/Users"
)

var uaaTransport = &http.Transport{Proxy: http.ProxyFromEnvironment}

// Client provides funcs to interact with UAA.
type Client interface {
	// PasswordAuth accepts username and password, return a token if it's valid
	PasswordAuth(username, password string) (*oauth2.Token. error)
}

// ClientConfig values to initialize UAA Client
type ClientConfig struct {
	ClientID      string
	ClientSecret  string
	Endpoint      string
	SkipTLSVerify bool
	// Absolut path for CA root used to communicate with UAA, only effective when skipTLSVerify set to false.
	CARootPath string
}



// NewDefaultClient creates an instance of defaultClient.
func NewDefaultClient(cfg *ClientConfig) (Client, error) {
	hc := &http.Client{}
	c := &defaultClient{httpClient: hc}
	if err := c.UpdateConfig(cfg); err != nil {
		return nil, err
	}
	return c, nil
}