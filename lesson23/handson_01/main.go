package main

import "html/template"

type user struct {
	UserName string
	Password string
	First    string
	Last     string
}

var tpl *template.Template
