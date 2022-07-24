package product

// 抽象产品
type IShirt interface {
	setLog(log string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) setLog(log string) {
	s.logo = log
}

func (s *Shirt) setSize(size int) {
	s.size = size
}

func (s *Shirt) getLog() string {
	return s.logo
}

func (s *Shirt) getSize() int {
	return s.size
}
