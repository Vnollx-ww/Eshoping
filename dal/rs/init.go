package rs

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var client *redis.Client

const (
	RedisLockKeyPrefix = "product_lock:"  // 锁的前缀，确保每个商品使用不同的锁
	LockTimeout        = 30 * time.Second // 锁的超时时间
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "my-redis:6379", // Redis 服务器地址
		Password: "",              // 没有密码则留空
		DB:       0,               // 使用默认数据库
	})
}
func AcquireLock(ctx context.Context, lockKey string) (bool, error) {
	// 使用 SETNX 命令来设置锁，锁的超时时间为 LockTimeout
	status := client.SetNX(ctx, lockKey, "locked", LockTimeout)
	return status.Result()
}
func ReleaseLock(ctx context.Context, lockKey string) error {
	// 删除 Redis 锁
	_, err := client.Del(ctx, lockKey).Result()
	return err
}
func SetKey(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return client.Set(ctx, key, value, expiration).Err()
}
func DelKey(ctx context.Context, key string) error {
	return client.Del(ctx, key).Err()
}
