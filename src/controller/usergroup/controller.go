package usergroup

import (
	"context"

	"github.com/cowk8s/harbor/src/pkg/usergroup/model"
)

var (
	Ctl 
)

// Controller manages the user group
type Controller interface {
	// Delete delete user group
	Delete(ctx context.Context, id int) error
	// Update update the user group name
	Update(ctx context.Context, id int, groupName string) error
	// Create create user group
	Create(ctx context.Context, group model.UserGroup) (int, error)

	Get(ctx context.Context, id int) (*model.UserGroup, error)

	Ensure(ctx context.Context, group *model.UserGroup) error

	Populate(ctx context.Context, userGroups []model.UserGroup) ([]int, error)

	List(ctx context.Context) ([]*model.UserGroup, error)

	Count(ctx context.Context) (int64, error)
}

type controller struct {
	mgr usergroup.Manager
}

func newController() Controller {
	return &controller{mgr: usergroup.Mgr}
}

