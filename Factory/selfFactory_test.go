package Factory

import (
	"fmt"
	"testing"
)

func TestCacheFactories_NewCache(t *testing.T) {
	var cacheFactories CacheFactories
	cache, err := cacheFactories.NewCache(mem)
	if err != nil {
		return
	}

	err = cache.Set("rds", "value")
	if err != nil {
		return
	}

	value := cache.Get("rds")
	fmt.Println(value)

}
