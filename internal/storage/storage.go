package storage

import (
	"context"
	"github.com/jumagaliev1/hash-generator/internal/model"
	"github.com/jumagaliev1/hash-generator/internal/storage/postgre"
)

type IHashRepository interface {
	Create(ctx context.Context, hash model.Hash) (*model.Hash, error)
	GetByID(ctx context.Context, ID uint) (*model.Hash, error)
	GetByValue(ctx context.Context, value string) (*model.Hash, error)
}

type Repository struct {
	Hash IHashRepository
}

func New(ctx context.Context) (*Repository, error) {
	db, err := postgre.Dial(ctx)
	if err != nil {
		return nil, err
	}

	hash := postgre.NewHashRepository(db)

	return &Repository{
		Hash: hash,
	}, nil
}
