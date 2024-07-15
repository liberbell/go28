package main

import (
	"fmt"
)

type secretAgent struct {
	person
	licenseToKill bool
}

type person struct {
	fname string
	lname string
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says "good morning, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `"says "Shaken, not stirred"`)
}

func main() {
	// x = 7
	// fmt.Printf("%T", x)
	// fmt.Println(x)
	// xi := []int{2, 4, 7, 9, 42}
	// fmt.Println(xi)

	// m := map[string]int{
	// 	"Todd": 45,
	// 	"Job":  42,
	// }
	// fmt.Println(m)

	p1 := person{
		"Miss",
		"Moneypenny",
	}
	// fmt.Println(p1)
	p1.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	sa1.speak()
	sa1.person.speak()
}
