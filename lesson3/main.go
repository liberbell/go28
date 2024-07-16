package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	name := "Tedd Baker"

	str := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello world</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`

	fmt.Println(str)
}
