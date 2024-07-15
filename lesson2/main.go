package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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
		log.Fatal("error creating index.html", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))

	fmt.Println(tpl)
}
