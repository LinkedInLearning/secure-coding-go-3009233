package main

import (
	"fmt"
	"log"
	"net/http"
)

func isAllowed(user, action, resource string) bool {
	if action == http.MethodGet {
		return true // anyone can read
	}

	return user == resource
}

func messagesHandler(w http.ResponseWriter, r *http.Request) {
	user := requestUser(r)
	resource := r.URL.Query().Get("user")
	if !isAllowed(user, r.Method, resource) {
		http.Error(w, "not allowed", http.StatusUnauthorized)
		return
	}

	// TODO:
	fmt.Fprintf(w, "OK\n")
}

func main() {
	http.HandleFunc("/messages", messagesHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
