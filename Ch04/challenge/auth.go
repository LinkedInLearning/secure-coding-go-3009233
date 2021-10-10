package main

import (
	"net/http"
	"strings"
)

type Role uint

// Roles (as bitmask)
const (
	Guest Role = 1 << iota
	Developer
	Admin
)

type User struct {
	Login string
	Roles Role // bitmask of roles
}

func (u *User) HasRole(r Role) bool {
	return u.Roles&r != 0
}

func RequestUser(r *http.Request) (User, bool) {
	token := strings.TrimPrefix("Bearer ", r.Header.Get("Authorization"))
	u, ok := users[token]
	return u, ok
}

// token -> user
// FIXME: Use a real and secure database
var users = map[string]User{
	"m4ur1c3": {"moss", Admin & Developer & Guest},
	"b4rb3r":  {"jen", Guest},
}
