package models

import (
	"time"
)

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

// ResourceLabelQuery : query parameters for the mapping relationships of resource and label
type ResourceLabelQuery struct {
	LabelID      int64
	ResourceID   int64
	ResourceName string
	ResourceType string
}
