package tag

import (
	"context"

	"github.com/cowk8s/harbor/src/lib/q"
)

var ()

type Manager interface {
	Count(ctx context.Context, query *q.Query) (total int64, err error)

	List(ctx context.Context, query *q.Query) (tags []*tag)
}
