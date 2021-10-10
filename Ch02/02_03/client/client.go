package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var message = []byte("slow and steady wins the race ")

type SlowReader struct {
	i    int
	max  int
	read int
}

// Read implements io.Reader
func (s *SlowReader) Read(p []byte) (int, error) {
	time.Sleep(100 * time.Millisecond)
	if s.i >= s.max {
		return 0, io.EOF
	}
	s.i++

	size := 0
	for size < len(p) {
		size += copy(p[size:], message)
	}
	s.read += len(p)
	return len(p), nil
}

func main() {
	body := &SlowReader{max: 10}
	start := time.Now()
	resp, err := http.Post("http://localhost:8080/", "text/plain", body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %d %s", resp.StatusCode, resp.Status)
	}

	duration := time.Since(start)
	fmt.Printf("sent %d bytes in %v\n", body.read, duration)
}
