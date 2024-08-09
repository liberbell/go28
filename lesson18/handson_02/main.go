package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "doggy doggy doggy")
}

type hotcat int

func (m hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "kitty kitty kitty")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()

	http.ListenAndServe(":8080", d)
}
