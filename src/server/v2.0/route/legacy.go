package route

import (
	"github.com/astaxie/beego"
	"github.com/cowk8s/harbor/src/core/api"
	"github.com/cowk8s/harbor/src/lib/config"
)

// RegisterRoutes for Harbor legacy APIs
func registerLegacyRoutes() {
	version := APIVersion
	beego.Router("/api/"+version+"/email/ping", &api.EmailAPI{}, "post:Ping")

	// APIs for chart repository
	if config.WithChartMuseum() {
		// Labels for chart
		chartLabelAPIType := &api.ChartLabelAPI{}
		beego.Router("/api/"+version+"/chartrepo/:repo/charts/:name/:version/labels", chartLabelAPIType, "get:GetLabels;post:MarkLabel")
		beego.Router("/api/"+version+"/chartrepo/:repo/charts/:name/:version/labels/:id([0-9]+)", chartLabelAPIType, "delete:RemoveLabel")
	}
}
