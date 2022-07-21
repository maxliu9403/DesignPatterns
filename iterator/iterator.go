package iterator

import (
	"github.com/DesignPatterns/iterator/bookshelf"
	"github.com/DesignPatterns/iterator/user"
)

type UserIterator interface {
	HasNext() bool
	GetNext() *user.User
}

type BookIterator interface {
	HasNext() bool
	GetNext() *bookshelf.Book
	AddBook(book *bookshelf.Book)
}
