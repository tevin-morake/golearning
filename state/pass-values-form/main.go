package main

import (
	"io"
	"net/http"
	"text/template"
)

var tpl *template.Template

type person struct {
	FirstName  string
	LastName   string
	Age        string
	Subscribed bool
}

func init() {
	tpl = template.Must(template.ParseGlob("*"))
}
func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		firstN := r.FormValue("first")
		lastN := r.FormValue("last")
		age := r.FormValue("age")
		subscr := r.FormValue("subscribed-bool") == "on"

		w.Header().Set("Content-Type", "text/html; charset=utf8")
		if err := tpl.ExecuteTemplate(w, "index.gohtml", person{firstN, lastN, age, subscr}); err != nil {
			// fmt.Fprintf(w, "<h1>Something Definitely went wrong : %s</h1>", err.Error())
			io.WriteString(w, err.Error())
		}
	})

	http.ListenAndServe(":8000", nil)
}
