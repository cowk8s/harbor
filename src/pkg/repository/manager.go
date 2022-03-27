package repository

import (
	"context"

	"github.com/cowk8s/harbor/src/lib/q"
)

// Mgr is the global repository manager instance
var Mgr = New()

type Manager interface {
	// Count returns the total count of repositories according to the query
	Count(ctx context.Context, query *q.Query) (total int64, err error)
	// List repositories according to the query
	List(ctx context.Context, query *q.Query) (repositories []*model.Re)
}
