package main

import (
	"io"
	"net/http"
)

type movie string

func (m movie) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Welcome Home Brother</h1>")

}
func main() {

	var classic movie
	classic = "License To Kill"

	mux := http.NewServeMux()
	mux.Handle("/", classic)

	http.ListenAndServe(":8000", mux)
}
