package factor

import product "github.com/DesignPatterns/abstractFactory/product"

type Nike struct {
}

func (n *Nike) makeShoe() *product.IShoe {
	return &product.IShoe{
		product.Shoe: product.Shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *Nike) makeShirt() *product.Shirt {
	return &product.Shirt{
		product.Shirt: product.Shirt{
			logo: "nike",
			size: 14,
		},
	}
}