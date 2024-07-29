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
		}
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
