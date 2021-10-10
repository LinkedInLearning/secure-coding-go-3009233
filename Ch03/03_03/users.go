package main

import (
	"fmt"
)

type User struct {
	Login   string
	Age     int
	Address string
}

func findUser(login string) *User {
	// TODO:
	return &User{
		Login:   login,
		Age:     21,
		Address: "23519 West, Civic Center Way, Malibu, CA 90265",
	}
}

func userFriends(u *User) ([]*User, error) {
	// TODO:
	return nil, &UserError{"friend not enabled", u}
}

type UserError struct {
	Reason string
	User   *User
}

func (e *UserError) Error() string {
	return fmt.Sprintf("%s (%#v)", e.Reason, e.User)
}
