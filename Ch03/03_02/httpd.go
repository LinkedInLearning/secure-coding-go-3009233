package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"
)

var htmlTemplate = `<!DOCTYPE html>
<html>
	<head>
		<title>%s's Messages</title>
	</head>
	<body>
	<h1>%s's Messages</h1>
	<p>
		You have %d messages.
	</p>
	%s
	</body>
</html>
`

type Message struct {
	Time    time.Time
	From    string
	Content string
}

func formateMessages(ms []Message) string {
	var buf bytes.Buffer
	for _, m := range ms {
		ts := m.Time.Format("2006-01-02T15:04")
		fmt.Fprintf(&buf, "<p>[%s %s] %s</p><hr />\n", ts, m.From, m.Content)
	}
	return buf.String()
}

func messagesHandler(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Path[1:] // "/frodo" -> "frodo"
	ms, err := loadMessages(user)
	if err != nil {
		http.Error(w, "can't get messages", http.StatusBadRequest)
		return
	}

	body := formateMessages(ms)
	fmt.Fprintf(w, htmlTemplate, user, user, len(ms), body)
}

func main() {
	http.HandleFunc("/", messagesHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func loadMessages(user string) ([]Message, error) {
	ms := []Message{
		{
			time.Date(2021, time.September, 23, 10, 33, 17, 0, time.UTC),
			"Pippin",
			"What about second breakfast?",
		},
		{
			time.Date(2021, time.September, 23, 14, 15, 32, 0, time.UTC),
			"Samy",
			"Where's the ring? <script>alert('Pwned!')</script>",
		},
	}

	return ms, nil
}
