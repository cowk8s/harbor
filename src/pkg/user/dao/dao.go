package dao

import (
	"context"

	commonmodels "github.com/cowk8s/harbor/src/common/models"
	"github.com/cowk8s/harbor/src/lib/errors"
	"github.com/cowk8s/harbor/src/lib/orm"
	"github.com/cowk8s/harbor/src/lib/q"
)

type DAO interface {
	Create(ctx context.Context, user *commonmodels.User) (int, error)
	List(ctx context.Context, query *q.Query) ([]*commonmodels.User, error)
	Count(ctx context.Context, query *q.Query) (int64, error)
	Update(ctx context.Context, user *commonmodels.User, props ...string) error
	Delete(ctx context.Context, userID int) error
}

func New() DAO {
	return &dao{}
}

type dao struct{}

func (d *dao) Delete(ctx context.Context, userID int) error {
	ormer, err := orm.FromContext(ctx)
	if err != nil {
		return err
	}
	_, err = ormer.Delete(&User{UserID: userID})
	return err
}

func (d *dao) Count(ctx context.Context, query *q.Query) (int64, error) {
	query = q.MustClone(query)
	query.Keywords["deleted"] = false
	qs, err := orm.QuerySetterForCount(ctx, &User{}, query)
	if err != nil {
		return 0, err
	}
	return qs.Count()
}

func (d *dao) Create(ctx context.Context, user *commonmodels.User) (int, error) {
	if user.UserID > 0 {
		return 0, errors.BadRequestError(nil).WithMessage("user ID is set when creating user: %d", user.UserID)
	}
	ormer, err := orm.FromContext(ctx)
	if err != nil {
		return 0, err
	}
	id, err := ormer.Insert(toDBUser(user))
	if err != nil {
		return 0, orm.WrapConflictError(err, "user %s or email %s already exists", user.Username, user.Email)
	}
	return int(id), nil
}

func (d *dao) Update(ctx context.Context, user *commonmodels.User, props ...string) error {
	ormer, err := orm.FromContext(ctx)
	if err != nil {
		return err
	}
	n, err := ormer.Update(toDBUser(user), props...)
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.NotFoundError(nil).WithMessage("user with id %d not found", user.UserID)
	}
	return nil
}

// List list users
func (d *dao) List(ctx context.Context, query *q.Query) ([]*commonmodels.User, error) {
	query = q.MustClone(query)
	query.Keywords["deleted"] = false

	qs, err := orm.QuerySetter(ctx, &User{}, query)
	if err != nil {
		return nil, err
	}

	var users []*User
	if _, err := qs.All(&users); err != nil {
		return nil, err
	}

	var retUsers []*commonmodels.User
	for _, u := range users {
		mU := toCommonUser(u)
		retUsers = append(retUsers, mU)
	}

	return retUsers, nil
}
