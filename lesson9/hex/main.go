package main

import (
	"math"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func doubule(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}
