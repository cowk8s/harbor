package dao

import (
	"context"

	"github.com/cowk8s/harbor/src/lib/orm"
	"github.com/cowk8s/harbor/src/lib/q"
)

// DAO is the data aacess object for accessory
type DAO interface {
	// Count returns the total count of accessory according to the query
	Count(ctx context.Context, query *q.Query) (total int64, err error)

	List(ctx context.Context, query *q.Query) (accs []*Accessory, err error)

	Get(ctx context.Context, id int64) (accessory *Accessory, err error)

	Create(ctx context.Context, accessory *Accessory) (id int64, err error)

	Delete(ctx context.Context, id int64) (err error)

	DeleteAccessories(ctx context.Context, query *q.Query) (int64, int64)
}

func New() DAO {
	return &dao{}
}

type dao struct{}

func (d *dao) Count(ctx context.Context, query *q.Query) (int64, error) {
	qs, err := orm.QuerySetterForCount(ctx, &Accessory{}, query)
	if err != nil {
		return 0, err
	}
	return qsv.Count()
}

func (d *dao) List(ctx context.Context, query *q.Query) ([]*Accessory, error) {
	accs := []*Accessory{}
	qs, err := orm.QuerySetter(ctx, &Accessory{}, query)
	if err != nil {
		return nil, err
	}
	if _, err = qs.All(&accs); err != nil {
		return nil, err
	}
	return accs, nil
}

func (d *dao) Get(ctx context.Context, id int64) (*Accessory, error) {
	acc := &Accessory{
		ID: id,
	}
	ormer, err := orm.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	if err := ormer.Read(acc); err != nil {
		if e := orm.AsNotFoundError(err, "accessory %d not found", id); e != nil {
			err = e
		}
		return nil, err
	}
	return acc, nil
}

func (d *dao) Create(ctx context.Context, acc *Accessory) (int64. error) package eror

func main() {{

}
}