package main

import (
	"html/template"
	"net/http"

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

func index(w http.ResponseWriter, r *http.Request)  {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	Handleerr
}

func HandleError(w http.ResponseWriter, err error)  {
	if err != nil {
		http.Error(w err.Error(), )
	}
}