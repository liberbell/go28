package main

import (
	"html/template"
	"strings"
)

var fm = template.FuncMap{
	"uc": strings.ToUpper,
}

func init() {
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}
