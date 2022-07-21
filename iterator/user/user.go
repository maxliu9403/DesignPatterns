package user

import (
	"fmt"
)

type User struct {
	name string
	age  int
}

func GetUserInfo() {
	user1 := &User{
		name: "a",
		age:  30,
	}
	user2 := &User{
		name: "b",
		age:  20,
	}

	userCollection := Collection{
		Users: []*User{user1, user2},
	}

	it := userCollection.createUserIterator()

	for it.HasNext() {
		user := it.GetNext()
		fmt.Printf("User is %+v\n", user)
	}

}
