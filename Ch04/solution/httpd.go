package main

import (
	"fmt"
	"log"
	"net/http"
)

func adminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You're in!\n")
}

// requireAdmin is a middleware allowing only users with Admin role to access the handler
func requireAdmin(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		u, ok := RequestUser(r)
		if !ok {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		if !u.HasRole(Admin) {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func main() {
	h := requireAdmin(http.HandlerFunc(adminHandler))
	http.Handle("/admin", h)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
