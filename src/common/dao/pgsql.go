package dao

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/cowk8s/harbor/src/common/models"
	"github.com/golang-migrate/migrate/v4"
	migrate "github.com/golang-migrate/migrate/v4"
	"xorm.io/xorm/migrate"
)

const defaultMigrationPath = "migrations/postgresql/"

type pgsql struct {
	host         string
	port         string
	usr          string
	pwd          string
	database     string
	sslmode      string
	maxIdleConns int
	maxOpenConns int
}

func (p *pgsql) Name() string {
	return "PostgreSQL"
}

func (p *pgsql) String() string {
	return fmt.Sprintf("type-%s host-%s port-%s database-%s sslmode-%q",
		p.Name(), p.host, p.port, p.database, p.sslmode)
}

func (p *pgsql) Register(alias ...string) error {

	if err := orm.RegisterDriver("pgx", orm.DRPostgres); err != nil {
		return err
	}

	an := "default"
	if len(alias) != 0 {
		an = alias[0]
	}
	info := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s timezone=UTC",
		p.host, p.port, p.usr, p.pwd, p.database, p.sslmode)

	if err := orm.RegisterDataBase(an, "pgx", info, p.maxIdleConns, p.maxOpenConns); err != nil {
		return err
	}

	// Due to the issues of beego v1.12.1 and v1.12.2, we set the max open conns ourselves.
	// See https://github.com/goharbor/harbor/issues/12403
	// and https://github.com/astaxie/beego/issues/4059 for more info.
	db, _ := orm.GetDB(an)
	db.SetMaxOpenConns(p.maxOpenConns)
	db.SetConnMaxLifetime(5 * time.Minute)

	return nil
}

func (p *pgsql) UpgradeSchema() error {
	port, err := strconv.Atoi(p.port)
	if err != nil {
		return err
	}
	m, err := NewMigrator(&models.PostGreSQL{
		Host:     p.host,
		Port:     port,
		Username: p.usr,
		Password: p.pwd,
		Database: p.database,
		SSLMode:  p.sslmode,
	})
	if err != nil {
		return err
	}
	defer func() {
		srcErr, dbErr := m.Close()
		if srcErr != nil || dbErr != nil {

		}
	}()
	err = m.Up()
	if err == migrate.ErrNoChange {

	} else if err != nil {
		return err
	}
	return nil
}

// NewMigrator creates a migrator base on the information
func NewMigrator(database *models.PostGreSQL) (*migrate.Migrate, error) {
	dbURL := url.URL{
		Scheme:   "pgx",
		User:     url.UserPassword(database.Username, database.Password),
		Host:     net.JoinHostPort(database.Host, strconv.Itoa(database.Port)),
		Path:     database.Database,
		RawQuery: fmt.Sprintf("sslmode=%s", database.SSLMode),
	}

	// For UT
	path := os.Getenv("POSTGRES_MIGRATION_SCRIPTS_PATH")
	if len(path) == 0 {
		path = defaultMigrationPath
	}
	srcURL := fmt.Sprintf("file://%s", path)
	m, err := migrate.New(srcURL, dbURL.String())
	if err != nil {
		return nil, err
	}
	return m, nil
}
