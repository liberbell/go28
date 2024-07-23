package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	xs := []string{"zero", "one", "two", "three", "four", "five"}

	// data := struct {
	// 	Words []string
	// 	Lname string
	// }{
	// 	xs,
	// 	"Mcleod",
	// }

	err := tpl.Execute(os.Stdout, xs)
	if err != nil {
		log.Fatalln(err)
	}
}
