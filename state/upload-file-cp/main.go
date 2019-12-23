package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var s string
		if r.Method == http.MethodPost {
			f, h, err := r.FormFile("simple-form")
			if err != nil {
				fmt.Fprintf(w, "Something went wrong %s", err.Error())
				return
			}
			defer f.Close()
			fmt.Println(f, h)

			bs, err := ioutil.ReadAll(f)
			if err != nil {
				fmt.Fprintf(w, "Something went wrong %s", err.Error())
				return
			}
			s = string(bs)
			// ioutil.WriteFile("dummy.txt", bs, 777)
			dst, err := os.Create("dummy.txt")
			if err != nil {
				fmt.Fprintf(w, "Eish : %s", err.Error())
				return
			}
			defer dst.Close()
			dst.Write(bs)

		}
		w.Header().Set("Content-Type", "text/html; charset=utf8")
		io.WriteString(w, `
			<form method="POST" enctype="multipart/form-data">
				<input type="file"  name="simple-form">
				<input type="submit">
			</form>
		`+s)

	})
	http.ListenAndServe(":8000", nil)
}
