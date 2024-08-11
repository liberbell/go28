package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/cat/", cat)
	http.HandleFunc("/data/", data)
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

func data(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("data.gohtml")
	if err != nil {
		log.Fatalln("parse error", err)
	}

	err = tpl.ExecuteTemplate(w, "data.gohtml", data)
	if err != nil {
		log.Fatalln("execute template error", err)
	}
}
