package mocked

import (
	"context"

	repositories "github.com/himynamej/api-test.git/repository"
	"github.com/himynamej/api-test.git/submetering"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (repo *UserRepository) FindByEmail(ctx context.Context, emailAddress string) (*submetering.User, error) {
	args := repo.Mock.Called(ctx, emailAddress)

	if v := args.Get(0); v != nil {
		return v.(*submetering.User), args.Error(1)
	}

	return nil, args.Error(1)
}

func (repo *UserRepository) Insert(ctx context.Context, user submetering.User) error {
	args := repo.Mock.Called(ctx, user)

	return args.Error(0)
}

func (repo *UserRepository) List(ctx context.Context, filters submetering.ListUsersFilters) ([]*submetering.User, error) {
	args := repo.Mock.Called(ctx, filters)

	if v := args.Get(0); v != nil {
		return v.([]*submetering.User), args.Error(1)
	}

	return nil, args.Error(1)
}

func (repo *UserRepository) Count(ctx context.Context, filters submetering.ListUsersFilters) (uint64, error) {
	args := repo.Mock.Called(ctx, filters)

	return args.Get(0).(uint64), args.Error(1)
}

func (repo *UserRepository) Find(ctx context.Context, uuid string) (*submetering.User, error) {
	args := repo.Mock.Called(ctx, uuid)

	if v := args.Get(0); v != nil {
		return v.(*submetering.User), args.Error(1)
	}

	return nil, args.Error(1)
}

func (repo *UserRepository) Replace(ctx context.Context, uuid string, user submetering.User) error {
	args := repo.Mock.Called(ctx, uuid, user)

	return args.Error(0)
}

var _ repositories.UserRepository = &UserRepository{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
