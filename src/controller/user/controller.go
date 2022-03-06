package user

import (
	"context"

	commonmodels "github.com/cowk8s/harbor/src/common/models"
	"github.com/cowk8s/harbor/src/pkg/user"
)

var (
	Ctl = NewController()
)

type Controller interface {
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
