package AbstractFactory

import (
	"fmt"
	"testing"
)

func TestRedisCacheFactory_NewRedisFactory(t *testing.T) {
	var redisCacheFactory RedisCacheFactory
	rds, err := redisCacheFactory.NewRedisFactory()
	if err != nil {
		return
	}

	err = rds.Set("key", "rds")
	if err != nil {
		return
	}

	err, v := rds.Get("key")
	if err != nil {
		return
	}

	fmt.Println(v)
}

func TestMemCacheFactory_NewMemCacheFactory(t *testing.T) {
	var memCacheFactory MemCacheFactory
	mem, err := memCacheFactory.NewMemCacheFactory()
	if err != nil {
		return
	}

	err = mem.Set("key", "mem")
	if err != nil {
		return
	}

	err, v := mem.Get("key")
	if err != nil {
		return
	}

	fmt.Println(v)
}
