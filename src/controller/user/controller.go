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
	// Get ...
	Get(ctx context.Context, id int, opt *Option) (*commonmodels.User, error)
	// GetByName gets the user model by username, it only supports getting the basic and does not support opt
	GetByName(ctx context.Context, username string) (*commonmodels.User, error)
}

func NewController() Controller {
	return &controller{}
}

// Option  option for getting User info
type Option struct {
	WithOIDCInfo bool
}

type controller struct {
	mgr user.Manager
}

func (c *controller) GetByName(ctx context.Context, username string) (*commonmodels.User, error) {
	return c.mgr.GetByName(ctx, username)
}

func (c *controller) Get(ctx context.Context, id int) (*commonmodels.User, error) {
	u, err := c.mgr.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return u, nil
}
