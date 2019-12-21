package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("img/league.png")
		defer f.Close()
		if err != nil {
			fmt.Fprintf(w, "<h1>Dude, theres a problem loading your picture</h1>")
		}

		io.Copy(w, f)
	})

	http.ListenAndServe(":8000", nil)
}
