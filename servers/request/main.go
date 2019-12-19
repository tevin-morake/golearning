package main

import (
	"net/http"
	"text/template"
)

type CoolDrink struct {
	Name    string
	Company string
}

var tpl *template.Template

func (c CoolDrink) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "simpleform.gohtml", r.Form)
}

func init() {
	tpl = template.Must(template.ParseFiles("simpleform.gohtml"))
}

func main() {
	coke := CoolDrink{
		Name:    "Coke Zero",
		Company: "Coca Cola",
	}

	http.ListenAndServe(":8081", coke)

}
