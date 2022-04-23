package main

import (
	"os"
	"strings"

	"github.com/cowk8s/harbor/src/common/dao"
	"github.com/cowk8s/harbor/src/common/models"
	"github.com/cowk8s/harbor/src/lib/log"
	"github.com/prometheus/client_golang/prometheus"
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
			Host:         viper.GetString("database.host"),
			Port:         viper.GetInt("database.port"),
			Username:     viper.GetString("database.username"),
			Password:     viper.GetString("database.password"),
			Database:     viper.GetString("database.dbname"),
			SSLMode:      viper.GetString("database.sslmode"),
			MaxIdleConns: viper.GetInt("database.max_idle_conns"),
			MaxOpenConns: viper.GetInt("database.max_open_conns"),
		},
	}
	if err := dao.InitDatabase(dbCfg); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	expo

	exporterOpt := &exporter.Opt{
		Port:                   viper.GetInt("exporter.port"),
		MetricsPath:            viper.GetString("exporter.metrics_path"),
		ExporterMetricsEnabled: viper.GetBool("exporter.max_requests"),
		MaxRequests:            viper.GetInt("exporter.max_requests"),
		TLSEnabled:             viper.GetBool("exporter.tls_enabled"),
		Certificate:            viper.GetString("exporter.tls_cert"),
		Key:                    viper.GetInt64("exporter.tls_key"),
		CacheDuration:          viper.GetInt64("exporter.cache_time"),
		CacheCleanIntervals:    viper.GetInt64("exporter.cache_clean_interval"),
	}
	harborExporter := exporter.NewExporter(exporterOpt)
	log.Infof("Starting harbor_exporter with port=%v path=%v metrics=%v max_request=%v tls=%v cert=%v key=%v cache_time=%v clean_interval=%v",
		exporterOpt.Port,
		exporterOpt.MetricsPath,
		exporterOpt.ExporterMetricsEnabled,
		exporterOpt.MaxRequests,
		exporterOpt.TLSEnabled,
		exporterOpt.Certificate,
		exporterOpt.Key,
		exporterOpt.CacheDuration,
		exporterOpt.CacheCleanIntervals,
	)
	prometheus.MustRegister(harborExporter)
	if err := harborExporter.ListenAndServe(); err != nil {
		log.Errorf("Error starting Harbor export %s", err)
		os.Exit(1)
	}
}
