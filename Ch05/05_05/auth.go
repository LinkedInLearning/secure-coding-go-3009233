package main

import "net/http"

type User struct {
	Login string
	Name  string
	// TODO: More fields
}

func loginUser(login, passwd string) (User, bool) {
	// FIXME: Use real auth database
	if login == "daffy" && passwd == "r4bb1ts3as0n" {
		return User{"daffy", "Daffy Duck"}, true
	}

	return User{}, false
}

func setUserCookie(w http.ResponseWriter, u User) {
	// TODO: JWT
}
