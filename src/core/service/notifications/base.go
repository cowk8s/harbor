package notifications

import "github.com/cowk8s/harbor/src/core/api"

// BaseHandler extracts the common funcs, all notification handlers should shadow this struct
type BaseHandler struct {
	api.BaseController
}
