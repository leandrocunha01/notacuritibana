package models

import (
	"fmt"
)

type User struct {
	Base
	Name string
}

func (user *User) NewUser() {
	fmt.Println(user.Name)
}
