package user

type Iterator struct {
	index int
	users []*User
}

// 用户集合
type Collection struct {
	Users []*User
}

// 创建用户迭代器
func (u *Collection) createUserIterator() *Iterator {
	return &Iterator{
		users: u.Users,
	}
}

func (u *Iterator) HasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false

}
func (u *Iterator) GetNext() *User {
	if u.HasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}
