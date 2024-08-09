package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "doggy doggy doggy")
}

func c(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "kitty kitty kitty")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/dog/", d)
	mux.HandleFunc("/cat", c)
	http.ListenAndServe(":8080", mux)
}
