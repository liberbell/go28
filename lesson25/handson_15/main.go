package main

type model struct {
	state    bool
	pictures []string
}

func main() {
	m := model{
		state: true,
		pictures: []string{
			"one.jpg",
			"two.jpg",
			"three.jpg",
		},
	}
}
