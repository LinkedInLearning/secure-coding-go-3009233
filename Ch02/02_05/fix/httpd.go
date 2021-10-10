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

func (p *Payment) Validate() error {
	if p.Amount <= 0 {
		return fmt.Errorf("bad Amount in %#v", p)
	}

	return nil
}

func paymenetHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)
	var p Payment
	if err := dec.Decode(&p); err != nil {
		http.Error(w, "bad JSON", http.StatusBadRequest)
		return
	}

	if err := p.Validate(); err != nil {
		log.Printf("error: paymenetHandler - %s", err)
		http.Error(w, "bad data", http.StatusBadRequest)
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
