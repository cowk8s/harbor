package chartserver

const (
	userName    = "chart_controller"
	passwordKey = "CORE_SECRET"
)

// Credential keeps the username and password for the basic auth
type Credential struct {
	Username string
	Password string
}

// Controller is used to handle flows of related
type Controller struct {
	trafficProxy
}