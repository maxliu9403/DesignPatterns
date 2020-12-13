package Factory

import (
	"errors"
)

type Cache interface {
	Set(string, string) error
	Get(string) (error, string)
}

type Redis struct {
	data map[string]string
}

type MemCache struct {
	data map[string]string
}

func (r *Redis) Set(key, value string) error {
	r.data[key] = value
	return nil
}

func (r *Redis) Get(key string) (error, string) {
	value := r.data[key]
	return nil, value
}

func (m *MemCache) Set(key, value string) error {
	m.data[key] = value
	return nil
}

func (m *MemCache) Get(key string) (error, string) {
	value := m.data[key]
	return nil, value
}

// 实现cache工厂
type CacheTypes int

const (
	rds CacheTypes = iota
	memCache
)

type CacheFactory struct {
}

func (c *CacheFactory) NewCacheFactory(CacheType CacheTypes) (Cache, error) {
	switch CacheType {
	case rds:
		return &Redis{data: map[string]string{}}, nil
	case memCache:
		return &MemCache{data: map[string]string{}}, nil
	default:
		return nil, errors.New("invalid CacheType！")
	}

}
