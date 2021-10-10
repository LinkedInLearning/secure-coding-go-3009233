package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func friendsHandler(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query().Get("user")
	if login == "" {
		http.Error(w, "bad path", http.StatusBadRequest)
		return
	}

	user := findUser(login)
	if user == nil {
		http.Error(w, "no such user", http.StatusBadRequest)
		return
	}

	friends, err := userFriends(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(friends)
}

func main() {
	http.HandleFunc("/friends", friendsHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
