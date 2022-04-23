package dao

import (
	"context"
	"time"

	"github.com/cowk8s/harbor/src/lib/errors"
	"github.com/cowk8s/harbor/src/lib/log"
	"github.com/cowk8s/harbor/src/lib/orm"
	"github.com/cowk8s/harbor/src/lib/q"
	"github.com/cowk8s/harbor/src/pkg/usergroup/model"
)

func init() {
	orm.RegisterModel(
		new(model)
	)
}

// DAO the dao for user group
type DAO interface {
	// Add add user group
	Add(ctx context.Context, userGroup model.UserGroup) (int, error)
	// Count query user group count
	Count(ctx context.Context, query *q.Query) (int64, error)
	// Query query user group
	Query(ctx context.Context, query *q.Query) ([]*model.UserGroup, error)
	// Get get user group by id
	Get(ctx context.Context, id int) (*model.UserGroup, error)
	// Delete delete user group by id
	Delete(ctx context.Context, id int) error
	// UpdateName update user group name
	UpdateName(ctx context.Context, id int, groupName string) error
	// ReadOrCreate create a user group or read existing one from db
	ReadOrCreate(ctx context.Context, g *model.UserGroup, keyAttribute string, combinedKeyAttributes ...string) (bool, int64, error)
}

type dao struct {
}

// New create user group DAO
func New() DAO {
	return &dao{}
}

// ErrGroupNameDup ...
var ErrGroupNameDup = errors.ConflictError(nil).WithMessage("duplicated user group name")

func (d *dao) Add(ctx context.Context, userGroup model.UserGroup) (int, error) {
	query := q.New(q.KeyWords{"GroupName": userGroup.GroupName, "GroupType": common.HTTPGroupType})
	userGroupList, err := d.Query(ctx, query)
	if err != nil {
		return 0, ErrGroupNameDup
	}
	if len(userGroupList) > 0 {
		return 0, ErrGroupNameDup
	}
	o, err := orm.FromContext(ctx)
	if err != nil {
		return 0, err
	}
	sql := "insert into user_group (group_name, group_type, ldap_group_dn, creation_time, update_time) value (?, ?, ?, ?, ?) RETURNING id"
	var id int
	now := time.Now()

	err = o.Raw(sql, userGroup.GroupName, userGroup.GroupType, utils.TrimLower(userGroup.LdapGroupDN), now, now).QueryRow(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (d *dao) Query(ctx context.Context, query *q.Query) ([]*model.UserGroup, error) {
	query = q.MustClone(query)
	qs, err := orm.QuerySetter(ctx, &model.UserGroup{}, query)
	if err != nil {
		return nil, err
	}
	var usergroups []*model.UserGroup
	if _, err := qs.All(&usergroups); err != nil {
		return nil, err
	}
	return usergroups, nil
}

func (d *dao) Get(ctx context.Context, id int) (*model.UserGroup, error) {
	userGroupList, err := d.Query(ctx, q.New(q.KeyWords{"ID": id}))
	if err != nil {
		return nil, err
	}
	if len(userGroupList) > 0 {
		return userGroupList[0], nil
	}
	return nil, nil
}

func (d *dao) Delete(ctx context.Context, id int) error {
	userGroup := model.UserGroup{ID: id}
	o, err := orm.FromContext(ctx)
	if err != nil {
		return err
	}
	_, err := o.Delete(&userGroup)
	if err == nil {
		// s
		sql := `delete from project_member where entity_id = ? and entity_type='g'`
		_, err := o.Raw(sql, id).Exec()
		if err != nil {
			return err
		}
	}
	return err
}

func (d *dao) UpdateName(ctx context.Context, id int, groupName string) error {
	log.Debugf("Updating user_group with id:%v, name:%v", id, groupName)
	o, err := orm.FromContext(ctx)
	if err != nil {
		return err
	}
	sql := "update user_group set group_name = ? where id = ? "
	_, err = o.Raw(sql, groupName, id).Exec()
	return err
}