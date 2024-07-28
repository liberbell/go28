package main

import "html/template"

type hotels struct {
	Name, Address, City, Zipcode, Region string
}

type hotel []hotels

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := hotels{
		hotel{
			Name:    "Hotel California",
			Address: "42 Sunset Boulevard",
			City:    "Los Angeles",
			Zip:     "95612",
			Region:  "southern",
		},
		hotel{
			Name:    "H",
			Address: "4",
			City:    "L",
			Zip:     "95612",
			Region:  "southern",
		},
	}
}
