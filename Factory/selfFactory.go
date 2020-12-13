package Factory

import (
	"errors"
)

type CacheSelf interface {
	Get(string) interface{}
	Set(string, string) error
}

type Rds struct {
	Data map[string]interface{}
}

type Mem struct {
	Data map[string]interface{}
}

func (r *Rds) Get(key string) interface{} {
	return r.Data[key]
}

func (r *Rds) Set(key, value string) error {
	r.Data[key] = value
	return nil

}

func (m *Mem) Get(key string) interface{} {
	return m.Data[key]
}

func (m *Mem) Set(key, value string) error {
	m.Data[key] = value
	return nil

}

// 定义cache工厂
type CacheFactories struct {
}

type CacheType int

const (
	redis CacheType = iota
	mem
)

func (c *CacheFactories) NewCache(cacheType CacheType) (CacheSelf, error) {
	switch cacheType {
	case redis:
		return &Rds{Data: map[string]interface{}{}}, nil
	case mem:
		return &Mem{Data: map[string]interface{}{}}, nil
	default:
		return nil, errors.New("error")
	}

}
