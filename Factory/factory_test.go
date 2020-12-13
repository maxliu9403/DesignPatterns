package Factory

import (
	"fmt"
	"testing"
)

func TestCacheFactory_NewCacheFactory(t *testing.T) {
	fmt.Println(11)
	var CacheFactory CacheFactory
	rds, err := CacheFactory.NewCacheFactory(redis)
	if err != nil {
		return
	}

	err = rds.Set("key", "value1")
	if err != nil {
		return
	}

	err, value := rds.Get("key")
	if err != nil {
		return
	}

	mem, err := CacheFactory.NewCacheFactory(memCache)
	if err != nil {
		return
	}

	err = mem.Set("key1", "value2")
	if err != nil {
		return
	}

	err, value1 := mem.Get("key1")
	if err != nil {
		return
	}

	fmt.Println(value, value1)

	return
}
