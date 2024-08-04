package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func logHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't read", http.StatusBadRequest)
		return
	}

	// TODO: Save in database

	fmt.Fprintf(w, "%d bytes stored\n", len(data))
}

func main() {
	http.HandleFunc("/log", logHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
