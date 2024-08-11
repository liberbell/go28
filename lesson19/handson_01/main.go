package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/cat", cat)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is an index.")
}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "doggy doggy doggy")
}

func cat(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "kitty kitty kitty")
}
