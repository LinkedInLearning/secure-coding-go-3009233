package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"
)

type Activity struct {
	User        string    `json:"user"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Description string    `json:"description"`
}

func processActivity(r io.Reader) error {
	var act Activity

	dec := json.NewDecoder(r)
	if err := dec.Decode(&act); err != nil {
		return err
	}

	log.Printf("activity: %#v", act)
	// TODO: Store in database

	return nil
}

func main() {
	if err := processActivity(os.Stdin); err != nil {
		log.Fatal(err)
	}
}
