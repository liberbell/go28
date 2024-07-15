package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	name := "Tedd baker"

	tpl := `
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

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating index.html")
	}

	fmt.Println(tpl)
}
