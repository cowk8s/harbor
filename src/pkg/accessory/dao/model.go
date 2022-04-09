package dao

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.RegisterModel(&Accessory{})
}

// Accessory model in database
type Accessory struct {
	ID                int64     `orm:"pk;auto;column(id)" json:"id"`
	ArtifactID        int64     `orm:"column(artifact_id)" json:"artifact_id"`
	SubjectArtifactID int64     `orm:"column(subject_artifact_id)" json:"subject_artifact_id"`
	Type              string    `orm:"column(type)" json:"type"`
	Size              int64     `orm:"column(size)" json:"size"`
	Digest            string    `orm:"column(digest)" json:"digest"`
	CreationTime      time.Time `orm:"column(creation_time);auto_now_add" json:"creation_time"`
}

// TableName for artifact reference
func (a *Accessory) TableName() string {
	return "artifact_accessory"
}
