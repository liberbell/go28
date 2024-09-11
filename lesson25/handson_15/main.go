package main

import (
	"encoding/json"
	"log"
	"os"
)

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

	bs, err := json.Marshal(m)
	if err != nil {
		log.Println("error: ", err)
	}
	os.Stdout.Write(bs)
}
