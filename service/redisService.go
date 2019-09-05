package service

import (
	"encoding/json"

	"github.com/go-redis/redis"
	"go.uber.org/dig"
)

type cacheService struct {
	Cache *redis.Client
}

type cacheParams struct {
	dig.In

	Cache *redis.Client
}

//CacheServiceInterface - interface
type CacheServiceInterface interface {
	CreateCache(key string, value interface{}, timeout int) error
	FindCache(key string, dest interface{}) error
	DeleteCache(key string) error
}

//NewCacheServiceInterface -
func NewCacheServiceInterface(c cacheParams) CacheServiceInterface {
	return cacheService{
		Cache: c.Cache,
	}
}

func (r cacheService) FindCache(key string, dest interface{}) error {
	res, err := r.Cache.Get(key).Result()
	if err != nil || err == redis.Nil {
		return err
	}
	return json.Unmarshal([]byte(res), dest)

}

func (r cacheService) DeleteCache(key string) error {
	return r.Cache.Del(key).Err()
}

func (r cacheService) CreateCache(key string, value interface{}, timeout int) error {
	encoded, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.Cache.SetNX(key, encoded, 0).Err()
}
