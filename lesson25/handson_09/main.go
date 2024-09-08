package main

import "net/http"

type person struct {
	Fname string
	Lname string
	Items []string
}

func bar(w http.ResponseWriter, r *http.Request) {
	s := `<!DOCTYPE html>
	      <head>
		  <meta charset="utf-8">
		  <title>F00</title>
		  </head>
		  <body>
		  <h1>You are at foo</h1>
		  `
}
