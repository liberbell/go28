package main

import (
	"encoding/json"
	"log"
	"os"
)

type model struct {
	State    bool
	Pictures []string
}

func main() {
	m := model{
		State: true,
		Pictures: []string{
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
