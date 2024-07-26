package main

import "text/template"

type person struct {
	Name string
	Age  int
}

var tpl *template.Template
