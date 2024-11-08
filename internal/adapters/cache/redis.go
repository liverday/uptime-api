package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"time"
)

type RedisCacheProvider struct {
	client *redis.Client
}

func NewCacheProvider() *RedisCacheProvider {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	log.Printf("Loading Redis cache provider with host: %s and port: %s\n", host, port)

	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
	})

	return &RedisCacheProvider{
		client: client,
	}
}

func (r *RedisCacheProvider) Set(ctx context.Context, key string, value interface{}) error {
	return r.client.Set(ctx, key, value, 0).Err()
}
func (r *RedisCacheProvider) Get(ctx context.Context, key string, target interface{}) error {
	val, err := r.client.Get(ctx, key).Result()

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(val), &target)

	if err != nil {
		log.Print("Error unmarshaling value from Redis: ", err)
		return err
	}

	return nil
}
func (r *RedisCacheProvider) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) bool {
	set, err := r.client.SetNX(ctx, key, value, expiration).Result()

	if err != nil {
		log.Print("Error setting value in Redis: ", err)
		return false
	}

	return set
}

func (r *RedisCacheProvider) Delete(ctx context.Context, key string) error {
	_, err := r.client.Del(ctx, key).Result()

	if err != nil {
		log.Print("Error deleting value from Redis: ", err)
		return err
	}

	return nil
}
