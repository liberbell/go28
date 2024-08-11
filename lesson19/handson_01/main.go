package main

import "net/http"

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/cat", cat)

}
