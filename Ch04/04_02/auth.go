package main

type User struct {
	Login string
}

func userFromToken(token string) *User {
	// FIXME: JWT, Oauth2 ...
	if token == "baz00ka" {
		return &User{"joe"}
	}
	return nil
}
