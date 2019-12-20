package main

import (
	"io"
	"net/http"
)

type numo int

func (n numo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/home":
		io.WriteString(w, "<h1>Submerge, for you are home</h1>")
	case "/":
		io.WriteString(w, "<h1>Fear not, for I am with you always</h1>")
	}
}

func main() {
	var x numo
	x = 3
	http.ListenAndServe(":8000", x)

}
