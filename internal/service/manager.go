package service

import (
	"errors"
	"github.com/jumagaliev1/hash-generator/internal/storage"
	"github.com/jumagaliev1/hash-generator/pkg/redis"
	"github.com/spf13/viper"
)

type Service struct {
	Hash IHashService
}

func New(repo *storage.Repository) (*Service, error) {
	if repo == nil {
		return nil, errors.New("no given repo")
	}

	redisClient := redis.New(viper.GetString("redis.uri"))

	hash := NewHashService(repo, redisClient)

	return &Service{
		Hash: hash,
	}, nil

}
