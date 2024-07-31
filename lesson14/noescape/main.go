package main

import "html/template"

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
	home = Page{
		Title:   "Nothin escape",
		Heading: "Nothing escaped with text/template",
		Input:   `<script>alert("YoW");</script>`,
	}
}
