package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

type course struct {
	Number, Name, Units, string
}

type semester struct {
	Term string
	Courses []course
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p1 := doubleZero{
		person{
			Name: "Ian Fleming",
			Age:  56,
		},
		true,
	}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
