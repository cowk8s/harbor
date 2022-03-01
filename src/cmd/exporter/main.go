package main

import (
	"strings"

	"github.com/cowk8s/harbor/src/common/models"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("harbor")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	dbCfg := &models.Database{
		Type: "postgresql",
		PostGreSQL: &models.PostGreSQL{
			Host: viper.GetString("database.host"),
		},
	}
	if err := 
}
