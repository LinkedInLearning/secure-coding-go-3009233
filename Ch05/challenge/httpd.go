package main

import (
	"expvar"
	"fmt"
	"log"
	"net/http"
)

var (
	numCalls = expvar.NewInt("messages.calls")
)

func messagesHandler(w http.ResponseWriter, r *http.Request) {
	numCalls.Add(1)

	// TODO:
	fmt.Fprintf(w, "TBD\n")
}

func main() {
	http.HandleFunc("/messages", messagesHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
