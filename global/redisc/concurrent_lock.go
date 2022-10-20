package redisc

import (
	"context"
	"time"
)

// RedisConcurrentLockInTime 获取redis并发锁，固定时间内禁止重发
func RedisConcurrentLockInTime(ctx context.Context, client RedisClientIOCInterface, lockKey string, lockTime time.Duration) (holdLock bool, err error) {
	rdsClient, err := client.GetClient()
	if err != nil {
		return false, err
	}
	holdLock, err = rdsClient.SetNX(ctx, "lock."+lockKey, time.Now(), lockTime).Result()
	if err != nil {
		return
	}
	return
}
