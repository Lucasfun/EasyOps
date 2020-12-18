package user

import "github.com/Lucasfun/EasyOps/gomockDemo/person"

type User struct {
	Person person.Male
}

func NewUser(p person.Male) *User {
	return &User{Person: p}
}

func (u *User) GetUser(id int64) error {
	return u.Person.Get(id)
}
