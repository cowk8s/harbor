package models

import "time"

type CVEAllowlist struct {
	ID           int64  `orm:"pk;auto;column(id)" json:"id,omitempty"`
	ProjectID    int64  `orm:"column(project_id)" json:"project_id"`
	ExpiresAt    *int64 `orm:"column(expires_at)" json:"expires_at,omitempty"`
	Items        []CVEAllowlistItem
	ItemsText    string
	CreationTime time.Time
	UpdateTime   time.Time
}

type CVEAllowlistItem struct {
	CVEID string `json:"cve_id"`
}

func (c *CVEAllowlist) TableName() string {
	return "cve_allowList"
}
