package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<h1>What kind of human do you take me for 'grammer ?</h1>")
	})

	http.ListenAndServe(":8000", nil)
}
