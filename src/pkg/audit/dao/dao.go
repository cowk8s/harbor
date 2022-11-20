package dao

import (
	"context"

	"github.com/cowk8s/harbor/src/lib/orm"
	"github.com/cowk8s/harbor/src/pkg/audit/model"
)

type DAO interface {
	Create(ctx context.Context, access *model.AuditLog) (id int64, err error)
}

func New() DAO {
	return &dao{}
}

type dao struct{}

// Create ...
func (d *dao) Create(ctx context.Context, audit *model.AuditLog) (int64, error) {
	ormer, err := orm.FromContext(ctx)
	if err != nil {
		return 0, err
	}
	if len(audit.Username) > 255 {
		audit.Username = audit.Username[:255] + "..."
	}
	id, err := ormer.Insert(audit)
	if err != nil {
		return 0, err
	}
	return id, err
}

