package main

import (
	"html/template"
	"log"
	"os"
)

type menu struct {
	Name, Description string
	Price             string
}

type time struct {
	Time  string
	Times []time
}

type Times []time

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := Times{
		time{
			Time: "Breakfast",
			Times: []hotel{
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
