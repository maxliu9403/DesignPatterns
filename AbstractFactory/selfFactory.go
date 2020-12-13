package AbstractFactory

/*
1. 定义具体实例子接口
2. 定义具体实例结构体
3. 实现接口
4. 定义抽象工厂，返回具体实例
5. 实现具体实例工厂
6. 创建具体的实例
*/

// 1. 定义具体实例子接口
type CacheSelf interface {
	Set(string, string) error
	Get(string) (error, interface{})
}

// 2. 定义具体实例结构体
type Rds struct {
	Data map[string]interface{}
}

// 3. 实现接口
func (r *Rds) Set(k, v string) error {
	r.Data[k] = v
	return nil
}

func (r *Rds) Get(k string) (error, interface{}) {
	return nil, r.Data[k]
}

// 4. 定义抽象工厂，返回具体实例
type CacheFactories interface {
	Create() CacheSelf
}

// 5. 实现具体实例工厂,rds
type RdsFactories struct {
}

func (r *RdsFactories) NewRdsFactories() CacheSelf {
	return &Rds{Data: map[string]interface{}{}}
}

// 6. 创建具体的实例
func (r *RdsFactories) Create() CacheSelf {
	rds := r.Create()
	return rds

}
