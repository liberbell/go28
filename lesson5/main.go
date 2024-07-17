package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln("one", err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("two", err)
	}

	// tpl, err = tpl.ParseFiles("two.gmao", "vespa.gmao")
	// if err != nil {
	// 	log.Fatalln("three", err)
	// }
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gmao", nil)
	if err != nil {
		log.Fatalln("four", err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
