package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	mux := httprouter.New()
	//These are the gets
	mux.GET("/introduction", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "<h1>I am</h1>")
	})

	//These are the posts
	http.ListenAndServe(":8000", mux)
}
