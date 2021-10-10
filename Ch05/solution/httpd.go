package main

import (
	"expvar"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var (
	numCalls = expvar.NewInt("messages.calls")
)

func messagesHandler(w http.ResponseWriter, r *http.Request) {
	numCalls.Add(1)

	// TODO:
	fmt.Fprintf(w, "TBD\n")
}

func checkAuth(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/debug") {
		login, passwd, ok := r.BasicAuth()
		if !ok || !isValidAuth(login, passwd) {
			http.Error(w, "bad auth", http.StatusUnauthorized)
			return
		}

	}

	http.DefaultServeMux.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/messages", messagesHandler)

	mux := http.HandlerFunc(checkAuth)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
