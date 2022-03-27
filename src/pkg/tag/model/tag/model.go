package tag

import (
	"time"

	"github.com/cowk8s/harbor/src/lib/q"
)

// Tag model in database
type Tag struct {
	ID           int64     `orm:"pk;auto;column(id)" json:"id"`
	RepositoryID int64     `orm:"column(repository_id)" json:"repository_id"` // tags are the resources of repository, one repository only contains one same name tag
	ArtifactID   int64     `orm:"column(artifact_id)" json:"artifact_id"`     // the artifact ID that the tag attaches to, it changes when pushing a same name but different digest artifact
	Name         string    `orm:"column(name)" json:"name"`
	PushTime     time.Time `orm:"column(push_time)" json:"push_time"`
	PullTime     time.Time `orm:"column(pull_time)" json:"pull_time"`
}

// GetDefaultSorts specifies the default sorts
func (t *Tag) GetDefaultSorts() []*q.Sort {
	return []*q.Sort{
		{
			Key:  "PushTime",
			DESC: true,
		},
		{
			Key:  "ID",
			DESC: true,
		},
	}
}
