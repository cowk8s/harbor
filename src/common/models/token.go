package models

// Token represents the json returned by registry token service
type Token struct {
	Token       string `json:"token"`
	AccessToken string `json:"access_token"` // the token returned by azure container registry is called "access_token"
	ExpiresIn   int    `json:"expires_in"`
	IssuedAt    string `json:"issued_at"`
}

// GetToken returns the content of the token
func (t *Token) GetToken() string {
	token := t.Token
	if len(token) == 0 {
		token = t.AccessToken
	}
	return token
}

// ResourceActions ...
type ResourceActions struct {
	Name    string   `json:"name"`
	Actions []string `json:"actions"`
}
