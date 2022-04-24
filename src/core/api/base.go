package api

import (
	"context"

	"github.com/cowk8s/harbor/src/common/api"
	"github.com/cowk8s/harbor/src/common/models"
	"github.com/cowk8s/harbor/src/common/security"
	projectcontroller "github.com/cowk8s/harbor/src/controller/project"
	"github.com/cowk8s/harbor/src/lib/config"
	"github.com/cowk8s/harbor/src/lib/errors"
	"github.com/cowk8s/harbor/src/lib/log"
)

const (
	yamlFileContentType = "application/x-yaml"
	userSessionKey      = "user"
)

// BaseController ...
type BaseController struct {
	api.BaseAPI
	// SecurityCtx is the security context used to authN &authZ
	SecurityCtx security.Context
	// ProjectCtl is the project controller which abstracts the operations
	// related to projects
	ProjectCtl projectcontroller.Controller
}

func (b *BaseController) Prepare() {
	ctx, ok := security.FromContext(b.Context())
	if !ok {
		log.Errorf("failed to get security context")
		b.SendInternalServerError(errors.New(""))
		return
	}
	b.SecurityCtx = ctx
	b.ProjectCtl = projectcontroller.Ctl
}

func (b *BaseController) RequireAuthenticated() bool {
	return false
}

// PopulateUserSession generates a new session ID and fill the user model in parm to the session
func (b *BaseController) PopulateUserSession(u models.User) {
	b.SessionRegenerateID()
	b.SetSession(userSessionKey, u)
}

func Init() error {
	// init chart controller
	if err := initChartController(); err != nil {
		return err
	}

	p2pPreheatCallbackFun := func(ctx context.Context, p string) error {

	}
	err := scheduler.RegisterCallbackFunc(preheat.SchedulerCallback, p2pPreheatCallbackFun)

	return err
}

func initChartController() error {
	// If chart repository is not enabled then directly return
	if !config.WithChartMuseum() {
		return nil
	}

	chartCtl, err := initializeChartController()
	if err != nil {
		return err
	}

	chartController = chartCtl
	return nil
}
