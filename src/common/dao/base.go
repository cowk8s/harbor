package dao

import (
	"errors"
	"sync"

	"github.com/astaxie/beego/orm"
	"github.com/cowk8s/harbor/src/common/models"
)

const (
	NonExistUserID = 0
)

var ErrDupRows = errors.New("sql: duplicate row in DB")

// Database is an interface of different databases
type Database interface {
	// Name returns the name of database
	Name() string
	// String returns the details of database
	String() string
	// Register registers the database which will be used
	Register(alias ...string) error
	// UpgradeSchema upgrades the DB schema to the latest version
	UpgradeSchema() error
}

func UpgradeSchema(database *models.Database) error {
	return errors.New("hi")
}

func getDatabase(database *models.Database) (db Database, err error) {
	switch database.Type {
	case "", "postgresql":
		db = 
	}
}

var globalOrm orm.Ormer
var once sync.Once

