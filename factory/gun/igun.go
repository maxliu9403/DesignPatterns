package gun

// 工厂方法
// 定义枪支的所有方法
// 将实例的生成交给子类
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}
