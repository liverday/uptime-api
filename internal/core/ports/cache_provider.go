package ports

import (
	"context"
	"time"
)

type CacheProvider interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string, target interface{}) error
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) bool
	Delete(ctx context.Context, key string) error
}
