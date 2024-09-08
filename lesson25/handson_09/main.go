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
		  </body>
		  </html>
		  `
	w.Write([]byte(s))
}

func encd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"James",
		"Bond",
		[]string{"Suite", "Gun", "Wry sense of humor"},
	}
}
