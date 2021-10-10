package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK\n")
}

func main() {
	http.HandleFunc("/health", healthHandler)

	if err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil); err != nil {
		log.Fatal(err)
	}
}
