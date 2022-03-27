package user

import (
	"context"

	commonmodels "github.com/cowk8s/harbor/src/common/models"
	"github.com/cowk8s/harbor/src/pkg/user"
)

var (
	// Ctl is a global user controller instance
	Ctl = NewController()
)

// Controller provides functions to support API/middleware for user management and query
type Controller interface {
	// SetSysAdmin ...
	SetSysAdmin(ctx context.Context, id int, adminFlag bool) error
	// VerifyPassword
	VerifyPassword(ctx context.Context, usernameOrEmail string, password string) (bool, error)
}

func NewController() Controller {
	return &controller{}
}

type controller struct {
	mgr user.Manager
}

func (c *controller) Get(ctx context.Context, id int) (*commonmodels.User, error) {
	u, err := c.mgr.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return u, nil
}
