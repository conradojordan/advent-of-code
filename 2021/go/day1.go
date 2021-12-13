package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const FilePath = "2021/inputs/day1.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile(FilePath)
	check(err)

	splitted_data := strings.Split(strings.TrimSpace(string(data)), "\n")
	last_measurement := math.Inf(1)
	var increased_count int

	for _, v := range splitted_data {
		measurement, err := strconv.ParseFloat(v, 64)
		check(err)

		if measurement > last_measurement {
			increased_count++
		}
		last_measurement = measurement
	}
	fmt.Printf("Number of times the measurement increased is %v\n", increased_count)
}
