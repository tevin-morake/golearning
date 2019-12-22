package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlFormValue := r.FormValue("q")
		// fmt.Fprintf(w, "Do my search : %s", urlFormValue)
		io.WriteString(w, "Do my search : "+urlFormValue)
	})

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8000", nil)
}
