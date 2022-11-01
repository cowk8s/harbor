package model

import (
	"time"

	"github.com/beego/beego/orm"
)

func init() {
	orm.RegisterModel(&Label{})
	orm.RegisterModel(&Reference{})
}

// Label holds information used for a label
type Label struct {
	ID           int64     `orm:"pk;auto;column(id)" json:"id"`
	Name         string    `orm:"column(name)" json:"name"`
	Description  string    `orm:"column(description)" json:"description"`
	Color        string    `orm:"column(color)" json:"color"`
	Level        string    `orm:"column(level)" json:"-"`
	Scope        string    `orm:"column(scope)" json:"scope"`
	ProjectID    int64     `orm:"column(project_id)" json:"project_id"`
	CreationTime time.Time `orm:"column(creation_time);auto_now_add" json:"creation_time"`
	UpdateTime   time.Time `orm:"column(update_time);auto_now" json:"update_time"`
	Deleted      bool      `orm:"column(deleted)" json:"deleted"`
}

// TableName ...
func (l *Label) TableName() string {
	return "harbor_label"
}

// Reference is the reference of label and artifact
type Reference struct {
	ID           int64     `orm:"pk;auto;column(id)"`
	LabelID      int64     `orm:"column(label_id)"`
	ArtifactID   int64     `orm:"column(artifact_id)"`
	CreationTime time.Time `orm:"column(creation_time);auto_now_add"`
	UpdateTime   time.Time `orm:"column(update_time);auto_now"`
}

// TableName defines the database table name
func (r *Reference) TableName() string {
	return "label_reference"
}

// ResourceLabel records the relationship between resource and label
type ResourceLabel struct {
	ID           int64     `orm:"pk;auto;column(id)"`
	LabelID      int64     `orm:"column(label_id)"`
	ResourceID   int64     `orm:"column(resource_id)"`
	ResourceName string    `orm:"column(resource_name)"`
	ResourceType string    `orm:"column(resource_type)"`
	CreationTime time.Time `orm:"column(creation_time);auto_now_add"`
	UpdateTime   time.Time `orm:"column(update_time);auto_now"`
}

// TableName ...
func (r *ResourceLabel) TableName() string {
	return "harbor_resource_label"
}