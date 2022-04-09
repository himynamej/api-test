package repositories

import (
	"context"

	"github.com/himynamej/api-test.git/submetering"
	"github.com/pkg/errors"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	FindByEmail(ctx context.Context, emailAddress string) (res *submetering.User, err error)
	Insert(ctx context.Context, user submetering.User) (err error)
	List(ctx context.Context, filters submetering.ListUsersFilters) (res []*submetering.User, err error)
	Count(ctx context.Context, filters submetering.ListUsersFilters) (res uint64, err error)
	Find(ctx context.Context, uuid string) (res *submetering.User, err error)
	Replace(ctx context.Context, uuid string, user submetering.User) (err error)
}
