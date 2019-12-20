package main

import (
	"io"
	"net/http"
)

type cars map[string]string

func (c cars) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Have you ever seen these car(s) ?"+c["Datsun"])
}

func main() {
	c := make(cars)
	c["Datsun"] = "<h1> Go Beyond your limits</h1>"

	http.Handle("/", c)

	http.ListenAndServe(":8000", nil)

}
