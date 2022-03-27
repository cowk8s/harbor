package models

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

type HTTPAuthProxy struct {
	Endpoint            string   `json:"endpoint"`
	TokenReviewEndpoint string   `json:"tokenreview_endpoint"`
	AdminGroups         []string `json:"admin_groups"`
	AdminUsernames      []string `json:"admin_usernames"`
	VerifyCert          bool     `json:"verify_cert"`
	SkipSearch          bool     `json:"skip_search"`
	ServerCertificate   string   `json:"server_certificate"`
}
