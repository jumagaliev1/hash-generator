package postgre

import (
	"context"
	"github.com/jumagaliev1/hash-generator/internal/model"
	"gorm.io/gorm"
)

type HashRepository struct {
	DB *gorm.DB
}

func NewHashRepository(DB *gorm.DB) *HashRepository {
	return &HashRepository{
		DB: DB,
	}
}

func (r *HashRepository) Create(ctx context.Context, hash model.Hash) (*model.Hash, error) {
	if err := r.DB.WithContext(ctx).Create(hash).Error; err != nil {
		return nil, err
	}

	return &hash, nil
}

func (r *HashRepository) GetByID(ctx context.Context, ID uint) (*model.Hash, error) {
	var hash model.Hash

	if err := r.DB.WithContext(ctx).Where("id = ?", ID).Find(&hash).Error; err != nil {
		return nil, err
	}

	return &hash, nil
}

func (r *HashRepository) GetByValue(ctx context.Context, value string) (*model.Hash, error) {
	var hash model.Hash

	if err := r.DB.WithContext(ctx).Where("value = ?", value).Find(&hash).Error; err != nil {
		return nil, err
	}

	return &hash, nil
}
