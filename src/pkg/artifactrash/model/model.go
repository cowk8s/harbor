package model

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.RegisterModel(&ArtifactTrash{})
}

// ArtifactTrash records the deleted artifact
type ArtifactTrash struct {
	ID                int64     `orm:"pk;auto;column(id)"`
	MediaType         string    `orm:"column(media_type)"`
	ManifestMediaType string    `orm:"column(manifest_media_type)"`
	RepositoryName    string    `orm:"column(repository_name)"`
	Digest            string    `orm:"column(digest)"`
	CreationTime      time.Time `orm:"column(creation_time);auto_now_add" json:"creation_time"`
}

// TableName for artifact trash
func (at *ArtifactTrash) TableName() string {
	return "artifact_trash"
}

func (at *ArtifactTrash) String() string {
	return fmt.Sprintf("ID-%d MediaType-%s ManifestMediaType-%s RepositoryName-%s Digest-%s CreationTime-%s",
		at.ID, at.MediaType, at.ManifestMediaType, at.RepositoryName, at.Digest, at.CreationTime.Format("2006-01-02 15:04:05"))
}