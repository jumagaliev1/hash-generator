package service

import (
	"context"
	b64 "encoding/base64"
	"github.com/jumagaliev1/hash-generator/internal/model"
	"github.com/jumagaliev1/hash-generator/internal/storage"
	"github.com/jumagaliev1/hash-generator/pkg/redis"
	"strings"
	"time"
)

type HashService struct {
	repo  *storage.Repository
	redis *redis.Client
}

var (
	letters = []byte("QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm1234567890/+")
)

func NewHashService(repo *storage.Repository, redis *redis.Client) *HashService {
	return &HashService{
		repo:  repo,
		redis: redis,
	}
}

func (s *HashService) Create(ctx context.Context) (*model.Hash, error) {
	enc := b64.StdEncoding.EncodeToString([]byte(string(time.Now().Unix())))

	hash := model.Hash{
		Value: enc,
	}

	return s.repo.Hash.Create(ctx, hash)
}

func (s *HashService) GetByID(ctx context.Context, ID uint) (*model.Hash, error) {
	panic("")
}

func (s *HashService) GetByValue(ctx context.Context, value string) (*model.Hash, error) {
	//TODO implement me
	panic("implement me")
}

func (s *HashService) Get(ctx context.Context) (string, error) {
	hash, err := s.redis.SRandMember(ctx, "myset")
	if err != nil {
		return "", err
	}
	err = s.redis.SRemove(ctx, "myset", hash)
	if err != nil {
		return "", err
	}

	return hash, nil
}

func (s *HashService) GenerateAll(ctx context.Context) error {
	for i1 := 0; i1 < 64; i1++ {
		for i2 := 0; i2 < 64; i2++ {
			for i3 := 0; i3 < 64; i3++ {
				for i4 := 0; i4 < 64; i4++ {
					for i5 := 0; i5 < 64; i5++ {
						s.generateString(context.Background(), i1, i2, i3, i4, i5)
					}
				}
			}
		}
	}

	return nil
}

func (s *HashService) generateString(ctx context.Context, i1, i2, i3, i4, i5 int) {
	var hash strings.Builder
	hash.WriteByte(letters[i1])
	hash.WriteByte(letters[i2])
	hash.WriteByte(letters[i3])
	hash.WriteByte(letters[i4])
	hash.WriteByte(letters[i5])
	//log.Info(hash.String())
	if err := s.redis.SAdd(ctx, "myset", hash.String()); err != nil {
		//log.Error(err.Error())
	}
}
