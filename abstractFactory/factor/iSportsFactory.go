package factor

import "github.com/DesignPatterns/abstractFactory/product"

// 抽象工厂接口
type ISportsFactory interface {
	makeShop() product.IShoe
	makeShirt() product.IShirt
}
