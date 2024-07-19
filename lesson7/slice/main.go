package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	sages := []string{"Gandhi", "MLK", "Budda", "Jesus", "Muhammad"}
	err := tpl.Execute(os.Stdout, sages)
}
