package cache

import (
	"context"
	"encoding/json"
	"test-task/domain"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	db             *redis.Client
	expirationTime time.Duration
	ctx            context.Context
}

func NewCache(ctx context.Context, address string, expirationTime time.Duration) *Cache {
	var db = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	return &Cache{db: db, expirationTime: expirationTime, ctx: ctx}
}

func (cache *Cache) CacheUsers(users []domain.User) error {
	usersBytes, err := json.Marshal(users)
	if err != nil {
		return err
	}

	err = cache.db.Set(cache.ctx, "users", usersBytes, cache.expirationTime).Err()
	if err != nil {
		return err
	}

	return nil
}

func (cache *Cache) GetCacheUsers() ([]domain.User, error) {
	val, err := cache.db.Get(cache.ctx, "users").Bytes()
	if err != nil {
		return nil, err
	}

	var users []domain.User
	err = json.Unmarshal(val, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
