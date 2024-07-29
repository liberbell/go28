package main

import (
	"html/template"
	"log"
	"os"
)

type item struct {
	Name, Description, Meal string
	Price                   float64
}

// type time struct {
// 	Time  string
// 	Times []time
// }

type items []item

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := items{
		item{
			Name:        "Bread",
			Description: "Hard bread",
			Meal:        "Breakfast",
			Price:       5.99,
		},
		item{
			Name:        "Salad",
			Description: "Normal salad",
			Meal:        "Breakfast",
			Price:       3.99,
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
