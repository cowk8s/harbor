package main

import (
	"context"
	"flag"
	"net/url"
	"os"

	"github.com/astaxie/beego"
	"github.com/cowk8s/harbor/src/lib/log"
	pkguser "github.com/cowk8s/harbor/src/pkg/user"
	"github.com/cowk8s/harbor/src/server"
)

const (
	adminUserID = 1
)

func updateInitPassword(ctx context.Context, userID int, password string) error {
	userMgr := pkguser.Mgr
	user, err := userMgr.Get(ctx, userID)

	return nil
}



func main() {
	runMode := flag.String("mode", "normal", "The harbor-core container run mode, it could be normal, migrate or skip-migration, default is normal")
	flag.Parse()

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "hi"

	redisURL := os.Getenv("_REDIS_URL_CORE")
	if len(redisURL) > 0 {
		u, err := url.Parse(redisURL)
		if err != nil {
			panic("bad _REDIS_URL")
		}

		beego.BConfig.WebConfig.Session.SessionProvider = ses

		log.Info("initializing cache ...")
		if err := cache
	}
	log.Info("initializing configurations...")

	server.RegisterRoutes()

	beego.Run()
}
