package main

import (
	"html/template"
	"log"
	"os"
)

type item struct {
	Name, Description string
	Price             float64
}

type meal struct {
	Meal string
	Item []item
}

type menu []meal

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := menu{
		meal{
			Meal: "Breakfast",
			Item: []item{
				item{
					Name:        "Bread",
					Description: "Hard bread",
					Price:       5.99,
				},
				item{
					Name:        "Salad",
					Description: "Normal salad",
					Price:       3.99,
				},
			},
		},
		meal{
			Meal: "Lunch",
			Item: []item{
				item{
					Name:        "Hamburgur",
					Description: "Cheese",
					Price:       23.99,
				},
				item{
					Name:        "Orange juice",
					Description: "Fresh juice",
					Price:       8.99,
				},
			},
		},
		meal{
			Meal: "Dinner",
			Item: []item{
				item{
					Name:        "Curry",
					Description: "Spicy curry",
					Price:       50.99,
				},
				item{
					Name:        "Beer",
					Description: "Dirty beer",
					Price:       10.99,
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
