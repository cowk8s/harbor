package token

import "github.com/golang-jwt/jwt/v4"

type Token struct {
	jwt.Token
	Opt   *Options
	Claim jwt.Claims
}

func New(opt *Options, claims jwt.Claims) (*Token, error) {
	err := claims.Valid()
	if err != nil {
		return nil, err
	}
	return &Token{
		Token: *jwt.NewWithClaims(opt.SignMethod, claims),
		Opt:   opt,
		Claim: claims,
	}, nil
}

func (tk *Token) Raw() (string, error) {

}
