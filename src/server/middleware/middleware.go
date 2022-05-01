package middleware

import "net/http"

// Middleware receives a handler and returns another handler.
// The returned handler can do some customized task according to
// the requirement
type Middleware func(http.Handler) http.Handler
