package product

// 抽象产品
type IShoe interface {
	setLog(log string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shoe struct {
	Logo string
	Size int
}

func (s *Shoe) setLog(log string) {
	s.Logo = log
}

func (s *Shoe) setSize(size int) {
	s.Size = size
}

func (s *Shoe) getLog() string {
	return s.Logo
}

func (s *Shoe) getSize() int {
	return s.Size
}
