package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

var (
	templateText = `<!DOCTYPE html>
<html>
	<head>
		<title>{{.Name}}'s Messages</title>
	</head>
	<body>
	<h1>{{.Name}}'s Messages</h1>
	<p>
		You have {{len .Messages}} messages.
	</p>
	{{range .Messages}}
	<p>[{{.Time.Format "2006-01-02T15:04"}} {{.From}}] {{.Content}}</p>
	{{end}}
	</body>
</html>
`
	htmlTemplate = template.Must(template.New("messages").Parse(templateText))
)

type Message struct {
	Time    time.Time
	From    string
	Content string
}

type Params struct {
	Name     string
	Messages []Message
}

func messagesHandler(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Path[1:] // "/frodo" -> "frodo"
	ms, err := loadMessages(user)
	if err != nil {
		http.Error(w, "can't get messages", http.StatusBadRequest)
		return
	}

	params := Params{user, ms}
	if err := htmlTemplate.Execute(w, params); err != nil {
		log.Printf("error: %s", err)
	}
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
