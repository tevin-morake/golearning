package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:  "ScriptWhizCookie",
			Value: "Doranbolt",
		})
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("ScriptWhizCookie")
		if err != nil {
			log.Fatal("This cookie don't taste right", err.Error())
		}
		fmt.Fprintf(w, "Cookie can be seen here : %s", c)
	})

	http.ListenAndServe(":8000", nil)
}
