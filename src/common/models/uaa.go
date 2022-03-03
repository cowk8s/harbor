package models

// UAASettings wraps the configuraations to access UAA service
type UAASettings struct {
	Endpoint     string
	ClientID     string
	ClientSecret string
	VerifyCert   bool
}
