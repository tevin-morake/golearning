package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if err := tpl.Execute(w, ""); err != nil {
			fmt.Fprintf(w, "<h1>Oh Dear, something went wrong...</h1>")
		}
	})

	http.ListenAndServe(":8000", mux)

}
