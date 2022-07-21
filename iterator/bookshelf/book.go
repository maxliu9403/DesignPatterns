package bookshelf

import "fmt"

type Book struct {
	Name string
	Auth string
	Num  int
}

func GetBoolInfo() {
	user1 := &Book{
		Name: "a",
	}
	user2 := &Book{
		Name: "b",
	}

	col := Collection{
		Books: []*Book{user1, user2},
	}

	it := col.createBoolIterator()
	it.AddBook(&Book{Name: "c"})
	for it.HasNext() {
		user := it.GetNext()
		fmt.Printf("Book is %+v\n", user)
	}


}
