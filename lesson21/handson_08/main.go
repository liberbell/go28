package main

import (
	"fmt"
	"net/http"
)

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
	})
	fmt.Fprintln(w, "Cookie Written - Check your Browser")
	fmt.Fprintln(w, "in chrome go to: dev tools / applications / cookies")
}

func read(w http.ResponseWriter, r *http.Request) {

}
