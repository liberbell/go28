package main

import (
	"net/http"
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
