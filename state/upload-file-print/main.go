package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var s string
		if r.Method == http.MethodPost {
			f, h, err := r.FormFile("simple-upload")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			defer f.Close()

			fmt.Println("\nFile :", f, "\nHeader :", h)

			bs, err := ioutil.ReadAll(f)
			if err != nil {
				log.Fatal("Sorry bruv", err)
			}
			s = string(bs)
		}
		w.Header().Set("Content-Type", "text/html; charset=utf8")
		io.WriteString(w, `
			<form method="POST" enctype="multipart/form-data">
				<input type="file" name="simple-upload">
				<input type="submit">
			</form>
		`+s)

	})

	http.ListenAndServe(":8000", nil)
}
