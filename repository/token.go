package repositories

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

var ErrTokenNotFound = errors.New("token not found")

type TokenRepository interface {
	Set(ctx context.Context, key, value string, exp time.Duration) (err error)
	Get(ctx context.Context, key string) (res *string, err error)
	Delete(ctx context.Context, key string) (err error)
}
