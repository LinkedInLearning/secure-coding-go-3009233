package main

import (
	"encoding/json"
	"fmt"
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

func (a *Activity) Validate() error {
	if len(a.User) == 0 {
		return fmt.Errorf("missing user in %#v", a)
	}

	if a.EndTime.Before(a.StartTime) {
		return fmt.Errorf("end before start in %#v", a)
	}

	if len(a.Description) == 0 {
		return fmt.Errorf("missing description in %#v", a)
	}

	return nil
}

func processActivity(r io.Reader) error {
	var act Activity
	const maxSize = 10 * 1024

	r = io.LimitReader(r, maxSize)
	dec := json.NewDecoder(r)
	if err := dec.Decode(&act); err != nil {
		return err
	}

	if err := act.Validate(); err != nil {
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
