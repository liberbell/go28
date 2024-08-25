package main

import "net/http"

func foo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("my-cookie")
	if err != nil {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
		}
	}
}
