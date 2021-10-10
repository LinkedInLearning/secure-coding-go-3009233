package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func auditsHandler(w http.ResponseWriter, r *http.Request) {
	user, passwd, ok := r.BasicAuth()
	if !ok || !authUser(user, passwd) {
		http.Error(w, "not authorized", http.StatusUnauthorized)
	}

	audits := loadAudits()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(audits)
}

func main() {
	http.HandleFunc("/audits", auditsHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
