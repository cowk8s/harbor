package allowlist

import "context"

type Manager interface {
	CreateEmpty(ctx context.Context, projectID int64) error
}