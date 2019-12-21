package main

import (
	"fmt"
	"net/http"
	"os"
)

type villain string

func main() {
	// var newVillan string
	// newVillan = "Shigaraki Tomura"
	http.Handle("/", new(villain))

	http.ListenAndServe(":8000", nil)
}

func (c villain) ServeHTTP(output http.ResponseWriter, input *http.Request) {
	f, err := os.Open("img/league.png")
	defer f.Close()
	if err != nil {
		fmt.Fprintf(output, "<h1>Well Serving Content with package http seems to have failed</h1>")
		return
	}

	/*
		 * Serve Content is not necessary for me. It's used in the e tag according to Todd.
		 Beats me. I'll use ServeFiles. Read the docs to understand benefit of ServeFiles over io.Copy etc
		fInfo, err := f.Stat()
		if err != nil {
			fmt.Fprintln(output, "<h1>Something definitely went wrong</h1>")
			return
		}
		http.ServeContent(output, input, fInfo.Name(), fInfo.ModTime(), f) */
	http.ServeFile(output, input, "img/league.png")
}
