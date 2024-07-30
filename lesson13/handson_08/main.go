package main

import (
	"os"
	"time"
)

type Record struct {
	Date time.Time
	Open float64
}

func main() {
	table_data, err := os.Open("table.csv", "r")

}
