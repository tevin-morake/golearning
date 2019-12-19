package main

import (
	"fmt"
	"net/http"
)

type bot map[string]string

func main() {
	var max = make(bot)
	max["Name"] = "Bot1"
	max["Function"] = "Listen & Serve"

	http.ListenAndServe(":8080", max)

}

func (s bot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is our first connection to a simple server in net/http")
}
