package main

import (
	"html/template"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/", index)
}
