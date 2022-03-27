package dao

import "context"

type DAO interface {
	Set(ctx context.Context, l model)
}
