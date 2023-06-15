package service

import (
	"context"
	"github.com/jumagaliev1/hash-generator/internal/model"
)

type IHashService interface {
	Create(ctx context.Context) (*model.Hash, error)
	GetByID(ctx context.Context, ID uint) (*model.Hash, error)
	GetByValue(ctx context.Context, value string) (*model.Hash, error)
	GenerateAll(ctx context.Context) error
	Get(ctx context.Context) (string, error)
}

type IRedisClient interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (interface{}, error)
	SAdd(ctx context.Context, set string, values ...interface{}) error
	SRandMember(ctx context.Context, set string) (string, error)
	SRemove(ctx context.Context, key string, value interface{}) error
}
