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

type restaurant struct {
	Name string
	Menu menu
}

type restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := restaurants{
		restaurant{
			Name: "Old turky",
			Menu: menu{
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
			},
		},
		restaurant{
			Name: "KFC",
			Menu: menu{
				meal{
					Meal: "Breakfast",
					Item: []item{
						item{
							Name:        "Hot dog",
							Description: "so so",
							Price:       7.99,
						},
						item{
							Name:        "Coffee",
							Description: "Too hot",
							Price:       2.99,
						},
					},
				},
				meal{
					Meal: "Lunch",
					Item: []item{
						item{
							Name:        "Chiken",
							Description: "Oily chiken",
							Price:       8.99,
						},
						item{
							Name:        "Old salad",
							Description: "Fresh color",
							Price:       5.99,
						},
					},
				},
				meal{
					Meal: "Dinner",
					Item: []item{
						item{
							Name:        "Potato",
							Description: "Not bad",
							Price:       4.99,
						},
						item{
							Name:        "Cola",
							Description: "Near sugar",
							Price:       1.99,
						},
					},
				},
			},
		},
	}
		
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
