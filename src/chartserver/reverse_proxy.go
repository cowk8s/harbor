package chartserver

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	hlog "github.com/cowk8s/harbor/src/lib/log"
)

const (
	agentHarbor         = "HARBOR"
	contentLengthHeader = "Content-Length"

	defaultRepo             = "library"
	rootUploadingEndpoint   = "/api/chartrepo/charts"
	rootIndexEndpoint       = "/chartrepo/index.html"
	chartRepoHealthEndpoint = "/api/chartrepo/health"
)

type ProxyEngine struct {
	// The backend target server the traffic will be forwarded to
	// Just in case we'sll use it
	backend *url.URL

	// Use go reverse proxy as engine
	engine http.Handler
}

func NewProxyEngine(target *url.URL, cred *Credential, middlewares ...func(http.Handler) http.Handler) *ProxyEngine {
	var engine http.Handler

	engine = &httputil.ReverseProxy{
		ErrorLog: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
		Director: func(req *http.Request) {
			directory
		},
		ModifyResponse: ,
		Transport: 
	}

	if len(middlewares) > 0 {
		hlog.Info("New chart server traffic proxy with middlewares")
		for i := len(middlewares) - 1; i >= 0; i-- {
			engine = middlewares[i](engine)
		}
	}

	return &ProxyEngine{
		backend: target,
		engine: engine,
	}
}

// ServeHTTP serves the incoming http requests
func (pe *ProxyEngine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	pe.engine.ServeHTTP(w, req)
}

func director(target *url.URL, cred *Credential, req *http.Request) {
	// Closure
	targetQuery := target.RawQuery

	// Overwite the request URL to the target path
	req.URL.Scheme = target.Scheme
	req.URL.Host = target.Host

}

// Rewrite the incoming RRL with the right backend URL pattern
// Remove 'chartrepo' fron the endpoints of the manipulation API
// Remove 'chartrepo' from the endpoints of the repository services

// Rewrite the incoming URL with the right backend URL pattern
// Remove 'chartrepo' from the endpoints of manipulation API
// Remove 'chartrepo' from the endpoints of repository services
func rewriteURLPath(req *http.Request) {
	incomingURLPath := req.URL.Path

	// Health check endpoint
	if incomingURLPath == chartRepoHealthEndpoint {
		req.URL.Path = "/health"
		return
	}

	// Root uploading endpoint
	if incomingURLPath == rootUploadingEndpoint {
		req.URL.Path = strings.Replace(incomingURLPath, "chartrepo", defaultRepo, 1)
		return
	}

	// Repository endpoints
	if strings.HasPrefix(incomingURLPath, "/chartrepo") {
		req.URL.Path = strings.TrimPrefix(incomingURLPath, "/chartrepo")
		return
	}

	// API endpoints
	if strings.HasPrefix(incomingURLPath, "/api/chartrepo") {
		req.URL.Path = strings.Replace(incomingURLPath, "/chartrepo", "", 1)
		return
	}
}

