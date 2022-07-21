package bookshelf

type Iterator struct {
	index int
	books []*Book
}

type Collection struct {
	Books []*Book
}

// 创建迭代器
func (u *Collection) createBoolIterator() *Iterator {
	return &Iterator{
		books: u.Books,
	}
}

func (u *Iterator) HasNext() bool {
	if u.index < len(u.books) {
		return true
	}
	return false
}

func (u *Iterator) GetNext() *Book {
	if u.HasNext() {
		user := u.books[u.index]
		u.index++
		return user
	}
	return nil
}

func (u *Iterator) AddBook(book *Book) {
	u.books = append(u.books, book)
}
