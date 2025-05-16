package db

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	ronce  sync.Once
)

// InitRedis initializes the singleton Redis client
func initRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Username: "default",
		Password: "hashxpos12345",
		DB:       0,
	})
}

// GetRedis returns the singleton Redis client instance
func GetRedis() *redis.Client {
	ronce.Do(initRedis)
	return client
}

// Set wraps redis SET operation
func RedisSet(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return GetRedis().Set(ctx, key, value, expiration).Err()
}

// Get wraps redis GET operation
func RedisGet(ctx context.Context, key string) (string, error) {
	return GetRedis().Get(ctx, key).Result()
}

func RedisRemoveKey(ctx context.Context, key string) (int64, error) {
	return GetRedis().Del(ctx, key).Result()
}

func CreateKey(s1 string, s2 string) string {
	return fmt.Sprintf("%s:%s", s1, s2)
}

// RedisHSet creates or updates multiple hash fields in Redis
func RedisHSet(ctx context.Context, key string, fields interface{}, expiration time.Duration) error {
	return GetRedis().HSet(ctx, key, fields, expiration).Err()
}

// RedisHGet retrieves a hash field value from Redis
func RedisHGet(ctx context.Context, key string, field string) (string, error) {
	return GetRedis().HGet(ctx, key, field).Result()
}

// RedisHGetAll retrieves all fields and values of a hash from Redis
func RedisHGetAll(ctx context.Context, key string) (map[string]string, error) {
	return GetRedis().HGetAll(ctx, key).Result()
}

// RedisHDel removes one or more hash fields from Redis
func RedisHDel(ctx context.Context, key string, fields ...string) (int64, error) {
	return GetRedis().HDel(ctx, key, fields...).Result()
}

// RedisHExists checks if a hash field exists in Redis
func RedisHExists(ctx context.Context, key string, field string) (bool, error) {
	return GetRedis().HExists(ctx, key, field).Result()
}
