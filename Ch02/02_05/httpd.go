package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Payment struct {
	Time   time.Time
	User   string
	To     string
	Amount float64
}

func paymenetHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)
	var p Payment
	if err := dec.Decode(&p); err != nil {
		log.Printf("error: %s", err)
		http.Error(w, "bad JSON", http.StatusBadRequest)
		return
	}

	// TODO: Process payment

	fmt.Fprintf(w, "OK\n")
}

func main() {
	http.HandleFunc("/payment", paymenetHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
