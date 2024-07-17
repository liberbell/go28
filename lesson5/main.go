package main

import (
	"html/template"
	"log"
)

func main() {
	tpl, err := template.ParseFiles("one.gmao")
	if err != nil {
		log.Fatalln(err)
	}
}
