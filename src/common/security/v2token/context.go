package v2token

import (
	"log"

	"github.com/cowk8s/harbor/src/pkg/permission/types"
)

// tokenSecutiryCtx is used for check permission of an internal signed token.
// The intention for this quy is only for support CLI push/pull. It should not be used in other scenario without careful review
// Each request should have a different instance of tokenSecurityCtx
type tokenSecurityCtx struct {
	logger *log.Logger
	name string
	accessMap map[string]map[types.Action]struct
}