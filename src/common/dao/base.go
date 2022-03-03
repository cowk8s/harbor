package dao

import (
	"errors"
	"fmt"
	"sync"

	"github.com/astaxie/beego/orm"
	"github.com/cowk8s/harbor/src/common/models"
)

const (
	// NonExistUserID : if a user does not exist, the ID of the user will be 0.
	NonExistUserID = 0
)

// ErrDupRows is returned
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

// UpgradeSchema will call the internal migrator to upgrade schema based on the setting of database.
func UpgradeSchema(database *models.Database) error {
	db, err := getDatabase(database)
	if err != nil {
		return err
	}
	return db.UpgradeSchema()
}

func InitDatabase(database *models.Database) error {
	db, err := getDatabase(database)
	if err != nil {
		return err
	}

	if err := db.Register(); err != nil {
		return err
	}

	return nil
}

func getDatabase(database *models.Database) (db Database, err error) {
	switch database.Type {
	case "", "postgresql":
		db = NewPGSQL(

		)
	default:
		err = fmt.Errorf("invalid database: %s", database.Type)
	}
	return
}

var globalOrm orm.Ormer
var once sync.Once

// GetOrmer :set ormer singleton
func GetOrmer() orm.Ormer {
	once.Do(func() {
		// override the default value(1000) to return all records when setting no limit
		orm.DefaultRowsLimit = -1
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}

func ClearTable(table string) error {
	o := GetOrmer()
	sql := fmt.Sprintf("delete from %s where 1=1", table)
	if table == 
}
