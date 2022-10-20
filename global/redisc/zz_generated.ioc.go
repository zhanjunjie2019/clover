//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package redisc

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
	"github.com/go-redis/redis/v9"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &redisClient_{}
		},
	})
	redisClientStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &RedisClient{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(redisClientStructDescriptor)
}

type redisClient_ struct {
	Create_    func(addr, password string, db uint8) error
	create_    func() error
	GetClient_ func() (*redis.Client, error)
}

func (r *redisClient_) Create(addr, password string, db uint8) error {
	return r.Create_(addr, password, db)
}

func (r *redisClient_) create() error {
	return r.create_()
}

func (r *redisClient_) GetClient() (*redis.Client, error) {
	return r.GetClient_()
}

type RedisClientIOCInterface interface {
	Create(addr, password string, db uint8) error
	create() error
	GetClient() (*redis.Client, error)
}

var _redisClientSDID string

func GetRedisClientSingleton() (*RedisClient, error) {
	if _redisClientSDID == "" {
		_redisClientSDID = util.GetSDIDByStructPtr(new(RedisClient))
	}
	i, err := singleton.GetImpl(_redisClientSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*RedisClient)
	return impl, nil
}

func GetRedisClientIOCInterfaceSingleton() (RedisClientIOCInterface, error) {
	if _redisClientSDID == "" {
		_redisClientSDID = util.GetSDIDByStructPtr(new(RedisClient))
	}
	i, err := singleton.GetImplWithProxy(_redisClientSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(RedisClientIOCInterface)
	return impl, nil
}

type ThisRedisClient struct {
}

func (t *ThisRedisClient) This() RedisClientIOCInterface {
	thisPtr, _ := GetRedisClientIOCInterfaceSingleton()
	return thisPtr
}
