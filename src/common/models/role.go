package models

// Role holds the details of a role.
type Role struct {
	RoleID   int    `orm:"pk;auto;column(role_id)" json:"role_id"`
	RoleCode string `orm:"column(role_code)" json:"role_code"`
	Name     string `orm:"column(name)" json:"role_name"`
	RoleMask int    `orm:"column(role_mask)" json:"role_mask"`
}
