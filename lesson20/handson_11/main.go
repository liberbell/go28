package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
}

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Fprintln(w, "go look at your terminal")
}
