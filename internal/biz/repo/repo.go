package repo

//go:generate mockgen -destination=./mock/repo.go -package=mock usm/internal/biz/repo Transaction

import (
	"context"
)

type Transaction interface {
	WithTx(context.Context, func(context.Context) error) error
}
