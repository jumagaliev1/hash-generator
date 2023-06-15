package redis

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	Client *redis.Client
}

func New(uri string) *Client {
	client := redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: "",
		DB:       0,
	})

	return &Client{
		Client: client,
	}
}

func (r *Client) Set(ctx context.Context, key string, value interface{}) error {
	err := r.Client.Set(ctx, key, value, 0).Err()

	return err
}

func (r *Client) Get(ctx context.Context, key string) (interface{}, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (r *Client) SRemove(ctx context.Context, key string, value interface{}) error {
	id, err := r.Client.SRem(ctx, key, value).Result()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	log.Info("SRem: ", id)
	return nil
}
func (r *Client) SAdd(ctx context.Context, set string, values ...interface{}) error {
	err := r.Client.SAdd(ctx, set, values).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Client) SRandMember(ctx context.Context, set string) (string, error) {
	result, err := r.Client.SRandMember(ctx, set).Result()
	if err != nil {
		return "", err
	}

	return result, nil
}
