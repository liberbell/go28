package main

import (
	"html/template"
	"net/http"
)

type hotdog int

var tpl *template.Template

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func init() {
	tpl := template.Must(template.ParseFiles("index.gohtml"))
}
