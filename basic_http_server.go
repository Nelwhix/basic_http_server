package main

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
	"io"
	"github.com/gorilla/handlers"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "5173"
	ADMIN_USER = "Nelwhix"
	ADMIN_PASSWORD = "admin"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "I love me!")
}

func BasicAuth(handler http.HandlerFunc, realm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(ADMIN_USER)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(ADMIN_PASSWORD)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("You are Unauthorized to access the application.\n"))
			return
    	}
    	handler(w, r)
	}
}

func main() {
  mux := http.NewServeMux()

  mux.HandleFunc("/", BasicAuth(helloWorldHandler, "Please enter your username and password"))

  log.Printf("Server starting on port %v\n", CONN_PORT)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(CONN_HOST + ":" + CONN_PORT), handlers.CompressHandler(mux)))
  
}


