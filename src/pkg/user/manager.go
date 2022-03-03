package user

import (
	"context"

	commonmodels "github.com/cowk8s/harbor/src/common/models"
)

type Manager interface {
	Get(ctx context.Context, id int) (*commonmodels.User, error)
}

type manager struct {
}

func (m *manager) GetByName(ctx context.Context, username string) (*commonmodels.User, error) {
	users, err := m.dap
}
