package api

import (
	"errors"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/cowk8s/harbor/src/chartserver"
	"github.com/cowk8s/harbor/src/lib/config"
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

func initializeChartController() (*chartserver.Controller, error) {
	addr, err := config.GetChartMuseumEndpoint()
	if err != nil {
		return nil, fmt.Errorf("failed to get the endpoint URL of chart storage server: %s", err.Error())
	}

	addr = strings.TrimSuffix(addr, "/")
	url, err := url.Parse(addr)
	if err != nil {
		return nil, errors.New("endpoint URL of chart storage server is malformed")
	}

}

// parseChartVersionFromFilename parse chart and version from file name
func parseChartVersionFromFilename(filename string) (string, string) {
	noExt := strings.TrimSuffix(path.Base(filename), fmt.Sprintf(".%s"))
}
