package models

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cowk8s/harbor/src/lib/orm"
	"github.com/cowk8s/harbor/src/pkg/allowlist"
)

const (
	// ProjectTable is the table name for project
	ProjectTable = "project"
	// ProjectPublic means project is public
	ProjectPublic = "public"
	// ProjectPrivate means project is private
	ProjectPrivate = "private"
)

func init() {
	orm.RegisterModel(&Project{})
}

// Project holds the details of a project.
type Project struct {
	ProjectID    int64                  `orm:"pk;auto;column(project_id)" json:"project_id"`
	OwnerID      int                    `orm:"column(owner_id)" json:"owner_id"`
	Name         string                 `orm:"column(name)" json:"name" sort:"default"`
	CreationTime time.Time              `orm:"column(creation_time);auto_now_add" json:"creation_time"`
	UpdateTime   time.Time              `orm:"column(update_time);auto_now" json:"update_time"`
	Deleted      bool                   `orm:"column(deleted)" json:"deleted"`
	OwnerName    string                 `orm:"-" json:"owner_name"`
	Role         int                    `orm:"-" json:"current_user_role_id"`
	RoleList     []int                  `orm:"-" json:"current_user_role_ids"`
	RepoCount    int64                  `orm:"-" json:"repo_count"`
	ChartCount   uint64                 `orm:"-" json:"chart_count"`
	Metadata     map[string]string      `orm:"-" json:"metadata"`
	CVEAllowlist allowlist.CVEAllowlist `orm:"-" json:"cve_allowlist"`
	RegistryID   int64                  `orm:"column(registry_id)" json:"registry_id"`
}

// NamesQuery ...
type NamesQuery struct {
	Names      []string // the names of project
	WithPublic bool     // include the public projects
}

// GetMetadata ...
func (p *Project) GetMetadata(key string) (string, bool) {
	if len(p.Metadata) == 0 {
		return "", false
	}
	value, exist := p.Metadata[key]
	return value, exist
}

func (p *Project) SetMetadata(key, value string) {
	if p.Metadata == nil {
		p.Metadata = map[string]string{}
	}
	p.Metadata[key] = value
}

func (p *Project) IsPublic() bool {
	public, exist := p.GetMetadata(ProMetaPublic)
	if !exist {
		return false
	}

	return isTrue(public)
}

func (p *Project) IsProxy() bool {
	return p.RegistryID > 0
}

func (p *Project) ContentTrustEnabled() bool {
	enabled, exist := p.GetMetadata(ProMetaEnableContentTrust)
	if !exist {
		return false
	}
	return isTrue(enabled)
}

func isTrue(i interface{}) bool {
	switch value := i.(type) {
	case bool:
		return value
	case string:
		v := strings.ToLower(value)
		return v == "true" || v == "1"
	default:
		return false
	}
}

func (p *Project) FilterByPublic(ctx context.Context, qs orm.QuerySeter, key string, value interface{}) orm.QuerySeter {
	subQuery := `SELECT project_id FROM project_metadata WHERE name = 'public' AND value = '%s'`
	if isTrue(value) {
		subQuery = fmt.Sprintf(subQuery, "true")
	} else {
		subQuery = fmt.Sprintf(subQuery, "false")
	}
	return qs.FilterRaw("project_id", fmt.Sprintf("IN (%s)", subQuery))
}

// TableName is required by beego orm to map Project to table project
func (p *Project) TableName() string {
	return ProjectTable
}

// Projects the connection for Project
type Projects []*Project

// OwnerIDs returns all the owner ids from the projects
func (projects Projects) OwnerIDs() []int {
	var ownerIDs []int
	for _, project := range projects {
		ownerIDs = append(ownerIDs, project.OwnerID)
	}
	return ownerIDs
}
