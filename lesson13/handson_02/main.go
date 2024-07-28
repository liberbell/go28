package main

import "html/template"

type hotels struct {
	Name    string
	Address string
	City    string
	Zipcode int
	Region  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
