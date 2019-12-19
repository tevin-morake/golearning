package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type CoolDrink struct {
	Name    string
	Company string
}

var tpl *template.Template

func (c CoolDrink) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "The Kingdom of Coke has failed to please your thirst sire : "+err.Error())
		return
	}
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
