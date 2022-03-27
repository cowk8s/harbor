package dao

import (
	"context"

	beego_orm "github.com/astaxie/beego/orm"
	"github.com/cowk8s/harbor/src/lib/orm"
	"github.com/cowk8s/harbor/src/lib/q"
	"github.com/cowk8s/harbor/src/pkg/tag/model/tag"
)

func init() {
	beego_orm.RegisterModel(&tag.Tag{})
}

// DAO is the data access object for tag
type DAO interface {
	// Count returns the total count of tags according to the query
	Count(ctx context.Context, query *q.Query) (total int64, err error)
	// List tags according to the query
	List(ctx context.Context, query *q.Query) (tags []*tag.Tag, err error)
	// Get the tag specified by ID
	Get(ctx context.Context, id int64) (tag *tag.Tag, err error)
	// Create the tag
	Create(ctx context.Context, tag *tag.Tag) (id int64, err error)
	// Update the tag. Only the properties specified by "props" will be updated if it is set
	Update(ctx context.Context, tag *tag.Tag, props ...string) (err error)
	// Delete the tag specified by ID
	Delete(ctx context.Context, id int64) (err error)
	// DeleteOfArtifact deletes all tags attached to the artifact
	DeleteOfArtifact(ctx context.Context, artifactID int64) (err error)
}

// New returns an instance of the default DAO
func New() DAO {
	return &dao{}
}

type dao struct{}

func (d *dao) Count(ctx context.Context, query *q.Query) (int64, error) {
	qs, err := orm.QuerySetterForCount(ctx, &tag.Tag{}, query)
	if err != nil {
		return 0, err
	}
	return qs.Count()
}
func (d *dao) List(ctx context.Context, query *q.Query) ([]*tag.Tag, error) {
	tags := []*tag.Tag{}
	return tags, nil
}
func (d *dao) Get(ctx context.Context, id int64) (*tag.Tag, error) {

}
