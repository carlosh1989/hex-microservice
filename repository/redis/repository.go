package redis

import (
	"fmt"
	"strconv"

	"github.com/carlosh1989/hex-microservice/shortener"
	"github.com/go-redis/redis"
	errs "github.com/pkg/errors"
)

type redisRepository struct {
	client *redis.Client
}

func newRedisClient(redisURL string) (*redis.Client, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)
	_, err = client.Ping().Result()
	return client, err
}

func NewRedisRepository(redisURL string) (shortener.RedirectRepository, error) {
	repo := &redisRepository{}
	client, err := newRedisClient(redisURL)
	if err != nil {
		return nil, errs.Wrap(err, "repository.NewRedisRepository")
	}
	repo.client = client
	return repo, nil
}

func (r *redisRepository) generateKey(code string) string {
	return fmt.Sprintf("redirect: %s", code)
}

func (r *redisRepository) Find(code string) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	key := r.generateKey(code)
	data, err := r.client.HGetAll(key).Result()
	if err != nil {
		return nil, errs.Wrap(err, "repository.Redirect.Find")
	}
	if len(data) == 0 {
		return nil, errs.Wrap(shortener.ErrRedirectNotFound, "repository.Redirect.Find")
	}

	CreatedAt, err := strconv.ParseInt(data["created_at"], 10, 64)
	if err != nil {
		return nil, errs.Wrap(err, "repository.Redirect.Find")
	}
	redirect.Code = data["code"]
	redirect.URL = data["url"]
	redirect.CreatedAt = CreatedAt
	return redirect, nil
}

func (r *redisRepository) Store(redirect *shortener.Redirect) error {
	key := r.generateKey(redirect.Code)
	data := map[string]interface{}{
		"code":       redirect.Code,
		"url":        redirect.URL,
		"created_at": redirect.CreatedAt,
	}
	_, err := r.client.HMSet(key, data).Result()
	if err != nil {
		return errs.Wrap(err, "repository.Redirect.Store")
	}
	return nil
}
