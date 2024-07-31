package main

import (
	"html/template"
	"log"
	"os"
)

type Page struct {
	Title   string
	Heading string
	Input   string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	home := Page{
		Title:   "Nothing escaped",
		Heading: "Nothing is escaped with text/template",
		Input:   `<script>alert("YoW");</script>`,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", home)
	if err != nil {
		log.Fatalln(err)
	}
}
