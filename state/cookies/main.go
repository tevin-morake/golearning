package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// Changed this to a counter that checks how many times a person visits the /set uri
	// to delete a cookie, just set MaxAge = -1. Done!
	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("ScriptWhizCookieSiteVisit")
		if err != nil {
			http.SetCookie(w, &http.Cookie{
				Name:  "ScriptWhizCookieSiteVisit",
				Value: "1",
			})
		} else {
			stringCookie := c.Value
			intCookieVal, err := strconv.Atoi(stringCookie)
			if err != nil {
				fmt.Println("Failed to convert string cookie to number")
			}
			intCookieVal++
			strCookie := strconv.Itoa(intCookieVal)
			http.SetCookie(w, &http.Cookie{
				Name:  "ScriptWhizCookieSiteVisit",
				Value: strCookie,
			})
		}
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
