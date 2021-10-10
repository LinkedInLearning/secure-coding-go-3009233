package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	apiKey string
	apiURL = "https://httpbin.org/basic-auth/key/l3tm3in"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		http.Error(w, "can't create request", http.StatusInternalServerError)
		return
	}
	req.SetBasicAuth("key", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(w, "can't call API", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "OK\n")
}

func loadConfig() error {
	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		return fmt.Errorf("missing API_KEY")
	}
	return nil
}

func main() {
	if err := loadConfig(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
