package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/", dogs)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("."))))
}

func dogs(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("Template did't execute.: ", err)
	}
}
