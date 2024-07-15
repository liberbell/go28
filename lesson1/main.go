package main

import "fmt"

type hotdog int
type person struct {
	fname string
	lname string
}

func main() {
	// x = 7
	// fmt.Printf("%T", x)
	// fmt.Println(x)
	xi := []int{2, 4, 7, 9, 42}
	fmt.Println(xi)

	m := map[string]int{
		"Todd": 45,
		"Job":  42,
	}
	fmt.Println(m)

	p1 := person{
		"Miss",
		"Moneypenny",
	}
	fmt.Println(p1)
}
