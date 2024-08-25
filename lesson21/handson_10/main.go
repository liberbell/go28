package main

import (
	"net/http"
	"strconv"
)

func foo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("my-cookie")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
		}
	}
	count, _ := strconv.Atoi(cookie.Value)
	count++
	cookie.Value = strconv.Itoa(count)
}
