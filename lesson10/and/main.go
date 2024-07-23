package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name  string
	Motto string
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	u1 := user{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
		Admin: false,
	}

	u2 := user{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
