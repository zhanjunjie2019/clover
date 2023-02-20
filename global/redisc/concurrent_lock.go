package redisc

import (
	"context"
	"time"
)

// RedisConcurrentLockInTime 获取redis并发锁，固定时间内禁止重发
func RedisConcurrentLockInTime(ctx context.Context, client RedisClientIOCInterface, lockKey string, lockTime time.Duration) (holdLock bool, err error) {
	holdLock, err = client.GetClient().SetNX(ctx, "lock."+lockKey, time.Now(), lockTime).Result()
	if err != nil {
		return
	}
	return
}
