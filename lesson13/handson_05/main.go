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
		item{
			Name:        "Hamburgur",
			Description: "Cheese",
			Meal:        "Lunch",
			Price:       23.99,
		},
		item{
			Name:        "Curry",
			Description: "Spicy curry",
			Meal:        "Dinner",
			Price:       50.99,
		},
		item{
			Name:        "Beer",
			Description: "Dirty beer",
			Meal:        "Dinner",
			Price:       10.99,
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
