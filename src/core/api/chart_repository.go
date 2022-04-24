package api

import (
	"strings"

	"github.com/cowk8s/harbor/src/chartserver"
)

const (
	namespaceParam          = ":repo"
	nameParam               = ":name"
	filenameParam           = ":filename"
	defaultRepo             = "library"
	rootUploadingEndpoint   = "/api/chartrepo/charts"
	rootIndexEndpoint       = "/chartrepo/index.yaml"
	chartRepoHealthEndpoint = "api/chartrepo/health"

	accessLevelPublic = iota
	accessLevelRead
	accessLevelWrite
	accessLevelAll
	accessLevelSystem

	formFieldNameForChart     = "chart"
	formFieldNameForProv      = "prov"
	headerContentType         = "Content-Type"
	contentTypeMultipart      = "multipart/form-data"
	chartPackageFileExtension = "tgz"
)

// chartController is a singleton instance
var chartController *chartserver.Controller

func GetChartController() *chartserver.Controller {
	return chartController
}

// ChartRepositoryAPI provides related API handlers for the chart repository APIs
type ChartRepositoryAPI struct {
	// The base controller to provide common utilities
	BaseController

	// For label management
	labelManager *label.BaseManager

	// Keep the namespace if exitsting
	namespace string
}

func (cra *ChartRepositoryAPI) Prepare() {
	// Call super prepare method
	cra.BaseController.Prepare()

	// Try to extract namespace for parameter of path
	cra.namespace = strings.TrimSpace(cra.GetStringFromPath(namespaceParam))

	// Check the existence of namespace
	// Exclude the following URI
	// -/index.yaml
	// -/api/chartserver/health
	incomingURI := cra.Ctx.Request.URL.Path
	if incomingURI == rootUploadingEndpoint {
		// Forward to the default repository
		cra.namespace = defaultRepo
	}

	if incomingURI != rootIndexEndpoint &&
		incomingURI != chartRepoHealthEndpoint {
		if !cra.requireNamespace(cra.namespace) {
			return
		}
	}

	// Init label manager
	cra.labelManager = &label.BaseManager{
		LabelMgr: pkg_label.Mgr,
	}
}

func (cra *ChartRepositoryAPI) requireAccess(action rbac.Action, subresource ...rbac.Resource) bool {
	if len(subresource) == 0 {
		subresource = append(subresource, rbac.ResourceHelmChart)
	}

	return cra.RequireProjectAccess(cra.namespace, action, subresource...)
}
