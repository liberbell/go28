package main

import (
	"encoding/json"
	"fmt"
)

type model struct {
	State    bool
	Pictures []string
}

func main() {
	m := model{}
	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(string(bs))
}
