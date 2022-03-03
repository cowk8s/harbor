package server

import (
	v2 "github.com/cowk8s/harbor/src/server/v2.0/route"
)

func RegisterRoutes() {
	v2.RegisterRoutes()
}
