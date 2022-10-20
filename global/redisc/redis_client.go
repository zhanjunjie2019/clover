package redisc

import (
	"context"
	"github.com/go-redis/redis/v9"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type RedisClient struct {
	addr     string
	password string
	db       uint8
	*redis.Client
}

func (r *RedisClient) Create(addr, password string, db uint8) error {
	r.addr = addr
	r.password = password
	r.db = db
	return r.create()
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
	if r.Client != nil {
		err = r.Client.Close()
		if err != nil {
			return err
		}
	}
	r.Client = client
	return nil
}

func (r *RedisClient) GetClient() (*redis.Client, error) {
	if r.Client == nil {
		err := r.create()
		if err != nil {
			return nil, err
		}
	}
	return r.Client, nil
}
