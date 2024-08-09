package main

import (
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "doggy doggy doggy")
}

func c(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "kitty kitty kitty")
}

func main() {
	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))
	http.ListenAndServe(":8080", nil)
}
