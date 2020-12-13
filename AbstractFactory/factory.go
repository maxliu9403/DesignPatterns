package AbstractFactory

/*
1. 定义具体实例子接口
2. 定义具体实例结构体
3. 实现接口
4. 定义抽象工厂，返回具体实例
5. 实现具体实例工厂
6. 创建具体的实例
*/

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

// 定义抽象的Cache工厂
type CacheFactory interface {
	Create() Cache
}

// 实现具体Redis Cache工厂
type RedisCacheFactory struct {
}

func (r *RedisCacheFactory) NewRedisFactory() (Cache, error) {
	return &Redis{data: map[string]string{}}, nil

}

// 创建具体的Redis
func (r *RedisCacheFactory) Create() Cache {
	cache, err := r.NewRedisFactory()
	if err != nil {
		return nil
	}

	return cache
}

// 实现实现具体的Mem Cache工厂
type MemCacheFactory struct {
}

func (m *MemCacheFactory) NewMemCacheFactory() (Cache, error) {
	return &MemCache{data: map[string]string{}}, nil

}

// 创建具体的MemCache
func (m *MemCacheFactory) Create() Cache {
	cache, err := m.NewMemCacheFactory()
	if err != nil {
		return nil
	}

	return cache

}
