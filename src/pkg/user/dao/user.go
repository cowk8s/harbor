package dao

import (
	"github.com/cowk8s/harbor/src/pkg/user/models"
)

type User struct {
	UserID   int    `orm:"pk;auto;column(user_id)" json:"user_id"`
	UserName string `orm:"pk;column(user_name)" json:"user" sort:"default"`
}

func (u *User) TableName() string {
	return models.UserTable
}
