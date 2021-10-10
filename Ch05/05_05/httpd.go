package main

import (
	"expvar"
	"fmt"
	"log"
	"net/http"
)

var (
	badLogins = expvar.NewInt("login.errors")
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
		return
	}

	login, passwd := r.FormValue("user"), r.FormValue("passwd")
	if len(login) == 0 || len(passwd) == 0 {
		http.Error(w, "missing auth", http.StatusUnauthorized)
		return

	}

	u, ok := loginUser(login, passwd)
	if !ok {
		badLogins.Add(1)
		log.Printf("bad %q login from %s", login, r.RemoteAddr)
		http.Error(w, "bad auth", http.StatusUnauthorized)
		return
	}

	log.Printf("%q login from %s", login, r.RemoteAddr)
	setUserCookie(w, u)

	// TODO: Redirect to main page
	fmt.Fprintf(w, "Welcome %s!\n", u.Name)
}

func main() {
	http.HandleFunc("/login", loginHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
