package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cowk8s/harbor/src/common/models"
	"github.com/cowk8s/harbor/src/lib/log"
	"github.com/cowk8s/harbor/src/pkg/usergroup/model"
)

// 1.5 seconds
const frozenTime time.Duration = 1500 * time.Millisecond

var lock = NewUserLock(frozenTime)

// ErrorUserNotExist ...
var ErrorUserNotExist = errors.New("user does not exist")

// ErrorGroupNotExist ...
var ErrorGroupNotExist = errors.New("group does not exist")

// ErrDuplicateLDAPGroup ...
var ErrDuplicateLDAPGroup = errors.New("a LDAP user group with same DN already exist")

// ErrInvalidLDAPGroupDN ...
var ErrInvalidLDAPGroupDN = errors.New("the LDAP group DN is invalid")

// ErrNotSupported ...
var ErrNotSupported = errors.New("not supported")

// ErrAuth is the type of error to indicate a failed authentication due to user's error.
type ErrAuth struct {
	details string
}

// Error ...
func (ea ErrAuth) Error() string {
	return fmt.Sprintf("Failed to authenticate user, due to error '%s'", ea.details)
}

// NewErrAuth ...
func NewErrAuth(msg string) ErrAuth {
	return ErrAuth{details: msg}
}

// AuthenticateHelper provides interface for user management in different auth modes.
type AuthenticateHelper interface {
	// Authenticate authenticate the user based on data in m.  Only when the error returned is an instance
	// of ErrAuth, it will be considered a bad credentials, other errors will be treated as server side error.
	Authenticate(ctx context.Context, m models.AuthModel) (*models.User, error)
	// OnBoardUser will check if a user exists in user table, if not insert the user and
	// put the id in the pointer of user model, if it does exist, fill in the user model based
	// on the data record of the user
	OnBoardUser(ctx context.Context, u *models.User) error
	// OnBoardGroup Create a group in harbor DB, if altGroupName is not empty, take the altGroupName as groupName in harbor DB.
	OnBoardGroup(ctx context.Context, g *model.UserGroup, altGroupName string) error
	// SearchUser Get user information from account repository
	SearchUser(ctx context.Context, username string) (*models.User, error)
	// SearchGroup Search a group based on specific authentication
	SearchGroup(ctx context.Context, groupDN string) (*model.UserGroup, error)
	// PostAuthenticate Update user information after authenticate, such as Onboard or sync info etc
	PostAuthenticate(ctx context.Context, u *models.User) error
}

// Login authenticates user credentials based on setting.
func Login(ctx context.Context, m models.AuthModel) (*models.User, error) {
	authMode, err := config.AuthMode(ctx)
	if err != nil {
		return nil, err
	}
	if authMode == "" || IsSuperUser(ctx, m.Principal) {
		authMode = common.DBAuth
	}
	log.Debug("Current AUTH_MODE is ", authMode)

	authenticator, ok := registry[authMode]
	if !ok {
		return nil, fmt.Errorf("unrecognized auth_mode: %s", authMode)
	}
	if lock.IsLocked(m.Principal) {
		log.Debugf("%s is locked due to login failure, login failed", m.Principal)
		return nil, nil
	}
	user, err := authenticator.Authenticate(ctx, m)
	if err != nil {
		if _, ok = err.(ErrAuth); ok {
			log.Debugf("Login failed, locking %s, and sleep for %v", m.Principal, frozenTime)
			lock.Lock(m.Principal)
			time.Sleep(frozenTime)
		}
		return nil, err
	}
	err = authenticator.PostAuthenticate(ctx, user)
	return user, err
}
