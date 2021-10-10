package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if err := checkHealth(); err != nil {
		http.Error(w, "health check failed", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "OK\n")
}

func messagesHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*User)
	if !ok {
		http.Error(w, "no user", http.StatusInternalServerError)
		return
	}
	log.Printf("user: %s", user)

	// FIXME:
	fmt.Fprint(w, "[]\n")
}

func authToken(r *http.Request) string {
	hdr := r.Header.Get("Authorization")
	return strings.TrimPrefix(hdr, "Bearer ")
}

func requireAuth(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		token := authToken(r)
		user := userFromToken(token)
		if user == nil {
			http.Error(w, "bad authentication", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	h := requireAuth(http.HandlerFunc(messagesHandler))
	http.Handle("/messages", h)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
