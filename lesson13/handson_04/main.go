package main

import (
	"html/template"
	"log"
	"os"
)

type hotel struct {
	Name, Address, City, Zip, Region string
}

type region struct {
	Region string
	Hotels []hotel
}

type Regions []region

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := Regions{
		region{
			Region: "Southern",
			Hotels: []hotel{
				hotel{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
				},
				hotel{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     "95612",
				},
			},
		},
		region{
			Region: "Nouthern",
			Hotels: []hotel{
				hotel{
					Name:    "Hotel alaska",
					Address: "someweare cold",
					City:    "Alaska",
					Zip:     "99999",
				},
				hotel{
					Name:    "C",
					Address: "10",
					City:    "A",
					Zip:     "99999",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
