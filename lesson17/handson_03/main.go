package main

import (
	"html/template"
	"net/http"
)

type hotdog int

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
