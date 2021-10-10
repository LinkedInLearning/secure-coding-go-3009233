package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	start := time.Now()
	n, err := io.Copy(io.Discard, r.Body)
	if err != nil {
		http.Error(w, "can't copy", http.StatusBadRequest)
		return
	}

	log.Printf("%d bytes in %v", n, time.Since(start))
	fmt.Fprintf(w, "%d bytes digested", n)
}

func main() {
	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
