package main

import "time"

type Audit struct {
	User   string
	Time   time.Time
	Action string
}

func loadAudits() []Audit {
	return []Audit{
		{
			"Elliot Alderson",
			time.Date(2015, time.June, 24, 23, 27, 23, 999, time.UTC),
			"access fsociety.dat",
		},
	}
}
