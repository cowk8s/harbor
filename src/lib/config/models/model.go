package models

import (
	"github.com/astaxie/beego/orm"
)

// Email ...
type Email struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	SSL      bool   `json:"ssl"`
	Identity string `json:"identity"`
	From     string `json:"from"`
	Insecure bool   `json:"insecure"`
}

// HTTPAuthProxy wraps the settings for HTTP auth proxy
type HTTPAuthProxy struct {
	Endpoint            string   `json:"endpoint"`
	TokenReviewEndpoint string   `json:"tokenreivew_endpoint"`
	AdminGroups         []string `json:"admin_groups"`
	AdminUsernames      []string `json:"admin_usernames"`
	VerifyCert          bool     `json:"verify_cert"`
	SkipSearch          bool     `json:"skip_search"`
	ServerCertificate   string   `json:"server_certificate"`
}

// OIDCSetting wraps the settings for OIDC auth endpoint
type OIDCSetting struct {
	Name               string            `json:"name"`
	Endpoint           string            `json:"endpoint"`
	VerifyCert         bool              `json:"verify_cert"`
	AutoOnboard        bool              `json:"auto_onboard"`
	ClientID           string            `json:"client_id"`
	ClientSecret       string            `json:"client_secret"`
	GroupsClaim        string            `json:"groups_claim"`
	AdminGroup         string            `json:"admin_group"`
	RedirectURL        string            `json:"redirect_url"`
	Scope              []string          `json:"scope"`
	UserClaim          string            `json:"user_claim"`
	ExtraRedirectParms map[string]string `json:"extra_redirect_parms"`
}
