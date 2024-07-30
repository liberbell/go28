package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

type Record struct {
	Date time.Time
	Open float64
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	records := pts("table.csv")
}

func pts(filepath string) []Record {
	src, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
	}
}
