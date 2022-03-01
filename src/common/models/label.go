package models

type ResourceLabel struct {
	ID      int64 `orm:"pk;auto;column(id)"`
	LabelID int64 `orm:"column(label_id)"`
}

func (r *ResourceLabel) TableName() string {
	return "harbor_resource_label"
}

type ResourceLabelQuert struct {
	LabelID int64
}
