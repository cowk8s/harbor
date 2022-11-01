package models

import "github.com/cowk8s/harbor/src/lib/orm"

const (
	ProjectTabel = "project"
	ProjectPublic = "public"
	ProjectPrivate = "private"
)

func init() {
	orm.RegisterModel(&Project{})
}

type Project struct {
	ProjectID int64
}