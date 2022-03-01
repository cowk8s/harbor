package main

import (
	"os"
	"strconv"

	"github.com/cowk8s/harbor/src/common/models"
)

// key: env var, value: default value
var defaultAttrs = map[string]string{
	"POSTGRESQL_HOST":     "localhost",
	"POSTGRESQL_PORT":     "5432",
	"POSTGRESQL_USERNAME": "postgres",
	"POSTGRESQL_PASSWORD": "password",
	"POSTGRESQL_DATABASE": "registry",
	"POSTGRESQL_SSLMODE":  "disable",
}

func main() {
	p, _ := strconv.Atoi(getAttr("POSTGRESQL_PORT"))
	db := &models.Database{
		Type: "postgresql",
		PostGreSQL: &models.PostGreSQL{
			Host:         getAttr("POSTGRESQL_HOST"),
			Port:         p,
			Username:     getAttr("POSTGRESQL_USERNAME"),
			Password:     getAttr("POSTGRESQL_PASSWORD"),
			Database:     getAttr("POSTGRESQL_DATABASE"),
			SSLMode:      getAttr("POSTGRESQL_SSLMODE"),
			MaxIdleConns: 5,
			MaxOpenConns: 5,
		},
	}

}

func getAttr(k string) string {
	v := os.Getenv(k)
	if len(v) > 0 {
		return v
	}
	return defaultAttrs[k]
}
