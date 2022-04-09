package mocked

import (
	"context"
	"time"

	repositories "github.com/himynamej/api-test.git/repository"
	"github.com/stretchr/testify/mock"
)

type TokenRepository struct {
	mock.Mock
}

func (repo *TokenRepository) Set(ctx context.Context, key, value string, exp time.Duration) error {
	args := repo.Mock.Called(ctx, key, value, exp)

	return args.Error(0)
}

func (repo *TokenRepository) Get(ctx context.Context, key string) (*string, error) {
	args := repo.Mock.Called(ctx, key)

	if v := args.Get(0); v != nil {
		return v.(*string), args.Error(1)
	}

	return nil, args.Error(1)
}

func (repo *TokenRepository) Delete(ctx context.Context, key string) error {
	args := repo.Mock.Called(ctx, key)

	return args.Error(0)
}

var _ repositories.TokenRepository = new(TokenRepository)

func NewTokenRepository() *TokenRepository {
	return new(TokenRepository)
}
