package user

import (
	"context"

	commonmodels "github.com/cowk8s/harbor/src/common/models"
	"github.com/cowk8s/harbor/src/lib/errors"
	"github.com/cowk8s/harbor/src/lib/q"
	"github.com/cowk8s/harbor/src/pkg/user/dao"
)

var (
	// Mgr is the global project manager
	Mgr = New()
)

// Manager is used for user management
type Manager interface {
	// Get get user by user id
	Get(ctx context.Context, id int) (*commonmodels.User, error)
}

// New returns a default implementation of Manager
func New() Manager {
	return &manager{dao: dao.New()}
}

type manager struct {
	dao dao.DAO
}

// Get get user by user id
func (m *manager) Get(ctx context.Context, id int) (*commonmodels.User, error) {
	users, err := m.dao.List(ctx, q.New(q.KeyWords{"user_id": id}))
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.NotFoundError(nil).WithMessage("user %d not found", id)
	}

	return users[0], nil
}
