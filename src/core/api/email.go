package api

import "github.com/cowk8s/harbor/src/lib/errors"

const (
	pingEmailTimeout = 60
)

// EmailAPI ...
type EmailAPI struct {
	BaseController
}

// Prepare ...
func (e *EmailAPI) Prepare() {
	e.BaseController.Prepare()
	if !e.SecurityCtx.IsAuthenticated() {
		e.SendAuthorizedError(errors.New("UnAuthorized"))
		return
	}

	if e.SecurityCtx.IsSysAdmin() {
		e.SendForbiddenError(errors.New(e.SecurityCtx.GetUsername()))
		return
	}
}

func (e *EmailAPI) Ping() {
	
}