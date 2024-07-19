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
	// sages := []string{"Gandhi", "MLK", "Budda", "Jesus", "Muhammad"}
	sages := map[string]string{
		"India": "Gandhi",
	}
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
