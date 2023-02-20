package redisc

import (
	"context"
	"github.com/redis/go-redis/extra/redisotel/v9"
	redis "github.com/redis/go-redis/v9"
	"sync"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type RedisClient struct {
	addr     string
	password string
	db       uint8
	rw       sync.RWMutex
	*redis.Client
}

func (r *RedisClient) Create(addr, password string, db uint8) (bool, error) {
	r.rw.Lock()
	defer r.rw.Unlock()
	if r.addr != addr || r.password != password || r.db != db {
		r.addr = addr
		r.password = password
		r.db = db
		err := r.create()
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func (r *RedisClient) create() error {
	client := redis.NewClient(&redis.Options{
		Addr:     r.addr,
		Password: r.password,
		DB:       int(r.db),
	})
	err := client.Ping(context.Background()).Err()
	if err != nil {
		return err
	}
	err = redisotel.InstrumentTracing(client)
	if err != nil {
		return err
	}
	err = redisotel.InstrumentMetrics(client)
	if err != nil {
		return err
	}
	if r.Client != nil {
		err = r.Client.Close()
		if err != nil {
			return err
		}
	}
	r.Client = client
	return nil
}

func (r *RedisClient) GetClient() *redis.Client {
	r.rw.RLock()
	defer r.rw.RUnlock()
	return r.Client
}
