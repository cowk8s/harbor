package token

import "github.com/golang-jwt/jwt/v4"

const (
	defaultIssuer       = "harbor-token-defaultIssuer"
	defaultSignedMethod = "RS256"
)

// Options ...
type Options struct {
	SignMethod jwt.SigningMethod
	PublicKey  []byte
	PrivateKey []byte
	Issuer     string
}

// GetKey
func (o *Options) GetKey() (interface{}, error) {

}
